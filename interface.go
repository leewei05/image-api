package image

import "time"

// Material is a struct for material images
type Material struct {
	// ID is Material's primary key.
	ID uint `gorm:"primary_key,AUTO_INCREMENT"`
	// Name is Material's name.
	Name  string  `gorm:"size:150" sql:"not null"`
	URL   string  `gorm:"size:150" sql:"not null"`
	Price float64 `sql:"not null"`
	// CreatedAt for gorm to insert create time.
	CreatedAt time.Time
	// UpdatedAt for gorm to insert update time.
	UpdatedAt time.Time
}
