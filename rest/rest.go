package rest

import (
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/leewei05/image-api/core"
)

// Rest is a struct
type rest struct {
	db   *gorm.DB
	rdb  *redis.Client
	gcs  *storage.Client
	core core.Core
}

// NewRest is a rest
func NewRest(
	db *gorm.DB,
	rdb *redis.Client,
	gcs *storage.Client,
) (Rest, error) {
	c, err := core.NewCore(db, rdb, gcs)
	if err != nil {
		return nil, err
	}

	return &rest{
		db:   db,
		rdb:  rdb,
		gcs:  gcs,
		core: c,
	}, nil
}

// GetProduct is
func (ri *rest) GetProduct(w http.ResponseWriter, r *http.Request) {
	res := ri.core.GetImage(0)

	w.Write([]byte(res))
}

// CreateProduct is
func (ri *rest) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

// UpdateProduct is
func (ri *rest) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

// DeleteProduct is
func (ri *rest) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
