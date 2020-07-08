package dao

import "time"

// Dao is an interface
type Dao interface {
	Get() (*Material, error)
	GetOne(string) (*Material, error)
	Create(*Material) error
	Update(*Material) error
	Delete(*Material) error
}

// Material is a struct for material images
type Material struct {
	// ID is Product's primary key.
	ID uint64 `gorm:"primary_key"`
	// Name is Product's name.
	Name  string  `gorm:"size:150" sql:"not null"`
	URL   string  `gorm:"size:150" sql:"not null"`
	Price float64 `sql:"not null"`
	// CreatedAt for gorm to insert create time.
	CreatedAt time.Time
	// UpdatedAt for gorm to insert update time.
	UpdatedAt time.Time
}
