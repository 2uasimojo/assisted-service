// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NodeCreateParams node create params
// swagger:model node-create-params
type NodeCreateParams struct {

	// hardware info
	// Required: true
	HardwareInfo *string `json:"hardware_info"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// serial
	// Required: true
	Serial *string `json:"serial"`
}

// Validate validates this node create params
func (m *NodeCreateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHardwareInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeCreateParams) validateHardwareInfo(formats strfmt.Registry) error {

	if err := validate.Required("hardware_info", "body", m.HardwareInfo); err != nil {
		return err
	}

	return nil
}

func (m *NodeCreateParams) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *NodeCreateParams) validateSerial(formats strfmt.Registry) error {

	if err := validate.Required("serial", "body", m.Serial); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeCreateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeCreateParams) UnmarshalBinary(b []byte) error {
	var res NodeCreateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
