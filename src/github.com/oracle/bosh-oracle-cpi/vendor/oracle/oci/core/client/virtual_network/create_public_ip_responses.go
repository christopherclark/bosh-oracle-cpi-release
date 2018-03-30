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

// CreatePublicIPReader is a Reader for the CreatePublicIP structure.
type CreatePublicIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePublicIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreatePublicIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePublicIPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreatePublicIPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreatePublicIPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreatePublicIPConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreatePublicIPInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreatePublicIPDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePublicIPOK creates a CreatePublicIPOK with default headers values
func NewCreatePublicIPOK() *CreatePublicIPOK {
	return &CreatePublicIPOK{}
}

/*CreatePublicIPOK handles this case with default header values.

The public IP was created.
*/
type CreatePublicIPOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.PublicIP
}

func (o *CreatePublicIPOK) Error() string {
	return fmt.Sprintf("[POST /publicIps][%d] createPublicIpOK  %+v", 200, o.Payload)
}

func (o *CreatePublicIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header etag
	o.Etag = response.GetHeader("etag")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.PublicIP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePublicIPBadRequest creates a CreatePublicIPBadRequest with default headers values
func NewCreatePublicIPBadRequest() *CreatePublicIPBadRequest {
	return &CreatePublicIPBadRequest{}
}

/*CreatePublicIPBadRequest handles this case with default header values.

Bad Request
*/
type CreatePublicIPBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreatePublicIPBadRequest) Error() string {
	return fmt.Sprintf("[POST /publicIps][%d] createPublicIpBadRequest  %+v", 400, o.Payload)
}

func (o *CreatePublicIPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePublicIPUnauthorized creates a CreatePublicIPUnauthorized with default headers values
func NewCreatePublicIPUnauthorized() *CreatePublicIPUnauthorized {
	return &CreatePublicIPUnauthorized{}
}

/*CreatePublicIPUnauthorized handles this case with default header values.

Unauthorized
*/
type CreatePublicIPUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreatePublicIPUnauthorized) Error() string {
	return fmt.Sprintf("[POST /publicIps][%d] createPublicIpUnauthorized  %+v", 401, o.Payload)
}

func (o *CreatePublicIPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePublicIPNotFound creates a CreatePublicIPNotFound with default headers values
func NewCreatePublicIPNotFound() *CreatePublicIPNotFound {
	return &CreatePublicIPNotFound{}
}

/*CreatePublicIPNotFound handles this case with default header values.

Not Found
*/
type CreatePublicIPNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreatePublicIPNotFound) Error() string {
	return fmt.Sprintf("[POST /publicIps][%d] createPublicIpNotFound  %+v", 404, o.Payload)
}

func (o *CreatePublicIPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePublicIPConflict creates a CreatePublicIPConflict with default headers values
func NewCreatePublicIPConflict() *CreatePublicIPConflict {
	return &CreatePublicIPConflict{}
}

/*CreatePublicIPConflict handles this case with default header values.

Conflict
*/
type CreatePublicIPConflict struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreatePublicIPConflict) Error() string {
	return fmt.Sprintf("[POST /publicIps][%d] createPublicIpConflict  %+v", 409, o.Payload)
}

func (o *CreatePublicIPConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePublicIPInternalServerError creates a CreatePublicIPInternalServerError with default headers values
func NewCreatePublicIPInternalServerError() *CreatePublicIPInternalServerError {
	return &CreatePublicIPInternalServerError{}
}

/*CreatePublicIPInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreatePublicIPInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *CreatePublicIPInternalServerError) Error() string {
	return fmt.Sprintf("[POST /publicIps][%d] createPublicIpInternalServerError  %+v", 500, o.Payload)
}

func (o *CreatePublicIPInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePublicIPDefault creates a CreatePublicIPDefault with default headers values
func NewCreatePublicIPDefault(code int) *CreatePublicIPDefault {
	return &CreatePublicIPDefault{
		_statusCode: code,
	}
}

/*CreatePublicIPDefault handles this case with default header values.

An error has occurred.
*/
type CreatePublicIPDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the create public Ip default response
func (o *CreatePublicIPDefault) Code() int {
	return o._statusCode
}

func (o *CreatePublicIPDefault) Error() string {
	return fmt.Sprintf("[POST /publicIps][%d] CreatePublicIp default  %+v", o._statusCode, o.Payload)
}

func (o *CreatePublicIPDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}