package addressv1

import "github.com/gin-gonic/gin"
import datatypes "github.com/tnyidea/ra-data-json-server/go/data/types"
import "net/http"

func GetList(c *gin.Context) {
	// GET http://my.api.url/posts?_sort=title&_order=ASC&_start=0&_end=24&title=bar
	var addresses []datatypes.Address
	c.IndentedJSON(http.StatusOK, addresses)
}
