package addressv1

import (
	"github.com/gin-gonic/gin"
	"github.com/tnyidea/ra-data-json-server/go/data/gorm/model"
	"net/http"
)

func (p *Handler) GetOne(c *gin.Context) {
	// GET http://my.api.url/post/123
	var address model.Address
	c.IndentedJSON(http.StatusOK, address)
}
