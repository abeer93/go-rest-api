// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	book "github.com/abeer93/go-rest-api/book/models"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddNewBook provides a mock function with given fields: _a0
func (_m *Service) AddNewBook(_a0 *book.Book) (book.Book, error) {
	ret := _m.Called(_a0)

	var r0 book.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(*book.Book) (book.Book, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*book.Book) book.Book); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(book.Book)
	}

	if rf, ok := ret.Get(1).(func(*book.Book) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllBooks provides a mock function with given fields:
func (_m *Service) GetAllBooks() ([]book.Book, error) {
	ret := _m.Called()

	var r0 []book.Book
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]book.Book, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []book.Book); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]book.Book)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveBook provides a mock function with given fields: _a0
func (_m *Service) RemoveBook(_a0 book.Book) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(book.Book) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
