package persistence

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
)

const (
	projectID  = "gomediasandbox"
	bucketName = "hall-of-fame"
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
}

var uploader *ClientUploader

func init() {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
	}

	uploader = &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
	}

}

func uploadFile(prefix string, filename string, file multipart.File) error {
	return uploader.upload(prefix+"/"+filename, file)
}

func downloadFile(prefix string, filename string) ([]byte, error) {
	return uploader.download(prefix + "/" + filename)
}

func (c *ClientUploader) upload(path string, file multipart.File) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(path).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}

func (c *ClientUploader) download(path string) ([]byte, error) {
	ctx := context.Background()
	rc, err := c.cl.Bucket(c.bucketName).Object(path).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	return ioutil.ReadAll(rc)
}
