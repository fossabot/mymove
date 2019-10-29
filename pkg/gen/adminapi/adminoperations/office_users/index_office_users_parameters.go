// Code generated by go-swagger; DO NOT EDIT.

package office_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewIndexOfficeUsersParams creates a new IndexOfficeUsersParams object
// no default values defined in spec.
func NewIndexOfficeUsersParams() IndexOfficeUsersParams {

	return IndexOfficeUsersParams{}
}

// IndexOfficeUsersParams contains all the bound params for the index office users operation
// typically these are obtained from a http.Request
//
// swagger:parameters indexOfficeUsers
type IndexOfficeUsersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	Filter []string
	/*
	  In: query
	*/
	Order *bool
	/*
	  In: query
	*/
	Page *int64
	/*
	  In: query
	*/
	PerPage *int64
	/*
	  In: query
	*/
	Sort *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewIndexOfficeUsersParams() beforehand.
func (o *IndexOfficeUsersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFilter, qhkFilter, _ := qs.GetOK("filter")
	if err := o.bindFilter(qFilter, qhkFilter, route.Formats); err != nil {
		res = append(res, err)
	}

	qOrder, qhkOrder, _ := qs.GetOK("order")
	if err := o.bindOrder(qOrder, qhkOrder, route.Formats); err != nil {
		res = append(res, err)
	}

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qPerPage, qhkPerPage, _ := qs.GetOK("perPage")
	if err := o.bindPerPage(qPerPage, qhkPerPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qSort, qhkSort, _ := qs.GetOK("sort")
	if err := o.bindSort(qSort, qhkSort, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFilter binds and validates array parameter Filter from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *IndexOfficeUsersParams) bindFilter(rawData []string, hasKey bool, formats strfmt.Registry) error {

	var qvFilter string
	if len(rawData) > 0 {
		qvFilter = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	filterIC := swag.SplitByFormat(qvFilter, "")
	if len(filterIC) == 0 {
		return nil
	}

	var filterIR []string
	for _, filterIV := range filterIC {
		filterI := filterIV

		filterIR = append(filterIR, filterI)
	}

	o.Filter = filterIR

	return nil
}

// bindOrder binds and validates parameter Order from query.
func (o *IndexOfficeUsersParams) bindOrder(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("order", "query", "bool", raw)
	}
	o.Order = &value

	return nil
}

// bindPage binds and validates parameter Page from query.
func (o *IndexOfficeUsersParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("page", "query", "int64", raw)
	}
	o.Page = &value

	return nil
}

// bindPerPage binds and validates parameter PerPage from query.
func (o *IndexOfficeUsersParams) bindPerPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("perPage", "query", "int64", raw)
	}
	o.PerPage = &value

	return nil
}

// bindSort binds and validates parameter Sort from query.
func (o *IndexOfficeUsersParams) bindSort(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Sort = &raw

	return nil
}
