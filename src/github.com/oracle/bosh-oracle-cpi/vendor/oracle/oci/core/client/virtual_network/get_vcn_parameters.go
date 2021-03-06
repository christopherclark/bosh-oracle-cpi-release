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

// NewGetVcnParams creates a new GetVcnParams object
// with the default values initialized.
func NewGetVcnParams() *GetVcnParams {
	var ()
	return &GetVcnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVcnParamsWithTimeout creates a new GetVcnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVcnParamsWithTimeout(timeout time.Duration) *GetVcnParams {
	var ()
	return &GetVcnParams{

		timeout: timeout,
	}
}

// NewGetVcnParamsWithContext creates a new GetVcnParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVcnParamsWithContext(ctx context.Context) *GetVcnParams {
	var ()
	return &GetVcnParams{

		Context: ctx,
	}
}

// NewGetVcnParamsWithHTTPClient creates a new GetVcnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVcnParamsWithHTTPClient(client *http.Client) *GetVcnParams {
	var ()
	return &GetVcnParams{
		HTTPClient: client,
	}
}

/*GetVcnParams contains all the parameters to send to the API endpoint
for the get vcn operation typically these are written to a http.Request
*/
type GetVcnParams struct {

	/*VcnID
	  The OCID of the VCN.

	*/
	VcnID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vcn params
func (o *GetVcnParams) WithTimeout(timeout time.Duration) *GetVcnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vcn params
func (o *GetVcnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vcn params
func (o *GetVcnParams) WithContext(ctx context.Context) *GetVcnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vcn params
func (o *GetVcnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vcn params
func (o *GetVcnParams) WithHTTPClient(client *http.Client) *GetVcnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vcn params
func (o *GetVcnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVcnID adds the vcnID to the get vcn params
func (o *GetVcnParams) WithVcnID(vcnID string) *GetVcnParams {
	o.SetVcnID(vcnID)
	return o
}

// SetVcnID adds the vcnId to the get vcn params
func (o *GetVcnParams) SetVcnID(vcnID string) {
	o.VcnID = vcnID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVcnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param vcnId
	if err := r.SetPathParam("vcnId", o.VcnID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
