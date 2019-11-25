// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	services "github.com/transcom/mymove/pkg/services"
)

// MoveListFetcher is an autogenerated mock type for the MoveListFetcher type
type MoveListFetcher struct {
	mock.Mock
}

// FetchMoveCount provides a mock function with given fields: filters
func (_m *MoveListFetcher) FetchMoveCount(filters []services.QueryFilter) (int, error) {
	ret := _m.Called(filters)

	var r0 int
	if rf, ok := ret.Get(0).(func([]services.QueryFilter) int); ok {
		r0 = rf(filters)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]services.QueryFilter) error); ok {
		r1 = rf(filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchMoveList provides a mock function with given fields: filters, associations, pagination, ordering
func (_m *MoveListFetcher) FetchMoveList(filters []services.QueryFilter, associations services.QueryAssociations, pagination services.Pagination, ordering services.QueryOrder) (models.Moves, error) {
	ret := _m.Called(filters, associations, pagination, ordering)

	var r0 models.Moves
	if rf, ok := ret.Get(0).(func([]services.QueryFilter, services.QueryAssociations, services.Pagination, services.QueryOrder) models.Moves); ok {
		r0 = rf(filters, associations, pagination, ordering)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.Moves)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]services.QueryFilter, services.QueryAssociations, services.Pagination, services.QueryOrder) error); ok {
		r1 = rf(filters, associations, pagination, ordering)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}