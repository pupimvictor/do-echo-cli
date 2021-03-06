// Code generated by go-swagger; DO NOT EDIT.

package echo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pupimvictor/do-echo-cli/models"
)

// EchoReader is a Reader for the Echo structure.
type EchoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EchoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEchoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewEchoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEchoOK creates a EchoOK with default headers values
func NewEchoOK() *EchoOK {
	return &EchoOK{}
}

/*EchoOK handles this case with default header values.

OK
*/
type EchoOK struct {
	Payload *models.Echo
}

func (o *EchoOK) Error() string {
	return fmt.Sprintf("[POST /echo][%d] echoOK  %+v", 200, o.Payload)
}

func (o *EchoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Echo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEchoDefault creates a EchoDefault with default headers values
func NewEchoDefault(code int) *EchoDefault {
	return &EchoDefault{
		_statusCode: code,
	}
}

/*EchoDefault handles this case with default header values.

error
*/
type EchoDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the echo default response
func (o *EchoDefault) Code() int {
	return o._statusCode
}

func (o *EchoDefault) Error() string {
	return fmt.Sprintf("[POST /echo][%d] echo default  %+v", o._statusCode, o.Payload)
}

func (o *EchoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
