// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	"github.com/cidverse/normalizeci/pkg/vcsrepository/vcsapi"
	"github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Check provides a mock function with given fields:
func (_m *Client) Check() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FindCommitByHash provides a mock function with given fields: hash, includeChanges
func (_m *Client) FindCommitByHash(hash string, includeChanges bool) (vcsapi.Commit, error) {
	ret := _m.Called(hash, includeChanges)

	var r0 vcsapi.Commit
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool) (vcsapi.Commit, error)); ok {
		return rf(hash, includeChanges)
	}
	if rf, ok := ret.Get(0).(func(string, bool) vcsapi.Commit); ok {
		r0 = rf(hash, includeChanges)
	} else {
		r0 = ret.Get(0).(vcsapi.Commit)
	}

	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(hash, includeChanges)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCommitsBetween provides a mock function with given fields: from, to, includeChanges, limit
func (_m *Client) FindCommitsBetween(from *vcsapi.VCSRef, to *vcsapi.VCSRef, includeChanges bool, limit int) ([]vcsapi.Commit, error) {
	ret := _m.Called(from, to, includeChanges, limit)

	var r0 []vcsapi.Commit
	var r1 error
	if rf, ok := ret.Get(0).(func(*vcsapi.VCSRef, *vcsapi.VCSRef, bool, int) ([]vcsapi.Commit, error)); ok {
		return rf(from, to, includeChanges, limit)
	}
	if rf, ok := ret.Get(0).(func(*vcsapi.VCSRef, *vcsapi.VCSRef, bool, int) []vcsapi.Commit); ok {
		r0 = rf(from, to, includeChanges, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]vcsapi.Commit)
		}
	}

	if rf, ok := ret.Get(1).(func(*vcsapi.VCSRef, *vcsapi.VCSRef, bool, int) error); ok {
		r1 = rf(from, to, includeChanges, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLatestRelease provides a mock function with given fields: stable
func (_m *Client) FindLatestRelease(stable bool) (vcsapi.VCSRelease, error) {
	ret := _m.Called(stable)

	var r0 vcsapi.VCSRelease
	var r1 error
	if rf, ok := ret.Get(0).(func(bool) (vcsapi.VCSRelease, error)); ok {
		return rf(stable)
	}
	if rf, ok := ret.Get(0).(func(bool) vcsapi.VCSRelease); ok {
		r0 = rf(stable)
	} else {
		r0 = ret.Get(0).(vcsapi.VCSRelease)
	}

	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(stable)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTags provides a mock function with given fields:
func (_m *Client) GetTags() []vcsapi.VCSRef {
	ret := _m.Called()

	var r0 []vcsapi.VCSRef
	if rf, ok := ret.Get(0).(func() []vcsapi.VCSRef); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]vcsapi.VCSRef)
		}
	}

	return r0
}

// GetTagsByHash provides a mock function with given fields: hash
func (_m *Client) GetTagsByHash(hash string) []vcsapi.VCSRef {
	ret := _m.Called(hash)

	var r0 []vcsapi.VCSRef
	if rf, ok := ret.Get(0).(func(string) []vcsapi.VCSRef); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]vcsapi.VCSRef)
		}
	}

	return r0
}

// VCSHead provides a mock function with given fields:
func (_m *Client) VCSHead() (vcsapi.VCSRef, error) {
	ret := _m.Called()

	var r0 vcsapi.VCSRef
	var r1 error
	if rf, ok := ret.Get(0).(func() (vcsapi.VCSRef, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() vcsapi.VCSRef); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(vcsapi.VCSRef)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VCSHostServer provides a mock function with given fields: remote
func (_m *Client) VCSHostServer(remote string) string {
	ret := _m.Called(remote)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(remote)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// VCSHostType provides a mock function with given fields: server
func (_m *Client) VCSHostType(server string) string {
	ret := _m.Called(server)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(server)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// VCSRefToInternalRef provides a mock function with given fields: ref
func (_m *Client) VCSRefToInternalRef(ref vcsapi.VCSRef) string {
	ret := _m.Called(ref)

	var r0 string
	if rf, ok := ret.Get(0).(func(vcsapi.VCSRef) string); ok {
		r0 = rf(ref)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// VCSRemote provides a mock function with given fields:
func (_m *Client) VCSRemote() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// VCSType provides a mock function with given fields:
func (_m *Client) VCSType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}