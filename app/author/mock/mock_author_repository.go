// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/author/author_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	author "github.com/ansidev/fiber-starter-project/domain/author"
	gomock "github.com/golang/mock/gomock"
)

// MockIAuthorRepository is a mock of IAuthorRepository interface.
type MockIAuthorRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthorRepositoryMockRecorder
}

// MockIAuthorRepositoryMockRecorder is the mock recorder for MockIAuthorRepository.
type MockIAuthorRepositoryMockRecorder struct {
	mock *MockIAuthorRepository
}

// NewMockIAuthorRepository creates a new mock instance.
func NewMockIAuthorRepository(ctrl *gomock.Controller) *MockIAuthorRepository {
	mock := &MockIAuthorRepository{ctrl: ctrl}
	mock.recorder = &MockIAuthorRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuthorRepository) EXPECT() *MockIAuthorRepositoryMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockIAuthorRepository) GetByID(id int64) (author.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(author.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIAuthorRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIAuthorRepository)(nil).GetByID), id)
}
