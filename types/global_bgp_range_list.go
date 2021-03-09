// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GlobalBGPRangeList global Bgp range list
//
// swagger:model GlobalBgpRangeList
type GlobalBGPRangeList struct {

	// global bgp ranges
	GlobalBGPRanges []*GlobalBGPRange `json:"global_bgp_ranges"`
}

// Validate validates this global Bgp range list
func (m *GlobalBGPRangeList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGlobalBGPRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GlobalBGPRangeList) validateGlobalBGPRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.GlobalBGPRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.GlobalBGPRanges); i++ {
		if swag.IsZero(m.GlobalBGPRanges[i]) { // not required
			continue
		}

		if m.GlobalBGPRanges[i] != nil {
			if err := m.GlobalBGPRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("global_bgp_ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this global Bgp range list based on the context it is used
func (m *GlobalBGPRangeList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGlobalBGPRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GlobalBGPRangeList) contextValidateGlobalBGPRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GlobalBGPRanges); i++ {

		if m.GlobalBGPRanges[i] != nil {
			if err := m.GlobalBGPRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("global_bgp_ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GlobalBGPRangeList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GlobalBGPRangeList) UnmarshalBinary(b []byte) error {
	var res GlobalBGPRangeList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}