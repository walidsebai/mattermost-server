// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"
import store "github.com/mattermost/mattermost-server/store"

// OAuthStore is an autogenerated mock type for the OAuthStore type
type OAuthStore struct {
	mock.Mock
}

// DeleteApp provides a mock function with given fields: id
func (_m *OAuthStore) DeleteApp(id string) store.StoreChannel {
	ret := _m.Called(id)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAccessData provides a mock function with given fields: token
func (_m *OAuthStore) GetAccessData(token string) store.StoreChannel {
	ret := _m.Called(token)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAccessDataByRefreshToken provides a mock function with given fields: token
func (_m *OAuthStore) GetAccessDataByRefreshToken(token string) store.StoreChannel {
	ret := _m.Called(token)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAccessDataByUserForApp provides a mock function with given fields: userId, clientId
func (_m *OAuthStore) GetAccessDataByUserForApp(userId string, clientId string) store.StoreChannel {
	ret := _m.Called(userId, clientId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(userId, clientId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetApp provides a mock function with given fields: id
func (_m *OAuthStore) GetApp(id string) store.StoreChannel {
	ret := _m.Called(id)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAppByUser provides a mock function with given fields: userId, offset, limit
func (_m *OAuthStore) GetAppByUser(userId string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(userId, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int) store.StoreChannel); ok {
		r0 = rf(userId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetApps provides a mock function with given fields: offset, limit
func (_m *OAuthStore) GetApps(offset int, limit int) store.StoreChannel {
	ret := _m.Called(offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int, int) store.StoreChannel); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAuthData provides a mock function with given fields: code
func (_m *OAuthStore) GetAuthData(code string) store.StoreChannel {
	ret := _m.Called(code)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAuthorizedApps provides a mock function with given fields: userId, offset, limit
func (_m *OAuthStore) GetAuthorizedApps(userId string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(userId, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int) store.StoreChannel); ok {
		r0 = rf(userId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetPreviousAccessData provides a mock function with given fields: userId, clientId
func (_m *OAuthStore) GetPreviousAccessData(userId string, clientId string) store.StoreChannel {
	ret := _m.Called(userId, clientId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(userId, clientId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteAuthDataByUser provides a mock function with given fields: userId
func (_m *OAuthStore) PermanentDeleteAuthDataByUser(userId string) store.StoreChannel {
	ret := _m.Called(userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// RemoveAccessData provides a mock function with given fields: token
func (_m *OAuthStore) RemoveAccessData(token string) store.StoreChannel {
	ret := _m.Called(token)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
func (_m *OAuthStore) RemoveAllAccessData() store.StoreChannel {
	var r0 store.StoreChannel
	return r0
}

// RemoveAuthData provides a mock function with given fields: code
func (_m *OAuthStore) RemoveAuthData(code string) store.StoreChannel {
	ret := _m.Called(code)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// SaveAccessData provides a mock function with given fields: accessData
func (_m *OAuthStore) SaveAccessData(accessData *model.AccessData) store.StoreChannel {
	ret := _m.Called(accessData)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.AccessData) store.StoreChannel); ok {
		r0 = rf(accessData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// SaveApp provides a mock function with given fields: app
func (_m *OAuthStore) SaveApp(app *model.OAuthApp) store.StoreChannel {
	ret := _m.Called(app)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.OAuthApp) store.StoreChannel); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// SaveAuthData provides a mock function with given fields: authData
func (_m *OAuthStore) SaveAuthData(authData *model.AuthData) store.StoreChannel {
	ret := _m.Called(authData)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.AuthData) store.StoreChannel); ok {
		r0 = rf(authData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateAccessData provides a mock function with given fields: accessData
func (_m *OAuthStore) UpdateAccessData(accessData *model.AccessData) store.StoreChannel {
	ret := _m.Called(accessData)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.AccessData) store.StoreChannel); ok {
		r0 = rf(accessData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateApp provides a mock function with given fields: app
func (_m *OAuthStore) UpdateApp(app *model.OAuthApp) store.StoreChannel {
	ret := _m.Called(app)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.OAuthApp) store.StoreChannel); ok {
		r0 = rf(app)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
