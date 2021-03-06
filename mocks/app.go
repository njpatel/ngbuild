package mocks

import "github.com/watchly/ngbuild/core"
import "github.com/stretchr/testify/mock"

// automatically generated by mockery

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// AppLocation provides a mock function with given fields:
func (_m *App) AppLocation() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Config provides a mock function with given fields: namespace, conf
func (_m *App) Config(namespace string, conf interface{}) error {
	ret := _m.Called(namespace, conf)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(namespace, conf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBuild provides a mock function with given fields: token
func (_m *App) GetBuild(token string) (core.Build, error) {
	ret := _m.Called(token)

	var r0 core.Build
	if rf, ok := ret.Get(0).(func(string) core.Build); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Build)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBuildHistory provides a mock function with given fields: group
func (_m *App) GetBuildHistory(group string) []core.Build {
	ret := _m.Called(group)

	var r0 []core.Build
	if rf, ok := ret.Get(0).(func(string) []core.Build); ok {
		r0 = rf(group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]core.Build)
		}
	}

	return r0
}

// GlobalConfig provides a mock function with given fields: conf
func (_m *App) GlobalConfig(conf interface{}) error {
	ret := _m.Called(conf)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(conf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Listen provides a mock function with given fields: event, listener
func (_m *App) Listen(event string, listener func(map[string]string)) core.EventHandler {
	ret := _m.Called(event, listener)

	var r0 core.EventHandler
	if rf, ok := ret.Get(0).(func(string, func(map[string]string)) core.EventHandler); ok {
		r0 = rf(event, listener)
	} else {
		r0 = ret.Get(0).(core.EventHandler)
	}

	return r0
}

// Logcritf provides a mock function with given fields: _a0, _a1
func (_m *App) Logcritf(_a0 string, _a1 ...interface{}) {
	_m.Called(_a0, _a1)
}

// Loginfof provides a mock function with given fields: _a0, _a1
func (_m *App) Loginfof(_a0 string, _a1 ...interface{}) {
	_m.Called(_a0, _a1)
}

// Logwarnf provides a mock function with given fields: _a0, _a1
func (_m *App) Logwarnf(_a0 string, _a1 ...interface{}) {
	_m.Called(_a0, _a1)
}

// Name provides a mock function with given fields:
func (_m *App) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewBuild provides a mock function with given fields: group, config
func (_m *App) NewBuild(group string, config *core.BuildConfig) (string, error) {
	ret := _m.Called(group, config)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, *core.BuildConfig) string); ok {
		r0 = rf(group, config)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *core.BuildConfig) error); ok {
		r1 = rf(group, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveEventHandler provides a mock function with given fields: _a0
func (_m *App) RemoveEventHandler(_a0 core.EventHandler) {
	_m.Called(_a0)
}

// SendEvent provides a mock function with given fields: event
func (_m *App) SendEvent(event string) {
	_m.Called(event)
}

// Shutdown provides a mock function with given fields:
func (_m *App) Shutdown() {
	_m.Called()
}
