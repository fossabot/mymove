// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdateMoveTaskOrderPostCounselingInformationParams creates a new UpdateMoveTaskOrderPostCounselingInformationParams object
// with the default values initialized.
func NewUpdateMoveTaskOrderPostCounselingInformationParams() *UpdateMoveTaskOrderPostCounselingInformationParams {
	var ()
	return &UpdateMoveTaskOrderPostCounselingInformationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMoveTaskOrderPostCounselingInformationParamsWithTimeout creates a new UpdateMoveTaskOrderPostCounselingInformationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateMoveTaskOrderPostCounselingInformationParamsWithTimeout(timeout time.Duration) *UpdateMoveTaskOrderPostCounselingInformationParams {
	var ()
	return &UpdateMoveTaskOrderPostCounselingInformationParams{

		timeout: timeout,
	}
}

// NewUpdateMoveTaskOrderPostCounselingInformationParamsWithContext creates a new UpdateMoveTaskOrderPostCounselingInformationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateMoveTaskOrderPostCounselingInformationParamsWithContext(ctx context.Context) *UpdateMoveTaskOrderPostCounselingInformationParams {
	var ()
	return &UpdateMoveTaskOrderPostCounselingInformationParams{

		Context: ctx,
	}
}

// NewUpdateMoveTaskOrderPostCounselingInformationParamsWithHTTPClient creates a new UpdateMoveTaskOrderPostCounselingInformationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateMoveTaskOrderPostCounselingInformationParamsWithHTTPClient(client *http.Client) *UpdateMoveTaskOrderPostCounselingInformationParams {
	var ()
	return &UpdateMoveTaskOrderPostCounselingInformationParams{
		HTTPClient: client,
	}
}

/*UpdateMoveTaskOrderPostCounselingInformationParams contains all the parameters to send to the API endpoint
for the update move task order post counseling information operation typically these are written to a http.Request
*/
type UpdateMoveTaskOrderPostCounselingInformationParams struct {

	/*Body*/
	Body UpdateMoveTaskOrderPostCounselingInformationBody
	/*MoveTaskOrderID
	  ID of move order to use

	*/
	MoveTaskOrderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update move task order post counseling information params
func (o *UpdateMoveTaskOrderPostCounselingInformationParams) WithTimeout(timeout time.Duration) *UpdateMoveTaskOrderPostCounselingInformationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update move task order post counseling information params
func (o *UpdateMoveTaskOrderPostCounselingInformationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update move task order post counseling information params
func (o *UpdateMoveTaskOrderPostCounselingInformationParams) WithContext(ctx context.Context) *UpdateMoveTaskOrderPostCounselingInformationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update move task order post counseling information params
func (o *UpdateMoveTaskOrderPostCounselingInformationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update move task order post counseling information params
func (o *UpdateMoveTaskOrderPostCounselingInformationParams) WithHTTPClient(client *http.Client) *UpdateMoveTaskOrderPostCounselingInformationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update move task order post counseling information params
func (o *UpdateMoveTaskOrderPostCounselingInformationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update move task order post counseling information params
func (o *UpdateMoveTaskOrderPostCounselingInformationParams) WithBody(body UpdateMoveTaskOrderPostCounselingInformationBody) *UpdateMoveTaskOrderPostCounselingInformationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update move task order post counseling information params
func (o *UpdateMoveTaskOrderPostCounselingInformationParams) SetBody(body UpdateMoveTaskOrderPostCounselingInformationBody) {
	o.Body = body
}

// WithMoveTaskOrderID adds the moveTaskOrderID to the update move task order post counseling information params
func (o *UpdateMoveTaskOrderPostCounselingInformationParams) WithMoveTaskOrderID(moveTaskOrderID string) *UpdateMoveTaskOrderPostCounselingInformationParams {
	o.SetMoveTaskOrderID(moveTaskOrderID)
	return o
}

// SetMoveTaskOrderID adds the moveTaskOrderId to the update move task order post counseling information params
func (o *UpdateMoveTaskOrderPostCounselingInformationParams) SetMoveTaskOrderID(moveTaskOrderID string) {
	o.MoveTaskOrderID = moveTaskOrderID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMoveTaskOrderPostCounselingInformationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param moveTaskOrderID
	if err := r.SetPathParam("moveTaskOrderID", o.MoveTaskOrderID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
