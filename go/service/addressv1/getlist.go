package addressv1

func GetList(c *gin.Context) {
    // GET http://my.api.url/posts?_sort=title&_order=ASC&_start=0&_end=24&title=bar
    var addresses []datatypes.Address
    c.IndentedJSON(http.StatusOK, addresses)
}