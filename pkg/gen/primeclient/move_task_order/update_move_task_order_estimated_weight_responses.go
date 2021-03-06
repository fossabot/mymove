// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	primemessages "github.com/transcom/mymove/pkg/gen/primemessages"
)

// UpdateMoveTaskOrderEstimatedWeightReader is a Reader for the UpdateMoveTaskOrderEstimatedWeight structure.
type UpdateMoveTaskOrderEstimatedWeightReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMoveTaskOrderEstimatedWeightReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMoveTaskOrderEstimatedWeightOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateMoveTaskOrderEstimatedWeightUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMoveTaskOrderEstimatedWeightForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateMoveTaskOrderEstimatedWeightNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateMoveTaskOrderEstimatedWeightUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateMoveTaskOrderEstimatedWeightInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateMoveTaskOrderEstimatedWeightOK creates a UpdateMoveTaskOrderEstimatedWeightOK with default headers values
func NewUpdateMoveTaskOrderEstimatedWeightOK() *UpdateMoveTaskOrderEstimatedWeightOK {
	return &UpdateMoveTaskOrderEstimatedWeightOK{}
}

/*UpdateMoveTaskOrderEstimatedWeightOK handles this case with default header values.

Successfully retrieved move task order
*/
type UpdateMoveTaskOrderEstimatedWeightOK struct {
	Payload *primemessages.MoveTaskOrder
}

func (o *UpdateMoveTaskOrderEstimatedWeightOK) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/prime-estimated-weight][%d] updateMoveTaskOrderEstimatedWeightOK  %+v", 200, o.Payload)
}

func (o *UpdateMoveTaskOrderEstimatedWeightOK) GetPayload() *primemessages.MoveTaskOrder {
	return o.Payload
}

func (o *UpdateMoveTaskOrderEstimatedWeightOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.MoveTaskOrder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMoveTaskOrderEstimatedWeightUnauthorized creates a UpdateMoveTaskOrderEstimatedWeightUnauthorized with default headers values
func NewUpdateMoveTaskOrderEstimatedWeightUnauthorized() *UpdateMoveTaskOrderEstimatedWeightUnauthorized {
	return &UpdateMoveTaskOrderEstimatedWeightUnauthorized{}
}

/*UpdateMoveTaskOrderEstimatedWeightUnauthorized handles this case with default header values.

The request was denied
*/
type UpdateMoveTaskOrderEstimatedWeightUnauthorized struct {
	Payload interface{}
}

func (o *UpdateMoveTaskOrderEstimatedWeightUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/prime-estimated-weight][%d] updateMoveTaskOrderEstimatedWeightUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateMoveTaskOrderEstimatedWeightUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMoveTaskOrderEstimatedWeightUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMoveTaskOrderEstimatedWeightForbidden creates a UpdateMoveTaskOrderEstimatedWeightForbidden with default headers values
func NewUpdateMoveTaskOrderEstimatedWeightForbidden() *UpdateMoveTaskOrderEstimatedWeightForbidden {
	return &UpdateMoveTaskOrderEstimatedWeightForbidden{}
}

/*UpdateMoveTaskOrderEstimatedWeightForbidden handles this case with default header values.

The request was denied
*/
type UpdateMoveTaskOrderEstimatedWeightForbidden struct {
	Payload interface{}
}

func (o *UpdateMoveTaskOrderEstimatedWeightForbidden) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/prime-estimated-weight][%d] updateMoveTaskOrderEstimatedWeightForbidden  %+v", 403, o.Payload)
}

func (o *UpdateMoveTaskOrderEstimatedWeightForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMoveTaskOrderEstimatedWeightForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMoveTaskOrderEstimatedWeightNotFound creates a UpdateMoveTaskOrderEstimatedWeightNotFound with default headers values
func NewUpdateMoveTaskOrderEstimatedWeightNotFound() *UpdateMoveTaskOrderEstimatedWeightNotFound {
	return &UpdateMoveTaskOrderEstimatedWeightNotFound{}
}

/*UpdateMoveTaskOrderEstimatedWeightNotFound handles this case with default header values.

The requested resource wasn't found
*/
type UpdateMoveTaskOrderEstimatedWeightNotFound struct {
	Payload interface{}
}

func (o *UpdateMoveTaskOrderEstimatedWeightNotFound) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/prime-estimated-weight][%d] updateMoveTaskOrderEstimatedWeightNotFound  %+v", 404, o.Payload)
}

func (o *UpdateMoveTaskOrderEstimatedWeightNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMoveTaskOrderEstimatedWeightNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMoveTaskOrderEstimatedWeightUnprocessableEntity creates a UpdateMoveTaskOrderEstimatedWeightUnprocessableEntity with default headers values
func NewUpdateMoveTaskOrderEstimatedWeightUnprocessableEntity() *UpdateMoveTaskOrderEstimatedWeightUnprocessableEntity {
	return &UpdateMoveTaskOrderEstimatedWeightUnprocessableEntity{}
}

/*UpdateMoveTaskOrderEstimatedWeightUnprocessableEntity handles this case with default header values.

The request payload is invalid
*/
type UpdateMoveTaskOrderEstimatedWeightUnprocessableEntity struct {
	Payload *primemessages.ValidationError
}

func (o *UpdateMoveTaskOrderEstimatedWeightUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/prime-estimated-weight][%d] updateMoveTaskOrderEstimatedWeightUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateMoveTaskOrderEstimatedWeightUnprocessableEntity) GetPayload() *primemessages.ValidationError {
	return o.Payload
}

func (o *UpdateMoveTaskOrderEstimatedWeightUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMoveTaskOrderEstimatedWeightInternalServerError creates a UpdateMoveTaskOrderEstimatedWeightInternalServerError with default headers values
func NewUpdateMoveTaskOrderEstimatedWeightInternalServerError() *UpdateMoveTaskOrderEstimatedWeightInternalServerError {
	return &UpdateMoveTaskOrderEstimatedWeightInternalServerError{}
}

/*UpdateMoveTaskOrderEstimatedWeightInternalServerError handles this case with default header values.

A server error occurred
*/
type UpdateMoveTaskOrderEstimatedWeightInternalServerError struct {
	Payload interface{}
}

func (o *UpdateMoveTaskOrderEstimatedWeightInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /move-task-orders/{moveTaskOrderID}/prime-estimated-weight][%d] updateMoveTaskOrderEstimatedWeightInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateMoveTaskOrderEstimatedWeightInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateMoveTaskOrderEstimatedWeightInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateMoveTaskOrderEstimatedWeightBody update move task order estimated weight body
swagger:model UpdateMoveTaskOrderEstimatedWeightBody
*/
type UpdateMoveTaskOrderEstimatedWeightBody struct {

	// prime estimated weight
	PrimeEstimatedWeight int64 `json:"primeEstimatedWeight,omitempty"`
}

// Validate validates this update move task order estimated weight body
func (o *UpdateMoveTaskOrderEstimatedWeightBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateMoveTaskOrderEstimatedWeightBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateMoveTaskOrderEstimatedWeightBody) UnmarshalBinary(b []byte) error {
	var res UpdateMoveTaskOrderEstimatedWeightBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
