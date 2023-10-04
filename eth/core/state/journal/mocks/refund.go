// Code generated by mockery v2.35.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	journal "pkg.berachain.dev/polaris/eth/core/state/journal"
)

// Refund is an autogenerated mock type for the Refund type
type Refund struct {
	mock.Mock
}

type Refund_Expecter struct {
	mock *mock.Mock
}

func (_m *Refund) EXPECT() *Refund_Expecter {
	return &Refund_Expecter{mock: &_m.Mock}
}

// AddRefund provides a mock function with given fields: gas
func (_m *Refund) AddRefund(gas uint64) {
	_m.Called(gas)
}

// Refund_AddRefund_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddRefund'
type Refund_AddRefund_Call struct {
	*mock.Call
}

// AddRefund is a helper method to define mock.On call
//   - gas uint64
func (_e *Refund_Expecter) AddRefund(gas interface{}) *Refund_AddRefund_Call {
	return &Refund_AddRefund_Call{Call: _e.mock.On("AddRefund", gas)}
}

func (_c *Refund_AddRefund_Call) Run(run func(gas uint64)) *Refund_AddRefund_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *Refund_AddRefund_Call) Return() *Refund_AddRefund_Call {
	_c.Call.Return()
	return _c
}

func (_c *Refund_AddRefund_Call) RunAndReturn(run func(uint64)) *Refund_AddRefund_Call {
	_c.Call.Return(run)
	return _c
}

// Clone provides a mock function with given fields:
func (_m *Refund) Clone() journal.Refund {
	ret := _m.Called()

	var r0 journal.Refund
	if rf, ok := ret.Get(0).(func() journal.Refund); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(journal.Refund)
		}
	}

	return r0
}

// Refund_Clone_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Clone'
type Refund_Clone_Call struct {
	*mock.Call
}

// Clone is a helper method to define mock.On call
func (_e *Refund_Expecter) Clone() *Refund_Clone_Call {
	return &Refund_Clone_Call{Call: _e.mock.On("Clone")}
}

func (_c *Refund_Clone_Call) Run(run func()) *Refund_Clone_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Refund_Clone_Call) Return(_a0 journal.Refund) *Refund_Clone_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Refund_Clone_Call) RunAndReturn(run func() journal.Refund) *Refund_Clone_Call {
	_c.Call.Return(run)
	return _c
}

// Finalize provides a mock function with given fields:
func (_m *Refund) Finalize() {
	_m.Called()
}

// Refund_Finalize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Finalize'
type Refund_Finalize_Call struct {
	*mock.Call
}

// Finalize is a helper method to define mock.On call
func (_e *Refund_Expecter) Finalize() *Refund_Finalize_Call {
	return &Refund_Finalize_Call{Call: _e.mock.On("Finalize")}
}

func (_c *Refund_Finalize_Call) Run(run func()) *Refund_Finalize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Refund_Finalize_Call) Return() *Refund_Finalize_Call {
	_c.Call.Return()
	return _c
}

func (_c *Refund_Finalize_Call) RunAndReturn(run func()) *Refund_Finalize_Call {
	_c.Call.Return(run)
	return _c
}

// GetRefund provides a mock function with given fields:
func (_m *Refund) GetRefund() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// Refund_GetRefund_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRefund'
type Refund_GetRefund_Call struct {
	*mock.Call
}

// GetRefund is a helper method to define mock.On call
func (_e *Refund_Expecter) GetRefund() *Refund_GetRefund_Call {
	return &Refund_GetRefund_Call{Call: _e.mock.On("GetRefund")}
}

func (_c *Refund_GetRefund_Call) Run(run func()) *Refund_GetRefund_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Refund_GetRefund_Call) Return(_a0 uint64) *Refund_GetRefund_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Refund_GetRefund_Call) RunAndReturn(run func() uint64) *Refund_GetRefund_Call {
	_c.Call.Return(run)
	return _c
}

// RegistryKey provides a mock function with given fields:
func (_m *Refund) RegistryKey() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Refund_RegistryKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegistryKey'
type Refund_RegistryKey_Call struct {
	*mock.Call
}

// RegistryKey is a helper method to define mock.On call
func (_e *Refund_Expecter) RegistryKey() *Refund_RegistryKey_Call {
	return &Refund_RegistryKey_Call{Call: _e.mock.On("RegistryKey")}
}

func (_c *Refund_RegistryKey_Call) Run(run func()) *Refund_RegistryKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Refund_RegistryKey_Call) Return(_a0 string) *Refund_RegistryKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Refund_RegistryKey_Call) RunAndReturn(run func() string) *Refund_RegistryKey_Call {
	_c.Call.Return(run)
	return _c
}

// RevertToSnapshot provides a mock function with given fields: _a0
func (_m *Refund) RevertToSnapshot(_a0 int) {
	_m.Called(_a0)
}

// Refund_RevertToSnapshot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevertToSnapshot'
type Refund_RevertToSnapshot_Call struct {
	*mock.Call
}

// RevertToSnapshot is a helper method to define mock.On call
//   - _a0 int
func (_e *Refund_Expecter) RevertToSnapshot(_a0 interface{}) *Refund_RevertToSnapshot_Call {
	return &Refund_RevertToSnapshot_Call{Call: _e.mock.On("RevertToSnapshot", _a0)}
}

func (_c *Refund_RevertToSnapshot_Call) Run(run func(_a0 int)) *Refund_RevertToSnapshot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Refund_RevertToSnapshot_Call) Return() *Refund_RevertToSnapshot_Call {
	_c.Call.Return()
	return _c
}

func (_c *Refund_RevertToSnapshot_Call) RunAndReturn(run func(int)) *Refund_RevertToSnapshot_Call {
	_c.Call.Return(run)
	return _c
}

// Snapshot provides a mock function with given fields:
func (_m *Refund) Snapshot() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Refund_Snapshot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Snapshot'
type Refund_Snapshot_Call struct {
	*mock.Call
}

// Snapshot is a helper method to define mock.On call
func (_e *Refund_Expecter) Snapshot() *Refund_Snapshot_Call {
	return &Refund_Snapshot_Call{Call: _e.mock.On("Snapshot")}
}

func (_c *Refund_Snapshot_Call) Run(run func()) *Refund_Snapshot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Refund_Snapshot_Call) Return(_a0 int) *Refund_Snapshot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Refund_Snapshot_Call) RunAndReturn(run func() int) *Refund_Snapshot_Call {
	_c.Call.Return(run)
	return _c
}

// SubRefund provides a mock function with given fields: gas
func (_m *Refund) SubRefund(gas uint64) {
	_m.Called(gas)
}

// Refund_SubRefund_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubRefund'
type Refund_SubRefund_Call struct {
	*mock.Call
}

// SubRefund is a helper method to define mock.On call
//   - gas uint64
func (_e *Refund_Expecter) SubRefund(gas interface{}) *Refund_SubRefund_Call {
	return &Refund_SubRefund_Call{Call: _e.mock.On("SubRefund", gas)}
}

func (_c *Refund_SubRefund_Call) Run(run func(gas uint64)) *Refund_SubRefund_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *Refund_SubRefund_Call) Return() *Refund_SubRefund_Call {
	_c.Call.Return()
	return _c
}

func (_c *Refund_SubRefund_Call) RunAndReturn(run func(uint64)) *Refund_SubRefund_Call {
	_c.Call.Return(run)
	return _c
}

// NewRefund creates a new instance of Refund. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRefund(t interface {
	mock.TestingT
	Cleanup(func())
}) *Refund {
	mock := &Refund{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
