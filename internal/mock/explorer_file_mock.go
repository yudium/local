// Code generated by MockGen. DO NOT EDIT.
// Source: internal/explorer/file.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	explorer "github.com/go-seidon/local/internal/explorer"
	gomock "github.com/golang/mock/gomock"
)

// MockFileManager is a mock of FileManager interface.
type MockFileManager struct {
	ctrl     *gomock.Controller
	recorder *MockFileManagerMockRecorder
}

// MockFileManagerMockRecorder is the mock recorder for MockFileManager.
type MockFileManagerMockRecorder struct {
	mock *MockFileManager
}

// NewMockFileManager creates a new mock instance.
func NewMockFileManager(ctrl *gomock.Controller) *MockFileManager {
	mock := &MockFileManager{ctrl: ctrl}
	mock.recorder = &MockFileManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileManager) EXPECT() *MockFileManagerMockRecorder {
	return m.recorder
}

// CreateDir mocks base method.
func (m *MockFileManager) CreateDir(ctx context.Context, p explorer.CreateDirParam) (*explorer.CreateDirResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDir", ctx, p)
	ret0, _ := ret[0].(*explorer.CreateDirResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDir indicates an expected call of CreateDir.
func (mr *MockFileManagerMockRecorder) CreateDir(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDir", reflect.TypeOf((*MockFileManager)(nil).CreateDir), ctx, p)
}

// IsFileExists mocks base method.
func (m *MockFileManager) IsFileExists(ctx context.Context, p explorer.IsFileExistsParam) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFileExists", ctx, p)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFileExists indicates an expected call of IsFileExists.
func (mr *MockFileManagerMockRecorder) IsFileExists(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFileExists", reflect.TypeOf((*MockFileManager)(nil).IsFileExists), ctx, p)
}

// OpenFile mocks base method.
func (m *MockFileManager) OpenFile(ctx context.Context, p explorer.OpenFileParam) (*explorer.OpenFileResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenFile", ctx, p)
	ret0, _ := ret[0].(*explorer.OpenFileResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenFile indicates an expected call of OpenFile.
func (mr *MockFileManagerMockRecorder) OpenFile(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenFile", reflect.TypeOf((*MockFileManager)(nil).OpenFile), ctx, p)
}

// ReadFile mocks base method.
func (m *MockFileManager) ReadFile(ctx context.Context, p explorer.ReadFileParam) (*explorer.ReadFileResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", ctx, p)
	ret0, _ := ret[0].(*explorer.ReadFileResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MockFileManagerMockRecorder) ReadFile(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockFileManager)(nil).ReadFile), ctx, p)
}

// RemoveFile mocks base method.
func (m *MockFileManager) RemoveFile(ctx context.Context, p explorer.RemoveFileParam) (*explorer.RemoveFileResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFile", ctx, p)
	ret0, _ := ret[0].(*explorer.RemoveFileResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveFile indicates an expected call of RemoveFile.
func (mr *MockFileManagerMockRecorder) RemoveFile(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFile", reflect.TypeOf((*MockFileManager)(nil).RemoveFile), ctx, p)
}

// SaveFile mocks base method.
func (m *MockFileManager) SaveFile(ctx context.Context, p explorer.SaveFileParam) (*explorer.SaveFileResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFile", ctx, p)
	ret0, _ := ret[0].(*explorer.SaveFileResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveFile indicates an expected call of SaveFile.
func (mr *MockFileManagerMockRecorder) SaveFile(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFile", reflect.TypeOf((*MockFileManager)(nil).SaveFile), ctx, p)
}
