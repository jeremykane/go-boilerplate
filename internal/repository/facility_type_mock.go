// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/facility_type.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/jeremykane/go-boilerplate/internal/entity"
	errorx "github.com/jeremykane/go-boilerplate/pkg/errorx"
)

// MockFacilityTypeRepository is a mock of FacilityTypeRepository interface.
type MockFacilityTypeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFacilityTypeRepositoryMockRecorder
}

// MockFacilityTypeRepositoryMockRecorder is the mock recorder for MockFacilityTypeRepository.
type MockFacilityTypeRepositoryMockRecorder struct {
	mock *MockFacilityTypeRepository
}

// NewMockFacilityTypeRepository creates a new mock instance.
func NewMockFacilityTypeRepository(ctrl *gomock.Controller) *MockFacilityTypeRepository {
	mock := &MockFacilityTypeRepository{ctrl: ctrl}
	mock.recorder = &MockFacilityTypeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFacilityTypeRepository) EXPECT() *MockFacilityTypeRepositoryMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockFacilityTypeRepository) GetAll(ctx context.Context) ([]entity.FacilityType, *errorx.CustomError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]entity.FacilityType)
	ret1, _ := ret[1].(*errorx.CustomError)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockFacilityTypeRepositoryMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockFacilityTypeRepository)(nil).GetAll), ctx)
}
