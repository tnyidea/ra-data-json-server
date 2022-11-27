package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDbSession(connstring string) (*gorm.DB, error) {
	// See: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
	// Example:
	// connstring := "host=localhost user=gorm password=gorm dbname=gorm port=55000 sslmode=disable"
	// connstring := "postgresql://gorm:gorm@localhost:55000/gorm?sslmode=disable"

	db, err := gorm.Open(postgres.Open(connstring), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
