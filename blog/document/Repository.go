package document

import (
	"cloud.google.com/go/storage"
	"context"
	"io"
)

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
