// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetEventsForDayParams creates a new GetEventsForDayParams object
//
// There are no default values defined in the spec.
func NewGetEventsForDayParams() GetEventsForDayParams {

	return GetEventsForDayParams{}
}

// GetEventsForDayParams contains all the bound params for the get events for day operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetEventsForDay
type GetEventsForDayParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	Day strfmt.Date
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetEventsForDayParams() beforehand.
func (o *GetEventsForDayParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDay, qhkDay, _ := qs.GetOK("day")
	if err := o.bindDay(qDay, qhkDay, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDay binds and validates parameter Day from query.
func (o *GetEventsForDayParams) bindDay(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("day", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("day", "query", raw); err != nil {
		return err
	}

	// Format: date
	value, err := formats.Parse("date", raw)
	if err != nil {
		return errors.InvalidType("day", "query", "strfmt.Date", raw)
	}
	o.Day = *(value.(*strfmt.Date))

	if err := o.validateDay(formats); err != nil {
		return err
	}

	return nil
}

// validateDay carries on validations for parameter Day
func (o *GetEventsForDayParams) validateDay(formats strfmt.Registry) error {

	if err := validate.FormatOf("day", "query", "date", o.Day.String(), formats); err != nil {
		return err
	}
	return nil
}
