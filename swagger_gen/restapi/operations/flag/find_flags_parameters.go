// Code generated by go-swagger; DO NOT EDIT.

package flag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindFlagsParams creates a new FindFlagsParams object
// with the default values initialized.
func NewFindFlagsParams() FindFlagsParams {

	var (
		// initialize parameters with default values

		limitDefault = int64(25)
	)

	return FindFlagsParams{
		Limit: &limitDefault,
	}
}

// FindFlagsParams contains all the bound params for the find flags operation
// typically these are obtained from a http.Request
//
// swagger:parameters findFlags
type FindFlagsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*return only flags matching given description
	  In: query
	*/
	Description *string
	/*return only flags having given enabled status
	  In: query
	*/
	Enabled *bool
	/*the numbers of flags to return
	  Minimum: 1
	  In: query
	  Default: 25
	*/
	Limit *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFindFlagsParams() beforehand.
func (o *FindFlagsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDescription, qhkDescription, _ := qs.GetOK("description")
	if err := o.bindDescription(qDescription, qhkDescription, route.Formats); err != nil {
		res = append(res, err)
	}

	qEnabled, qhkEnabled, _ := qs.GetOK("enabled")
	if err := o.bindEnabled(qEnabled, qhkEnabled, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindFlagsParams) bindDescription(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Description = &raw

	return nil
}

func (o *FindFlagsParams) bindEnabled(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("enabled", "query", "bool", raw)
	}
	o.Enabled = &value

	return nil
}

func (o *FindFlagsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewFindFlagsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

func (o *FindFlagsParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("limit", "query", int64(*o.Limit), 1, false); err != nil {
		return err
	}

	return nil
}
