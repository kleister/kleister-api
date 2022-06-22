// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TeamPack team pack
//
// swagger:model team_pack
type TeamPack struct {

	// created at
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// pack
	// Read Only: true
	Pack *Pack `json:"pack,omitempty"`

	// pack id
	// Required: true
	// Format: uuid
	PackID *strfmt.UUID `json:"pack_id"`

	// perm
	// Required: true
	// Enum: [user admin owner]
	Perm *string `json:"perm"`

	// team
	// Read Only: true
	Team *Team `json:"team,omitempty"`

	// team id
	// Required: true
	// Format: uuid
	TeamID *strfmt.UUID `json:"team_id"`

	// updated at
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this team pack
func (m *TeamPack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeamID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamPack) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TeamPack) validatePack(formats strfmt.Registry) error {
	if swag.IsZero(m.Pack) { // not required
		return nil
	}

	if m.Pack != nil {
		if err := m.Pack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pack")
			}
			return err
		}
	}

	return nil
}

func (m *TeamPack) validatePackID(formats strfmt.Registry) error {

	if err := validate.Required("pack_id", "body", m.PackID); err != nil {
		return err
	}

	if err := validate.FormatOf("pack_id", "body", "uuid", m.PackID.String(), formats); err != nil {
		return err
	}

	return nil
}

var teamPackTypePermPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user","admin","owner"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		teamPackTypePermPropEnum = append(teamPackTypePermPropEnum, v)
	}
}

const (

	// TeamPackPermUser captures enum value "user"
	TeamPackPermUser string = "user"

	// TeamPackPermAdmin captures enum value "admin"
	TeamPackPermAdmin string = "admin"

	// TeamPackPermOwner captures enum value "owner"
	TeamPackPermOwner string = "owner"
)

// prop value enum
func (m *TeamPack) validatePermEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, teamPackTypePermPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TeamPack) validatePerm(formats strfmt.Registry) error {

	if err := validate.Required("perm", "body", m.Perm); err != nil {
		return err
	}

	// value enum
	if err := m.validatePermEnum("perm", "body", *m.Perm); err != nil {
		return err
	}

	return nil
}

func (m *TeamPack) validateTeam(formats strfmt.Registry) error {
	if swag.IsZero(m.Team) { // not required
		return nil
	}

	if m.Team != nil {
		if err := m.Team.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

func (m *TeamPack) validateTeamID(formats strfmt.Registry) error {

	if err := validate.Required("team_id", "body", m.TeamID); err != nil {
		return err
	}

	if err := validate.FormatOf("team_id", "body", "uuid", m.TeamID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TeamPack) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this team pack based on the context it is used
func (m *TeamPack) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamPack) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *TeamPack) contextValidatePack(ctx context.Context, formats strfmt.Registry) error {

	if m.Pack != nil {
		if err := m.Pack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pack")
			}
			return err
		}
	}

	return nil
}

func (m *TeamPack) contextValidateTeam(ctx context.Context, formats strfmt.Registry) error {

	if m.Team != nil {
		if err := m.Team.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

func (m *TeamPack) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamPack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamPack) UnmarshalBinary(b []byte) error {
	var res TeamPack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
