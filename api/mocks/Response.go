// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	handlers "github.com/richmondwang/golang-wallet-api/pkg/handlers"

	mock "github.com/stretchr/testify/mock"
)

// Response is an autogenerated mock type for the Response type
type Response struct {
	mock.Mock
}

// Render provides a mock function with given fields: w, r
func (_m *Response) Render(w http.ResponseWriter, r *http.Request) error {
	ret := _m.Called(w, r)

	if len(ret) == 0 {
		panic("no return value specified for Render")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(http.ResponseWriter, *http.Request) error); ok {
		r0 = rf(w, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithCode provides a mock function with given fields: code
func (_m *Response) WithCode(code int) handlers.Response {
	ret := _m.Called(code)

	if len(ret) == 0 {
		panic("no return value specified for WithCode")
	}

	var r0 handlers.Response
	if rf, ok := ret.Get(0).(func(int) handlers.Response); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(handlers.Response)
		}
	}

	return r0
}

// NewResponse creates a new instance of Response. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewResponse(t interface {
	mock.TestingT
	Cleanup(func())
}) *Response {
	mock := &Response{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}