package mocks

import (
	"io"

	"github.com/stretchr/testify/mock"
	"github.com/watchly/ngbuild/core"
)

import "time"

// automatically generated by mockery

// Build is an autogenerated mock type for the Build type
type Build struct {
	mock.Mock
}

// Artifact provides a mock function with given fields: name
func (_m *Build) Artifact(name string) []string {
	ret := _m.Called(name)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// BuildTime provides a mock function with given fields:
func (_m *Build) BuildTime() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// Config provides a mock function with given fields:
func (_m *Build) Config() core.BuildConfig {
	ret := _m.Called()

	var r0 core.BuildConfig
	if rf, ok := ret.Get(0).(func() core.BuildConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(core.BuildConfig)
	}

	return r0
}

// ExitCode provides a mock function with given fields:
func (_m *Build) ExitCode() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Group provides a mock function with given fields:
func (_m *Build) Group() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// History provides a mock function with given fields:
func (_m *Build) History() []core.Build {
	ret := _m.Called()

	var r0 []core.Build
	if rf, ok := ret.Get(0).(func() []core.Build); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]core.Build)
		}
	}

	return r0
}

// NewBuild provides a mock function with given fields:
func (_m *Build) NewBuild() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ref provides a mock function with given fields:
func (_m *Build) Ref() {
	_m.Called()
}

// Start provides a mock function with given fields:
func (_m *Build) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stderr provides a mock function with given fields:
func (_m *Build) Stderr() (io.Reader, error) {
	ret := _m.Called()

	var r0 io.Reader
	if rf, ok := ret.Get(0).(func() io.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stdout provides a mock function with given fields:
func (_m *Build) Stdout() (io.Reader, error) {
	ret := _m.Called()

	var r0 io.Reader
	if rf, ok := ret.Get(0).(func() io.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stop provides a mock function with given fields:
func (_m *Build) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Token provides a mock function with given fields:
func (_m *Build) Token() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Unref provides a mock function with given fields:
func (_m *Build) Unref() {
	_m.Called()
}
