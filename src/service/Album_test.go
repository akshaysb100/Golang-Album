package service

import (
	"gin/src/model"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type AlbumsServiceTestSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
	context  *gin.Context
	recorder *httptest.ResponseRecorder
	albums   Albums
}

func TestAlbumServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AlbumsServiceTestSuite))
}

func (suite *AlbumsServiceTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	suite.albums = NewAlbums()

}

func (suite *AlbumsServiceTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *AlbumsServiceTestSuite) TestHealthCheckReturnsResponseOnSuccess() {
	//suite.context.Request, _ = http.NewRequest(http.MethodGet, "/health", nil)

	var album = []model.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	actuleAlbum := suite.albums.GetAlbum()
	suite.Equal(album, actuleAlbum)
	//suite.Equal(http.StatusOK, suite.recorder.Code)
}
