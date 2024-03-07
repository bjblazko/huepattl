package index

import (
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"huepattl.de/blog"
	"log"
	"time"
)

type RepositoryProperties struct {
	ProjectId  string // Google Cloud project ID
	Collection string // Firestore collection
}

// Lists all blog entries from Cloud Firestore
// sorted by date, newest first.
func List(repo RepositoryProperties) ([]blog.Entry, error) {
	var results []blog.Entry

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

		var entry blog.Entry
		if err := doc.DataTo(&entry); err != nil {
			return nil, err
		}
		results = append(results, entry)
	}

	return results, nil
}

// Returns a blog entry using a given id.
func FindById(repo RepositoryProperties, id string) (*blog.Entry, error) {
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

	var entry blog.Entry
	if err := doc.DataTo(&entry); err != nil {
		return nil, err
	}

	return &entry, nil
}

// Persists a blog entry to Cloud Firestore index
func Save(repo RepositoryProperties, blog *blog.Entry) error {
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
