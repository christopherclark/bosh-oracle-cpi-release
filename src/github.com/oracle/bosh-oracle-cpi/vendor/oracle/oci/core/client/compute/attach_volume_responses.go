// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "oracle/oci/core/models"
)

// AttachVolumeReader is a Reader for the AttachVolume structure.
type AttachVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttachVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAttachVolumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAttachVolumeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAttachVolumeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAttachVolumeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewAttachVolumeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAttachVolumeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAttachVolumeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAttachVolumeOK creates a AttachVolumeOK with default headers values
func NewAttachVolumeOK() *AttachVolumeOK {
	return &AttachVolumeOK{}
}

/*AttachVolumeOK handles this case with default header values.

The volume is being attached.
*/
type AttachVolumeOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload models.VolumeAttachment
}

func (o *AttachVolumeOK) Error() string {
	return fmt.Sprintf("[POST /volumeAttachments/][%d] attachVolumeOK  %+v", 200, o.Payload)
}

func (o *AttachVolumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	// response payload as interface type
	payload, err := models.UnmarshalVolumeAttachment(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewAttachVolumeBadRequest creates a AttachVolumeBadRequest with default headers values
func NewAttachVolumeBadRequest() *AttachVolumeBadRequest {
	return &AttachVolumeBadRequest{}
}

/*AttachVolumeBadRequest handles this case with default header values.

Bad Request
*/
type AttachVolumeBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *AttachVolumeBadRequest) Error() string {
	return fmt.Sprintf("[POST /volumeAttachments/][%d] attachVolumeBadRequest  %+v", 400, o.Payload)
}

func (o *AttachVolumeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachVolumeUnauthorized creates a AttachVolumeUnauthorized with default headers values
func NewAttachVolumeUnauthorized() *AttachVolumeUnauthorized {
	return &AttachVolumeUnauthorized{}
}

/*AttachVolumeUnauthorized handles this case with default header values.

Unauthorized
*/
type AttachVolumeUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *AttachVolumeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /volumeAttachments/][%d] attachVolumeUnauthorized  %+v", 401, o.Payload)
}

func (o *AttachVolumeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachVolumeNotFound creates a AttachVolumeNotFound with default headers values
func NewAttachVolumeNotFound() *AttachVolumeNotFound {
	return &AttachVolumeNotFound{}
}

/*AttachVolumeNotFound handles this case with default header values.

Not Found
*/
type AttachVolumeNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *AttachVolumeNotFound) Error() string {
	return fmt.Sprintf("[POST /volumeAttachments/][%d] attachVolumeNotFound  %+v", 404, o.Payload)
}

func (o *AttachVolumeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachVolumeConflict creates a AttachVolumeConflict with default headers values
func NewAttachVolumeConflict() *AttachVolumeConflict {
	return &AttachVolumeConflict{}
}

/*AttachVolumeConflict handles this case with default header values.

Conflict
*/
type AttachVolumeConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *AttachVolumeConflict) Error() string {
	return fmt.Sprintf("[POST /volumeAttachments/][%d] attachVolumeConflict  %+v", 409, o.Payload)
}

func (o *AttachVolumeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachVolumeInternalServerError creates a AttachVolumeInternalServerError with default headers values
func NewAttachVolumeInternalServerError() *AttachVolumeInternalServerError {
	return &AttachVolumeInternalServerError{}
}

/*AttachVolumeInternalServerError handles this case with default header values.

Internal Server Error
*/
type AttachVolumeInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *AttachVolumeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /volumeAttachments/][%d] attachVolumeInternalServerError  %+v", 500, o.Payload)
}

func (o *AttachVolumeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAttachVolumeDefault creates a AttachVolumeDefault with default headers values
func NewAttachVolumeDefault(code int) *AttachVolumeDefault {
	return &AttachVolumeDefault{
		_statusCode: code,
	}
}

/*AttachVolumeDefault handles this case with default header values.

An error has occurred.
*/
type AttachVolumeDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the attach volume default response
func (o *AttachVolumeDefault) Code() int {
	return o._statusCode
}

func (o *AttachVolumeDefault) Error() string {
	return fmt.Sprintf("[POST /volumeAttachments/][%d] AttachVolume default  %+v", o._statusCode, o.Payload)
}

func (o *AttachVolumeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
