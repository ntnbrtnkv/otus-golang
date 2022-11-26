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

// PostUpdateEventMaxParseMemory sets the maximum size in bytes for
// the multipart form parser for this operation.
//
// The default value is 32 MB.
// The multipart parser stores up to this + 10MB.
var PostUpdateEventMaxParseMemory int64 = 32 << 20

// NewPostUpdateEventParams creates a new PostUpdateEventParams object
//
// There are no default values defined in the spec.
func NewPostUpdateEventParams() PostUpdateEventParams {

	return PostUpdateEventParams{}
}

// PostUpdateEventParams contains all the bound params for the post update event operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostUpdateEvent
type PostUpdateEventParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: formData
	*/
	Description string
	/*
	  Required: true
	  In: formData
	*/
	ID string
	/*
	  Required: true
	  In: formData
	*/
	TimeEnd strfmt.DateTime
	/*
	  Required: true
	  In: formData
	*/
	TimeStart strfmt.DateTime
	/*
	  Required: true
	  In: formData
	*/
	Title string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostUpdateEventParams() beforehand.
func (o *PostUpdateEventParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(PostUpdateEventMaxParseMemory); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	fdDescription, fdhkDescription, _ := fds.GetOK("description")
	if err := o.bindDescription(fdDescription, fdhkDescription, route.Formats); err != nil {
		res = append(res, err)
	}

	fdID, fdhkID, _ := fds.GetOK("id")
	if err := o.bindID(fdID, fdhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTimeEnd, fdhkTimeEnd, _ := fds.GetOK("time_end")
	if err := o.bindTimeEnd(fdTimeEnd, fdhkTimeEnd, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTimeStart, fdhkTimeStart, _ := fds.GetOK("time_start")
	if err := o.bindTimeStart(fdTimeStart, fdhkTimeStart, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTitle, fdhkTitle, _ := fds.GetOK("title")
	if err := o.bindTitle(fdTitle, fdhkTitle, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDescription binds and validates parameter Description from formData.
func (o *PostUpdateEventParams) bindDescription(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("description", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("description", "formData", raw); err != nil {
		return err
	}
	o.Description = raw

	return nil
}

// bindID binds and validates parameter ID from formData.
func (o *PostUpdateEventParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("id", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("id", "formData", raw); err != nil {
		return err
	}
	o.ID = raw

	return nil
}

// bindTimeEnd binds and validates parameter TimeEnd from formData.
func (o *PostUpdateEventParams) bindTimeEnd(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("time_end", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("time_end", "formData", raw); err != nil {
		return err
	}

	// Format: date-time
	value, err := formats.Parse("date-time", raw)
	if err != nil {
		return errors.InvalidType("time_end", "formData", "strfmt.DateTime", raw)
	}
	o.TimeEnd = *(value.(*strfmt.DateTime))

	if err := o.validateTimeEnd(formats); err != nil {
		return err
	}

	return nil
}

// validateTimeEnd carries on validations for parameter TimeEnd
func (o *PostUpdateEventParams) validateTimeEnd(formats strfmt.Registry) error {

	if err := validate.FormatOf("time_end", "formData", "date-time", o.TimeEnd.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindTimeStart binds and validates parameter TimeStart from formData.
func (o *PostUpdateEventParams) bindTimeStart(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("time_start", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("time_start", "formData", raw); err != nil {
		return err
	}

	// Format: date-time
	value, err := formats.Parse("date-time", raw)
	if err != nil {
		return errors.InvalidType("time_start", "formData", "strfmt.DateTime", raw)
	}
	o.TimeStart = *(value.(*strfmt.DateTime))

	if err := o.validateTimeStart(formats); err != nil {
		return err
	}

	return nil
}

// validateTimeStart carries on validations for parameter TimeStart
func (o *PostUpdateEventParams) validateTimeStart(formats strfmt.Registry) error {

	if err := validate.FormatOf("time_start", "formData", "date-time", o.TimeStart.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindTitle binds and validates parameter Title from formData.
func (o *PostUpdateEventParams) bindTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("title", "formData", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("title", "formData", raw); err != nil {
		return err
	}
	o.Title = raw

	return nil
}