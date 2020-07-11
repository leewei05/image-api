package rest

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Rest is a struct
type Rest struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewRest is a rest
func NewRest(
	db *gorm.DB,
	rdb *redis.Client,
) *Rest {
	return &Rest{
		db:  db,
		rdb: rdb,
	}
}

// GetProduct is
func (ri *Rest) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

// CreateProduct is
func (ri *Rest) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

// UpdateProduct is
func (ri *Rest) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

// DeleteProduct is
func (ri *Rest) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
