package main

import (
	"github.com/gin-gonic/gin"
    "github.com/tnyidea/ra-data-json-server/go/service/addressv1"
	"log"
)

func main() {
	router := gin.Default()
    router.GET("/v1/address", addressv1.GetMany)
	router.GET("/v1/address/:id", addressv1.GetOne)

	log.Fatal(router.Run("localhost:8080"))
}
