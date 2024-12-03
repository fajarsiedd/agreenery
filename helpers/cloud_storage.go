package helpers

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

type UploaderParams struct {
	File         multipart.File
	OldObjectURL string
}

func UploadFile(params UploaderParams) (string, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	bucketName := "agreenery"

	bucket := client.Bucket(bucketName)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	// DELETE OLD OBJECT IF EXIST
	if params.OldObjectURL != "" {
		bucket.Object("uploads/" + params.OldObjectURL).Delete(ctx)
	}

	// CREATE NEW OBJECT
	obj := bucket.Object("uploads/" + uuid.New().String())

	writer := obj.NewWriter(ctx)
	if _, err := io.Copy(writer, params.File); err != nil {
		return "", err
	}
	if err := writer.Close(); err != nil {
		return "", err
	}

	url := fmt.Sprintf("http://storage.googleapis.com/%s/%s", obj.BucketName(), obj.ObjectName())

	return url, nil
}
