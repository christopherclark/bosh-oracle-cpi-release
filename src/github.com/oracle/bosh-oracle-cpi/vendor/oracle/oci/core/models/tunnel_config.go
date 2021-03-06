// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TunnelConfig Specific connection details for an IPSec tunnel.
// swagger:model TunnelConfig
type TunnelConfig struct {

	// The IP address of Oracle's VPN headend.
	//
	// Example: `129.146.17.50`
	//
	// Required: true
	// Max Length: 46
	// Min Length: 1
	IPAddress *string `json:"ipAddress"`

	// The shared secret of the IPSec tunnel.
	//
	// Example: `vFG2IF6TWq4UToUiLSRDoJEUs6j1c.p8G.dVQxiMfMO0yXMLi.lZTbYIWhGu4V8o`
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	SharedSecret *string `json:"sharedSecret"`

	// The date and time the IPSec connection was created, in the format defined by RFC3339.
	//
	// Example: `2016-08-25T21:10:29.600Z`
	//
	TimeCreated strfmt.DateTime `json:"timeCreated,omitempty"`
}

// Validate validates this tunnel config
func (m *TunnelConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSharedSecret(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TunnelConfig) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ipAddress", "body", m.IPAddress); err != nil {
		return err
	}

	if err := validate.MinLength("ipAddress", "body", string(*m.IPAddress), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("ipAddress", "body", string(*m.IPAddress), 46); err != nil {
		return err
	}

	return nil
}

func (m *TunnelConfig) validateSharedSecret(formats strfmt.Registry) error {

	if err := validate.Required("sharedSecret", "body", m.SharedSecret); err != nil {
		return err
	}

	if err := validate.MinLength("sharedSecret", "body", string(*m.SharedSecret), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("sharedSecret", "body", string(*m.SharedSecret), 255); err != nil {
		return err
	}

	return nil
}

func (m *TunnelConfig) validateTimeCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("timeCreated", "body", "date-time", m.TimeCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TunnelConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TunnelConfig) UnmarshalBinary(b []byte) error {
	var res TunnelConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
