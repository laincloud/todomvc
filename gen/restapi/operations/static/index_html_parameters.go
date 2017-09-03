// Code generated by go-swagger; DO NOT EDIT.

package static

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewIndexHTMLParams creates a new IndexHTMLParams object
// with the default values initialized.
func NewIndexHTMLParams() IndexHTMLParams {
	var ()
	return IndexHTMLParams{}
}

// IndexHTMLParams contains all the bound params for the index HTML operation
// typically these are obtained from a http.Request
//
// swagger:parameters indexHTML
type IndexHTMLParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *IndexHTMLParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
