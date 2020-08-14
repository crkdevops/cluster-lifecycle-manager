// Code generated by go-swagger; DO NOT EDIT.

package node_pool_config_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zalando-incubator/cluster-lifecycle-manager/pkg/cluster-registry/models"
)

// DeleteNodePoolConfigItemReader is a Reader for the DeleteNodePoolConfigItem structure.
type DeleteNodePoolConfigItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNodePoolConfigItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNodePoolConfigItemNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteNodePoolConfigItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteNodePoolConfigItemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteNodePoolConfigItemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteNodePoolConfigItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteNodePoolConfigItemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNodePoolConfigItemNoContent creates a DeleteNodePoolConfigItemNoContent with default headers values
func NewDeleteNodePoolConfigItemNoContent() *DeleteNodePoolConfigItemNoContent {
	return &DeleteNodePoolConfigItemNoContent{}
}

/*DeleteNodePoolConfigItemNoContent handles this case with default header values.

Config item deleted.
*/
type DeleteNodePoolConfigItemNoContent struct {
}

func (o *DeleteNodePoolConfigItemNoContent) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}/config-items/{config_key}][%d] deleteNodePoolConfigItemNoContent ", 204)
}

func (o *DeleteNodePoolConfigItemNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodePoolConfigItemBadRequest creates a DeleteNodePoolConfigItemBadRequest with default headers values
func NewDeleteNodePoolConfigItemBadRequest() *DeleteNodePoolConfigItemBadRequest {
	return &DeleteNodePoolConfigItemBadRequest{}
}

/*DeleteNodePoolConfigItemBadRequest handles this case with default header values.

Invalid request
*/
type DeleteNodePoolConfigItemBadRequest struct {
	Payload *models.Error
}

func (o *DeleteNodePoolConfigItemBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}/config-items/{config_key}][%d] deleteNodePoolConfigItemBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteNodePoolConfigItemBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNodePoolConfigItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodePoolConfigItemUnauthorized creates a DeleteNodePoolConfigItemUnauthorized with default headers values
func NewDeleteNodePoolConfigItemUnauthorized() *DeleteNodePoolConfigItemUnauthorized {
	return &DeleteNodePoolConfigItemUnauthorized{}
}

/*DeleteNodePoolConfigItemUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteNodePoolConfigItemUnauthorized struct {
}

func (o *DeleteNodePoolConfigItemUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}/config-items/{config_key}][%d] deleteNodePoolConfigItemUnauthorized ", 401)
}

func (o *DeleteNodePoolConfigItemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodePoolConfigItemForbidden creates a DeleteNodePoolConfigItemForbidden with default headers values
func NewDeleteNodePoolConfigItemForbidden() *DeleteNodePoolConfigItemForbidden {
	return &DeleteNodePoolConfigItemForbidden{}
}

/*DeleteNodePoolConfigItemForbidden handles this case with default header values.

Forbidden
*/
type DeleteNodePoolConfigItemForbidden struct {
}

func (o *DeleteNodePoolConfigItemForbidden) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}/config-items/{config_key}][%d] deleteNodePoolConfigItemForbidden ", 403)
}

func (o *DeleteNodePoolConfigItemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodePoolConfigItemNotFound creates a DeleteNodePoolConfigItemNotFound with default headers values
func NewDeleteNodePoolConfigItemNotFound() *DeleteNodePoolConfigItemNotFound {
	return &DeleteNodePoolConfigItemNotFound{}
}

/*DeleteNodePoolConfigItemNotFound handles this case with default header values.

Config item not found
*/
type DeleteNodePoolConfigItemNotFound struct {
}

func (o *DeleteNodePoolConfigItemNotFound) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}/config-items/{config_key}][%d] deleteNodePoolConfigItemNotFound ", 404)
}

func (o *DeleteNodePoolConfigItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodePoolConfigItemInternalServerError creates a DeleteNodePoolConfigItemInternalServerError with default headers values
func NewDeleteNodePoolConfigItemInternalServerError() *DeleteNodePoolConfigItemInternalServerError {
	return &DeleteNodePoolConfigItemInternalServerError{}
}

/*DeleteNodePoolConfigItemInternalServerError handles this case with default header values.

Unexpected error
*/
type DeleteNodePoolConfigItemInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteNodePoolConfigItemInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /kubernetes-clusters/{cluster_id}/node-pools/{node_pool_name}/config-items/{config_key}][%d] deleteNodePoolConfigItemInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteNodePoolConfigItemInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNodePoolConfigItemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
