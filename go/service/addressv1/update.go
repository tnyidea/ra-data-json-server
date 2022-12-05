package addressv1

import (
	"context"
	"github.com/gin-gonic/gin"
	gormmodel "github.com/tnyidea/ra-data-json-server/go/data/gorm/model"
	mongodbmodel "github.com/tnyidea/ra-data-json-server/go/data/mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func (p *Handler) Update(c *gin.Context) {
	if p.serviceDb == ServiceDbMongoDb {
		p.mongodbUpdate(c)
		return
	}

	if p.serviceDb == ServiceDbPostgres {
		p.gormUpdate(c)
		return
	}
}

func (p *Handler) gormUpdate(c *gin.Context) {
	// Will handle either
	// update PUT http://my.api.url/post/123
	// updateMany PUT http://my.api.url/post/123, PUT http://my.api.url/post/456, PUT http://my.api.url/post/789
	var address gormmodel.Address
	err := c.BindJSON(&address)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, address)
		return
	}

	tx := p.gormDb.Model(&gormmodel.Address{}).
		Where("id = ?", c.Param("id")).Updates(&address).First(&address)
	if tx.Error != nil {
		log.Println(tx.Error)

		if tx.Error == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, address)
}

func (p *Handler) mongodbUpdate(c *gin.Context) {
	// PUT http://my.api.url/post

	var requestBody map[string]any
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, requestBody)
		return
	}

	// Execute the Query
	collection := p.mongoDb.Collection("address")
	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	_, err = collection.UpdateOne(context.TODO(),
		bson.M{"_id": objectId},
		bson.D{{
			"$set",
			requestBody,
		}},
	)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var address mongodbmodel.Address
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&address)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, address)
}
