package core

import (
	"errors"
	"path"

	"cloud.google.com/go/storage"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/leewei05/image-api/dao"
)

type core struct {
	db   *gorm.DB
	rdb  *redis.Client
	stor *storage.Client

	pg    dao.PostgresDao
	redis dao.RedisDao
	gcs   dao.StorageDao
}

// NewCore defines a new instance of core
func NewCore(
	db *gorm.DB,
	rdb *redis.Client,
	stor *storage.Client,
) (Core, error) {
	pg, err := dao.NewPostgres(db)
	if err != nil {
		return nil, err
	}

	redis := dao.NewRedisDao(rdb)
	gcs := dao.NewStorage(stor)

	return &core{
		db:    db,
		rdb:   rdb,
		stor:  stor,
		pg:    pg,
		redis: redis,
		gcs:   gcs,
	}, nil
}

func (c *core) CreateImage(bucketName, dst, imageName, contentType string, bs []byte) error {
	path := path.Join(dst, imageName)
	if ok, err := c.gcs.CheckExists(bucketName, path); ok {
		if err != nil {
			return err
		}

		return errors.New("Object exists")
	}

	err := c.gcs.WriteObject(bucketName, dst, contentType, bs)
	if err != nil {
		return nil
	}

	err = c.redis.Set(imageName, path)
	if err != nil {
		return err
	}

	return nil
}

func (c *core) GetImage(id int) string {
	return "hey bro"
}
