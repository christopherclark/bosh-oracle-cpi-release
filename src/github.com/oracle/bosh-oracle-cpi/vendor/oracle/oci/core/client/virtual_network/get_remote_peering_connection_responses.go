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

// GetRemotePeeringConnectionReader is a Reader for the GetRemotePeeringConnection structure.
type GetRemotePeeringConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRemotePeeringConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRemotePeeringConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetRemotePeeringConnectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRemotePeeringConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetRemotePeeringConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetRemotePeeringConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRemotePeeringConnectionOK creates a GetRemotePeeringConnectionOK with default headers values
func NewGetRemotePeeringConnectionOK() *GetRemotePeeringConnectionOK {
	return &GetRemotePeeringConnectionOK{}
}

/*GetRemotePeeringConnectionOK handles this case with default header values.

The RPC was retrieved.
*/
type GetRemotePeeringConnectionOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.RemotePeeringConnection
}

func (o *GetRemotePeeringConnectionOK) Error() string {
	return fmt.Sprintf("[GET /remotePeeringConnections/{remotePeeringConnectionId}][%d] getRemotePeeringConnectionOK  %+v", 200, o.Payload)
}

func (o *GetRemotePeeringConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.RemotePeeringConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRemotePeeringConnectionUnauthorized creates a GetRemotePeeringConnectionUnauthorized with default headers values
func NewGetRemotePeeringConnectionUnauthorized() *GetRemotePeeringConnectionUnauthorized {
	return &GetRemotePeeringConnectionUnauthorized{}
}

/*GetRemotePeeringConnectionUnauthorized handles this case with default header values.

Unauthorized
*/
type GetRemotePeeringConnectionUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetRemotePeeringConnectionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /remotePeeringConnections/{remotePeeringConnectionId}][%d] getRemotePeeringConnectionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRemotePeeringConnectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRemotePeeringConnectionNotFound creates a GetRemotePeeringConnectionNotFound with default headers values
func NewGetRemotePeeringConnectionNotFound() *GetRemotePeeringConnectionNotFound {
	return &GetRemotePeeringConnectionNotFound{}
}

/*GetRemotePeeringConnectionNotFound handles this case with default header values.

Not Found
*/
type GetRemotePeeringConnectionNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetRemotePeeringConnectionNotFound) Error() string {
	return fmt.Sprintf("[GET /remotePeeringConnections/{remotePeeringConnectionId}][%d] getRemotePeeringConnectionNotFound  %+v", 404, o.Payload)
}

func (o *GetRemotePeeringConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRemotePeeringConnectionInternalServerError creates a GetRemotePeeringConnectionInternalServerError with default headers values
func NewGetRemotePeeringConnectionInternalServerError() *GetRemotePeeringConnectionInternalServerError {
	return &GetRemotePeeringConnectionInternalServerError{}
}

/*GetRemotePeeringConnectionInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetRemotePeeringConnectionInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetRemotePeeringConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /remotePeeringConnections/{remotePeeringConnectionId}][%d] getRemotePeeringConnectionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRemotePeeringConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRemotePeeringConnectionDefault creates a GetRemotePeeringConnectionDefault with default headers values
func NewGetRemotePeeringConnectionDefault(code int) *GetRemotePeeringConnectionDefault {
	return &GetRemotePeeringConnectionDefault{
		_statusCode: code,
	}
}

/*GetRemotePeeringConnectionDefault handles this case with default header values.

An error has occurred.
*/
type GetRemotePeeringConnectionDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get remote peering connection default response
func (o *GetRemotePeeringConnectionDefault) Code() int {
	return o._statusCode
}

func (o *GetRemotePeeringConnectionDefault) Error() string {
	return fmt.Sprintf("[GET /remotePeeringConnections/{remotePeeringConnectionId}][%d] GetRemotePeeringConnection default  %+v", o._statusCode, o.Payload)
}

func (o *GetRemotePeeringConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}