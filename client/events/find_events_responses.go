// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/t0mk/gometal/models"
)

// FindEventsReader is a Reader for the FindEvents structure.
type FindEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFindEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFindEventsOK creates a FindEventsOK with default headers values
func NewFindEventsOK() *FindEventsOK {
	return &FindEventsOK{}
}

/*FindEventsOK handles this case with default header values.

ok
*/
type FindEventsOK struct {
	Payload *models.EventList
}

func (o *FindEventsOK) Error() string {
	return fmt.Sprintf("[GET /events][%d] findEventsOK  %+v", 200, o.Payload)
}

func (o *FindEventsOK) GetPayload() *models.EventList {
	return o.Payload
}

func (o *FindEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindEventsUnauthorized creates a FindEventsUnauthorized with default headers values
func NewFindEventsUnauthorized() *FindEventsUnauthorized {
	return &FindEventsUnauthorized{}
}

/*FindEventsUnauthorized handles this case with default header values.

unauthorized
*/
type FindEventsUnauthorized struct {
}

func (o *FindEventsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /events][%d] findEventsUnauthorized ", 401)
}

func (o *FindEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}