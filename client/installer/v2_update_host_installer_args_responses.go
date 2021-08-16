// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2UpdateHostInstallerArgsReader is a Reader for the V2UpdateHostInstallerArgs structure.
type V2UpdateHostInstallerArgsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2UpdateHostInstallerArgsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV2UpdateHostInstallerArgsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV2UpdateHostInstallerArgsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV2UpdateHostInstallerArgsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2UpdateHostInstallerArgsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2UpdateHostInstallerArgsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2UpdateHostInstallerArgsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV2UpdateHostInstallerArgsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2UpdateHostInstallerArgsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewV2UpdateHostInstallerArgsNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2UpdateHostInstallerArgsCreated creates a V2UpdateHostInstallerArgsCreated with default headers values
func NewV2UpdateHostInstallerArgsCreated() *V2UpdateHostInstallerArgsCreated {
	return &V2UpdateHostInstallerArgsCreated{}
}

/*V2UpdateHostInstallerArgsCreated handles this case with default header values.

Success.
*/
type V2UpdateHostInstallerArgsCreated struct {
	Payload *models.Host
}

func (o *V2UpdateHostInstallerArgsCreated) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/installer-args][%d] v2UpdateHostInstallerArgsCreated  %+v", 201, o.Payload)
}

func (o *V2UpdateHostInstallerArgsCreated) GetPayload() *models.Host {
	return o.Payload
}

func (o *V2UpdateHostInstallerArgsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Host)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostInstallerArgsBadRequest creates a V2UpdateHostInstallerArgsBadRequest with default headers values
func NewV2UpdateHostInstallerArgsBadRequest() *V2UpdateHostInstallerArgsBadRequest {
	return &V2UpdateHostInstallerArgsBadRequest{}
}

/*V2UpdateHostInstallerArgsBadRequest handles this case with default header values.

Error.
*/
type V2UpdateHostInstallerArgsBadRequest struct {
	Payload *models.Error
}

func (o *V2UpdateHostInstallerArgsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/installer-args][%d] v2UpdateHostInstallerArgsBadRequest  %+v", 400, o.Payload)
}

func (o *V2UpdateHostInstallerArgsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostInstallerArgsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostInstallerArgsUnauthorized creates a V2UpdateHostInstallerArgsUnauthorized with default headers values
func NewV2UpdateHostInstallerArgsUnauthorized() *V2UpdateHostInstallerArgsUnauthorized {
	return &V2UpdateHostInstallerArgsUnauthorized{}
}

/*V2UpdateHostInstallerArgsUnauthorized handles this case with default header values.

Unauthorized.
*/
type V2UpdateHostInstallerArgsUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2UpdateHostInstallerArgsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/installer-args][%d] v2UpdateHostInstallerArgsUnauthorized  %+v", 401, o.Payload)
}

func (o *V2UpdateHostInstallerArgsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2UpdateHostInstallerArgsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostInstallerArgsForbidden creates a V2UpdateHostInstallerArgsForbidden with default headers values
func NewV2UpdateHostInstallerArgsForbidden() *V2UpdateHostInstallerArgsForbidden {
	return &V2UpdateHostInstallerArgsForbidden{}
}

/*V2UpdateHostInstallerArgsForbidden handles this case with default header values.

Forbidden.
*/
type V2UpdateHostInstallerArgsForbidden struct {
	Payload *models.InfraError
}

func (o *V2UpdateHostInstallerArgsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/installer-args][%d] v2UpdateHostInstallerArgsForbidden  %+v", 403, o.Payload)
}

func (o *V2UpdateHostInstallerArgsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2UpdateHostInstallerArgsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostInstallerArgsNotFound creates a V2UpdateHostInstallerArgsNotFound with default headers values
func NewV2UpdateHostInstallerArgsNotFound() *V2UpdateHostInstallerArgsNotFound {
	return &V2UpdateHostInstallerArgsNotFound{}
}

/*V2UpdateHostInstallerArgsNotFound handles this case with default header values.

Error.
*/
type V2UpdateHostInstallerArgsNotFound struct {
	Payload *models.Error
}

func (o *V2UpdateHostInstallerArgsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/installer-args][%d] v2UpdateHostInstallerArgsNotFound  %+v", 404, o.Payload)
}

func (o *V2UpdateHostInstallerArgsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostInstallerArgsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostInstallerArgsMethodNotAllowed creates a V2UpdateHostInstallerArgsMethodNotAllowed with default headers values
func NewV2UpdateHostInstallerArgsMethodNotAllowed() *V2UpdateHostInstallerArgsMethodNotAllowed {
	return &V2UpdateHostInstallerArgsMethodNotAllowed{}
}

/*V2UpdateHostInstallerArgsMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type V2UpdateHostInstallerArgsMethodNotAllowed struct {
	Payload *models.Error
}

func (o *V2UpdateHostInstallerArgsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/installer-args][%d] v2UpdateHostInstallerArgsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2UpdateHostInstallerArgsMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostInstallerArgsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostInstallerArgsConflict creates a V2UpdateHostInstallerArgsConflict with default headers values
func NewV2UpdateHostInstallerArgsConflict() *V2UpdateHostInstallerArgsConflict {
	return &V2UpdateHostInstallerArgsConflict{}
}

/*V2UpdateHostInstallerArgsConflict handles this case with default header values.

Error.
*/
type V2UpdateHostInstallerArgsConflict struct {
	Payload *models.Error
}

func (o *V2UpdateHostInstallerArgsConflict) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/installer-args][%d] v2UpdateHostInstallerArgsConflict  %+v", 409, o.Payload)
}

func (o *V2UpdateHostInstallerArgsConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostInstallerArgsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostInstallerArgsInternalServerError creates a V2UpdateHostInstallerArgsInternalServerError with default headers values
func NewV2UpdateHostInstallerArgsInternalServerError() *V2UpdateHostInstallerArgsInternalServerError {
	return &V2UpdateHostInstallerArgsInternalServerError{}
}

/*V2UpdateHostInstallerArgsInternalServerError handles this case with default header values.

Error.
*/
type V2UpdateHostInstallerArgsInternalServerError struct {
	Payload *models.Error
}

func (o *V2UpdateHostInstallerArgsInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/installer-args][%d] v2UpdateHostInstallerArgsInternalServerError  %+v", 500, o.Payload)
}

func (o *V2UpdateHostInstallerArgsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostInstallerArgsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateHostInstallerArgsNotImplemented creates a V2UpdateHostInstallerArgsNotImplemented with default headers values
func NewV2UpdateHostInstallerArgsNotImplemented() *V2UpdateHostInstallerArgsNotImplemented {
	return &V2UpdateHostInstallerArgsNotImplemented{}
}

/*V2UpdateHostInstallerArgsNotImplemented handles this case with default header values.

Not implemented.
*/
type V2UpdateHostInstallerArgsNotImplemented struct {
	Payload *models.Error
}

func (o *V2UpdateHostInstallerArgsNotImplemented) Error() string {
	return fmt.Sprintf("[PATCH /v2/infra-envs/{infra_env_id}/hosts/{host_id}/installer-args][%d] v2UpdateHostInstallerArgsNotImplemented  %+v", 501, o.Payload)
}

func (o *V2UpdateHostInstallerArgsNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateHostInstallerArgsNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}