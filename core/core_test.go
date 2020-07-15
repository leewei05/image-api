package core

import (
	"testing"

	"github.com/c2fo/testify/suite"
)

type TestCoreSuite struct {
	suite.Suite
}

func NewCoreTestSuite(t *testing.T) {
	suite.Run(t, new(TestCoreSuite))
}
