package addressv1

import (
	"github.com/gin-gonic/gin"
	"github.com/tnyidea/ra-data-json-server/go/data/gorm/model"
	"log"
	"net/http"
)

func (p *Handler) Create(c *gin.Context) {
	// POST http://my.api.url/post

	var address model.Address
	err := c.BindJSON(&address)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, address)
		return
	}

	tx := p.db.Model(&model.Address{}).Create(&address)
	if tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, address)
}
