package addressv1

func GetOne(c *gin.Context) {
    // GET http://my.api.url/posts/123
    var address datatypes.Address
    c.IndentedJSON(http.StatusOK, address)
}