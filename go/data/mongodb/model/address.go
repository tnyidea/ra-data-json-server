package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Address struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
	DeletedAt   time.Time          `bson:"deletedAt,omitempty" json:"deletedAt"`
	FirstName   string             `bson:"firstName,omitempty" json:"firstName"`
	LastName    string             `bson:"lastName,omitempty" json:"lastName"`
	CompanyName string             `bson:"companyName,omitempty" json:"companyName"`
	Address     string             `bson:"address,omitempty" json:"address"`
	City        string             `bson:"city,omitempty" json:"city"`
	County      string             `bson:"county,omitempty" json:"county"`
	State       string             `bson:"state,omitempty" json:"state"`
	Zip         string             `bson:"zip,omitempty" json:"zip"`
	Phone1      string             `bson:"phone1,omitempty" json:"phone1"`
	Phone2      string             `bson:"phone2,omitempty" json:"phone2"`
	Email       string             `bson:"email,omitempty" json:"email"`
	Web         string             `bson:"web,omitempty" json:"web"`
}
