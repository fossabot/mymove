// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	ghcmessages "github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// CreateMTOServiceItemCreatedCode is the HTTP code returned for type CreateMTOServiceItemCreated
const CreateMTOServiceItemCreatedCode int = 201

/*CreateMTOServiceItemCreated Successfully created service item for move task order

swagger:response createMTOServiceItemCreated
*/
type CreateMTOServiceItemCreated struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.MTOServiceItem `json:"body,omitempty"`
}

// NewCreateMTOServiceItemCreated creates CreateMTOServiceItemCreated with default headers values
func NewCreateMTOServiceItemCreated() *CreateMTOServiceItemCreated {

	return &CreateMTOServiceItemCreated{}
}

// WithPayload adds the payload to the create m t o service item created response
func (o *CreateMTOServiceItemCreated) WithPayload(payload *ghcmessages.MTOServiceItem) *CreateMTOServiceItemCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create m t o service item created response
func (o *CreateMTOServiceItemCreated) SetPayload(payload *ghcmessages.MTOServiceItem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMTOServiceItemCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMTOServiceItemNotFoundCode is the HTTP code returned for type CreateMTOServiceItemNotFound
const CreateMTOServiceItemNotFoundCode int = 404

/*CreateMTOServiceItemNotFound The requested resource wasn't found

swagger:response createMTOServiceItemNotFound
*/
type CreateMTOServiceItemNotFound struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.ClientError `json:"body,omitempty"`
}

// NewCreateMTOServiceItemNotFound creates CreateMTOServiceItemNotFound with default headers values
func NewCreateMTOServiceItemNotFound() *CreateMTOServiceItemNotFound {

	return &CreateMTOServiceItemNotFound{}
}

// WithPayload adds the payload to the create m t o service item not found response
func (o *CreateMTOServiceItemNotFound) WithPayload(payload *ghcmessages.ClientError) *CreateMTOServiceItemNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create m t o service item not found response
func (o *CreateMTOServiceItemNotFound) SetPayload(payload *ghcmessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMTOServiceItemNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMTOServiceItemUnprocessableEntityCode is the HTTP code returned for type CreateMTOServiceItemUnprocessableEntity
const CreateMTOServiceItemUnprocessableEntityCode int = 422

/*CreateMTOServiceItemUnprocessableEntity Validation error

swagger:response createMTOServiceItemUnprocessableEntity
*/
type CreateMTOServiceItemUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.ValidationError `json:"body,omitempty"`
}

// NewCreateMTOServiceItemUnprocessableEntity creates CreateMTOServiceItemUnprocessableEntity with default headers values
func NewCreateMTOServiceItemUnprocessableEntity() *CreateMTOServiceItemUnprocessableEntity {

	return &CreateMTOServiceItemUnprocessableEntity{}
}

// WithPayload adds the payload to the create m t o service item unprocessable entity response
func (o *CreateMTOServiceItemUnprocessableEntity) WithPayload(payload *ghcmessages.ValidationError) *CreateMTOServiceItemUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create m t o service item unprocessable entity response
func (o *CreateMTOServiceItemUnprocessableEntity) SetPayload(payload *ghcmessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMTOServiceItemUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMTOServiceItemInternalServerErrorCode is the HTTP code returned for type CreateMTOServiceItemInternalServerError
const CreateMTOServiceItemInternalServerErrorCode int = 500

/*CreateMTOServiceItemInternalServerError A server error occurred

swagger:response createMTOServiceItemInternalServerError
*/
type CreateMTOServiceItemInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewCreateMTOServiceItemInternalServerError creates CreateMTOServiceItemInternalServerError with default headers values
func NewCreateMTOServiceItemInternalServerError() *CreateMTOServiceItemInternalServerError {

	return &CreateMTOServiceItemInternalServerError{}
}

// WithPayload adds the payload to the create m t o service item internal server error response
func (o *CreateMTOServiceItemInternalServerError) WithPayload(payload interface{}) *CreateMTOServiceItemInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create m t o service item internal server error response
func (o *CreateMTOServiceItemInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMTOServiceItemInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
