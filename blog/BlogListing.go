package blog

import (
	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
)

type BigQueryTable struct {
	ProjectId string
	Dataset   string
	Table     string
}

// A blog entry
type Entry struct {
	DateTime   civil.DateTime `bigquery:"date"`
	AuthorName string         `bigquery:"author"`
	Title      string         `bigquery:"title"`
	Preview    string         `bigquery:"preview"`
	Bucket     string         `bigquery:"bucket"`
	Filename   string         `bigquery:"filename"`
}

// Lists all blog entries from the blog index, sorted from newest
// at the beginning.
func List(table BigQueryTable) ([]Entry, error) {
	var results []Entry

	log.Println("load all blogs...")
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, table.ProjectId)
	if err != nil {
		return nil, fmt.Errorf("failed to create BigQuery client: %v", err)
	}
	defer client.Close()

	query := client.Query(fmt.Sprintf("SELECT * FROM %s.%s", table.Dataset, table.Table))
	it, err := query.Read(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to run query: %v", err)
	}

	for {
		var entry Entry
		err := it.Next(&entry)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error fetching results: %v", err)
		}
		results = append(results, entry)
	}

	return results, nil
}

func Insert(entries []*Entry, table BigQueryTable) (int, error) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, table.ProjectId)
	if err != nil {
		return 0, fmt.Errorf("bigquery.NewClient: %w", err)
	}
	defer client.Close()

	inserter := client.Dataset(table.Dataset).Table(table.Table).Inserter()
	if err := inserter.Put(ctx, entries); err != nil {
		return 0, err
	}

	return len(entries), nil
}

func DeleteTable(table BigQueryTable) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, table.ProjectId)
	if err != nil {
		return err
	}
	err = client.Dataset(table.Dataset).Table(table.Table).Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}

func CreateTable(table BigQueryTable) error {
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, table.ProjectId)
	if err != nil {
		return err
	}
	defer client.Close()

	schema := bigquery.Schema{
		{Name: "date", Type: bigquery.DateFieldType, Required: true},
		{Name: "author", Type: bigquery.StringFieldType, Required: true},
		{Name: "title", Type: bigquery.StringFieldType, Required: true},
		{Name: "preview", Type: bigquery.StringFieldType, Required: true},
		{Name: "bucket", Type: bigquery.StringFieldType, Required: true},
		{Name: "filename", Type: bigquery.StringFieldType, Required: true},
	}

	// Create table metadata
	metadata := &bigquery.TableMetadata{
		Schema: schema,
	}

	if err := client.Dataset(table.Dataset).Table(table.Table).Create(ctx, metadata); err != nil {
		return err
	}
	return nil
}
