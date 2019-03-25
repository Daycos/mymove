// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import afero "github.com/spf13/afero"
import mock "github.com/stretchr/testify/mock"

// FileStorer is an autogenerated mock type for the FileStorer type
type FileStorer struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *FileStorer) Create(_a0 string) (afero.File, error) {
	ret := _m.Called(_a0)

	var r0 afero.File
	if rf, ok := ret.Get(0).(func(string) afero.File); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(afero.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
