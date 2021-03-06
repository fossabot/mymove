// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MTOShipment m t o shipment
// swagger:model MTOShipment
type MTOShipment struct {

	// created at
	// Format: datetime
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// customer remarks
	CustomerRemarks *string `json:"customerRemarks,omitempty"`

	// destination address
	DestinationAddress *Address `json:"destinationAddress,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// move task order ID
	// Format: uuid
	MoveTaskOrderID strfmt.UUID `json:"moveTaskOrderID,omitempty"`

	// pickup address
	PickupAddress *Address `json:"pickupAddress,omitempty"`

	// rejection reason
	RejectionReason *string `json:"rejectionReason,omitempty"`

	// requested pickup date
	// Format: date
	RequestedPickupDate strfmt.Date `json:"requestedPickupDate,omitempty"`

	// scheduled pickup date
	// Format: date
	ScheduledPickupDate strfmt.Date `json:"scheduledPickupDate,omitempty"`

	// secondary delivery address
	SecondaryDeliveryAddress *Address `json:"secondaryDeliveryAddress,omitempty"`

	// secondary pickup address
	SecondaryPickupAddress *Address `json:"secondaryPickupAddress,omitempty"`

	// shipment type
	// Enum: [HHG INTERNATIONAL_HHG INTERNATIONAL_UB]
	ShipmentType interface{} `json:"shipmentType,omitempty"`

	// status
	// Enum: [APPROVED SUBMITTED REJECTED]
	Status string `json:"status,omitempty"`

	// updated at
	// Format: datetime
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this m t o shipment
func (m *MTOShipment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveTaskOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePickupAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedPickupDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledPickupDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryDeliveryAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryPickupAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MTOShipment) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "datetime", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateDestinationAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.DestinationAddress) { // not required
		return nil
	}

	if m.DestinationAddress != nil {
		if err := m.DestinationAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destinationAddress")
			}
			return err
		}
	}

	return nil
}

func (m *MTOShipment) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateMoveTaskOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveTaskOrderID) { // not required
		return nil
	}

	if err := validate.FormatOf("moveTaskOrderID", "body", "uuid", m.MoveTaskOrderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validatePickupAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.PickupAddress) { // not required
		return nil
	}

	if m.PickupAddress != nil {
		if err := m.PickupAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pickupAddress")
			}
			return err
		}
	}

	return nil
}

func (m *MTOShipment) validateRequestedPickupDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedPickupDate) { // not required
		return nil
	}

	if err := validate.FormatOf("requestedPickupDate", "body", "date", m.RequestedPickupDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateScheduledPickupDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheduledPickupDate) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduledPickupDate", "body", "date", m.ScheduledPickupDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateSecondaryDeliveryAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.SecondaryDeliveryAddress) { // not required
		return nil
	}

	if m.SecondaryDeliveryAddress != nil {
		if err := m.SecondaryDeliveryAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondaryDeliveryAddress")
			}
			return err
		}
	}

	return nil
}

func (m *MTOShipment) validateSecondaryPickupAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.SecondaryPickupAddress) { // not required
		return nil
	}

	if m.SecondaryPickupAddress != nil {
		if err := m.SecondaryPickupAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondaryPickupAddress")
			}
			return err
		}
	}

	return nil
}

var mTOShipmentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["APPROVED","SUBMITTED","REJECTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mTOShipmentTypeStatusPropEnum = append(mTOShipmentTypeStatusPropEnum, v)
	}
}

const (

	// MTOShipmentStatusAPPROVED captures enum value "APPROVED"
	MTOShipmentStatusAPPROVED string = "APPROVED"

	// MTOShipmentStatusSUBMITTED captures enum value "SUBMITTED"
	MTOShipmentStatusSUBMITTED string = "SUBMITTED"

	// MTOShipmentStatusREJECTED captures enum value "REJECTED"
	MTOShipmentStatusREJECTED string = "REJECTED"
)

// prop value enum
func (m *MTOShipment) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, mTOShipmentTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MTOShipment) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *MTOShipment) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "datetime", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MTOShipment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MTOShipment) UnmarshalBinary(b []byte) error {
	var res MTOShipment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
