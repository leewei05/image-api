package rest

import (
	"github.com/leewei05/image-api"
)

// PostgresDao is an interface
type PostgresDao interface {
	Get() (*[]image.Material, error)
	GetOne(uint) (*image.Material, error)
	Create(*image.Material) error
	Update(*image.Material) error
	Delete(uint) error
}
