package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/leewei05/image-api"
	"github.com/leewei05/image-api/rest"
)

type postgresDao struct {
	db *gorm.DB
}

// NewPostgres is a function that defines a new dao instance
func NewPostgres(db *gorm.DB) rest.PostgresDao {
	return &postgresDao{
		db: db,
	}
}

func (p *postgresDao) ensureSchema() error {
	return p.db.AutoMigrate(image.Material{}).Error
}

func (p *postgresDao) Get() (*image.Material, error) {
	var m image.Material

	if err := p.db.Find(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (p *postgresDao) GetOne(n string) (*image.Material, error) {
	var m image.Material
	sql := "name = ?"

	if err := p.db.Where(sql, n).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (p *postgresDao) Create(m *image.Material) error {
	return p.db.Create(&m).Error
}

func (p *postgresDao) Update(m *image.Material) error {
	return p.db.Save(&m).Error
}

func (p *postgresDao) Delete(m *image.Material) error {
	return p.db.Delete(&m).Error
}
