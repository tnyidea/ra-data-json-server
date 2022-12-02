package addressv1

import (
	"github.com/gin-gonic/gin"
	"github.com/tnyidea/ra-data-json-server/go/data/gorm/model"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func (p *Handler) GetOne(c *gin.Context) {
	// GET http://my.api.url/post/123

	var address model.Address
	tx := p.gormDb.Model(model.Address{}).Where("id = ?", c.Param("id")).First(&address)
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
