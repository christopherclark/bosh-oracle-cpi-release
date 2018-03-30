// Code generated by go-swagger; DO NOT EDIT.

package blockstorage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "oracle/oci/core/models"
)

// GetVolumeBackupPolicyAssetAssignmentReader is a Reader for the GetVolumeBackupPolicyAssetAssignment structure.
type GetVolumeBackupPolicyAssetAssignmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVolumeBackupPolicyAssetAssignmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVolumeBackupPolicyAssetAssignmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetVolumeBackupPolicyAssetAssignmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetVolumeBackupPolicyAssetAssignmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetVolumeBackupPolicyAssetAssignmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetVolumeBackupPolicyAssetAssignmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetVolumeBackupPolicyAssetAssignmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVolumeBackupPolicyAssetAssignmentOK creates a GetVolumeBackupPolicyAssetAssignmentOK with default headers values
func NewGetVolumeBackupPolicyAssetAssignmentOK() *GetVolumeBackupPolicyAssetAssignmentOK {
	return &GetVolumeBackupPolicyAssetAssignmentOK{}
}

/*GetVolumeBackupPolicyAssetAssignmentOK handles this case with default header values.

The list is being retrieved.
*/
type GetVolumeBackupPolicyAssetAssignmentOK struct {
	/*For pagination of a list of items. When paging through a list, if this header appears in the response,
	then a partial list might have been returned. Include this value as the `page` parameter for the
	subsequent GET request to get the next batch of items.

	*/
	OpcNextPage string
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload []*models.VolumeBackupPolicyAssignment
}

func (o *GetVolumeBackupPolicyAssetAssignmentOK) Error() string {
	return fmt.Sprintf("[GET /volumeBackupPolicyAssignments][%d] getVolumeBackupPolicyAssetAssignmentOK  %+v", 200, o.Payload)
}

func (o *GetVolumeBackupPolicyAssetAssignmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVolumeBackupPolicyAssetAssignmentBadRequest creates a GetVolumeBackupPolicyAssetAssignmentBadRequest with default headers values
func NewGetVolumeBackupPolicyAssetAssignmentBadRequest() *GetVolumeBackupPolicyAssetAssignmentBadRequest {
	return &GetVolumeBackupPolicyAssetAssignmentBadRequest{}
}

/*GetVolumeBackupPolicyAssetAssignmentBadRequest handles this case with default header values.

Bad Request
*/
type GetVolumeBackupPolicyAssetAssignmentBadRequest struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetVolumeBackupPolicyAssetAssignmentBadRequest) Error() string {
	return fmt.Sprintf("[GET /volumeBackupPolicyAssignments][%d] getVolumeBackupPolicyAssetAssignmentBadRequest  %+v", 400, o.Payload)
}

func (o *GetVolumeBackupPolicyAssetAssignmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeBackupPolicyAssetAssignmentUnauthorized creates a GetVolumeBackupPolicyAssetAssignmentUnauthorized with default headers values
func NewGetVolumeBackupPolicyAssetAssignmentUnauthorized() *GetVolumeBackupPolicyAssetAssignmentUnauthorized {
	return &GetVolumeBackupPolicyAssetAssignmentUnauthorized{}
}

/*GetVolumeBackupPolicyAssetAssignmentUnauthorized handles this case with default header values.

Unauthorized
*/
type GetVolumeBackupPolicyAssetAssignmentUnauthorized struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetVolumeBackupPolicyAssetAssignmentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /volumeBackupPolicyAssignments][%d] getVolumeBackupPolicyAssetAssignmentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVolumeBackupPolicyAssetAssignmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeBackupPolicyAssetAssignmentNotFound creates a GetVolumeBackupPolicyAssetAssignmentNotFound with default headers values
func NewGetVolumeBackupPolicyAssetAssignmentNotFound() *GetVolumeBackupPolicyAssetAssignmentNotFound {
	return &GetVolumeBackupPolicyAssetAssignmentNotFound{}
}

/*GetVolumeBackupPolicyAssetAssignmentNotFound handles this case with default header values.

Not Found
*/
type GetVolumeBackupPolicyAssetAssignmentNotFound struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetVolumeBackupPolicyAssetAssignmentNotFound) Error() string {
	return fmt.Sprintf("[GET /volumeBackupPolicyAssignments][%d] getVolumeBackupPolicyAssetAssignmentNotFound  %+v", 404, o.Payload)
}

func (o *GetVolumeBackupPolicyAssetAssignmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeBackupPolicyAssetAssignmentInternalServerError creates a GetVolumeBackupPolicyAssetAssignmentInternalServerError with default headers values
func NewGetVolumeBackupPolicyAssetAssignmentInternalServerError() *GetVolumeBackupPolicyAssetAssignmentInternalServerError {
	return &GetVolumeBackupPolicyAssetAssignmentInternalServerError{}
}

/*GetVolumeBackupPolicyAssetAssignmentInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetVolumeBackupPolicyAssetAssignmentInternalServerError struct {
	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

func (o *GetVolumeBackupPolicyAssetAssignmentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /volumeBackupPolicyAssignments][%d] getVolumeBackupPolicyAssetAssignmentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVolumeBackupPolicyAssetAssignmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeBackupPolicyAssetAssignmentDefault creates a GetVolumeBackupPolicyAssetAssignmentDefault with default headers values
func NewGetVolumeBackupPolicyAssetAssignmentDefault(code int) *GetVolumeBackupPolicyAssetAssignmentDefault {
	return &GetVolumeBackupPolicyAssetAssignmentDefault{
		_statusCode: code,
	}
}

/*GetVolumeBackupPolicyAssetAssignmentDefault handles this case with default header values.

An error has occurred.
*/
type GetVolumeBackupPolicyAssetAssignmentDefault struct {
	_statusCode int

	/*Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	a particular request, please provide the request ID.

	*/
	OpcRequestID string

	Payload *models.Error
}

// Code gets the status code for the get volume backup policy asset assignment default response
func (o *GetVolumeBackupPolicyAssetAssignmentDefault) Code() int {
	return o._statusCode
}

func (o *GetVolumeBackupPolicyAssetAssignmentDefault) Error() string {
	return fmt.Sprintf("[GET /volumeBackupPolicyAssignments][%d] GetVolumeBackupPolicyAssetAssignment default  %+v", o._statusCode, o.Payload)
}

func (o *GetVolumeBackupPolicyAssetAssignmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header opc-request-id
	o.OpcRequestID = response.GetHeader("opc-request-id")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}