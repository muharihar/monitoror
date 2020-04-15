// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/monitoror/monitoror/api/config/models"
	mock "github.com/stretchr/testify/mock"
)

// VariantMetadataExplorer is an autogenerated mock type for the VariantMetadataExplorer type
type VariantMetadataExplorer struct {
	mock.Mock
}

// GetValidator provides a mock function with given fields:
func (_m *VariantMetadataExplorer) GetValidator() models.ParamsValidator {
	ret := _m.Called()

	var r0 models.ParamsValidator
	if rf, ok := ret.Get(0).(func() models.ParamsValidator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.ParamsValidator)
		}
	}

	return r0
}

// IsEnabled provides a mock function with given fields:
func (_m *VariantMetadataExplorer) IsEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}