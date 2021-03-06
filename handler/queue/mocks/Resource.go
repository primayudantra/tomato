// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Resource is an autogenerated mock type for the Resource type
type Resource struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: target
func (_m *Resource) Fetch(target string) ([][]byte, error) {
	ret := _m.Called(target)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func(string) [][]byte); ok {
		r0 = rf(target)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(target)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Listen provides a mock function with given fields: target
func (_m *Resource) Listen(target string) error {
	ret := _m.Called(target)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(target)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Publish provides a mock function with given fields: target, payload
func (_m *Resource) Publish(target string, payload []byte) error {
	ret := _m.Called(target, payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte) error); ok {
		r0 = rf(target, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *Resource) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reset provides a mock function with given fields:
func (_m *Resource) Reset() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
