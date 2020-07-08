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
	if err := d.db.Create(&m).Error; err != nil {
		return err
	}
	return nil
}

func (d *dao) Update() {}

func (d *dao) Delete() {}
