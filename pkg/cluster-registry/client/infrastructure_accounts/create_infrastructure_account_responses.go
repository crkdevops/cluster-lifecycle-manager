// Code generated by go-swagger; DO NOT EDIT.

package infrastructure_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/models"
)

// CreateInfrastructureAccountReader is a Reader for the CreateInfrastructureAccount structure.
type CreateInfrastructureAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInfrastructureAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateInfrastructureAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateInfrastructureAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateInfrastructureAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateInfrastructureAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateInfrastructureAccountConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateInfrastructureAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateInfrastructureAccountCreated creates a CreateInfrastructureAccountCreated with default headers values
func NewCreateInfrastructureAccountCreated() *CreateInfrastructureAccountCreated {
	return &CreateInfrastructureAccountCreated{}
}

/*CreateInfrastructureAccountCreated handles this case with default header values.

Infrastructure account was scheduled for creation.
*/
type CreateInfrastructureAccountCreated struct {
	Payload *models.InfrastructureAccount
}

func (o *CreateInfrastructureAccountCreated) Error() string {
	return fmt.Sprintf("[POST /infrastructure-accounts][%d] createInfrastructureAccountCreated  %+v", 201, o.Payload)
}

func (o *CreateInfrastructureAccountCreated) GetPayload() *models.InfrastructureAccount {
	return o.Payload
}

func (o *CreateInfrastructureAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfrastructureAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInfrastructureAccountBadRequest creates a CreateInfrastructureAccountBadRequest with default headers values
func NewCreateInfrastructureAccountBadRequest() *CreateInfrastructureAccountBadRequest {
	return &CreateInfrastructureAccountBadRequest{}
}

/*CreateInfrastructureAccountBadRequest handles this case with default header values.

Invalid parameters
*/
type CreateInfrastructureAccountBadRequest struct {
}

func (o *CreateInfrastructureAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /infrastructure-accounts][%d] createInfrastructureAccountBadRequest ", 400)
}

func (o *CreateInfrastructureAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInfrastructureAccountUnauthorized creates a CreateInfrastructureAccountUnauthorized with default headers values
func NewCreateInfrastructureAccountUnauthorized() *CreateInfrastructureAccountUnauthorized {
	return &CreateInfrastructureAccountUnauthorized{}
}

/*CreateInfrastructureAccountUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateInfrastructureAccountUnauthorized struct {
}

func (o *CreateInfrastructureAccountUnauthorized) Error() string {
	return fmt.Sprintf("[POST /infrastructure-accounts][%d] createInfrastructureAccountUnauthorized ", 401)
}

func (o *CreateInfrastructureAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInfrastructureAccountForbidden creates a CreateInfrastructureAccountForbidden with default headers values
func NewCreateInfrastructureAccountForbidden() *CreateInfrastructureAccountForbidden {
	return &CreateInfrastructureAccountForbidden{}
}

/*CreateInfrastructureAccountForbidden handles this case with default header values.

Forbidden
*/
type CreateInfrastructureAccountForbidden struct {
}

func (o *CreateInfrastructureAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /infrastructure-accounts][%d] createInfrastructureAccountForbidden ", 403)
}

func (o *CreateInfrastructureAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInfrastructureAccountConflict creates a CreateInfrastructureAccountConflict with default headers values
func NewCreateInfrastructureAccountConflict() *CreateInfrastructureAccountConflict {
	return &CreateInfrastructureAccountConflict{}
}

/*CreateInfrastructureAccountConflict handles this case with default header values.

Conflict, already existing.
*/
type CreateInfrastructureAccountConflict struct {
}

func (o *CreateInfrastructureAccountConflict) Error() string {
	return fmt.Sprintf("[POST /infrastructure-accounts][%d] createInfrastructureAccountConflict ", 409)
}

func (o *CreateInfrastructureAccountConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInfrastructureAccountInternalServerError creates a CreateInfrastructureAccountInternalServerError with default headers values
func NewCreateInfrastructureAccountInternalServerError() *CreateInfrastructureAccountInternalServerError {
	return &CreateInfrastructureAccountInternalServerError{}
}

/*CreateInfrastructureAccountInternalServerError handles this case with default header values.

Unexpected error
*/
type CreateInfrastructureAccountInternalServerError struct {
	Payload *models.Error
}

func (o *CreateInfrastructureAccountInternalServerError) Error() string {
	return fmt.Sprintf("[POST /infrastructure-accounts][%d] createInfrastructureAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateInfrastructureAccountInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateInfrastructureAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
