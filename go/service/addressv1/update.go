package addressv1

import (
	"github.com/gin-gonic/gin"
	"github.com/tnyidea/ra-data-json-server/go/data/gorm/model"
	"net/http"
)

func (p *Handler) Update(c *gin.Context) {
	// Will handle either
	// update PUT http://my.api.url/post/123
	// updateMany PUT http://my.api.url/post/123, PUT http://my.api.url/post/456, PUT http://my.api.url/post/789
	var address model.Address
	c.IndentedJSON(http.StatusOK, address)
}
