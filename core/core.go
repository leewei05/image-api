package core

import (
	"cloud.google.com/go/storage"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type core struct {
	db  *gorm.DB
	rdb *redis.Client
	gcs *storage.Client
}

// NewCore defines a new instance of core
func NewCore(
	db *gorm.DB,
	rdb *redis.Client,
	gcs *storage.Client,
) Core {
	return &core{
		db:  db,
		rdb: rdb,
		gcs: gcs,
	}
}
