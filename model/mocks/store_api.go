package mocks

import "github.com/solderapp/solder-api/model"
import "github.com/stretchr/testify/mock"

// StoreAPI describes a store API.
type StoreAPI struct {
	mock.Mock
}

// GetBuilds provides a mock function with given fields: _a0
func (_m *StoreAPI) GetBuilds(_a0 int) (*model.Builds, error) {
	ret := _m.Called(_a0)

	var r0 *model.Builds
	if rf, ok := ret.Get(0).(func(int) *model.Builds); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Builds)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBuild provides a mock function with given fields: _a0, _a1
func (_m *StoreAPI) GetBuild(_a0 int, _a1 string) (*model.Build, *model.Store) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Build
	if rf, ok := ret.Get(0).(func(int, string) *model.Build); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Build)
		}
	}

	var r1 *model.Store
	if rf, ok := ret.Get(1).(func(int, string) *model.Store); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.Store)
		}
	}

	return r0, r1
}

// GetClients provides a mock function with given fields:
func (_m *StoreAPI) GetClients() (*model.Clients, error) {
	ret := _m.Called()

	var r0 *model.Clients
	if rf, ok := ret.Get(0).(func() *model.Clients); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Clients)
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

// GetClient provides a mock function with given fields: _a0
func (_m *StoreAPI) GetClient(_a0 string) (*model.Client, *model.Store) {
	ret := _m.Called(_a0)

	var r0 *model.Client
	if rf, ok := ret.Get(0).(func(string) *model.Client); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Client)
		}
	}

	var r1 *model.Store
	if rf, ok := ret.Get(1).(func(string) *model.Store); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.Store)
		}
	}

	return r0, r1
}

// GetForges provides a mock function with given fields:
func (_m *StoreAPI) GetForges() (*model.Forges, error) {
	ret := _m.Called()

	var r0 *model.Forges
	if rf, ok := ret.Get(0).(func() *model.Forges); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Forges)
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

// GetForge provides a mock function with given fields: _a0
func (_m *StoreAPI) GetForge(_a0 string) (*model.Forge, *model.Store) {
	ret := _m.Called(_a0)

	var r0 *model.Forge
	if rf, ok := ret.Get(0).(func(string) *model.Forge); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Forge)
		}
	}

	var r1 *model.Store
	if rf, ok := ret.Get(1).(func(string) *model.Store); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.Store)
		}
	}

	return r0, r1
}

// GetKeys provides a mock function with given fields:
func (_m *StoreAPI) GetKeys() (*model.Keys, error) {
	ret := _m.Called()

	var r0 *model.Keys
	if rf, ok := ret.Get(0).(func() *model.Keys); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Keys)
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

// GetKey provides a mock function with given fields: _a0
func (_m *StoreAPI) GetKey(_a0 string) (*model.Key, *model.Store) {
	ret := _m.Called(_a0)

	var r0 *model.Key
	if rf, ok := ret.Get(0).(func(string) *model.Key); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Key)
		}
	}

	var r1 *model.Store
	if rf, ok := ret.Get(1).(func(string) *model.Store); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.Store)
		}
	}

	return r0, r1
}

// GetMinecrafts provides a mock function with given fields:
func (_m *StoreAPI) GetMinecrafts() (*model.Minecrafts, error) {
	ret := _m.Called()

	var r0 *model.Minecrafts
	if rf, ok := ret.Get(0).(func() *model.Minecrafts); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Minecrafts)
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

// GetMinecraft provides a mock function with given fields: _a0
func (_m *StoreAPI) GetMinecraft(_a0 string) (*model.Minecraft, *model.Store) {
	ret := _m.Called(_a0)

	var r0 *model.Minecraft
	if rf, ok := ret.Get(0).(func(string) *model.Minecraft); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Minecraft)
		}
	}

	var r1 *model.Store
	if rf, ok := ret.Get(1).(func(string) *model.Store); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.Store)
		}
	}

	return r0, r1
}

// GetMods provides a mock function with given fields:
func (_m *StoreAPI) GetMods() (*model.Mods, error) {
	ret := _m.Called()

	var r0 *model.Mods
	if rf, ok := ret.Get(0).(func() *model.Mods); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Mods)
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

// GetMod provides a mock function with given fields: _a0
func (_m *StoreAPI) GetMod(_a0 string) (*model.Mod, *model.Store) {
	ret := _m.Called(_a0)

	var r0 *model.Mod
	if rf, ok := ret.Get(0).(func(string) *model.Mod); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Mod)
		}
	}

	var r1 *model.Store
	if rf, ok := ret.Get(1).(func(string) *model.Store); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.Store)
		}
	}

	return r0, r1
}

// GetPacks provides a mock function with given fields:
func (_m *StoreAPI) GetPacks() (*model.Packs, error) {
	ret := _m.Called()

	var r0 *model.Packs
	if rf, ok := ret.Get(0).(func() *model.Packs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Packs)
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

// GetPack provides a mock function with given fields: _a0
func (_m *StoreAPI) GetPack(_a0 string) (*model.Pack, *model.Store) {
	ret := _m.Called(_a0)

	var r0 *model.Pack
	if rf, ok := ret.Get(0).(func(string) *model.Pack); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Pack)
		}
	}

	var r1 *model.Store
	if rf, ok := ret.Get(1).(func(string) *model.Store); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.Store)
		}
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields:
func (_m *StoreAPI) GetUsers() (*model.Users, error) {
	ret := _m.Called()

	var r0 *model.Users
	if rf, ok := ret.Get(0).(func() *model.Users); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Users)
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

// GetUser provides a mock function with given fields: _a0
func (_m *StoreAPI) GetUser(_a0 string) (*model.User, *model.Store) {
	ret := _m.Called(_a0)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 *model.Store
	if rf, ok := ret.Get(1).(func(string) *model.Store); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.Store)
		}
	}

	return r0, r1
}

// GetVersions provides a mock function with given fields: _a0
func (_m *StoreAPI) GetVersions(_a0 int) (*model.Versions, error) {
	ret := _m.Called(_a0)

	var r0 *model.Versions
	if rf, ok := ret.Get(0).(func(int) *model.Versions); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Versions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVersion provides a mock function with given fields: _a0, _a1
func (_m *StoreAPI) GetVersion(_a0 int, _a1 string) (*model.Version, *model.Store) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Version
	if rf, ok := ret.Get(0).(func(int, string) *model.Version); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Version)
		}
	}

	var r1 *model.Store
	if rf, ok := ret.Get(1).(func(int, string) *model.Store); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.Store)
		}
	}

	return r0, r1
}
