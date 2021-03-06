// Code generated by go-swagger; DO NOT EDIT.

package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDhcpOptionsParams creates a new GetDhcpOptionsParams object
// with the default values initialized.
func NewGetDhcpOptionsParams() *GetDhcpOptionsParams {
	var ()
	return &GetDhcpOptionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDhcpOptionsParamsWithTimeout creates a new GetDhcpOptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDhcpOptionsParamsWithTimeout(timeout time.Duration) *GetDhcpOptionsParams {
	var ()
	return &GetDhcpOptionsParams{

		timeout: timeout,
	}
}

// NewGetDhcpOptionsParamsWithContext creates a new GetDhcpOptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDhcpOptionsParamsWithContext(ctx context.Context) *GetDhcpOptionsParams {
	var ()
	return &GetDhcpOptionsParams{

		Context: ctx,
	}
}

// NewGetDhcpOptionsParamsWithHTTPClient creates a new GetDhcpOptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDhcpOptionsParamsWithHTTPClient(client *http.Client) *GetDhcpOptionsParams {
	var ()
	return &GetDhcpOptionsParams{
		HTTPClient: client,
	}
}

/*GetDhcpOptionsParams contains all the parameters to send to the API endpoint
for the get dhcp options operation typically these are written to a http.Request
*/
type GetDhcpOptionsParams struct {

	/*DhcpID
	  The OCID for the set of DHCP options.

	*/
	DhcpID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get dhcp options params
func (o *GetDhcpOptionsParams) WithTimeout(timeout time.Duration) *GetDhcpOptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dhcp options params
func (o *GetDhcpOptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dhcp options params
func (o *GetDhcpOptionsParams) WithContext(ctx context.Context) *GetDhcpOptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dhcp options params
func (o *GetDhcpOptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dhcp options params
func (o *GetDhcpOptionsParams) WithHTTPClient(client *http.Client) *GetDhcpOptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dhcp options params
func (o *GetDhcpOptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDhcpID adds the dhcpID to the get dhcp options params
func (o *GetDhcpOptionsParams) WithDhcpID(dhcpID string) *GetDhcpOptionsParams {
	o.SetDhcpID(dhcpID)
	return o
}

// SetDhcpID adds the dhcpId to the get dhcp options params
func (o *GetDhcpOptionsParams) SetDhcpID(dhcpID string) {
	o.DhcpID = dhcpID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDhcpOptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dhcpId
	if err := r.SetPathParam("dhcpId", o.DhcpID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
