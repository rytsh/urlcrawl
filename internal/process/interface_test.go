// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=interface_test.go -package=process Storage Resource
//

// Package process is a generated GoMock package.
package process

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
	isgomock struct{}
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockStorage) Get(path string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", path)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStorageMockRecorder) Get(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStorage)(nil).Get), path)
}

// Has mocks base method.
func (m *MockStorage) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockStorageMockRecorder) Has(path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockStorage)(nil).Has), path)
}

// Set mocks base method.
func (m *MockStorage) Set(path string, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", path, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockStorageMockRecorder) Set(path, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockStorage)(nil).Set), path, data)
}

// MockResource is a mock of Resource interface.
type MockResource struct {
	ctrl     *gomock.Controller
	recorder *MockResourceMockRecorder
	isgomock struct{}
}

// MockResourceMockRecorder is the mock recorder for MockResource.
type MockResourceMockRecorder struct {
	mock *MockResource
}

// NewMockResource creates a new mock instance.
func NewMockResource(ctrl *gomock.Controller) *MockResource {
	mock := &MockResource{ctrl: ctrl}
	mock.recorder = &MockResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResource) EXPECT() *MockResourceMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockResource) Fetch(ctx context.Context, url string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, url)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockResourceMockRecorder) Fetch(ctx, url any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockResource)(nil).Fetch), ctx, url)
}

// Links mocks base method.
func (m *MockResource) Links(data []byte) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Links", data)
	ret0, _ := ret[0].([]string)
	return ret0
}

// Links indicates an expected call of Links.
func (mr *MockResourceMockRecorder) Links(data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Links", reflect.TypeOf((*MockResource)(nil).Links), data)
}
