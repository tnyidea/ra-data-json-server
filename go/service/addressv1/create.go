package addressv1

import (
	"github.com/gin-gonic/gin"
	"github.com/tnyidea/ra-data-json-server/go/data/model"
	"net/http"
)

func (p *Handler) Create(c *gin.Context) {
	// POST http://my.api.url/post
	var address model.Address
	c.IndentedJSON(http.StatusOK, address)
}
