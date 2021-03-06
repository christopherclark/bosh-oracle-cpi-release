// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RouteTable A collection of `RouteRule` objects, which are used to route packets
// based on destination IP to a particular network entity. For more information, see
// [Overview of the Networking Service](/Content/Network/Concepts/overview.htm).
//
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).
//
// swagger:model RouteTable
type RouteTable struct {

	// The OCID of the compartment containing the route table.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	CompartmentID *string `json:"compartmentId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see [Resource Tags](/Content/General/Concepts/resourcetags.htm).
	//
	// Example: `{"Operations": {"CostCenter": "42"}}`
	//
	DefinedTags map[string]map[string]interface{} `json:"definedTags,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	//
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// [Resource Tags](/Content/General/Concepts/resourcetags.htm).
	//
	// Example: `{"Department": "Finance"}`
	//
	FreeformTags map[string]string `json:"freeformTags,omitempty"`

	// The route table's Oracle ID (OCID).
	// Required: true
	// Max Length: 255
	// Min Length: 1
	ID *string `json:"id"`

	// The route table's current state.
	// Required: true
	LifecycleState *string `json:"lifecycleState"`

	// The collection of rules for routing destination IPs to network devices.
	// Required: true
	RouteRules []*RouteRule `json:"routeRules"`

	// The date and time the route table was created, in the format defined by RFC3339.
	//
	// Example: `2016-08-25T21:10:29.600Z`
	//
	TimeCreated strfmt.DateTime `json:"timeCreated,omitempty"`

	// The OCID of the VCN the route table list belongs to.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	VcnID *string `json:"vcnId"`
}

// Validate validates this route table
func (m *RouteTable) Validate(formats strfmt.Registry) error {
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

	if err := m.validateRouteRules(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVcnID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RouteTable) validateCompartmentID(formats strfmt.Registry) error {

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

func (m *RouteTable) validateDisplayName(formats strfmt.Registry) error {

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

func (m *RouteTable) validateID(formats strfmt.Registry) error {

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

var routeTableTypeLifecycleStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PROVISIONING","AVAILABLE","TERMINATING","TERMINATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeTableTypeLifecycleStatePropEnum = append(routeTableTypeLifecycleStatePropEnum, v)
	}
}

const (

	// RouteTableLifecycleStatePROVISIONING captures enum value "PROVISIONING"
	RouteTableLifecycleStatePROVISIONING string = "PROVISIONING"

	// RouteTableLifecycleStateAVAILABLE captures enum value "AVAILABLE"
	RouteTableLifecycleStateAVAILABLE string = "AVAILABLE"

	// RouteTableLifecycleStateTERMINATING captures enum value "TERMINATING"
	RouteTableLifecycleStateTERMINATING string = "TERMINATING"

	// RouteTableLifecycleStateTERMINATED captures enum value "TERMINATED"
	RouteTableLifecycleStateTERMINATED string = "TERMINATED"
)

// prop value enum
func (m *RouteTable) validateLifecycleStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, routeTableTypeLifecycleStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RouteTable) validateLifecycleState(formats strfmt.Registry) error {

	if err := validate.Required("lifecycleState", "body", m.LifecycleState); err != nil {
		return err
	}

	// value enum
	if err := m.validateLifecycleStateEnum("lifecycleState", "body", *m.LifecycleState); err != nil {
		return err
	}

	return nil
}

func (m *RouteTable) validateRouteRules(formats strfmt.Registry) error {

	if err := validate.Required("routeRules", "body", m.RouteRules); err != nil {
		return err
	}

	for i := 0; i < len(m.RouteRules); i++ {

		if swag.IsZero(m.RouteRules[i]) { // not required
			continue
		}

		if m.RouteRules[i] != nil {

			if err := m.RouteRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routeRules" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *RouteTable) validateTimeCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("timeCreated", "body", "date-time", m.TimeCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RouteTable) validateVcnID(formats strfmt.Registry) error {

	if err := validate.Required("vcnId", "body", m.VcnID); err != nil {
		return err
	}

	if err := validate.MinLength("vcnId", "body", string(*m.VcnID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("vcnId", "body", string(*m.VcnID), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RouteTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouteTable) UnmarshalBinary(b []byte) error {
	var res RouteTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
