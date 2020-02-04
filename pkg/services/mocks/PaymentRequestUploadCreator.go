// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	uuid "github.com/gofrs/uuid"
)

// PaymentRequestUploadCreator is an autogenerated mock type for the PaymentRequestUploadCreator type
type PaymentRequestUploadCreator struct {
	mock.Mock
}

// CreateUpload provides a mock function with given fields: file, paymentRequestID, userID
func (_m *PaymentRequestUploadCreator) CreateUpload(file io.ReadCloser, paymentRequestID uuid.UUID, userID uuid.UUID) (*models.Upload, error) {
	ret := _m.Called(file, paymentRequestID, userID)

	var r0 *models.Upload
	if rf, ok := ret.Get(0).(func(io.ReadCloser, uuid.UUID, uuid.UUID) *models.Upload); ok {
		r0 = rf(file, paymentRequestID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Upload)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.ReadCloser, uuid.UUID, uuid.UUID) error); ok {
		r1 = rf(file, paymentRequestID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
