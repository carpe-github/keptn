// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutProjectProjectNameStageStageNameServiceServiceNameHandlerFunc turns a function with the right signature into a put project project name stage stage name service service name handler
type PutProjectProjectNameStageStageNameServiceServiceNameHandlerFunc func(PutProjectProjectNameStageStageNameServiceServiceNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutProjectProjectNameStageStageNameServiceServiceNameHandlerFunc) Handle(params PutProjectProjectNameStageStageNameServiceServiceNameParams) middleware.Responder {
	return fn(params)
}

// PutProjectProjectNameStageStageNameServiceServiceNameHandler interface for that can handle valid put project project name stage stage name service service name params
type PutProjectProjectNameStageStageNameServiceServiceNameHandler interface {
	Handle(PutProjectProjectNameStageStageNameServiceServiceNameParams) middleware.Responder
}

// NewPutProjectProjectNameStageStageNameServiceServiceName creates a new http.Handler for the put project project name stage stage name service service name operation
func NewPutProjectProjectNameStageStageNameServiceServiceName(ctx *middleware.Context, handler PutProjectProjectNameStageStageNameServiceServiceNameHandler) *PutProjectProjectNameStageStageNameServiceServiceName {
	return &PutProjectProjectNameStageStageNameServiceServiceName{Context: ctx, Handler: handler}
}

/* PutProjectProjectNameStageStageNameServiceServiceName swagger:route PUT /project/{projectName}/stage/{stageName}/service/{serviceName} Service putProjectProjectNameStageStageNameServiceServiceName

INTERNAL Endpoint: Update the specified service

*/
type PutProjectProjectNameStageStageNameServiceServiceName struct {
	Context *middleware.Context
	Handler PutProjectProjectNameStageStageNameServiceServiceNameHandler
}

func (o *PutProjectProjectNameStageStageNameServiceServiceName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutProjectProjectNameStageStageNameServiceServiceNameParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
