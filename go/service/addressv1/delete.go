package addressv1

import (
	"context"
	"github.com/gin-gonic/gin"
	gormmodel "github.com/tnyidea/ra-data-json-server/go/data/gorm/model"
	mongomodel "github.com/tnyidea/ra-data-json-server/go/data/mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
)

func (p *Handler) Delete(c *gin.Context) {
	if p.serviceDb == ServiceDbMongoDb {
		p.mongodbDelete(c)
		return
	}

	if p.serviceDb == ServiceDbPostgres {
		p.gormDelete(c)
		return
	}
}

func (p *Handler) gormDelete(c *gin.Context) {
	// DELETE http://my.api.url/post/123

	var address gormmodel.Address
	tx := p.gormDb.Model(&gormmodel.Address{}).Clauses(clause.Returning{}).Where("id = ?", c.Param("id")).Delete(&address)
	log.Println(tx.RowsAffected)
	if tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if tx.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, address)
}

func (p *Handler) mongodbDelete(c *gin.Context) {
	log.Println(c.Param("id"))
	// Execute the Query
	collection := p.mongoDb.Collection("address")
	objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var address mongomodel.Address
	err = collection.FindOneAndDelete(context.TODO(), bson.M{"_id": objectId}).Decode(&address)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, address)
}
