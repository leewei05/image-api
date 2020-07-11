package rest

import (
	"github.com/leewei05/image-api"
)

// ImageDao is an interface
type ImageDao interface {
	Get() (*image.Material, error)
	GetOne(string) (*image.Material, error)
	Create(*image.Material) error
	Update(*image.Material) error
	Delete(*image.Material) error
}
