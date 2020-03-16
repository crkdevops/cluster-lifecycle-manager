// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterUpdate cluster update
//
// swagger:model ClusterUpdate
type ClusterUpdate struct {

	// Human readable alias for the Kubernetes cluster. The alias is unique
	// but can be changed.
	//
	Alias string `json:"alias,omitempty"`

	// URL of the Kubernetes API server endpoint
	APIServerURL string `json:"api_server_url,omitempty"`

	// A version channel for the cluster. Possible values are "alpha", "stable"
	Channel string `json:"channel,omitempty"`

	// Configuration items unique to the cluster. E.g. custom API key used
	// by one of the cluster services.
	//
	ConfigItems map[string]string `json:"config_items,omitempty"`

	// Level of criticality as defined by tech controlling. 1 is non
	// critical, 2 is standard production, 3 is PCI.
	//
	CriticalityLevel int32 `json:"criticality_level,omitempty"`

	// The environment in which the cluster run.
	//
	Environment string `json:"environment,omitempty"`

	// Globally unique ID of the Kubernetes cluster
	ID string `json:"id,omitempty"`

	// The identifier of the infrastructure account in which the cluster will
	// live in
	//
	InfrastructureAccount string `json:"infrastructure_account,omitempty"`

	// Status of the cluster.
	// Enum: [requested creating ready decommission-requested decommissioned]
	LifecycleStatus string `json:"lifecycle_status,omitempty"`

	// Cluster identifier which is local to the region
	LocalID string `json:"local_id,omitempty"`

	// The provider of the cluster. Possible values are "zalando-aws", "GKE", ...
	Provider string `json:"provider,omitempty"`

	// The region of the cluster
	Region string `json:"region,omitempty"`

	// status
	Status *ClusterStatus `json:"status,omitempty"`
}

// Validate validates this cluster update
func (m *ClusterUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLifecycleStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var clusterUpdateTypeLifecycleStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["requested","creating","ready","decommission-requested","decommissioned"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterUpdateTypeLifecycleStatusPropEnum = append(clusterUpdateTypeLifecycleStatusPropEnum, v)
	}
}

const (

	// ClusterUpdateLifecycleStatusRequested captures enum value "requested"
	ClusterUpdateLifecycleStatusRequested string = "requested"

	// ClusterUpdateLifecycleStatusCreating captures enum value "creating"
	ClusterUpdateLifecycleStatusCreating string = "creating"

	// ClusterUpdateLifecycleStatusReady captures enum value "ready"
	ClusterUpdateLifecycleStatusReady string = "ready"

	// ClusterUpdateLifecycleStatusDecommissionRequested captures enum value "decommission-requested"
	ClusterUpdateLifecycleStatusDecommissionRequested string = "decommission-requested"

	// ClusterUpdateLifecycleStatusDecommissioned captures enum value "decommissioned"
	ClusterUpdateLifecycleStatusDecommissioned string = "decommissioned"
)

// prop value enum
func (m *ClusterUpdate) validateLifecycleStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterUpdateTypeLifecycleStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterUpdate) validateLifecycleStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.LifecycleStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateLifecycleStatusEnum("lifecycle_status", "body", m.LifecycleStatus); err != nil {
		return err
	}

	return nil
}

func (m *ClusterUpdate) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterUpdate) UnmarshalBinary(b []byte) error {
	var res ClusterUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
