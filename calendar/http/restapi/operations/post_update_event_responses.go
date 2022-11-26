// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ntnbrtnkv/otus-golang/calendar/http/models"
)

// PostUpdateEventOKCode is the HTTP code returned for type PostUpdateEventOK
const PostUpdateEventOKCode int = 200

/*
PostUpdateEventOK Updated successfully

swagger:response postUpdateEventOK
*/
type PostUpdateEventOK struct {
}

// NewPostUpdateEventOK creates PostUpdateEventOK with default headers values
func NewPostUpdateEventOK() *PostUpdateEventOK {

	return &PostUpdateEventOK{}
}

// WriteResponse to the client
func (o *PostUpdateEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostUpdateEventBadRequestCode is the HTTP code returned for type PostUpdateEventBadRequest
const PostUpdateEventBadRequestCode int = 400

/*
PostUpdateEventBadRequest Failed to update event

swagger:response postUpdateEventBadRequest
*/
type PostUpdateEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostUpdateEventBadRequest creates PostUpdateEventBadRequest with default headers values
func NewPostUpdateEventBadRequest() *PostUpdateEventBadRequest {

	return &PostUpdateEventBadRequest{}
}

// WithPayload adds the payload to the post update event bad request response
func (o *PostUpdateEventBadRequest) WithPayload(payload *models.Error) *PostUpdateEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post update event bad request response
func (o *PostUpdateEventBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUpdateEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}