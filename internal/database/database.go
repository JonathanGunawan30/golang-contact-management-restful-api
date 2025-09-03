package database

import "gorm.io/gorm"

type Database interface {
	Connect(dsn string) error
	GetDB() *gorm.DB
	Close() error
}
