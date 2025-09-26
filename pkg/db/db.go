package db

import (
	"go-links-shorter/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDB(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn))
	if err != nil {
		panic(err)
	}
	return &Db{db}
}
