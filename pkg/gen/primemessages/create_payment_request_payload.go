// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreatePaymentRequestPayload create payment request payload
// swagger:model CreatePaymentRequestPayload
type CreatePaymentRequestPayload struct {

	// is final
	IsFinal *bool `json:"isFinal,omitempty"`

	// move task order ID
	// Format: uuid
	MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

	// proof of service package
	ProofOfServicePackage *ProofOfServicePackage `json:"proofOfServicePackage,omitempty"`
}

// Validate validates this create payment request payload
func (m *CreatePaymentRequestPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMoveTaskOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProofOfServicePackage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatePaymentRequestPayload) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveTaskOrderID) { // not required
		return nil
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreatePaymentRequestPayload) validateProofOfServicePackage(formats strfmt.Registry) error {

	if swag.IsZero(m.ProofOfServicePackage) { // not required
		return nil
	}

	if m.ProofOfServicePackage != nil {
		if err := m.ProofOfServicePackage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proofOfServicePackage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreatePaymentRequestPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePaymentRequestPayload) UnmarshalBinary(b []byte) error {
	var res CreatePaymentRequestPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
