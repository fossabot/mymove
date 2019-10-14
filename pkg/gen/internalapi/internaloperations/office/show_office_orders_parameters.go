// Code generated by go-swagger; DO NOT EDIT.

package office

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewShowOfficeOrdersParams creates a new ShowOfficeOrdersParams object
// no default values defined in spec.
func NewShowOfficeOrdersParams() ShowOfficeOrdersParams {

	return ShowOfficeOrdersParams{}
}

// ShowOfficeOrdersParams contains all the bound params for the show office orders operation
// typically these are obtained from a http.Request
//
// swagger:parameters showOfficeOrders
type ShowOfficeOrdersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*UUID of the move
	  Required: true
	  In: path
	*/
	MoveID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewShowOfficeOrdersParams() beforehand.
func (o *ShowOfficeOrdersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rMoveID, rhkMoveID, _ := route.Params.GetOK("moveId")
	if err := o.bindMoveID(rMoveID, rhkMoveID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMoveID binds and validates parameter MoveID from path.
func (o *ShowOfficeOrdersParams) bindMoveID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("moveId", "path", "strfmt.UUID", raw)
	}
	o.MoveID = *(value.(*strfmt.UUID))

	if err := o.validateMoveID(formats); err != nil {
		return err
	}

	return nil
}

// validateMoveID carries on validations for parameter MoveID
func (o *ShowOfficeOrdersParams) validateMoveID(formats strfmt.Registry) error {

	if err := validate.FormatOf("moveId", "path", "uuid", o.MoveID.String(), formats); err != nil {
		return err
	}
	return nil
}