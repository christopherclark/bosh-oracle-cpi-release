// Code generated by go-swagger; DO NOT EDIT.

package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "oracle/oci/core/models"
)

// GetIPSecConnectionDeviceStatusReader is a Reader for the GetIPSecConnectionDeviceStatus structure.
type GetIPSecConnectionDeviceStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPSecConnectionDeviceStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIPSecConnectionDeviceStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetIPSecConnectionDeviceStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetIPSecConnectionDeviceStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetIPSecConnectionDeviceStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetIPSecConnectionDeviceStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIPSecConnectionDeviceStatusOK creates a GetIPSecConnectionDeviceStatusOK with default headers values
func NewGetIPSecConnectionDeviceStatusOK() *GetIPSecConnectionDeviceStatusOK {
	return &GetIPSecConnectionDeviceStatusOK{}
}

/*GetIPSecConnectionDeviceStatusOK handles this case with default header values.

The status was retrieved.
*/
type GetIPSecConnectionDeviceStatusOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.IPSecConnectionDeviceStatus
}

func (o *GetIPSecConnectionDeviceStatusOK) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections/{ipscId}/deviceStatus][%d] getIpSecConnectionDeviceStatusOK  %+v", 200, o.Payload)
}

func (o *GetIPSecConnectionDeviceStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.IPSecConnectionDeviceStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPSecConnectionDeviceStatusUnauthorized creates a GetIPSecConnectionDeviceStatusUnauthorized with default headers values
func NewGetIPSecConnectionDeviceStatusUnauthorized() *GetIPSecConnectionDeviceStatusUnauthorized {
	return &GetIPSecConnectionDeviceStatusUnauthorized{}
}

/*GetIPSecConnectionDeviceStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type GetIPSecConnectionDeviceStatusUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetIPSecConnectionDeviceStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections/{ipscId}/deviceStatus][%d] getIpSecConnectionDeviceStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIPSecConnectionDeviceStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPSecConnectionDeviceStatusNotFound creates a GetIPSecConnectionDeviceStatusNotFound with default headers values
func NewGetIPSecConnectionDeviceStatusNotFound() *GetIPSecConnectionDeviceStatusNotFound {
	return &GetIPSecConnectionDeviceStatusNotFound{}
}

/*GetIPSecConnectionDeviceStatusNotFound handles this case with default header values.

Not Found
*/
type GetIPSecConnectionDeviceStatusNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetIPSecConnectionDeviceStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections/{ipscId}/deviceStatus][%d] getIpSecConnectionDeviceStatusNotFound  %+v", 404, o.Payload)
}

func (o *GetIPSecConnectionDeviceStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPSecConnectionDeviceStatusInternalServerError creates a GetIPSecConnectionDeviceStatusInternalServerError with default headers values
func NewGetIPSecConnectionDeviceStatusInternalServerError() *GetIPSecConnectionDeviceStatusInternalServerError {
	return &GetIPSecConnectionDeviceStatusInternalServerError{}
}

/*GetIPSecConnectionDeviceStatusInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetIPSecConnectionDeviceStatusInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetIPSecConnectionDeviceStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections/{ipscId}/deviceStatus][%d] getIpSecConnectionDeviceStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIPSecConnectionDeviceStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPSecConnectionDeviceStatusDefault creates a GetIPSecConnectionDeviceStatusDefault with default headers values
func NewGetIPSecConnectionDeviceStatusDefault(code int) *GetIPSecConnectionDeviceStatusDefault {
	return &GetIPSecConnectionDeviceStatusDefault{
		_statusCode: code,
	}
}

/*GetIPSecConnectionDeviceStatusDefault handles this case with default header values.

An error has occurred.
*/
type GetIPSecConnectionDeviceStatusDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get IP sec connection device status default response
func (o *GetIPSecConnectionDeviceStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetIPSecConnectionDeviceStatusDefault) Error() string {
	return fmt.Sprintf("[GET /ipsecConnections/{ipscId}/deviceStatus][%d] GetIPSecConnectionDeviceStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetIPSecConnectionDeviceStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
