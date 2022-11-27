package addressv1

import (
	"github.com/gin-gonic/gin"
	"github.com/tnyidea/ra-data-json-server/go/data/gorm/model"
	"log"
	"net/http"
)

func (p *Handler) GetMany(c *gin.Context) {
	// Determine if the request is one of:
	// getList GET http://my.api.url/post?_sort=title&_order=ASC&_start=0&_end=24&title=bar
	// getMany GET http://my.api.url/post?id=123&id=456&id=789
	// getManyReference GET http://my.api.url/post?author_id=345

	var addresses []model.Address
	tx := p.db.Model(&model.Address{}).Find(&addresses)
	if tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, addresses)
}
