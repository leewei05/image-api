package core

import (
	"github.com/leewei05/image-api"
)

// PostgresDao is an interface for accessing PostgreSQL
type PostgresDao interface {
	Get() (*[]image.Material, error)
	GetOne(uint) (*image.Material, error)
	Create(*image.Material) error
	Update(*image.Material) error
	Delete(uint) error
}

// RedisDao is an interface for accessing Redis
type RedisDao interface {
	Set(string, string) error
	Get(string) (string, error)
	Flush() error
}

// StorageDao is an interface for managing Cloud Storage
type StorageDao interface {
	CheckExists(string, string) (bool, error)
	GetObject(string, string) ([]byte, error)
	WriteObject(string, string, string, []byte) error
	RemoveObject(string, string) (bool, error)
}

// Core is an interface of managing business logic
type Core interface {
}
