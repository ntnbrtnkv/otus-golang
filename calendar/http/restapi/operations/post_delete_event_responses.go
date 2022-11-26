// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ntnbrtnkv/otus-golang/calendar/http/models"
)

// PostDeleteEventOKCode is the HTTP code returned for type PostDeleteEventOK
const PostDeleteEventOKCode int = 200

/*
PostDeleteEventOK Deleted successfully

swagger:response postDeleteEventOK
*/
type PostDeleteEventOK struct {
}

// NewPostDeleteEventOK creates PostDeleteEventOK with default headers values
func NewPostDeleteEventOK() *PostDeleteEventOK {

	return &PostDeleteEventOK{}
}

// WriteResponse to the client
func (o *PostDeleteEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostDeleteEventBadRequestCode is the HTTP code returned for type PostDeleteEventBadRequest
const PostDeleteEventBadRequestCode int = 400

/*
PostDeleteEventBadRequest Failed to delete event

swagger:response postDeleteEventBadRequest
*/
type PostDeleteEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostDeleteEventBadRequest creates PostDeleteEventBadRequest with default headers values
func NewPostDeleteEventBadRequest() *PostDeleteEventBadRequest {

	return &PostDeleteEventBadRequest{}
}

// WithPayload adds the payload to the post delete event bad request response
func (o *PostDeleteEventBadRequest) WithPayload(payload *models.Error) *PostDeleteEventBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post delete event bad request response
func (o *PostDeleteEventBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDeleteEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}