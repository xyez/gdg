// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"github.com/esnet/gdg/internal/service/domain"
	"github.com/esnet/gdg/internal/service/filters"
	mock "github.com/stretchr/testify/mock"
)

// newOrganizationCrudApi creates a new instance of organizationCrudApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newOrganizationCrudApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *organizationCrudApi {
	mock := &organizationCrudApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// organizationCrudApi is an autogenerated mock type for the organizationCrudApi type
type organizationCrudApi struct {
	mock.Mock
}

type organizationCrudApi_Expecter struct {
	mock *mock.Mock
}

func (_m *organizationCrudApi) EXPECT() *organizationCrudApi_Expecter {
	return &organizationCrudApi_Expecter{mock: &_m.Mock}
}

// DownloadOrganizations provides a mock function for the type organizationCrudApi
func (_mock *organizationCrudApi) DownloadOrganizations(filter filters.V2Filter) []string {
	ret := _mock.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for DownloadOrganizations")
	}

	var r0 []string
	if returnFunc, ok := ret.Get(0).(func(filters.V2Filter) []string); ok {
		r0 = returnFunc(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}
	return r0
}

// organizationCrudApi_DownloadOrganizations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DownloadOrganizations'
type organizationCrudApi_DownloadOrganizations_Call struct {
	*mock.Call
}

// DownloadOrganizations is a helper method to define mock.On call
//   - filter filters.V2Filter
func (_e *organizationCrudApi_Expecter) DownloadOrganizations(filter interface{}) *organizationCrudApi_DownloadOrganizations_Call {
	return &organizationCrudApi_DownloadOrganizations_Call{Call: _e.mock.On("DownloadOrganizations", filter)}
}

func (_c *organizationCrudApi_DownloadOrganizations_Call) Run(run func(filter filters.V2Filter)) *organizationCrudApi_DownloadOrganizations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 filters.V2Filter
		if args[0] != nil {
			arg0 = args[0].(filters.V2Filter)
		}
		run(
			arg0,
		)
	})
	return _c
}

func (_c *organizationCrudApi_DownloadOrganizations_Call) Return(strings []string) *organizationCrudApi_DownloadOrganizations_Call {
	_c.Call.Return(strings)
	return _c
}

func (_c *organizationCrudApi_DownloadOrganizations_Call) RunAndReturn(run func(filter filters.V2Filter) []string) *organizationCrudApi_DownloadOrganizations_Call {
	_c.Call.Return(run)
	return _c
}

// ListOrganizations provides a mock function for the type organizationCrudApi
func (_mock *organizationCrudApi) ListOrganizations(filter filters.V2Filter, withPreferences bool) []*domain.OrgsDTOWithPreferences {
	ret := _mock.Called(filter, withPreferences)

	if len(ret) == 0 {
		panic("no return value specified for ListOrganizations")
	}

	var r0 []*domain.OrgsDTOWithPreferences
	if returnFunc, ok := ret.Get(0).(func(filters.V2Filter, bool) []*domain.OrgsDTOWithPreferences); ok {
		r0 = returnFunc(filter, withPreferences)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.OrgsDTOWithPreferences)
		}
	}
	return r0
}

// organizationCrudApi_ListOrganizations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListOrganizations'
type organizationCrudApi_ListOrganizations_Call struct {
	*mock.Call
}

// ListOrganizations is a helper method to define mock.On call
//   - filter filters.V2Filter
//   - withPreferences bool
func (_e *organizationCrudApi_Expecter) ListOrganizations(filter interface{}, withPreferences interface{}) *organizationCrudApi_ListOrganizations_Call {
	return &organizationCrudApi_ListOrganizations_Call{Call: _e.mock.On("ListOrganizations", filter, withPreferences)}
}

func (_c *organizationCrudApi_ListOrganizations_Call) Run(run func(filter filters.V2Filter, withPreferences bool)) *organizationCrudApi_ListOrganizations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 filters.V2Filter
		if args[0] != nil {
			arg0 = args[0].(filters.V2Filter)
		}
		var arg1 bool
		if args[1] != nil {
			arg1 = args[1].(bool)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *organizationCrudApi_ListOrganizations_Call) Return(orgsDTOWithPreferencess []*domain.OrgsDTOWithPreferences) *organizationCrudApi_ListOrganizations_Call {
	_c.Call.Return(orgsDTOWithPreferencess)
	return _c
}

func (_c *organizationCrudApi_ListOrganizations_Call) RunAndReturn(run func(filter filters.V2Filter, withPreferences bool) []*domain.OrgsDTOWithPreferences) *organizationCrudApi_ListOrganizations_Call {
	_c.Call.Return(run)
	return _c
}

// UploadOrganizations provides a mock function for the type organizationCrudApi
func (_mock *organizationCrudApi) UploadOrganizations(filter filters.V2Filter) []string {
	ret := _mock.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for UploadOrganizations")
	}

	var r0 []string
	if returnFunc, ok := ret.Get(0).(func(filters.V2Filter) []string); ok {
		r0 = returnFunc(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}
	return r0
}

// organizationCrudApi_UploadOrganizations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadOrganizations'
type organizationCrudApi_UploadOrganizations_Call struct {
	*mock.Call
}

// UploadOrganizations is a helper method to define mock.On call
//   - filter filters.V2Filter
func (_e *organizationCrudApi_Expecter) UploadOrganizations(filter interface{}) *organizationCrudApi_UploadOrganizations_Call {
	return &organizationCrudApi_UploadOrganizations_Call{Call: _e.mock.On("UploadOrganizations", filter)}
}

func (_c *organizationCrudApi_UploadOrganizations_Call) Run(run func(filter filters.V2Filter)) *organizationCrudApi_UploadOrganizations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 filters.V2Filter
		if args[0] != nil {
			arg0 = args[0].(filters.V2Filter)
		}
		run(
			arg0,
		)
	})
	return _c
}

func (_c *organizationCrudApi_UploadOrganizations_Call) Return(strings []string) *organizationCrudApi_UploadOrganizations_Call {
	_c.Call.Return(strings)
	return _c
}

func (_c *organizationCrudApi_UploadOrganizations_Call) RunAndReturn(run func(filter filters.V2Filter) []string) *organizationCrudApi_UploadOrganizations_Call {
	_c.Call.Return(run)
	return _c
}
