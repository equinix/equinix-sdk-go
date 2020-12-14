// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/models"
)

// FindProjectDevicesReader is a Reader for the FindProjectDevices structure.
type FindProjectDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindProjectDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindProjectDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindProjectDevicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFindProjectDevicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFindProjectDevicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindProjectDevicesOK creates a FindProjectDevicesOK with default headers values
func NewFindProjectDevicesOK() *FindProjectDevicesOK {
	return &FindProjectDevicesOK{}
}

/*FindProjectDevicesOK handles this case with default header values.

ok
*/
type FindProjectDevicesOK struct {
	Payload *models.DeviceList
}

func (o *FindProjectDevicesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{id}/devices][%d] findProjectDevicesOK  %+v", 200, o.Payload)
}

func (o *FindProjectDevicesOK) GetPayload() *models.DeviceList {
	return o.Payload
}

func (o *FindProjectDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindProjectDevicesUnauthorized creates a FindProjectDevicesUnauthorized with default headers values
func NewFindProjectDevicesUnauthorized() *FindProjectDevicesUnauthorized {
	return &FindProjectDevicesUnauthorized{}
}

/*FindProjectDevicesUnauthorized handles this case with default header values.

unauthorized
*/
type FindProjectDevicesUnauthorized struct {
}

func (o *FindProjectDevicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{id}/devices][%d] findProjectDevicesUnauthorized ", 401)
}

func (o *FindProjectDevicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindProjectDevicesForbidden creates a FindProjectDevicesForbidden with default headers values
func NewFindProjectDevicesForbidden() *FindProjectDevicesForbidden {
	return &FindProjectDevicesForbidden{}
}

/*FindProjectDevicesForbidden handles this case with default header values.

forbidden
*/
type FindProjectDevicesForbidden struct {
}

func (o *FindProjectDevicesForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{id}/devices][%d] findProjectDevicesForbidden ", 403)
}

func (o *FindProjectDevicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindProjectDevicesNotFound creates a FindProjectDevicesNotFound with default headers values
func NewFindProjectDevicesNotFound() *FindProjectDevicesNotFound {
	return &FindProjectDevicesNotFound{}
}

/*FindProjectDevicesNotFound handles this case with default header values.

not found
*/
type FindProjectDevicesNotFound struct {
}

func (o *FindProjectDevicesNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{id}/devices][%d] findProjectDevicesNotFound ", 404)
}

func (o *FindProjectDevicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}