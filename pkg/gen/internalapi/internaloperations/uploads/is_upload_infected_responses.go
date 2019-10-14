// Code generated by go-swagger; DO NOT EDIT.

package uploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// IsUploadInfectedOKCode is the HTTP code returned for type IsUploadInfectedOK
const IsUploadInfectedOKCode int = 200

/*IsUploadInfectedOK Upload is infected

swagger:response isUploadInfectedOK
*/
type IsUploadInfectedOK struct {

	/*
	  In: Body
	*/
	Payload bool `json:"body,omitempty"`
}

// NewIsUploadInfectedOK creates IsUploadInfectedOK with default headers values
func NewIsUploadInfectedOK() *IsUploadInfectedOK {

	return &IsUploadInfectedOK{}
}

// WithPayload adds the payload to the is upload infected o k response
func (o *IsUploadInfectedOK) WithPayload(payload bool) *IsUploadInfectedOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the is upload infected o k response
func (o *IsUploadInfectedOK) SetPayload(payload bool) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *IsUploadInfectedOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// IsUploadInfectedBadRequestCode is the HTTP code returned for type IsUploadInfectedBadRequest
const IsUploadInfectedBadRequestCode int = 400

/*IsUploadInfectedBadRequest invalid request

swagger:response isUploadInfectedBadRequest
*/
type IsUploadInfectedBadRequest struct {
}

// NewIsUploadInfectedBadRequest creates IsUploadInfectedBadRequest with default headers values
func NewIsUploadInfectedBadRequest() *IsUploadInfectedBadRequest {

	return &IsUploadInfectedBadRequest{}
}

// WriteResponse to the client
func (o *IsUploadInfectedBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// IsUploadInfectedInternalServerErrorCode is the HTTP code returned for type IsUploadInfectedInternalServerError
const IsUploadInfectedInternalServerErrorCode int = 500

/*IsUploadInfectedInternalServerError server error

swagger:response isUploadInfectedInternalServerError
*/
type IsUploadInfectedInternalServerError struct {
}

// NewIsUploadInfectedInternalServerError creates IsUploadInfectedInternalServerError with default headers values
func NewIsUploadInfectedInternalServerError() *IsUploadInfectedInternalServerError {

	return &IsUploadInfectedInternalServerError{}
}

// WriteResponse to the client
func (o *IsUploadInfectedInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
