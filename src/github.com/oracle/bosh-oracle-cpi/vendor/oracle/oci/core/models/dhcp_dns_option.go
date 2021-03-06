// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DhcpDNSOption DHCP option for specifying how DNS (hostname resolution) is handled in the subnets in the VCN.
// For more information, see
// [DNS in Your Virtual Cloud Network](/Content/Network/Concepts/dns.htm).
//
// swagger:model DhcpDnsOption
type DhcpDNSOption struct {
	typeField string

	customDnsServersField []string

	serverTypeField *string
}

func (m *DhcpDNSOption) Type() string {
	// return m.typeField
	return DiscriminatorTypeValues["DhcpDnsOption"]
}
func (m *DhcpDNSOption) SetType(val string) {
	m.typeField = val
}

func (m *DhcpDNSOption) CustomDNSServers() []string {
	return m.customDnsServersField
}
func (m *DhcpDNSOption) SetCustomDNSServers(val []string) {
	m.customDnsServersField = val
}

func (m *DhcpDNSOption) ServerType() *string {
	return m.serverTypeField
}
func (m *DhcpDNSOption) SetServerType(val *string) {
	m.serverTypeField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *DhcpDNSOption) UnmarshalJSON(raw []byte) error {
	var data struct {
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Type string `json:"type"`

		CustomDNSServers []string `json:"customDnsServers"`

		ServerType *string `json:"serverType"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result DhcpDNSOption

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.customDnsServersField = base.CustomDNSServers

	result.serverTypeField = base.ServerType

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m DhcpDNSOption) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
	}{},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`

		CustomDNSServers []string `json:"customDnsServers,omitempty"`

		ServerType *string `json:"serverType"`
	}{

		Type: m.Type(),

		CustomDNSServers: m.CustomDNSServers(),

		ServerType: m.ServerType(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this dhcp Dns option
func (m *DhcpDNSOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomDNSServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DhcpDNSOption) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type())); err != nil {
		return err
	}

	return nil
}

func (m *DhcpDNSOption) validateCustomDNSServers(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomDNSServers()) { // not required
		return nil
	}

	iCustomDNSServersSize := int64(len(m.CustomDNSServers()))

	if err := validate.MaxItems("customDnsServers", "body", iCustomDNSServersSize, 3); err != nil {
		return err
	}

	return nil
}

var dhcpDnsOptionTypeServerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VcnLocal","VcnLocalPlusInternet","CustomDnsServer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dhcpDnsOptionTypeServerTypePropEnum = append(dhcpDnsOptionTypeServerTypePropEnum, v)
	}
}

// property enum
func (m *DhcpDNSOption) validateServerTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, dhcpDnsOptionTypeServerTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DhcpDNSOption) validateServerType(formats strfmt.Registry) error {

	if err := validate.Required("serverType", "body", m.ServerType()); err != nil {
		return err
	}

	// value enum
	if err := m.validateServerTypeEnum("serverType", "body", *m.ServerType()); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DhcpDNSOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DhcpDNSOption) UnmarshalBinary(b []byte) error {
	var res DhcpDNSOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
