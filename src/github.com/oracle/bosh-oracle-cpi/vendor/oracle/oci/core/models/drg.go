// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Drg A Dynamic Routing Gateway (DRG), which is a virtual router that provides a path for private
// network traffic between your VCN and your existing network. You use it with other Networking
// Service components to create an IPSec VPN or a connection that uses
// Oracle Cloud Infrastructure FastConnect. For more information, see
// [Overview of the Networking Service](/Content/Network/Concepts/overview.htm).
//
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).
//
// swagger:model Drg
type Drg struct {

	// The OCID of the compartment containing the DRG.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	CompartmentID *string `json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	//
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// The DRG's Oracle ID (OCID).
	// Required: true
	// Max Length: 255
	// Min Length: 1
	ID *string `json:"id"`

	// The DRG's current state.
	// Required: true
	LifecycleState *string `json:"lifecycleState"`

	// The date and time the DRG was created, in the format defined by RFC3339.
	//
	// Example: `2016-08-25T21:10:29.600Z`
	//
	TimeCreated strfmt.DateTime `json:"timeCreated,omitempty"`
}

// Validate validates this drg
func (m *Drg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLifecycleState(formats); err != nil {
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

func (m *Drg) validateCompartmentID(formats strfmt.Registry) error {

	if err := validate.Required("compartmentId", "body", m.CompartmentID); err != nil {
		return err
	}

	if err := validate.MinLength("compartmentId", "body", string(*m.CompartmentID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("compartmentId", "body", string(*m.CompartmentID), 255); err != nil {
		return err
	}

	return nil
}

func (m *Drg) validateDisplayName(formats strfmt.Registry) error {

	if swag.IsZero(m.DisplayName) { // not required
		return nil
	}

	if err := validate.MinLength("displayName", "body", string(m.DisplayName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", string(m.DisplayName), 255); err != nil {
		return err
	}

	return nil
}

func (m *Drg) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID), 255); err != nil {
		return err
	}

	return nil
}

var drgTypeLifecycleStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PROVISIONING","AVAILABLE","TERMINATING","TERMINATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		drgTypeLifecycleStatePropEnum = append(drgTypeLifecycleStatePropEnum, v)
	}
}

const (

	// DrgLifecycleStatePROVISIONING captures enum value "PROVISIONING"
	DrgLifecycleStatePROVISIONING string = "PROVISIONING"

	// DrgLifecycleStateAVAILABLE captures enum value "AVAILABLE"
	DrgLifecycleStateAVAILABLE string = "AVAILABLE"

	// DrgLifecycleStateTERMINATING captures enum value "TERMINATING"
	DrgLifecycleStateTERMINATING string = "TERMINATING"

	// DrgLifecycleStateTERMINATED captures enum value "TERMINATED"
	DrgLifecycleStateTERMINATED string = "TERMINATED"
)

// prop value enum
func (m *Drg) validateLifecycleStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, drgTypeLifecycleStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Drg) validateLifecycleState(formats strfmt.Registry) error {

	if err := validate.Required("lifecycleState", "body", m.LifecycleState); err != nil {
		return err
	}

	// value enum
	if err := m.validateLifecycleStateEnum("lifecycleState", "body", *m.LifecycleState); err != nil {
		return err
	}

	return nil
}

func (m *Drg) validateTimeCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("timeCreated", "body", "date-time", m.TimeCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Drg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Drg) UnmarshalBinary(b []byte) error {
	var res Drg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
