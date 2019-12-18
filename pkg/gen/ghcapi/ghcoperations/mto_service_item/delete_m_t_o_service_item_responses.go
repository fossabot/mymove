// Code generated by go-swagger; DO NOT EDIT.

package mto_service_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	ghcmessages "github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// DeleteMTOServiceItemOKCode is the HTTP code returned for type DeleteMTOServiceItemOK
const DeleteMTOServiceItemOKCode int = 200

/*DeleteMTOServiceItemOK Successfully deleted move task order

swagger:response deleteMTOServiceItemOK
*/
type DeleteMTOServiceItemOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.MoveTaskOrder `json:"body,omitempty"`
}

// NewDeleteMTOServiceItemOK creates DeleteMTOServiceItemOK with default headers values
func NewDeleteMTOServiceItemOK() *DeleteMTOServiceItemOK {

	return &DeleteMTOServiceItemOK{}
}

// WithPayload adds the payload to the delete m t o service item o k response
func (o *DeleteMTOServiceItemOK) WithPayload(payload *ghcmessages.MoveTaskOrder) *DeleteMTOServiceItemOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete m t o service item o k response
func (o *DeleteMTOServiceItemOK) SetPayload(payload *ghcmessages.MoveTaskOrder) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMTOServiceItemOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteMTOServiceItemBadRequestCode is the HTTP code returned for type DeleteMTOServiceItemBadRequest
const DeleteMTOServiceItemBadRequestCode int = 400

/*DeleteMTOServiceItemBadRequest The request payload is invalid

swagger:response deleteMTOServiceItemBadRequest
*/
type DeleteMTOServiceItemBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteMTOServiceItemBadRequest creates DeleteMTOServiceItemBadRequest with default headers values
func NewDeleteMTOServiceItemBadRequest() *DeleteMTOServiceItemBadRequest {

	return &DeleteMTOServiceItemBadRequest{}
}

// WithPayload adds the payload to the delete m t o service item bad request response
func (o *DeleteMTOServiceItemBadRequest) WithPayload(payload interface{}) *DeleteMTOServiceItemBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete m t o service item bad request response
func (o *DeleteMTOServiceItemBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMTOServiceItemBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteMTOServiceItemUnauthorizedCode is the HTTP code returned for type DeleteMTOServiceItemUnauthorized
const DeleteMTOServiceItemUnauthorizedCode int = 401

/*DeleteMTOServiceItemUnauthorized The request was denied

swagger:response deleteMTOServiceItemUnauthorized
*/
type DeleteMTOServiceItemUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteMTOServiceItemUnauthorized creates DeleteMTOServiceItemUnauthorized with default headers values
func NewDeleteMTOServiceItemUnauthorized() *DeleteMTOServiceItemUnauthorized {

	return &DeleteMTOServiceItemUnauthorized{}
}

// WithPayload adds the payload to the delete m t o service item unauthorized response
func (o *DeleteMTOServiceItemUnauthorized) WithPayload(payload interface{}) *DeleteMTOServiceItemUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete m t o service item unauthorized response
func (o *DeleteMTOServiceItemUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMTOServiceItemUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteMTOServiceItemForbiddenCode is the HTTP code returned for type DeleteMTOServiceItemForbidden
const DeleteMTOServiceItemForbiddenCode int = 403

/*DeleteMTOServiceItemForbidden The request was denied

swagger:response deleteMTOServiceItemForbidden
*/
type DeleteMTOServiceItemForbidden struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteMTOServiceItemForbidden creates DeleteMTOServiceItemForbidden with default headers values
func NewDeleteMTOServiceItemForbidden() *DeleteMTOServiceItemForbidden {

	return &DeleteMTOServiceItemForbidden{}
}

// WithPayload adds the payload to the delete m t o service item forbidden response
func (o *DeleteMTOServiceItemForbidden) WithPayload(payload interface{}) *DeleteMTOServiceItemForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete m t o service item forbidden response
func (o *DeleteMTOServiceItemForbidden) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMTOServiceItemForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteMTOServiceItemNotFoundCode is the HTTP code returned for type DeleteMTOServiceItemNotFound
const DeleteMTOServiceItemNotFoundCode int = 404

/*DeleteMTOServiceItemNotFound The requested resource wasn't found

swagger:response deleteMTOServiceItemNotFound
*/
type DeleteMTOServiceItemNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteMTOServiceItemNotFound creates DeleteMTOServiceItemNotFound with default headers values
func NewDeleteMTOServiceItemNotFound() *DeleteMTOServiceItemNotFound {

	return &DeleteMTOServiceItemNotFound{}
}

// WithPayload adds the payload to the delete m t o service item not found response
func (o *DeleteMTOServiceItemNotFound) WithPayload(payload interface{}) *DeleteMTOServiceItemNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete m t o service item not found response
func (o *DeleteMTOServiceItemNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMTOServiceItemNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteMTOServiceItemInternalServerErrorCode is the HTTP code returned for type DeleteMTOServiceItemInternalServerError
const DeleteMTOServiceItemInternalServerErrorCode int = 500

/*DeleteMTOServiceItemInternalServerError A server error occurred

swagger:response deleteMTOServiceItemInternalServerError
*/
type DeleteMTOServiceItemInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewDeleteMTOServiceItemInternalServerError creates DeleteMTOServiceItemInternalServerError with default headers values
func NewDeleteMTOServiceItemInternalServerError() *DeleteMTOServiceItemInternalServerError {

	return &DeleteMTOServiceItemInternalServerError{}
}

// WithPayload adds the payload to the delete m t o service item internal server error response
func (o *DeleteMTOServiceItemInternalServerError) WithPayload(payload interface{}) *DeleteMTOServiceItemInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete m t o service item internal server error response
func (o *DeleteMTOServiceItemInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteMTOServiceItemInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}