// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/post/post_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	post "github.com/ansidev/fiber-starter-project/domain/post"
	gomock "github.com/golang/mock/gomock"
)

// MockIPostRepository is a mock of IPostRepository interface.
type MockIPostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPostRepositoryMockRecorder
}

// MockIPostRepositoryMockRecorder is the mock recorder for MockIPostRepository.
type MockIPostRepositoryMockRecorder struct {
	mock *MockIPostRepository
}

// NewMockIPostRepository creates a new mock instance.
func NewMockIPostRepository(ctrl *gomock.Controller) *MockIPostRepository {
	mock := &MockIPostRepository{ctrl: ctrl}
	mock.recorder = &MockIPostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPostRepository) EXPECT() *MockIPostRepositoryMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockIPostRepository) GetByID(id int64) (post.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(post.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIPostRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIPostRepository)(nil).GetByID), id)
}
