// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Volume A detachable block volume device that allows you to dynamically expand
// the storage capacity of an instance. For more information, see
// [Overview of Cloud Volume Storage](/Content/Block/Concepts/overview.htm).
//
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).
//
// swagger:model Volume
type Volume struct {

	// The Availability Domain of the volume.
	//
	// Example: `Uocm:PHX-AD-1`
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	AvailabilityDomain *string `json:"availabilityDomain"`

	// The OCID of the compartment that contains the volume.
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
	// Required: true
	// Max Length: 255
	// Min Length: 1
	DisplayName *string `json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// [Resource Tags](/Content/General/Concepts/resourcetags.htm).
	//
	// Example: `{"Department": "Finance"}`
	//
	FreeformTags map[string]string `json:"freeformTags,omitempty"`

	// The OCID of the volume.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	ID *string `json:"id"`

	// Specifies whether the cloned volume's data has finished copying from the source volume or backup.
	IsHydrated bool `json:"isHydrated,omitempty"`

	// The current state of a volume.
	// Required: true
	LifecycleState *string `json:"lifecycleState"`

	// The size of the volume in GBs.
	SizeInGBs int64 `json:"sizeInGBs,omitempty"`

	// The size of the volume in MBs. This field is deprecated. Use sizeInGBs instead.
	// Required: true
	SizeInMBs *int64 `json:"sizeInMBs"`

	sourceDetailsField VolumeSourceDetails

	// The date and time the volume was created. Format defined by RFC3339.
	// Required: true
	TimeCreated *strfmt.DateTime `json:"timeCreated"`
}

func (m *Volume) SourceDetails() VolumeSourceDetails {
	return m.sourceDetailsField
}
func (m *Volume) SetSourceDetails(val VolumeSourceDetails) {
	m.sourceDetailsField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *Volume) UnmarshalJSON(raw []byte) error {
	var data struct {
		AvailabilityDomain *string `json:"availabilityDomain"`

		CompartmentID *string `json:"compartmentId"`

		DefinedTags map[string]map[string]interface{} `json:"definedTags,omitempty"`

		DisplayName *string `json:"displayName"`

		FreeformTags map[string]string `json:"freeformTags,omitempty"`

		ID *string `json:"id"`

		IsHydrated bool `json:"isHydrated,omitempty"`

		LifecycleState *string `json:"lifecycleState"`

		SizeInGBs int64 `json:"sizeInGBs,omitempty"`

		SizeInMBs *int64 `json:"sizeInMBs"`

		SourceDetails json.RawMessage `json:"sourceDetails,omitempty"`

		TimeCreated *strfmt.DateTime `json:"timeCreated"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	sourceDetails, err := UnmarshalVolumeSourceDetails(bytes.NewBuffer(data.SourceDetails), runtime.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	var result Volume

	// availabilityDomain
	result.AvailabilityDomain = data.AvailabilityDomain

	// compartmentId
	result.CompartmentID = data.CompartmentID

	// definedTags
	result.DefinedTags = data.DefinedTags

	// displayName
	result.DisplayName = data.DisplayName

	// freeformTags
	result.FreeformTags = data.FreeformTags

	// id
	result.ID = data.ID

	// isHydrated
	result.IsHydrated = data.IsHydrated

	// lifecycleState
	result.LifecycleState = data.LifecycleState

	// sizeInGBs
	result.SizeInGBs = data.SizeInGBs

	// sizeInMBs
	result.SizeInMBs = data.SizeInMBs

	// sourceDetails
	result.sourceDetailsField = sourceDetails

	// timeCreated
	result.TimeCreated = data.TimeCreated

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m Volume) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		AvailabilityDomain *string `json:"availabilityDomain"`

		CompartmentID *string `json:"compartmentId"`

		DefinedTags map[string]map[string]interface{} `json:"definedTags,omitempty"`

		DisplayName *string `json:"displayName"`

		FreeformTags map[string]string `json:"freeformTags,omitempty"`

		ID *string `json:"id"`

		IsHydrated bool `json:"isHydrated,omitempty"`

		LifecycleState *string `json:"lifecycleState"`

		SizeInGBs int64 `json:"sizeInGBs,omitempty"`

		SizeInMBs *int64 `json:"sizeInMBs"`

		TimeCreated *strfmt.DateTime `json:"timeCreated"`
	}{

		AvailabilityDomain: m.AvailabilityDomain,

		CompartmentID: m.CompartmentID,

		DefinedTags: m.DefinedTags,

		DisplayName: m.DisplayName,

		FreeformTags: m.FreeformTags,

		ID: m.ID,

		IsHydrated: m.IsHydrated,

		LifecycleState: m.LifecycleState,

		SizeInGBs: m.SizeInGBs,

		SizeInMBs: m.SizeInMBs,

		TimeCreated: m.TimeCreated,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		SourceDetails VolumeSourceDetails `json:"sourceDetails,omitempty"`
	}{

		SourceDetails: m.sourceDetailsField,
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this volume
func (m *Volume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailabilityDomain(formats); err != nil {
		// prop
		res = append(res, err)
	}

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

	if err := m.validateSizeInMBs(formats); err != nil {
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

func (m *Volume) validateAvailabilityDomain(formats strfmt.Registry) error {

	if err := validate.Required("availabilityDomain", "body", m.AvailabilityDomain); err != nil {
		return err
	}

	if err := validate.MinLength("availabilityDomain", "body", string(*m.AvailabilityDomain), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("availabilityDomain", "body", string(*m.AvailabilityDomain), 255); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateCompartmentID(formats strfmt.Registry) error {

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

func (m *Volume) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	if err := validate.MinLength("displayName", "body", string(*m.DisplayName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", string(*m.DisplayName), 255); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateID(formats strfmt.Registry) error {

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

var volumeTypeLifecycleStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PROVISIONING","RESTORING","AVAILABLE","TERMINATING","TERMINATED","FAULTY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		volumeTypeLifecycleStatePropEnum = append(volumeTypeLifecycleStatePropEnum, v)
	}
}

const (

	// VolumeLifecycleStatePROVISIONING captures enum value "PROVISIONING"
	VolumeLifecycleStatePROVISIONING string = "PROVISIONING"

	// VolumeLifecycleStateRESTORING captures enum value "RESTORING"
	VolumeLifecycleStateRESTORING string = "RESTORING"

	// VolumeLifecycleStateAVAILABLE captures enum value "AVAILABLE"
	VolumeLifecycleStateAVAILABLE string = "AVAILABLE"

	// VolumeLifecycleStateTERMINATING captures enum value "TERMINATING"
	VolumeLifecycleStateTERMINATING string = "TERMINATING"

	// VolumeLifecycleStateTERMINATED captures enum value "TERMINATED"
	VolumeLifecycleStateTERMINATED string = "TERMINATED"

	// VolumeLifecycleStateFAULTY captures enum value "FAULTY"
	VolumeLifecycleStateFAULTY string = "FAULTY"
)

// prop value enum
func (m *Volume) validateLifecycleStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, volumeTypeLifecycleStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Volume) validateLifecycleState(formats strfmt.Registry) error {

	if err := validate.Required("lifecycleState", "body", m.LifecycleState); err != nil {
		return err
	}

	// value enum
	if err := m.validateLifecycleStateEnum("lifecycleState", "body", *m.LifecycleState); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateSizeInMBs(formats strfmt.Registry) error {

	if err := validate.Required("sizeInMBs", "body", m.SizeInMBs); err != nil {
		return err
	}

	return nil
}

func (m *Volume) validateTimeCreated(formats strfmt.Registry) error {

	if err := validate.Required("timeCreated", "body", m.TimeCreated); err != nil {
		return err
	}

	if err := validate.FormatOf("timeCreated", "body", "date-time", m.TimeCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Volume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Volume) UnmarshalBinary(b []byte) error {
	var res Volume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
