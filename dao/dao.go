package dao

import "github.com/jinzhu/gorm"

type dao struct {
	db *gorm.DB
}

// NewDao is a function that defines a new dao instance
func NewDao(db *gorm.DB) Dao {
	return &dao{
		db: db,
	}
}

func (d *dao) ensureSchema() error {
	return d.db.AutoMigrate(Material{}).Error
}

func (d *dao) Get() (*Material, error) {
	var m Material

	if err := d.db.Find(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (d *dao) GetOne(n string) (*Material, error) {
	var m Material
	sql := "name = ?"

	if err := d.db.Where(sql, n).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (d *dao) Create(m *Material) error {
	return d.db.Create(&m).Error
}

func (d *dao) Update(m *Material) error {
	return d.db.Save(&m).Error
}

func (d *dao) Delete(m *Material) error {
	return d.db.Delete(&m).Error
}
