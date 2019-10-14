// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ServiceMemberRank Rank
// swagger:model ServiceMemberRank
type ServiceMemberRank string

const (

	// ServiceMemberRankE1 captures enum value "E_1"
	ServiceMemberRankE1 ServiceMemberRank = "E_1"

	// ServiceMemberRankE2 captures enum value "E_2"
	ServiceMemberRankE2 ServiceMemberRank = "E_2"

	// ServiceMemberRankE3 captures enum value "E_3"
	ServiceMemberRankE3 ServiceMemberRank = "E_3"

	// ServiceMemberRankE4 captures enum value "E_4"
	ServiceMemberRankE4 ServiceMemberRank = "E_4"

	// ServiceMemberRankE5 captures enum value "E_5"
	ServiceMemberRankE5 ServiceMemberRank = "E_5"

	// ServiceMemberRankE6 captures enum value "E_6"
	ServiceMemberRankE6 ServiceMemberRank = "E_6"

	// ServiceMemberRankE7 captures enum value "E_7"
	ServiceMemberRankE7 ServiceMemberRank = "E_7"

	// ServiceMemberRankE8 captures enum value "E_8"
	ServiceMemberRankE8 ServiceMemberRank = "E_8"

	// ServiceMemberRankE9 captures enum value "E_9"
	ServiceMemberRankE9 ServiceMemberRank = "E_9"

	// ServiceMemberRankO1ACADEMYGRADUATE captures enum value "O_1_ACADEMY_GRADUATE"
	ServiceMemberRankO1ACADEMYGRADUATE ServiceMemberRank = "O_1_ACADEMY_GRADUATE"

	// ServiceMemberRankO2 captures enum value "O_2"
	ServiceMemberRankO2 ServiceMemberRank = "O_2"

	// ServiceMemberRankO3 captures enum value "O_3"
	ServiceMemberRankO3 ServiceMemberRank = "O_3"

	// ServiceMemberRankO4 captures enum value "O_4"
	ServiceMemberRankO4 ServiceMemberRank = "O_4"

	// ServiceMemberRankO5 captures enum value "O_5"
	ServiceMemberRankO5 ServiceMemberRank = "O_5"

	// ServiceMemberRankO6 captures enum value "O_6"
	ServiceMemberRankO6 ServiceMemberRank = "O_6"

	// ServiceMemberRankO7 captures enum value "O_7"
	ServiceMemberRankO7 ServiceMemberRank = "O_7"

	// ServiceMemberRankO8 captures enum value "O_8"
	ServiceMemberRankO8 ServiceMemberRank = "O_8"

	// ServiceMemberRankO9 captures enum value "O_9"
	ServiceMemberRankO9 ServiceMemberRank = "O_9"

	// ServiceMemberRankO10 captures enum value "O_10"
	ServiceMemberRankO10 ServiceMemberRank = "O_10"

	// ServiceMemberRankW1 captures enum value "W_1"
	ServiceMemberRankW1 ServiceMemberRank = "W_1"

	// ServiceMemberRankW2 captures enum value "W_2"
	ServiceMemberRankW2 ServiceMemberRank = "W_2"

	// ServiceMemberRankW3 captures enum value "W_3"
	ServiceMemberRankW3 ServiceMemberRank = "W_3"

	// ServiceMemberRankW4 captures enum value "W_4"
	ServiceMemberRankW4 ServiceMemberRank = "W_4"

	// ServiceMemberRankW5 captures enum value "W_5"
	ServiceMemberRankW5 ServiceMemberRank = "W_5"

	// ServiceMemberRankAVIATIONCADET captures enum value "AVIATION_CADET"
	ServiceMemberRankAVIATIONCADET ServiceMemberRank = "AVIATION_CADET"

	// ServiceMemberRankCIVILIANEMPLOYEE captures enum value "CIVILIAN_EMPLOYEE"
	ServiceMemberRankCIVILIANEMPLOYEE ServiceMemberRank = "CIVILIAN_EMPLOYEE"

	// ServiceMemberRankACADEMYCADET captures enum value "ACADEMY_CADET"
	ServiceMemberRankACADEMYCADET ServiceMemberRank = "ACADEMY_CADET"

	// ServiceMemberRankMIDSHIPMAN captures enum value "MIDSHIPMAN"
	ServiceMemberRankMIDSHIPMAN ServiceMemberRank = "MIDSHIPMAN"
)

// for schema
var serviceMemberRankEnum []interface{}

func init() {
	var res []ServiceMemberRank
	if err := json.Unmarshal([]byte(`["E_1","E_2","E_3","E_4","E_5","E_6","E_7","E_8","E_9","O_1_ACADEMY_GRADUATE","O_2","O_3","O_4","O_5","O_6","O_7","O_8","O_9","O_10","W_1","W_2","W_3","W_4","W_5","AVIATION_CADET","CIVILIAN_EMPLOYEE","ACADEMY_CADET","MIDSHIPMAN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceMemberRankEnum = append(serviceMemberRankEnum, v)
	}
}

func (m ServiceMemberRank) validateServiceMemberRankEnum(path, location string, value ServiceMemberRank) error {
	if err := validate.Enum(path, location, value, serviceMemberRankEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this service member rank
func (m ServiceMemberRank) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateServiceMemberRankEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}