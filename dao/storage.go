package dao

import (
	"context"
	"io/ioutil"
	"time"

	"cloud.google.com/go/storage"
)

var (
	gcsAPITimeout = 5 * time.Minute
)

// NewStorage creates a new Cloud Storage Manager
func NewStorage(c *storage.Client) StorageDao {
	return &storageDao{
		client: c,
	}
}

type storageDao struct {
	client *storage.Client
}

func (s *storageDao) CheckExists(bucketName, object string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), gcsAPITimeout)
	defer cancel()

	_, err := s.client.Bucket(bucketName).Object(object).Attrs(ctx)
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

func (s *storageDao) WriteObject(bucketName, dstFile, contentType string, b []byte) error {
	ctx := context.Background()
	w := s.client.Bucket(bucketName).Object(dstFile).NewWriter(ctx)
	w.ContentType = contentType

	_, err := w.Write(b)
	if err != nil {
		return err
	}
	defer w.Close()

	return err
}

func (s *storageDao) RemoveObject(bucketName, path string) (bool, error) {
	deleteCtx, deleteCancel := context.WithTimeout(context.Background(), gcsAPITimeout)
	defer deleteCancel()

	err := s.client.Bucket(bucketName).Object(path).Delete(deleteCtx)
	if err == storage.ErrObjectNotExist {
		return false, err
	}
	if err != nil {
		return true, err
	}

	return true, nil
}
