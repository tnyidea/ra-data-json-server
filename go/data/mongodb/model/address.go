package model

import (
	"gorm.io/gorm"
	"time"
)

type Address struct {
	Id          string         `bson:"_id" json:"id"`
	CreatedAt   time.Time      `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time      `bson:"updatedAt" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `bson:"deletedAt" json:"deletedAt"`
	FirstName   string         `bson:"firstName" json:"firstName"`
	LastName    string         `bson:"lastName" json:"lastName"`
	CompanyName string         `bson:"companyName" json:"companyName"`
	Address     string         `bson:"address" json:"address"`
	City        string         `bson:"city" json:"city"`
	County      string         `bson:"county" json:"county"`
	State       string         `bson:"state" json:"state"`
	Zip         string         `bson:"zip" json:"zip"`
	Phone1      string         `bson:"phone1" json:"phone1"`
	Phone2      string         `bson:"phone2" json:"phone2"`
	Email       string         `bson:"email" json:"email"`
	Web         string         `bson:"web" json:"web"`
}
