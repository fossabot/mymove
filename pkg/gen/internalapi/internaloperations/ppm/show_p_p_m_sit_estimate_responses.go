// Code generated by go-swagger; DO NOT EDIT.

package ppm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	internalmessages "github.com/transcom/mymove/pkg/gen/internalmessages"
)

// ShowPPMSitEstimateOKCode is the HTTP code returned for type ShowPPMSitEstimateOK
const ShowPPMSitEstimateOKCode int = 200

/*ShowPPMSitEstimateOK show PPM SIT estimate

swagger:response showPPMSitEstimateOK
*/
type ShowPPMSitEstimateOK struct {

	/*
	  In: Body
	*/
	Payload *internalmessages.PPMSitEstimate `json:"body,omitempty"`
}

// NewShowPPMSitEstimateOK creates ShowPPMSitEstimateOK with default headers values
func NewShowPPMSitEstimateOK() *ShowPPMSitEstimateOK {

	return &ShowPPMSitEstimateOK{}
}

// WithPayload adds the payload to the show p p m sit estimate o k response
func (o *ShowPPMSitEstimateOK) WithPayload(payload *internalmessages.PPMSitEstimate) *ShowPPMSitEstimateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the show p p m sit estimate o k response
func (o *ShowPPMSitEstimateOK) SetPayload(payload *internalmessages.PPMSitEstimate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ShowPPMSitEstimateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ShowPPMSitEstimateBadRequestCode is the HTTP code returned for type ShowPPMSitEstimateBadRequest
const ShowPPMSitEstimateBadRequestCode int = 400

/*ShowPPMSitEstimateBadRequest invalid request

swagger:response showPPMSitEstimateBadRequest
*/
type ShowPPMSitEstimateBadRequest struct {
}

// NewShowPPMSitEstimateBadRequest creates ShowPPMSitEstimateBadRequest with default headers values
func NewShowPPMSitEstimateBadRequest() *ShowPPMSitEstimateBadRequest {

	return &ShowPPMSitEstimateBadRequest{}
}

// WriteResponse to the client
func (o *ShowPPMSitEstimateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ShowPPMSitEstimateUnauthorizedCode is the HTTP code returned for type ShowPPMSitEstimateUnauthorized
const ShowPPMSitEstimateUnauthorizedCode int = 401

/*ShowPPMSitEstimateUnauthorized request requires user authentication

swagger:response showPPMSitEstimateUnauthorized
*/
type ShowPPMSitEstimateUnauthorized struct {
}

// NewShowPPMSitEstimateUnauthorized creates ShowPPMSitEstimateUnauthorized with default headers values
func NewShowPPMSitEstimateUnauthorized() *ShowPPMSitEstimateUnauthorized {

	return &ShowPPMSitEstimateUnauthorized{}
}

// WriteResponse to the client
func (o *ShowPPMSitEstimateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ShowPPMSitEstimateForbiddenCode is the HTTP code returned for type ShowPPMSitEstimateForbidden
const ShowPPMSitEstimateForbiddenCode int = 403

/*ShowPPMSitEstimateForbidden user is not authorized

swagger:response showPPMSitEstimateForbidden
*/
type ShowPPMSitEstimateForbidden struct {
}

// NewShowPPMSitEstimateForbidden creates ShowPPMSitEstimateForbidden with default headers values
func NewShowPPMSitEstimateForbidden() *ShowPPMSitEstimateForbidden {

	return &ShowPPMSitEstimateForbidden{}
}

// WriteResponse to the client
func (o *ShowPPMSitEstimateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// ShowPPMSitEstimateConflictCode is the HTTP code returned for type ShowPPMSitEstimateConflict
const ShowPPMSitEstimateConflictCode int = 409

/*ShowPPMSitEstimateConflict distance is less than 50 miles (no short haul moves)

swagger:response showPPMSitEstimateConflict
*/
type ShowPPMSitEstimateConflict struct {
}

// NewShowPPMSitEstimateConflict creates ShowPPMSitEstimateConflict with default headers values
func NewShowPPMSitEstimateConflict() *ShowPPMSitEstimateConflict {

	return &ShowPPMSitEstimateConflict{}
}

// WriteResponse to the client
func (o *ShowPPMSitEstimateConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// ShowPPMSitEstimateInternalServerErrorCode is the HTTP code returned for type ShowPPMSitEstimateInternalServerError
const ShowPPMSitEstimateInternalServerErrorCode int = 500

/*ShowPPMSitEstimateInternalServerError internal server error

swagger:response showPPMSitEstimateInternalServerError
*/
type ShowPPMSitEstimateInternalServerError struct {
}

// NewShowPPMSitEstimateInternalServerError creates ShowPPMSitEstimateInternalServerError with default headers values
func NewShowPPMSitEstimateInternalServerError() *ShowPPMSitEstimateInternalServerError {

	return &ShowPPMSitEstimateInternalServerError{}
}

// WriteResponse to the client
func (o *ShowPPMSitEstimateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
