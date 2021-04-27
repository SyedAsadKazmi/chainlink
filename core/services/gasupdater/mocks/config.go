// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	big "math/big"

	mock "github.com/stretchr/testify/mock"
)

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// EthFinalityDepth provides a mock function with given fields:
func (_m *Config) EthFinalityDepth() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// EthMaxGasPriceWei provides a mock function with given fields:
func (_m *Config) EthMaxGasPriceWei() *big.Int {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GasUpdaterBatchSize provides a mock function with given fields:
func (_m *Config) GasUpdaterBatchSize() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// GasUpdaterBlockDelay provides a mock function with given fields:
func (_m *Config) GasUpdaterBlockDelay() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// GasUpdaterBlockHistorySize provides a mock function with given fields:
func (_m *Config) GasUpdaterBlockHistorySize() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// GasUpdaterTransactionPercentile provides a mock function with given fields:
func (_m *Config) GasUpdaterTransactionPercentile() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// SetEthGasPriceDefault provides a mock function with given fields: value
func (_m *Config) SetEthGasPriceDefault(value *big.Int) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(*big.Int) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
