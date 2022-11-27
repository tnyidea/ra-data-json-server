package model

import (
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type DB struct {
	db *gorm.DB
}

func NewDatabaseConnection(dsn string) (DB, error) {
	// Example
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return DB{}, err
	}

	return DB{
		db: db,
	}, nil
}

func (p *DB) LoadSampleData(filename string) error {
	err := p.db.AutoMigrate(&Address{})
	if err != nil {
		return err
	}

	b, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var addressList []Address
	err = json.Unmarshal(b, &addressList)
	if err != nil {
		return err
	}

	for _, address := range addressList {
		p.db.Create(&address)
	}

	return nil
}
