// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ClientIPParam IP address of the end user. Use if you're sending activities server-side for a client application.
//
// swagger:model ClientIpParam
type ClientIPParam string

// Validate validates this client Ip param
func (m ClientIPParam) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this client Ip param based on context it is used
func (m ClientIPParam) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}