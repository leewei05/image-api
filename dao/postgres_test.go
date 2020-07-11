package dao

import (
	"fmt"
	"log"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/leewei05/image-api"
	"github.com/stretchr/testify/suite"

	_ "github.com/lib/pq"
)

var (
	testImages = []image.Material{
		{
			Name:  "Dog",
			URL:   "gs://image-api-v1/dog.png",
			Price: 30,
		},
		{
			Name:  "Cat",
			URL:   "gs://image-api-v1/cat.png",
			Price: 20,
		},
	}

	dbHost = "localhost"
	dbPort = 5432
	dbUser = "postgres"
	dbPwd  = "123456"
	dbName = "db"
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
	var err error
	pgStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPwd, dbName)

	p.db, err = gorm.Open("postgres", pgStr)
	if err != nil {
		log.Panic("Cannot open testing PostgreSQL database")
	}
	p.NoError(err)

	pg, err := NewPostgres(p.db)
	p.NoError(err)

	p.impl = pg.(*postgresDao)

	p.insertTestEntries()
}

func (p *postgresTestSuite) insertTestEntries() {
	for _, i := range testImages {
		err := p.impl.Create(&i)
		p.NoError(err)
	}
}

func (p *postgresTestSuite) TestGet() {
	_, err := p.impl.Get()
	p.NoError(err)
}

func (p *postgresTestSuite) TestGetOneEntity() {
	_, err := p.impl.GetOne(1)
	p.NoError(err)
}

func (p *postgresTestSuite) TestUpdate() {
	updateImage := &image.Material{
		ID:    1,
		Name:  "Turtle",
		URL:   "gs://image-api-v1/turtle.png",
		Price: 100,
	}

	err := p.impl.Update(updateImage)
	p.NoError(err)
}

func (p *postgresTestSuite) TestDelete() {
	err := p.impl.Delete(2)
	p.NoError(err)

	err = p.impl.Delete(0)
	p.Error(err)
}
