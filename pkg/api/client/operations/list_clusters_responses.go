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

// ListClustersReader is a Reader for the ListClusters structure.
type ListClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListClustersOK creates a ListClustersOK with default headers values
func NewListClustersOK() *ListClustersOK {
	return &ListClustersOK{}
}

/*ListClustersOK handles this case with default header values.

OK
*/
type ListClustersOK struct {
	Payload []*models.Kluster
}

func (o *ListClustersOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/clusters][%d] listClustersOK  %+v", 200, o.Payload)
}

func (o *ListClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersDefault creates a ListClustersDefault with default headers values
func NewListClustersDefault(code int) *ListClustersDefault {
	return &ListClustersDefault{
		_statusCode: code,
	}
}

/*ListClustersDefault handles this case with default header values.

Error
*/
type ListClustersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list clusters default response
func (o *ListClustersDefault) Code() int {
	return o._statusCode
}

func (o *ListClustersDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/clusters][%d] ListClusters default  %+v", o._statusCode, o.Payload)
}

func (o *ListClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
