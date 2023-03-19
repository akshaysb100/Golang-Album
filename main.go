package main

import (
	"fmt"
	"gin/src/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	albumController := controller.NewAlbum()
	router.GET("/albums", albumController.GetAlbums)
	router.GET("/album/:id", albumController.GetAlbumByID)
	router.POST("/albums", albumController.PostAlbums)

	//router.Run("localhost:8000")
	err := router.Run(":8000")
	if err != nil {
		message := fmt.Sprintf("could not start server due to %s", err.Error())
		panic(message)
	}
}
