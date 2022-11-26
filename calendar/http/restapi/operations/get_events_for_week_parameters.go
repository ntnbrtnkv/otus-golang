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

// NewGetEventsForWeekParams creates a new GetEventsForWeekParams object
//
// There are no default values defined in the spec.
func NewGetEventsForWeekParams() GetEventsForWeekParams {

	return GetEventsForWeekParams{}
}

// GetEventsForWeekParams contains all the bound params for the get events for week operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetEventsForWeek
type GetEventsForWeekParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	Week strfmt.Date
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetEventsForWeekParams() beforehand.
func (o *GetEventsForWeekParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qWeek, qhkWeek, _ := qs.GetOK("week")
	if err := o.bindWeek(qWeek, qhkWeek, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindWeek binds and validates parameter Week from query.
func (o *GetEventsForWeekParams) bindWeek(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("week", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("week", "query", raw); err != nil {
		return err
	}

	// Format: date
	value, err := formats.Parse("date", raw)
	if err != nil {
		return errors.InvalidType("week", "query", "strfmt.Date", raw)
	}
	o.Week = *(value.(*strfmt.Date))

	if err := o.validateWeek(formats); err != nil {
		return err
	}

	return nil
}

// validateWeek carries on validations for parameter Week
func (o *GetEventsForWeekParams) validateWeek(formats strfmt.Registry) error {

	if err := validate.FormatOf("week", "query", "date", o.Week.String(), formats); err != nil {
		return err
	}
	return nil
}