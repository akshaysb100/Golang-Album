// Code generated by MockGen. DO NOT EDIT.
// Source: src/service/Album.go

// Package mocks is a generated GoMock package.
package mocks

import (
	model "gin/src/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAlbums is a mock of Albums interface.
type MockAlbums struct {
	ctrl     *gomock.Controller
	recorder *MockAlbumsMockRecorder
}

// MockAlbumsMockRecorder is the mock recorder for MockAlbums.
type MockAlbumsMockRecorder struct {
	mock *MockAlbums
}

// NewMockAlbums creates a new mock instance.
func NewMockAlbums(ctrl *gomock.Controller) *MockAlbums {
	mock := &MockAlbums{ctrl: ctrl}
	mock.recorder = &MockAlbumsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlbums) EXPECT() *MockAlbumsMockRecorder {
	return m.recorder
}

// GetAlbum mocks base method.
func (m *MockAlbums) GetAlbum() []model.Album {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlbum")
	ret0, _ := ret[0].([]model.Album)
	return ret0
}

// GetAlbum indicates an expected call of GetAlbum.
func (mr *MockAlbumsMockRecorder) GetAlbum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlbum", reflect.TypeOf((*MockAlbums)(nil).GetAlbum))
}