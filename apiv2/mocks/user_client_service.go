// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	user "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/user"
)

// MockUserClientService is an autogenerated mock type for the ClientService type
type MockUserClientService struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: params, authInfo
func (_m *MockUserClientService) CreateUser(params *user.CreateUserParams, authInfo runtime.ClientAuthInfoWriter) (*user.CreateUserCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *user.CreateUserCreated
	if rf, ok := ret.Get(0).(func(*user.CreateUserParams, runtime.ClientAuthInfoWriter) *user.CreateUserCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.CreateUserCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.CreateUserParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: params, authInfo
func (_m *MockUserClientService) DeleteUser(params *user.DeleteUserParams, authInfo runtime.ClientAuthInfoWriter) (*user.DeleteUserOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *user.DeleteUserOK
	if rf, ok := ret.Get(0).(func(*user.DeleteUserParams, runtime.ClientAuthInfoWriter) *user.DeleteUserOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.DeleteUserOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.DeleteUserParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrentUserInfo provides a mock function with given fields: params, authInfo
func (_m *MockUserClientService) GetCurrentUserInfo(params *user.GetCurrentUserInfoParams, authInfo runtime.ClientAuthInfoWriter) (*user.GetCurrentUserInfoOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *user.GetCurrentUserInfoOK
	if rf, ok := ret.Get(0).(func(*user.GetCurrentUserInfoParams, runtime.ClientAuthInfoWriter) *user.GetCurrentUserInfoOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.GetCurrentUserInfoOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.GetCurrentUserInfoParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrentUserPermissions provides a mock function with given fields: params, authInfo
func (_m *MockUserClientService) GetCurrentUserPermissions(params *user.GetCurrentUserPermissionsParams, authInfo runtime.ClientAuthInfoWriter) (*user.GetCurrentUserPermissionsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *user.GetCurrentUserPermissionsOK
	if rf, ok := ret.Get(0).(func(*user.GetCurrentUserPermissionsParams, runtime.ClientAuthInfoWriter) *user.GetCurrentUserPermissionsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.GetCurrentUserPermissionsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.GetCurrentUserPermissionsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: params, authInfo
func (_m *MockUserClientService) GetUser(params *user.GetUserParams, authInfo runtime.ClientAuthInfoWriter) (*user.GetUserOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *user.GetUserOK
	if rf, ok := ret.Get(0).(func(*user.GetUserParams, runtime.ClientAuthInfoWriter) *user.GetUserOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.GetUserOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.GetUserParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUsers provides a mock function with given fields: params, authInfo
func (_m *MockUserClientService) ListUsers(params *user.ListUsersParams, authInfo runtime.ClientAuthInfoWriter) (*user.ListUsersOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *user.ListUsersOK
	if rf, ok := ret.Get(0).(func(*user.ListUsersParams, runtime.ClientAuthInfoWriter) *user.ListUsersOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.ListUsersOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.ListUsersParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchUsers provides a mock function with given fields: params, authInfo
func (_m *MockUserClientService) SearchUsers(params *user.SearchUsersParams, authInfo runtime.ClientAuthInfoWriter) (*user.SearchUsersOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *user.SearchUsersOK
	if rf, ok := ret.Get(0).(func(*user.SearchUsersParams, runtime.ClientAuthInfoWriter) *user.SearchUsersOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.SearchUsersOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.SearchUsersParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetCliSecret provides a mock function with given fields: params, authInfo
func (_m *MockUserClientService) SetCliSecret(params *user.SetCliSecretParams, authInfo runtime.ClientAuthInfoWriter) (*user.SetCliSecretOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *user.SetCliSecretOK
	if rf, ok := ret.Get(0).(func(*user.SetCliSecretParams, runtime.ClientAuthInfoWriter) *user.SetCliSecretOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.SetCliSecretOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.SetCliSecretParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockUserClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// SetUserSysAdmin provides a mock function with given fields: params, authInfo
func (_m *MockUserClientService) SetUserSysAdmin(params *user.SetUserSysAdminParams, authInfo runtime.ClientAuthInfoWriter) (*user.SetUserSysAdminOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *user.SetUserSysAdminOK
	if rf, ok := ret.Get(0).(func(*user.SetUserSysAdminParams, runtime.ClientAuthInfoWriter) *user.SetUserSysAdminOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.SetUserSysAdminOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.SetUserSysAdminParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserPassword provides a mock function with given fields: params, authInfo
func (_m *MockUserClientService) UpdateUserPassword(params *user.UpdateUserPasswordParams, authInfo runtime.ClientAuthInfoWriter) (*user.UpdateUserPasswordOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *user.UpdateUserPasswordOK
	if rf, ok := ret.Get(0).(func(*user.UpdateUserPasswordParams, runtime.ClientAuthInfoWriter) *user.UpdateUserPasswordOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.UpdateUserPasswordOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.UpdateUserPasswordParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserProfile provides a mock function with given fields: params, authInfo
func (_m *MockUserClientService) UpdateUserProfile(params *user.UpdateUserProfileParams, authInfo runtime.ClientAuthInfoWriter) (*user.UpdateUserProfileOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *user.UpdateUserProfileOK
	if rf, ok := ret.Get(0).(func(*user.UpdateUserProfileParams, runtime.ClientAuthInfoWriter) *user.UpdateUserProfileOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.UpdateUserProfileOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.UpdateUserProfileParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}