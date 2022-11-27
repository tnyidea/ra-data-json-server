package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tnyidea/ra-data-json-server/go/service/addressv1"
	"github.com/tnyidea/typeutils"
	"log"
	"os"
)

func main() {
	env := map[string]string{
		"POSTGRES_URL": os.Getenv("POSTGRES_URL"),
	}
	err := typeutils.MapNoEmptyValues(env)
	if err != nil {
		log.Fatal(err)
	}

	addressv1handler, err := addressv1.NewHandler(env)
	if err != nil {
		return
	}

	router := gin.Default()
	router.GET("/v1/address", addressv1handler.GetMany)
	router.GET("/v1/address/:id", addressv1handler.GetOne)
	router.POST("/v1/address", addressv1handler.Create)
	router.PUT("/v1/address/:id", addressv1handler.Update)
	router.DELETE("/v1/address/:id", addressv1handler.Delete)

	log.Fatal(router.Run("localhost:8080"))
}
