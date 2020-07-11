package dao

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/leewei05/image-api"
	"github.com/leewei05/image-api/rest"
)

type postgresDao struct {
	db *gorm.DB
}

// Automatically migrate your schema, to keep your schema update to date.
func (p *postgresDao) ensureSchema() error {
	return p.db.AutoMigrate(image.Material{}).Error
}

// NewPostgres is a function that defines a new dao instance
func NewPostgres(db *gorm.DB) (rest.PostgresDao, error) {
	dao := &postgresDao{
		db: db,
	}

	if err := dao.ensureSchema(); err != nil {
		return nil, err
	}

	return dao, nil
}

func (p *postgresDao) Get() (*[]image.Material, error) {
	var materials []image.Material

	if err := p.db.Find(&materials).Error; err != nil {
		return nil, err
	}

	return &materials, nil
}

func (p *postgresDao) GetOne(n uint) (*image.Material, error) {
	if n == 0 {
		return nil, errors.New("ID cannot be empty")
	}
	var material image.Material
	sql := "id = ?"

	if err := p.db.Where(sql, n).First(&material).Error; err != nil {
		return nil, err
	}

	return &material, nil
}

func (p *postgresDao) Create(m *image.Material) error {
	if m.Name == "" {
		return errors.New("Name cannot be empty")
	}

	return p.db.Create(&m).Error
}

func (p *postgresDao) Update(m *image.Material) error {
	if m.ID == 0 {
		return errors.New("ID cannot be empty")
	}

	return p.db.Save(&m).Error
}

func (p *postgresDao) Delete(n uint) error {
	if n == 0 {
		return errors.New("ID cannot be empty")
	}
	var material image.Material
	material.ID = n

	return p.db.Delete(&material).Error
}
