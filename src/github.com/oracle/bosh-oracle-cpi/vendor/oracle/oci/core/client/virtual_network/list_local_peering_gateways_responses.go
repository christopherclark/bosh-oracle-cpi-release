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

// ListLocalPeeringGatewaysReader is a Reader for the ListLocalPeeringGateways structure.
type ListLocalPeeringGatewaysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLocalPeeringGatewaysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListLocalPeeringGatewaysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListLocalPeeringGatewaysBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListLocalPeeringGatewaysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListLocalPeeringGatewaysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListLocalPeeringGatewaysInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListLocalPeeringGatewaysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListLocalPeeringGatewaysOK creates a ListLocalPeeringGatewaysOK with default headers values
func NewListLocalPeeringGatewaysOK() *ListLocalPeeringGatewaysOK {
	return &ListLocalPeeringGatewaysOK{}
}

/*ListLocalPeeringGatewaysOK handles this case with default header values.

The list is being retrieved.
*/
type ListLocalPeeringGatewaysOK struct {
	/*For pagination of a list of items. When paging through a list, if this header appears in the response,
	then a partial list might have been returned. Include this value as the `page` parameter for the
	subsequent GET request to get the next batch of items.

	*/
	OpcNextPage string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload []*models.LocalPeeringGateway
}

func (o *ListLocalPeeringGatewaysOK) Error() string {
	return fmt.Sprintf("[GET /localPeeringGateways][%d] listLocalPeeringGatewaysOK  %+v", 200, o.Payload)
}

func (o *ListLocalPeeringGatewaysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-next-page
	o.OpcNextPage = response.GetHeader("opc-next-page")

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocalPeeringGatewaysBadRequest creates a ListLocalPeeringGatewaysBadRequest with default headers values
func NewListLocalPeeringGatewaysBadRequest() *ListLocalPeeringGatewaysBadRequest {
	return &ListLocalPeeringGatewaysBadRequest{}
}

/*ListLocalPeeringGatewaysBadRequest handles this case with default header values.

Bad Request
*/
type ListLocalPeeringGatewaysBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListLocalPeeringGatewaysBadRequest) Error() string {
	return fmt.Sprintf("[GET /localPeeringGateways][%d] listLocalPeeringGatewaysBadRequest  %+v", 400, o.Payload)
}

func (o *ListLocalPeeringGatewaysBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocalPeeringGatewaysUnauthorized creates a ListLocalPeeringGatewaysUnauthorized with default headers values
func NewListLocalPeeringGatewaysUnauthorized() *ListLocalPeeringGatewaysUnauthorized {
	return &ListLocalPeeringGatewaysUnauthorized{}
}

/*ListLocalPeeringGatewaysUnauthorized handles this case with default header values.

Unauthorized
*/
type ListLocalPeeringGatewaysUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListLocalPeeringGatewaysUnauthorized) Error() string {
	return fmt.Sprintf("[GET /localPeeringGateways][%d] listLocalPeeringGatewaysUnauthorized  %+v", 401, o.Payload)
}

func (o *ListLocalPeeringGatewaysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocalPeeringGatewaysNotFound creates a ListLocalPeeringGatewaysNotFound with default headers values
func NewListLocalPeeringGatewaysNotFound() *ListLocalPeeringGatewaysNotFound {
	return &ListLocalPeeringGatewaysNotFound{}
}

/*ListLocalPeeringGatewaysNotFound handles this case with default header values.

Not Found
*/
type ListLocalPeeringGatewaysNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListLocalPeeringGatewaysNotFound) Error() string {
	return fmt.Sprintf("[GET /localPeeringGateways][%d] listLocalPeeringGatewaysNotFound  %+v", 404, o.Payload)
}

func (o *ListLocalPeeringGatewaysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocalPeeringGatewaysInternalServerError creates a ListLocalPeeringGatewaysInternalServerError with default headers values
func NewListLocalPeeringGatewaysInternalServerError() *ListLocalPeeringGatewaysInternalServerError {
	return &ListLocalPeeringGatewaysInternalServerError{}
}

/*ListLocalPeeringGatewaysInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListLocalPeeringGatewaysInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *ListLocalPeeringGatewaysInternalServerError) Error() string {
	return fmt.Sprintf("[GET /localPeeringGateways][%d] listLocalPeeringGatewaysInternalServerError  %+v", 500, o.Payload)
}

func (o *ListLocalPeeringGatewaysInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocalPeeringGatewaysDefault creates a ListLocalPeeringGatewaysDefault with default headers values
func NewListLocalPeeringGatewaysDefault(code int) *ListLocalPeeringGatewaysDefault {
	return &ListLocalPeeringGatewaysDefault{
		_statusCode: code,
	}
}

/*ListLocalPeeringGatewaysDefault handles this case with default header values.

An error has occurred.
*/
type ListLocalPeeringGatewaysDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the list local peering gateways default response
func (o *ListLocalPeeringGatewaysDefault) Code() int {
	return o._statusCode
}

func (o *ListLocalPeeringGatewaysDefault) Error() string {
	return fmt.Sprintf("[GET /localPeeringGateways][%d] ListLocalPeeringGateways default  %+v", o._statusCode, o.Payload)
}

func (o *ListLocalPeeringGatewaysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}