// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ntnbrtnkv/otus-golang/calendar/http/models"
)

// PostCreateEventCreatedCode is the HTTP code returned for type PostCreateEventCreated
const PostCreateEventCreatedCode int = 201

/*
PostCreateEventCreated Created successfully

swagger:response postCreateEventCreated
*/
type PostCreateEventCreated struct {
}

// NewPostCreateEventCreated creates PostCreateEventCreated with default headers values
func NewPostCreateEventCreated() *PostCreateEventCreated {

	return &PostCreateEventCreated{}
}

// WriteResponse to the client
func (o *PostCreateEventCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostCreateEventBadRequestCode is the HTTP code returned for type PostCreateEventBadRequest
const PostCreateEventBadRequestCode int = 400

/*
PostCreateEventBadRequest Failed to create event

swagger:response postCreateEventBadRequest
*/
type PostCreateEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostCreateEventBadRequest creates PostCreateEventBadRequest with default headers values
func NewPostCreateEventBadRequest() *PostCreateEventBadRequest {

	return &PostCreateEventBadRequest{}
}

// WithPayload adds the payload to the post create event bad request response
func (o *PostCreateEventBadRequest) WithPayload(payload *models.Error) *PostCreateEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post create event bad request response
func (o *PostCreateEventBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCreateEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
