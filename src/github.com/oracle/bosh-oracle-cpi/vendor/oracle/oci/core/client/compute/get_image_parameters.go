// Code generated by go-swagger; DO NOT EDIT.

package compute

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

// NewGetImageParams creates a new GetImageParams object
// with the default values initialized.
func NewGetImageParams() *GetImageParams {
	var ()
	return &GetImageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImageParamsWithTimeout creates a new GetImageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImageParamsWithTimeout(timeout time.Duration) *GetImageParams {
	var ()
	return &GetImageParams{

		timeout: timeout,
	}
}

// NewGetImageParamsWithContext creates a new GetImageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImageParamsWithContext(ctx context.Context) *GetImageParams {
	var ()
	return &GetImageParams{

		Context: ctx,
	}
}

// NewGetImageParamsWithHTTPClient creates a new GetImageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImageParamsWithHTTPClient(client *http.Client) *GetImageParams {
	var ()
	return &GetImageParams{
		HTTPClient: client,
	}
}

/*GetImageParams contains all the parameters to send to the API endpoint
for the get image operation typically these are written to a http.Request
*/
type GetImageParams struct {

	/*ImageID
	  The OCID of the image.

	*/
	ImageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get image params
func (o *GetImageParams) WithTimeout(timeout time.Duration) *GetImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get image params
func (o *GetImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get image params
func (o *GetImageParams) WithContext(ctx context.Context) *GetImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get image params
func (o *GetImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get image params
func (o *GetImageParams) WithHTTPClient(client *http.Client) *GetImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get image params
func (o *GetImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageID adds the imageID to the get image params
func (o *GetImageParams) WithImageID(imageID string) *GetImageParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the get image params
func (o *GetImageParams) SetImageID(imageID string) {
	o.ImageID = imageID
}

// WriteToRequest writes these params to a swagger request
func (o *GetImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param imageId
	if err := r.SetPathParam("imageId", o.ImageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
