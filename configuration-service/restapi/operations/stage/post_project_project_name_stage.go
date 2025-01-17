// Code generated by go-swagger; DO NOT EDIT.

package stage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostProjectProjectNameStageHandlerFunc turns a function with the right signature into a post project project name stage handler
type PostProjectProjectNameStageHandlerFunc func(PostProjectProjectNameStageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostProjectProjectNameStageHandlerFunc) Handle(params PostProjectProjectNameStageParams) middleware.Responder {
	return fn(params)
}

// PostProjectProjectNameStageHandler interface for that can handle valid post project project name stage params
type PostProjectProjectNameStageHandler interface {
	Handle(PostProjectProjectNameStageParams) middleware.Responder
}

// NewPostProjectProjectNameStage creates a new http.Handler for the post project project name stage operation
func NewPostProjectProjectNameStage(ctx *middleware.Context, handler PostProjectProjectNameStageHandler) *PostProjectProjectNameStage {
	return &PostProjectProjectNameStage{Context: ctx, Handler: handler}
}

/* PostProjectProjectNameStage swagger:route POST /project/{projectName}/stage Stage postProjectProjectNameStage

INTERNAL Endpoint: Create a new stage by stage name

*/
type PostProjectProjectNameStage struct {
	Context *middleware.Context
	Handler PostProjectProjectNameStageHandler
}

func (o *PostProjectProjectNameStage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostProjectProjectNameStageParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
