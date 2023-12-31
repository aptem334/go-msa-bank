// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Client client
//
// swagger:model Client
type Client struct {

	// birth date
	// Required: true
	// Format: date
	BirthDate *strfmt.Date `json:"birthDate"`

	// first name
	// Required: true
	FirstName *string `json:"firstName"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// last name
	// Required: true
	LastName *string `json:"lastName"`
}

// Validate validates this client
func (m *Client) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBirthDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Client) validateBirthDate(formats strfmt.Registry) error {

	if err := validate.Required("birthDate", "body", m.BirthDate); err != nil {
		return err
	}

	if err := validate.FormatOf("birthDate", "body", "date", m.BirthDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Client) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *Client) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Client) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this client based on context it is used
func (m *Client) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Client) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Client) UnmarshalBinary(b []byte) error {
	var res Client
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
