// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	assets "github.com/smartcontractkit/chainlink/core/assets"

	gas "github.com/smartcontractkit/chainlink/core/chains/evm/gas"

	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink/core/chains/evm/types"
)

// Estimator is an autogenerated mock type for the Estimator type
type Estimator struct {
	mock.Mock
}

// BumpDynamicFee provides a mock function with given fields: original, gasLimit, maxGasPriceWei, attempts
func (_m *Estimator) BumpDynamicFee(original gas.DynamicFee, gasLimit uint32, maxGasPriceWei *assets.Wei, attempts []gas.PriorAttempt) (gas.DynamicFee, uint32, error) {
	ret := _m.Called(original, gasLimit, maxGasPriceWei, attempts)

	var r0 gas.DynamicFee
	if rf, ok := ret.Get(0).(func(gas.DynamicFee, uint32, *assets.Wei, []gas.PriorAttempt) gas.DynamicFee); ok {
		r0 = rf(original, gasLimit, maxGasPriceWei, attempts)
	} else {
		r0 = ret.Get(0).(gas.DynamicFee)
	}

	var r1 uint32
	if rf, ok := ret.Get(1).(func(gas.DynamicFee, uint32, *assets.Wei, []gas.PriorAttempt) uint32); ok {
		r1 = rf(original, gasLimit, maxGasPriceWei, attempts)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(gas.DynamicFee, uint32, *assets.Wei, []gas.PriorAttempt) error); ok {
		r2 = rf(original, gasLimit, maxGasPriceWei, attempts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// BumpLegacyGas provides a mock function with given fields: originalGasPrice, gasLimit, maxGasPriceWei, attempts
func (_m *Estimator) BumpLegacyGas(originalGasPrice *assets.Wei, gasLimit uint32, maxGasPriceWei *assets.Wei, attempts []gas.PriorAttempt) (*assets.Wei, uint32, error) {
	ret := _m.Called(originalGasPrice, gasLimit, maxGasPriceWei, attempts)

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func(*assets.Wei, uint32, *assets.Wei, []gas.PriorAttempt) *assets.Wei); ok {
		r0 = rf(originalGasPrice, gasLimit, maxGasPriceWei, attempts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	var r1 uint32
	if rf, ok := ret.Get(1).(func(*assets.Wei, uint32, *assets.Wei, []gas.PriorAttempt) uint32); ok {
		r1 = rf(originalGasPrice, gasLimit, maxGasPriceWei, attempts)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*assets.Wei, uint32, *assets.Wei, []gas.PriorAttempt) error); ok {
		r2 = rf(originalGasPrice, gasLimit, maxGasPriceWei, attempts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Close provides a mock function with given fields:
func (_m *Estimator) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDynamicFee provides a mock function with given fields: gasLimit, maxGasPriceWei
func (_m *Estimator) GetDynamicFee(gasLimit uint32, maxGasPriceWei *assets.Wei) (gas.DynamicFee, uint32, error) {
	ret := _m.Called(gasLimit, maxGasPriceWei)

	var r0 gas.DynamicFee
	if rf, ok := ret.Get(0).(func(uint32, *assets.Wei) gas.DynamicFee); ok {
		r0 = rf(gasLimit, maxGasPriceWei)
	} else {
		r0 = ret.Get(0).(gas.DynamicFee)
	}

	var r1 uint32
	if rf, ok := ret.Get(1).(func(uint32, *assets.Wei) uint32); ok {
		r1 = rf(gasLimit, maxGasPriceWei)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint32, *assets.Wei) error); ok {
		r2 = rf(gasLimit, maxGasPriceWei)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetLegacyGas provides a mock function with given fields: calldata, gasLimit, maxGasPriceWei, opts
func (_m *Estimator) GetLegacyGas(calldata []byte, gasLimit uint32, maxGasPriceWei *assets.Wei, opts ...gas.Opt) (*assets.Wei, uint32, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, calldata, gasLimit, maxGasPriceWei)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func([]byte, uint32, *assets.Wei, ...gas.Opt) *assets.Wei); ok {
		r0 = rf(calldata, gasLimit, maxGasPriceWei, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	var r1 uint32
	if rf, ok := ret.Get(1).(func([]byte, uint32, *assets.Wei, ...gas.Opt) uint32); ok {
		r1 = rf(calldata, gasLimit, maxGasPriceWei, opts...)
	} else {
		r1 = ret.Get(1).(uint32)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]byte, uint32, *assets.Wei, ...gas.Opt) error); ok {
		r2 = rf(calldata, gasLimit, maxGasPriceWei, opts...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// OnNewLongestChain provides a mock function with given fields: _a0, _a1
func (_m *Estimator) OnNewLongestChain(_a0 context.Context, _a1 *types.Head) {
	_m.Called(_a0, _a1)
}

// Start provides a mock function with given fields: _a0
func (_m *Estimator) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewEstimator interface {
	mock.TestingT
	Cleanup(func())
}

// NewEstimator creates a new instance of Estimator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEstimator(t mockConstructorTestingTNewEstimator) *Estimator {
	mock := &Estimator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
