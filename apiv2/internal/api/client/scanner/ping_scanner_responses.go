// Code generated by go-swagger; DO NOT EDIT.

package scanner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v4/apiv2/model"
)

// PingScannerReader is a Reader for the PingScanner structure.
type PingScannerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PingScannerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPingScannerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPingScannerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPingScannerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPingScannerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPingScannerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPingScannerOK creates a PingScannerOK with default headers values
func NewPingScannerOK() *PingScannerOK {
	return &PingScannerOK{}
}

/* PingScannerOK describes a response with status code 200, with default header values.

Success
*/
type PingScannerOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *PingScannerOK) Error() string {
	return fmt.Sprintf("[POST /scanners/ping][%d] pingScannerOK ", 200)
}

func (o *PingScannerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewPingScannerBadRequest creates a PingScannerBadRequest with default headers values
func NewPingScannerBadRequest() *PingScannerBadRequest {
	return &PingScannerBadRequest{}
}

/* PingScannerBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PingScannerBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *PingScannerBadRequest) Error() string {
	return fmt.Sprintf("[POST /scanners/ping][%d] pingScannerBadRequest  %+v", 400, o.Payload)
}
func (o *PingScannerBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *PingScannerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingScannerUnauthorized creates a PingScannerUnauthorized with default headers values
func NewPingScannerUnauthorized() *PingScannerUnauthorized {
	return &PingScannerUnauthorized{}
}

/* PingScannerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PingScannerUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *PingScannerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /scanners/ping][%d] pingScannerUnauthorized  %+v", 401, o.Payload)
}
func (o *PingScannerUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *PingScannerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingScannerForbidden creates a PingScannerForbidden with default headers values
func NewPingScannerForbidden() *PingScannerForbidden {
	return &PingScannerForbidden{}
}

/* PingScannerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PingScannerForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *PingScannerForbidden) Error() string {
	return fmt.Sprintf("[POST /scanners/ping][%d] pingScannerForbidden  %+v", 403, o.Payload)
}
func (o *PingScannerForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *PingScannerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingScannerInternalServerError creates a PingScannerInternalServerError with default headers values
func NewPingScannerInternalServerError() *PingScannerInternalServerError {
	return &PingScannerInternalServerError{}
}

/* PingScannerInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PingScannerInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *PingScannerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /scanners/ping][%d] pingScannerInternalServerError  %+v", 500, o.Payload)
}
func (o *PingScannerInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *PingScannerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}