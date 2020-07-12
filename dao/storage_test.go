package dao

import (
	"context"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/stretchr/testify/suite"
)

var (
	testBucket         = "lee-image-test"
	testImageNotExists = "dog.png"
	testImageExists    = "cat.jpg"
)

type storageTestSuite struct {
	suite.Suite
	impl *storageDao
}

func TestStorageTestSuit(t *testing.T) {
	suite.Run(t, new(storageTestSuite))
}

func (s *storageTestSuite) SetupTest() {
	client, err := storage.NewClient(context.Background())
	s.NoError(err)

	gcs := NewStorage(client)

	s.impl = gcs.(*storageDao)
}

func (s *storageTestSuite) TestCheckExists() {
	b, err := s.impl.CheckExists(testBucket, testImageExists)
	s.NoError(err)
	s.Equal(true, b)

	b, err = s.impl.CheckExists(testBucket, testImageNotExists)
	s.NoError(err)
	s.Equal(false, b)
}

func (s *storageTestSuite) TestUploadFile() {

}
