package addressv1

import (
	"context"
	"github.com/gin-gonic/gin"
	gormmodel "github.com/tnyidea/ra-data-json-server/go/data/mongodb/model"
	mongomodel "github.com/tnyidea/ra-data-json-server/go/data/mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func (p *Handler) Create(c *gin.Context) {
	if p.serviceDb == ServiceDbMongoDb {
		p.mongodbCreate(c)
		return
	}

	if p.serviceDb == ServiceDbPostgres {
		p.gormCreate(c)
		return
	}
}

func (p *Handler) gormCreate(c *gin.Context) {
	// POST http://my.api.url/post

	var address gormmodel.Address
	err := c.BindJSON(&address)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, address)
		return
	}

	tx := p.gormDb.Model(&gormmodel.Address{}).Create(&address)
	if tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, address)
}

func (p *Handler) mongodbCreate(c *gin.Context) {
	// POST http://my.api.url/post

	var requestBody mongomodel.Address
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, requestBody)
		return
	}
	requestBody.Id = primitive.NewObjectID()
	b, err := bson.Marshal(&requestBody)
	if err != nil {
		return
	}

	// Execute the Query
	collection := p.mongoDb.Collection("address")
	_, err = collection.InsertOne(context.TODO(), b)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, requestBody)
}
