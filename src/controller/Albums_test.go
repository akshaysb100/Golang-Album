package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	httptest "net/http/httptest"
	"reflect"
	"testing"
)

func TestOfferInputForBeverageItems(t *testing.T) {
	t.Run("Test_albums_GetAlbums", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		recorder := &httptest.ResponseRecorder{}
		recorder = httptest.NewRecorder()
		context := &gin.Context{}
		context, _ = gin.CreateTestContext(recorder)
		albums := NewAlbum()
		albums.GetAlbums(context)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestNewAlbum(t *testing.T) {
	tests := []struct {
		name string
		want Albums
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAlbum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAlbum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_albums_GetAlbumByID(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := albums{}
			a.GetAlbumByID(tt.args.c)
		})
	}
}

func Test_albums_GetAlbums(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := albums{}
			a.GetAlbums(tt.args.c)
		})
	}
}

func Test_albums_PostAlbums(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := albums{}
			a.PostAlbums(tt.args.c)
		})
	}
}
