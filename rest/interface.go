package rest

import "net/http"

// Rest is an interface for rest API
type Rest interface {
	GetProduct(http.ResponseWriter, *http.Request)
	CreateProduct(http.ResponseWriter, *http.Request)
	UpdateProduct(http.ResponseWriter, *http.Request)
	DeleteProduct(http.ResponseWriter, *http.Request)
}
