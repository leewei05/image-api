package rest

import (
	"database/sql"
	"net/http"
)

// Rest is a struct
type Rest struct {
	db *sql.DB
}

// NewRest is a rest
func NewRest(
	db *sql.DB,
) *Rest {
	return &Rest{
		db: db,
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
