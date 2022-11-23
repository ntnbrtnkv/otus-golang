// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostCreateEventHandlerFunc turns a function with the right signature into a post create event handler
type PostCreateEventHandlerFunc func(PostCreateEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostCreateEventHandlerFunc) Handle(params PostCreateEventParams) middleware.Responder {
	return fn(params)
}

// PostCreateEventHandler interface for that can handle valid post create event params
type PostCreateEventHandler interface {
	Handle(PostCreateEventParams) middleware.Responder
}

// NewPostCreateEvent creates a new http.Handler for the post create event operation
func NewPostCreateEvent(ctx *middleware.Context, handler PostCreateEventHandler) *PostCreateEvent {
	return &PostCreateEvent{Context: ctx, Handler: handler}
}

/*
	PostCreateEvent swagger:route POST /create_event postCreateEvent

Creates event
*/
type PostCreateEvent struct {
	Context *middleware.Context
	Handler PostCreateEventHandler
}

func (o *PostCreateEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostCreateEventParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
