package dao

import (
	"context"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/stretchr/testify/suite"
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
