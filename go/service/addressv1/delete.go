package addressv1

import (
	"github.com/gin-gonic/gin"
	"github.com/tnyidea/ra-data-json-server/go/data/model"
	"net/http"
)

func (p *Handler) Delete(c *gin.Context) {
	// DELETE http://my.api.url/post/123
	var address model.Address
	c.IndentedJSON(http.StatusOK, address)
}
