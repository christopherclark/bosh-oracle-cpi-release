package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"oracle/oci/core/models"
)

// DeletePrivateIPReader is a Reader for the DeletePrivateIP structure.
type DeletePrivateIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePrivateIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeletePrivateIPNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeletePrivateIPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeletePrivateIPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeletePrivateIPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeletePrivateIPPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeletePrivateIPInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeletePrivateIPDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePrivateIPNoContent creates a DeletePrivateIPNoContent with default headers values
func NewDeletePrivateIPNoContent() *DeletePrivateIPNoContent {
	return &DeletePrivateIPNoContent{}
}

/*DeletePrivateIPNoContent handles this case with default header values.

The private IP is being deleted.
*/
type DeletePrivateIPNoContent struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string
}

func (o *DeletePrivateIPNoContent) Error() string {
	return fmt.Sprintf("[DELETE /privateIps/{privateIpId}][%d] deletePrivateIpNoContent ", 204)
}

func (o *DeletePrivateIPNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	return nil
}

// NewDeletePrivateIPBadRequest creates a DeletePrivateIPBadRequest with default headers values
func NewDeletePrivateIPBadRequest() *DeletePrivateIPBadRequest {
	return &DeletePrivateIPBadRequest{}
}

/*DeletePrivateIPBadRequest handles this case with default header values.

Bad Request
*/
type DeletePrivateIPBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeletePrivateIPBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /privateIps/{privateIpId}][%d] deletePrivateIpBadRequest  %+v", 400, o.Payload)
}

func (o *DeletePrivateIPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateIPUnauthorized creates a DeletePrivateIPUnauthorized with default headers values
func NewDeletePrivateIPUnauthorized() *DeletePrivateIPUnauthorized {
	return &DeletePrivateIPUnauthorized{}
}

/*DeletePrivateIPUnauthorized handles this case with default header values.

Unauthorized
*/
type DeletePrivateIPUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeletePrivateIPUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /privateIps/{privateIpId}][%d] deletePrivateIpUnauthorized  %+v", 401, o.Payload)
}

func (o *DeletePrivateIPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateIPNotFound creates a DeletePrivateIPNotFound with default headers values
func NewDeletePrivateIPNotFound() *DeletePrivateIPNotFound {
	return &DeletePrivateIPNotFound{}
}

/*DeletePrivateIPNotFound handles this case with default header values.

Not Found
*/
type DeletePrivateIPNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeletePrivateIPNotFound) Error() string {
	return fmt.Sprintf("[DELETE /privateIps/{privateIpId}][%d] deletePrivateIpNotFound  %+v", 404, o.Payload)
}

func (o *DeletePrivateIPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateIPPreconditionFailed creates a DeletePrivateIPPreconditionFailed with default headers values
func NewDeletePrivateIPPreconditionFailed() *DeletePrivateIPPreconditionFailed {
	return &DeletePrivateIPPreconditionFailed{}
}

/*DeletePrivateIPPreconditionFailed handles this case with default header values.

Precondition Failed
*/
type DeletePrivateIPPreconditionFailed struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeletePrivateIPPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /privateIps/{privateIpId}][%d] deletePrivateIpPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeletePrivateIPPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateIPInternalServerError creates a DeletePrivateIPInternalServerError with default headers values
func NewDeletePrivateIPInternalServerError() *DeletePrivateIPInternalServerError {
	return &DeletePrivateIPInternalServerError{}
}

/*DeletePrivateIPInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeletePrivateIPInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *DeletePrivateIPInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /privateIps/{privateIpId}][%d] deletePrivateIpInternalServerError  %+v", 500, o.Payload)
}

func (o *DeletePrivateIPInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePrivateIPDefault creates a DeletePrivateIPDefault with default headers values
func NewDeletePrivateIPDefault(code int) *DeletePrivateIPDefault {
	return &DeletePrivateIPDefault{
		_statusCode: code,
	}
}

/*DeletePrivateIPDefault handles this case with default header values.

An error has occurred.
*/
type DeletePrivateIPDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the delete private Ip default response
func (o *DeletePrivateIPDefault) Code() int {
	return o._statusCode
}

func (o *DeletePrivateIPDefault) Error() string {
	return fmt.Sprintf("[DELETE /privateIps/{privateIpId}][%d] DeletePrivateIp default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePrivateIPDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}