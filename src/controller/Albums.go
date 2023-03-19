package controller

import (
	"gin/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Albums interface {
	GetAlbums(c *gin.Context)
	GetAlbumByID(c *gin.Context)
	PostAlbums(c *gin.Context)
}

type albums struct {
}

func NewAlbum() Albums {
	return &albums{}
}

func (a albums) GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, album)
}

func (a albums) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range album {
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

	// Add the new album to the slice.
	album = append(album, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

var album = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
