// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostCreateEventCreatedCode is the HTTP code returned for type PostCreateEventCreated
const PostCreateEventCreatedCode int = 201

/*
PostCreateEventCreated created successfully

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
