package addressv1

import (
	"github.com/gin-gonic/gin"
	"github.com/tnyidea/ra-data-json-server/go/data/gorm/model"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func (p *Handler) Update(c *gin.Context) {
	// Will handle either
	// update PUT http://my.api.url/post/123
	// updateMany PUT http://my.api.url/post/123, PUT http://my.api.url/post/456, PUT http://my.api.url/post/789
	var address model.Address
	err := c.BindJSON(&address)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, address)
		return
	}

	tx := p.db.Model(&model.Address{}).
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
