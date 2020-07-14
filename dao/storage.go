package dao

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/leewei05/image-api/rest"
)

var (
	gcsAPITimeout = 5 * time.Minute
)

// NewStorage creates a new Cloud Storage Manager
func NewStorage(c *storage.Client) rest.StorageDao {
	return &storageDao{
		client: c,
	}
}

type storageDao struct {
	client *storage.Client
}

func (s *storageDao) CheckExists(bucket, object string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), gcsAPITimeout)
	defer cancel()

	_, err := s.client.Bucket(bucket).Object(object).Attrs(ctx)
	if err == storage.ErrObjectNotExist {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *storageDao) GetObject(bucketName, object string) ([]byte, error) {
	ctx := context.Background()
	r, err := s.client.
		Bucket(bucketName).
		Object(object).
		NewReader(ctx)

	if err != nil {
		return []byte{}, err
	}
	defer r.Close()

	return ioutil.ReadAll(r)
}

func (s *storageDao) WriteContentTo(bucket, dstFile, contentType string, b []byte) error {
	ctx := context.Background()
	w := s.client.Bucket(bucket).Object(dstFile).NewWriter(ctx)
	w.ContentType = contentType

	_, err := w.Write(b)
	if err != nil {
		return err
	}
	defer w.Close()

	return err
}

func (s *storageDao) UploadFile(bucketName, src, dst string) error {
	ctx, cancel := context.WithTimeout(context.Background(), gcsAPITimeout)
	defer cancel()

	writer := s.client.
		Bucket(bucketName).
		Object(dst).
		NewWriter(ctx)

	reader, err := os.Open(src)
	if err != nil {
		return err
	}
	defer writer.Close()
	defer reader.Close()

	return s.fileCopy(writer, reader)
}

func (s *storageDao) fileCopy(w io.Writer, r io.Reader) error {
	_, err := io.Copy(w, r)
	return err
}
