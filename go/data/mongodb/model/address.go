package model

import (
	"gorm.io/gorm"
	"time"
)

type Address struct {
	Id          string         `bson:"id"`
	CreatedAt   time.Time      `bson:"createdAt"`
	UpdatedAt   time.Time      `bson:"updatedAt"`
	DeletedAt   gorm.DeletedAt `bson:"deletedAt"`
	FirstName   string         `bson:"firstName"`
	LastName    string         `bson:"lastName"`
	CompanyName string         `bson:"companyName"`
	Address     string         `bson:"address"`
	City        string         `bson:"city"`
	County      string         `bson:"county"`
	State       string         `bson:"state"`
	Zip         string         `bson:"zip"`
	Phone1      string         `bson:"phone1"`
	Phone2      string         `bson:"phone2"`
	Email       string         `bson:"email"`
	Web         string         `bson:"web"`
}
