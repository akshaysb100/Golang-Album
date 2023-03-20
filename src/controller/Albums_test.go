package controller

import (
	"gin/src/mocks"
	"gin/src/model"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOfferInputForBeverageItems(t *testing.T) {
	t.Run("Test_albums_GetAlbums", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockAlbums := mocks.NewMockAlbums(mockCtrl)
		recorder := &httptest.ResponseRecorder{}
		recorder = httptest.NewRecorder()
		context := &gin.Context{}
		context, _ = gin.CreateTestContext(recorder)
		albums := NewAlbum(mockAlbums)
		var album = []model.Album{
			{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
			{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		}
		mockAlbums.EXPECT().GetAlbum().Return(album)
		albums.GetAlbums(context)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
