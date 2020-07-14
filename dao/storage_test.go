package dao

import (
	"context"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/stretchr/testify/suite"
)

var (
	testBucket         = "lee-image-test"
	testImageNotExists = "dog.png"
	testImageExists    = "cat.jpg"
	testImageWrite     = "turtle.jpg"

	contentTypeImg = "image/jpg"
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

func (s *storageTestSuite) TestGetObject() {
	b, err := s.impl.GetObject(testBucket, testImageExists)
	s.NoError(err)

	cat, err := fileToBytes("../assets/cat.jpg")
	s.NoError(err)
	s.Equal(cat, b)
}

func (s *storageTestSuite) TestWriteContentToGCS() {
	filePath := path.Join("gcs_write_test", "turtle.jpg")
	turtle, err := fileToBytes("../assets/turtle.jpg")
	s.NoError(err)

	err = s.impl.WriteContentTo(testBucket, filePath, contentTypeImg, turtle)
	s.NoError(err)

	imageWritePath := path.Join("gcs_write_test", testImageWrite)
	t, err := s.impl.GetObject(testBucket, imageWritePath)
	s.NoError(err)
	s.Equal(turtle, t)
}

func fileToBytes(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return b, nil
}
