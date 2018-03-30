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

// GetPublicIPByPrivateIPIDReader is a Reader for the GetPublicIPByPrivateIPID structure.
type GetPublicIPByPrivateIPIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicIPByPrivateIPIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicIPByPrivateIPIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPublicIPByPrivateIPIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetPublicIPByPrivateIPIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPublicIPByPrivateIPIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetPublicIPByPrivateIPIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetPublicIPByPrivateIPIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPublicIPByPrivateIPIDOK creates a GetPublicIPByPrivateIPIDOK with default headers values
func NewGetPublicIPByPrivateIPIDOK() *GetPublicIPByPrivateIPIDOK {
	return &GetPublicIPByPrivateIPIDOK{}
}

/*GetPublicIPByPrivateIPIDOK handles this case with default header values.

The public IP was retrieved.
*/
type GetPublicIPByPrivateIPIDOK struct {
	/*For optimistic concurrency control. See `if-match`.
	 */
	Etag string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.PublicIP
}

func (o *GetPublicIPByPrivateIPIDOK) Error() string {
	return fmt.Sprintf("[POST /publicIps/actions/getByPrivateIpId][%d] getPublicIpByPrivateIpIdOK  %+v", 200, o.Payload)
}

func (o *GetPublicIPByPrivateIPIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPublicIPByPrivateIPIDBadRequest creates a GetPublicIPByPrivateIPIDBadRequest with default headers values
func NewGetPublicIPByPrivateIPIDBadRequest() *GetPublicIPByPrivateIPIDBadRequest {
	return &GetPublicIPByPrivateIPIDBadRequest{}
}

/*GetPublicIPByPrivateIPIDBadRequest handles this case with default header values.

Bad Request
*/
type GetPublicIPByPrivateIPIDBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetPublicIPByPrivateIPIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /publicIps/actions/getByPrivateIpId][%d] getPublicIpByPrivateIpIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetPublicIPByPrivateIPIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPublicIPByPrivateIPIDUnauthorized creates a GetPublicIPByPrivateIPIDUnauthorized with default headers values
func NewGetPublicIPByPrivateIPIDUnauthorized() *GetPublicIPByPrivateIPIDUnauthorized {
	return &GetPublicIPByPrivateIPIDUnauthorized{}
}

/*GetPublicIPByPrivateIPIDUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPublicIPByPrivateIPIDUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetPublicIPByPrivateIPIDUnauthorized) Error() string {
	return fmt.Sprintf("[POST /publicIps/actions/getByPrivateIpId][%d] getPublicIpByPrivateIpIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPublicIPByPrivateIPIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPublicIPByPrivateIPIDNotFound creates a GetPublicIPByPrivateIPIDNotFound with default headers values
func NewGetPublicIPByPrivateIPIDNotFound() *GetPublicIPByPrivateIPIDNotFound {
	return &GetPublicIPByPrivateIPIDNotFound{}
}

/*GetPublicIPByPrivateIPIDNotFound handles this case with default header values.

Not Found
*/
type GetPublicIPByPrivateIPIDNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetPublicIPByPrivateIPIDNotFound) Error() string {
	return fmt.Sprintf("[POST /publicIps/actions/getByPrivateIpId][%d] getPublicIpByPrivateIpIdNotFound  %+v", 404, o.Payload)
}

func (o *GetPublicIPByPrivateIPIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPublicIPByPrivateIPIDInternalServerError creates a GetPublicIPByPrivateIPIDInternalServerError with default headers values
func NewGetPublicIPByPrivateIPIDInternalServerError() *GetPublicIPByPrivateIPIDInternalServerError {
	return &GetPublicIPByPrivateIPIDInternalServerError{}
}

/*GetPublicIPByPrivateIPIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetPublicIPByPrivateIPIDInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetPublicIPByPrivateIPIDInternalServerError) Error() string {
	return fmt.Sprintf("[POST /publicIps/actions/getByPrivateIpId][%d] getPublicIpByPrivateIpIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPublicIPByPrivateIPIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPublicIPByPrivateIPIDDefault creates a GetPublicIPByPrivateIPIDDefault with default headers values
func NewGetPublicIPByPrivateIPIDDefault(code int) *GetPublicIPByPrivateIPIDDefault {
	return &GetPublicIPByPrivateIPIDDefault{
		_statusCode: code,
	}
}

/*GetPublicIPByPrivateIPIDDefault handles this case with default header values.

An error has occurred.
*/
type GetPublicIPByPrivateIPIDDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get public Ip by private Ip Id default response
func (o *GetPublicIPByPrivateIPIDDefault) Code() int {
	return o._statusCode
}

func (o *GetPublicIPByPrivateIPIDDefault) Error() string {
	return fmt.Sprintf("[POST /publicIps/actions/getByPrivateIpId][%d] GetPublicIpByPrivateIpId default  %+v", o._statusCode, o.Payload)
}

func (o *GetPublicIPByPrivateIPIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}