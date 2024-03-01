package blog

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/iterator"
	"io"
	"log"
	"time"
)

type BlogEntry struct {
	Id         string    `firestore:"id"`
	Date       time.Time `firestore:"date"`
	Author     string    `firestore:"author"`
	Title      string    `firestore:"title"`
	Lead       string    `firestore:"lead"`
	Tags       string    `firestore:"tags"`
	DocumentId string    `firestore:"documentId"`
}

type RepositoryProperties struct {
	ProjectId  string
	Collection string
}

func Save(repo RepositoryProperties, blog *BlogEntry) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, repo.ProjectId)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
		return err
	}
	defer client.Close()

	_, err = client.Collection(repo.Collection).Doc(blog.Id).Set(ctx, blog)
	if err != nil {
		log.Fatalf("Failed to save Firestore document: %v", err)
		return err
	}

	return nil
}

func List(repo RepositoryProperties) ([]BlogEntry, error) {
	var results []BlogEntry

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	client, err := firestore.NewClient(ctx, repo.ProjectId)
	if err != nil {
		log.Printf("Failed to create Firestore client: %v", err)
		return nil, err
	}
	defer client.Close()

	query := client.Collection(repo.Collection).
		//Where("author_name", "==", authorName).
		OrderBy("date", firestore.Desc)

	iter := query.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var entry BlogEntry
		if err := doc.DataTo(&entry); err != nil {
			return nil, err
		}
		results = append(results, entry)
	}

	return results, nil
}

func FindById(repo RepositoryProperties, id string) (*BlogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	client, err := firestore.NewClient(ctx, repo.ProjectId)
	if err != nil {
		log.Printf("Failed to create Firestore client: %v", err)
		return nil, err
	}
	defer client.Close()

	query := client.Collection(repo.Collection).
		Where("id", "==", id).
		OrderBy("date", firestore.Desc)

	iter := query.Documents(ctx)

	doc, err := iter.Next()
	if err != nil {
		return nil, err
	}

	var entry BlogEntry
	if err := doc.DataTo(&entry); err != nil {
		return nil, err
	}

	return &entry, nil
}

func GetDocument(id string) (string, error) {
	bucketName := "huepattl-de-blogs"
	objectName := id
	ctx := context.Background()

	// create a client
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "[not found]", err
	}
	defer client.Close()

	// open the object from GCS
	reader, err := client.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		return "[not found]", err
	}
	defer reader.Close()

	// read the content of the object into a string
	content, err := io.ReadAll(reader)
	if err != nil {
		return "[not found]", err
	}

	return string(content), nil
}
