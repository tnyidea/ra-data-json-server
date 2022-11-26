package addressv1

import "github.com/gin-gonic/gin"
import datatypes "github.com/tnyidea/ra-data-json-server/go/data/types"
import "net/http"

func GetMany(c *gin.Context) {
    // Determine if the request is one of:
	// getList GET http://my.api.url/posts?_sort=title&_order=ASC&_start=0&_end=24&title=bar
    // getMany GET http://my.api.url/post?id=123&id=456&id=789
    // getManyReference GET http://my.api.url/post?author_id=345
    
	var addresses []datatypes.Address
	c.IndentedJSON(http.StatusOK, addresses)
}
