package controller

import (
	"gin/src/model"
	"gin/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Albums interface {
	GetAlbums(c *gin.Context)
	GetAlbumByID(c *gin.Context)
	PostAlbums(c *gin.Context)
}

type albums struct {
	service service.Albums
}

func NewAlbum(service service.Albums) Albums {
	return &albums{
		service: service,
	}
}

func (a albums) GetAlbums(c *gin.Context) {
	getAlbum := a.service.GetAlbum()
	c.IndentedJSON(http.StatusOK, getAlbum)
}

func (a albums) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	getAlbum := a.service.GetAlbum()
	for _, a := range getAlbum {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func (a albums) PostAlbums(c *gin.Context) {
	var newAlbum model.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	getAlbum := a.service.GetAlbum()
	// Add the new getAlbum to the slice.
	getAlbum = append(getAlbum, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
