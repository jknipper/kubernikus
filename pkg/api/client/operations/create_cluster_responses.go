// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sapcc/kubernikus/pkg/api/models"
)

// CreateClusterReader is a Reader for the CreateCluster structure.
type CreateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateClusterCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateClusterCreated creates a CreateClusterCreated with default headers values
func NewCreateClusterCreated() *CreateClusterCreated {
	return &CreateClusterCreated{}
}

/*CreateClusterCreated handles this case with default header values.

OK
*/
type CreateClusterCreated struct {
	Payload *models.Kluster
}

func (o *CreateClusterCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/clusters][%d] createClusterCreated  %+v", 201, o.Payload)
}

func (o *CreateClusterCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Kluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterDefault creates a CreateClusterDefault with default headers values
func NewCreateClusterDefault(code int) *CreateClusterDefault {
	return &CreateClusterDefault{
		_statusCode: code,
	}
}

/*CreateClusterDefault handles this case with default header values.

Error
*/
type CreateClusterDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create cluster default response
func (o *CreateClusterDefault) Code() int {
	return o._statusCode
}

func (o *CreateClusterDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/clusters][%d] CreateCluster default  %+v", o._statusCode, o.Payload)
}

func (o *CreateClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
