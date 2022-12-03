package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tnyidea/ra-data-json-server/go/service/addressv1"
	"github.com/tnyidea/typeutils"
	"log"
	"os"
)

func main() {
	env := map[string]string{
		"SERVICE_DB":   os.Getenv("SERVICE_DB"),
		"MONGODB_URL":  os.Getenv("MONGODB_URL"),
		"POSTGRES_URL": os.Getenv("POSTGRES_URL"),
	}
	err := typeutils.MapNoEmptyValues(env)
	if err != nil {
		log.Fatal(err)
	}

	addressv1handler, err := addressv1.NewHandler(env)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("here")
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddExposeHeaders("Content-Range", "X-Total-Count")
	router.Use(cors.New(corsConfig))

	router.GET("/v1/address", addressv1handler.GetMany)
	router.GET("/v1/address/:id", addressv1handler.GetOne)
	router.POST("/v1/address", addressv1handler.Create)
	router.PUT("/v1/address/:id", addressv1handler.Update)
	router.DELETE("/v1/address/:id", addressv1handler.Delete)

	log.Fatal(router.Run("localhost:8080"))
}
