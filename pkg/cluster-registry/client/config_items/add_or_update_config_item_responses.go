// Code generated by go-swagger; DO NOT EDIT.

package config_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/models"
)

// AddOrUpdateConfigItemReader is a Reader for the AddOrUpdateConfigItem structure.
type AddOrUpdateConfigItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOrUpdateConfigItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddOrUpdateConfigItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddOrUpdateConfigItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddOrUpdateConfigItemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddOrUpdateConfigItemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddOrUpdateConfigItemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddOrUpdateConfigItemOK creates a AddOrUpdateConfigItemOK with default headers values
func NewAddOrUpdateConfigItemOK() *AddOrUpdateConfigItemOK {
	return &AddOrUpdateConfigItemOK{}
}

/*AddOrUpdateConfigItemOK handles this case with default header values.

The config items add/update request is accepted.
*/
type AddOrUpdateConfigItemOK struct {
	Payload *models.ConfigValue
}

func (o *AddOrUpdateConfigItemOK) Error() string {
	return fmt.Sprintf("[PUT /kubernetes-clusters/{cluster_id}/config-items/{config_key}][%d] addOrUpdateConfigItemOK  %+v", 200, o.Payload)
}

func (o *AddOrUpdateConfigItemOK) GetPayload() *models.ConfigValue {
	return o.Payload
}

func (o *AddOrUpdateConfigItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrUpdateConfigItemBadRequest creates a AddOrUpdateConfigItemBadRequest with default headers values
func NewAddOrUpdateConfigItemBadRequest() *AddOrUpdateConfigItemBadRequest {
	return &AddOrUpdateConfigItemBadRequest{}
}

/*AddOrUpdateConfigItemBadRequest handles this case with default header values.

Invalid request
*/
type AddOrUpdateConfigItemBadRequest struct {
	Payload *models.Error
}

func (o *AddOrUpdateConfigItemBadRequest) Error() string {
	return fmt.Sprintf("[PUT /kubernetes-clusters/{cluster_id}/config-items/{config_key}][%d] addOrUpdateConfigItemBadRequest  %+v", 400, o.Payload)
}

func (o *AddOrUpdateConfigItemBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddOrUpdateConfigItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOrUpdateConfigItemUnauthorized creates a AddOrUpdateConfigItemUnauthorized with default headers values
func NewAddOrUpdateConfigItemUnauthorized() *AddOrUpdateConfigItemUnauthorized {
	return &AddOrUpdateConfigItemUnauthorized{}
}

/*AddOrUpdateConfigItemUnauthorized handles this case with default header values.

Unauthorized
*/
type AddOrUpdateConfigItemUnauthorized struct {
}

func (o *AddOrUpdateConfigItemUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /kubernetes-clusters/{cluster_id}/config-items/{config_key}][%d] addOrUpdateConfigItemUnauthorized ", 401)
}

func (o *AddOrUpdateConfigItemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddOrUpdateConfigItemForbidden creates a AddOrUpdateConfigItemForbidden with default headers values
func NewAddOrUpdateConfigItemForbidden() *AddOrUpdateConfigItemForbidden {
	return &AddOrUpdateConfigItemForbidden{}
}

/*AddOrUpdateConfigItemForbidden handles this case with default header values.

Forbidden
*/
type AddOrUpdateConfigItemForbidden struct {
}

func (o *AddOrUpdateConfigItemForbidden) Error() string {
	return fmt.Sprintf("[PUT /kubernetes-clusters/{cluster_id}/config-items/{config_key}][%d] addOrUpdateConfigItemForbidden ", 403)
}

func (o *AddOrUpdateConfigItemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddOrUpdateConfigItemInternalServerError creates a AddOrUpdateConfigItemInternalServerError with default headers values
func NewAddOrUpdateConfigItemInternalServerError() *AddOrUpdateConfigItemInternalServerError {
	return &AddOrUpdateConfigItemInternalServerError{}
}

/*AddOrUpdateConfigItemInternalServerError handles this case with default header values.

Unexpected error
*/
type AddOrUpdateConfigItemInternalServerError struct {
	Payload *models.Error
}

func (o *AddOrUpdateConfigItemInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /kubernetes-clusters/{cluster_id}/config-items/{config_key}][%d] addOrUpdateConfigItemInternalServerError  %+v", 500, o.Payload)
}

func (o *AddOrUpdateConfigItemInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddOrUpdateConfigItemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
