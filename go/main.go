package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tnyidea/react-admin/dataprovider/go/service/endpointsv1"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/v1/address", endpointsv1.GetAllAddresses)

	log.Fatal(router.Run("localhost:8080"))
}
