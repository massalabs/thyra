// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// MgmtWalletImportHandlerFunc turns a function with the right signature into a mgmt wallet import handler
type MgmtWalletImportHandlerFunc func(MgmtWalletImportParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MgmtWalletImportHandlerFunc) Handle(params MgmtWalletImportParams) middleware.Responder {
	return fn(params)
}

// MgmtWalletImportHandler interface for that can handle valid mgmt wallet import params
type MgmtWalletImportHandler interface {
	Handle(MgmtWalletImportParams) middleware.Responder
}

// NewMgmtWalletImport creates a new http.Handler for the mgmt wallet import operation
func NewMgmtWalletImport(ctx *middleware.Context, handler MgmtWalletImportHandler) *MgmtWalletImport {
	return &MgmtWalletImport{Context: ctx, Handler: handler}
}

/* MgmtWalletImport swagger:route PUT /mgmt/wallet mgmtWalletImport

MgmtWalletImport mgmt wallet import API

*/
type MgmtWalletImport struct {
	Context *middleware.Context
	Handler MgmtWalletImportHandler
}

func (o *MgmtWalletImport) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewMgmtWalletImportParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}