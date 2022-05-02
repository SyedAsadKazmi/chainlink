// Code generated by mockery v2.10.1. DO NOT EDIT.

package mocks

import (
	context "context"

	headtrackertypes "github.com/smartcontractkit/chainlink/core/chains/evm/headtracker/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink/core/chains/evm/types"
)

// HeadBroadcaster is an autogenerated mock type for the HeadBroadcaster type
type HeadBroadcaster struct {
	mock.Mock
}

// BroadcastNewLongestChain provides a mock function with given fields: head
func (_m *HeadBroadcaster) BroadcastNewLongestChain(head *types.Head) {
	_m.Called(head)
}

// Close provides a mock function with given fields:
func (_m *HeadBroadcaster) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Healthy provides a mock function with given fields:
func (_m *HeadBroadcaster) Healthy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *HeadBroadcaster) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *HeadBroadcaster) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: callback
func (_m *HeadBroadcaster) Subscribe(callback headtrackertypes.HeadTrackable) (*types.Head, func()) {
	ret := _m.Called(callback)

	var r0 *types.Head
	if rf, ok := ret.Get(0).(func(headtrackertypes.HeadTrackable) *types.Head); ok {
		r0 = rf(callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Head)
		}
	}

	var r1 func()
	if rf, ok := ret.Get(1).(func(headtrackertypes.HeadTrackable) func()); ok {
		r1 = rf(callback)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(func())
		}
	}

	return r0, r1
}
