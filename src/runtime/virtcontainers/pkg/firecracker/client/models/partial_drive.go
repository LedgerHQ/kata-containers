// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PartialDrive partial drive
// swagger:model PartialDrive
type PartialDrive struct {

	// drive id
	// Required: true
	DriveID *string `json:"drive_id"`

	// Host level path for the guest drive
	PathOnHost string `json:"path_on_host,omitempty"`

	// rate limiter
	RateLimiter *RateLimiter `json:"rate_limiter,omitempty"`
}

// Validate validates this partial drive
func (m *PartialDrive) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDriveID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRateLimiter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartialDrive) validateDriveID(formats strfmt.Registry) error {

	if err := validate.Required("drive_id", "body", m.DriveID); err != nil {
		return err
	}

	return nil
}

func (m *PartialDrive) validateRateLimiter(formats strfmt.Registry) error {

	if swag.IsZero(m.RateLimiter) { // not required
		return nil
	}

	if m.RateLimiter != nil {
		if err := m.RateLimiter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rate_limiter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartialDrive) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartialDrive) UnmarshalBinary(b []byte) error {
	var res PartialDrive
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
