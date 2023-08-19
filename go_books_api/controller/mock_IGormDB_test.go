// Code generated by mockery v2.20.0. DO NOT EDIT.

package controller_test

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// MockIGormDB is an autogenerated mock type for the IGormDB type
type MockIGormDB struct {
	mock.Mock
}

// AutoMigrate provides a mock function with given fields: dst
func (_m *MockIGormDB) AutoMigrate(dst ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, dst...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(dst...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: dst
func (_m *MockIGormDB) Create(dst interface{}) *gorm.DB {
	ret := _m.Called(dst)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(dst)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: dst, conds
func (_m *MockIGormDB) Delete(dst interface{}, conds ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, dst)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(dst, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Find provides a mock function with given fields: dst, conds
func (_m *MockIGormDB) Find(dst interface{}, conds ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, dst)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(dst, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// First provides a mock function with given fields: dst, conds
func (_m *MockIGormDB) First(dst interface{}, conds ...interface{}) *gorm.DB {
	var _ca []interface{}
	_ca = append(_ca, dst)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) *gorm.DB); ok {
		r0 = rf(dst, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Save provides a mock function with given fields: dst
func (_m *MockIGormDB) Save(dst interface{}) *gorm.DB {
	ret := _m.Called(dst)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(interface{}) *gorm.DB); ok {
		r0 = rf(dst)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

type mockConstructorTestingTNewMockIGormDB interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIGormDB creates a new instance of MockIGormDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIGormDB(t mockConstructorTestingTNewMockIGormDB) *MockIGormDB {
	mock := &MockIGormDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}