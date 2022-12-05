package addressv1

import (
	"context"
	"github.com/gin-gonic/gin"
	gormmodel "github.com/tnyidea/ra-data-json-server/go/data/gorm/model"
	mongomodel "github.com/tnyidea/ra-data-json-server/go/data/mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func (p *Handler) GetOne(c *gin.Context) {
	if p.serviceDb == ServiceDbMongoDb {
		p.mongodbGetOne(c)
		return
	}

	if p.serviceDb == ServiceDbPostgres {
		p.gormGetOne(c)
		return
	}
}

func (p *Handler) gormGetOne(c *gin.Context) {
	// GET http://my.api.url/post/123

	var address gormmodel.Address
	tx := p.gormDb.Model(gormmodel.Address{}).Where("id = ?", c.Param("id")).First(&address)
	if tx.Error == gorm.ErrRecordNotFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, address)
}

func (p *Handler) mongodbGetOne(c *gin.Context) {
	// Execute the Query
	collection := p.mongoDb.Collection("address")
	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var address mongomodel.Address
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&address)
	if err == mongo.ErrNoDocuments {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, address)
}
