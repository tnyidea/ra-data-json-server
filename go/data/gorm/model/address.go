package model

import (
	"gorm.io/gorm"
	"time"
)

type Address struct {
	//gorm.Model -- overriding this with custom model for primary key
	Id          string         `gorm:"primary_key;type:text;default:gen_random_uuid()" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	FirstName   string         `json:"firstName"`
	LastName    string         `json:"lastName"`
	CompanyName string         `json:"companyName"`
	Address     string         `json:"address"`
	City        string         `json:"city"`
	County      string         `json:"county"`
	State       string         `json:"state"`
	Zip         string         `json:"zip"`
	Phone1      string         `json:"phone1"`
	Phone2      string         `json:"phone2"`
	Email       string         `json:"email"`
	Web         string         `json:"web"`
}
