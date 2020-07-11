package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/leewei05/image-api"
	"github.com/leewei05/image-api/rest"
)

type postgres struct {
	db *gorm.DB
}

// NewDao is a function that defines a new dao instance
func NewDao(db *gorm.DB) rest.ImageDao {
	return &postgres{
		db: db,
	}
}

func (p *postgres) ensureSchema() error {
	return p.db.AutoMigrate(image.Material{}).Error
}

func (p *postgres) Get() (*image.Material, error) {
	var m image.Material

	if err := p.db.Find(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (p *postgres) GetOne(n string) (*image.Material, error) {
	var m image.Material
	sql := "name = ?"

	if err := p.db.Where(sql, n).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (p *postgres) Create(m *image.Material) error {
	return p.db.Create(&m).Error
}

func (p *postgres) Update(m *image.Material) error {
	return p.db.Save(&m).Error
}

func (p *postgres) Delete(m *image.Material) error {
	return p.db.Delete(&m).Error
}
