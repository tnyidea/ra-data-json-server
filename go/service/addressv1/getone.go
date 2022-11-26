package addressv1

import (
    "github.com/gin-gonic/gin"
    datatypes "github.com/tnyidea/ra-data-json-server/go/data/types"
    "net/http"
)

func GetOne(c *gin.Context) {
	// GET http://my.api.url/posts/123`
	var address datatypes.Address
	c.IndentedJSON(http.StatusOK, address)
}
