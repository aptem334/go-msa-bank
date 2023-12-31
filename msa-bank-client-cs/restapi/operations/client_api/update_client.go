// Code generated by go-swagger; DO NOT EDIT.

package client_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateClientHandlerFunc turns a function with the right signature into a update client handler
type UpdateClientHandlerFunc func(UpdateClientParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateClientHandlerFunc) Handle(params UpdateClientParams) middleware.Responder {
	return fn(params)
}

// UpdateClientHandler interface for that can handle valid update client params
type UpdateClientHandler interface {
	Handle(UpdateClientParams) middleware.Responder
}

// NewUpdateClient creates a new http.Handler for the update client operation
func NewUpdateClient(ctx *middleware.Context, handler UpdateClientHandler) *UpdateClient {
	return &UpdateClient{Context: ctx, Handler: handler}
}

/*
	UpdateClient swagger:route PUT /client client-api updateClient

Update an existing client
*/
type UpdateClient struct {
	Context *middleware.Context
	Handler UpdateClientHandler
}

func (o *UpdateClient) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateClientParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
