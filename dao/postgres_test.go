package dao

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
)

type postgresTestSuite struct {
	suite.Suite
	impl *postgresDao
	db   *gorm.DB
}

func TestPostgresDao(t *testing.T) {
	suite.Run(t, new(postgresTestSuite))
}

func (p *postgresTestSuite) SetupTest() {
	pg := NewPostgres(p.db)

	p.impl = pg.(*postgresDao)
}
