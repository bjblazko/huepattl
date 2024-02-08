package blog

import (
	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
)

type Connection struct {
	Client  *bigquery.Client
	Context context.Context
}

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
	Text       string         `bigquery:"text"`
}

func CreateClient(projectId string) (Connection, error) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectId)
	if err != nil {
		//return nil, err
	}
	conn := Connection{
		Context: ctx,
		Client:  client,
	}
	defer client.Close()
	return conn, nil
}

// Lists all blog entries from the blog index, sorted from newest
// at the beginning.
func List(connection Connection, table BigQueryTable) ([]Entry, error) {
	var results []Entry

	query := connection.Client.Query(fmt.Sprintf(
		"SELECT * FROM `%s.%s` ORDER BY date DESC LIMIT 1000",
		table.Dataset, table.Table),
	)
	it, err := query.Read(connection.Context)
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

func Insert(entries []*Entry, connection Connection, table BigQueryTable) (int, error) {
	inserter := connection.Client.Dataset(table.Dataset).Table(table.Table).Inserter()
	if err := inserter.Put(connection.Context, entries); err != nil {
		return 0, err
	}

	return len(entries), nil
}

func DeleteTable(connection Connection, table BigQueryTable) error {
	err := connection.Client.Dataset(table.Dataset).Table(table.Table).
		Delete(connection.Context)
	if err != nil {
		return err
	}

	return nil
}

func CreateTable(connection Connection, table BigQueryTable) error {
	schema := bigquery.Schema{
		{Name: "date", Type: bigquery.DateTimeFieldType, Required: true},
		{Name: "author", Type: bigquery.StringFieldType, Required: true},
		{Name: "title", Type: bigquery.StringFieldType, Required: true},
		{Name: "text", Type: bigquery.StringFieldType, Required: true},
	}

	// Create table metadata
	metadata := &bigquery.TableMetadata{
		Schema: schema,
	}

	if err := connection.Client.Dataset(table.Dataset).Table(table.Table).
		Create(connection.Context, metadata); err != nil {
		return err
	}
	return nil
}
