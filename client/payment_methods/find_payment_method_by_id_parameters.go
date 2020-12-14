// Code generated by go-swagger; DO NOT EDIT.

package payment_methods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewFindPaymentMethodByIDParams creates a new FindPaymentMethodByIDParams object
// with the default values initialized.
func NewFindPaymentMethodByIDParams() *FindPaymentMethodByIDParams {
	var ()
	return &FindPaymentMethodByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindPaymentMethodByIDParamsWithTimeout creates a new FindPaymentMethodByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindPaymentMethodByIDParamsWithTimeout(timeout time.Duration) *FindPaymentMethodByIDParams {
	var ()
	return &FindPaymentMethodByIDParams{

		timeout: timeout,
	}
}

// NewFindPaymentMethodByIDParamsWithContext creates a new FindPaymentMethodByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindPaymentMethodByIDParamsWithContext(ctx context.Context) *FindPaymentMethodByIDParams {
	var ()
	return &FindPaymentMethodByIDParams{

		Context: ctx,
	}
}

// NewFindPaymentMethodByIDParamsWithHTTPClient creates a new FindPaymentMethodByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindPaymentMethodByIDParamsWithHTTPClient(client *http.Client) *FindPaymentMethodByIDParams {
	var ()
	return &FindPaymentMethodByIDParams{
		HTTPClient: client,
	}
}

/*FindPaymentMethodByIDParams contains all the parameters to send to the API endpoint
for the find payment method by Id operation typically these are written to a http.Request
*/
type FindPaymentMethodByIDParams struct {

	/*ID
	  Payment Method UUID

	*/
	ID strfmt.UUID
	/*Include
	  related attributes to include

	*/
	Include *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find payment method by Id params
func (o *FindPaymentMethodByIDParams) WithTimeout(timeout time.Duration) *FindPaymentMethodByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find payment method by Id params
func (o *FindPaymentMethodByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find payment method by Id params
func (o *FindPaymentMethodByIDParams) WithContext(ctx context.Context) *FindPaymentMethodByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find payment method by Id params
func (o *FindPaymentMethodByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find payment method by Id params
func (o *FindPaymentMethodByIDParams) WithHTTPClient(client *http.Client) *FindPaymentMethodByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find payment method by Id params
func (o *FindPaymentMethodByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find payment method by Id params
func (o *FindPaymentMethodByIDParams) WithID(id strfmt.UUID) *FindPaymentMethodByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find payment method by Id params
func (o *FindPaymentMethodByIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithInclude adds the include to the find payment method by Id params
func (o *FindPaymentMethodByIDParams) WithInclude(include *string) *FindPaymentMethodByIDParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the find payment method by Id params
func (o *FindPaymentMethodByIDParams) SetInclude(include *string) {
	o.Include = include
}

// WriteToRequest writes these params to a swagger request
func (o *FindPaymentMethodByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.Include != nil {

		// query param include
		var qrInclude string
		if o.Include != nil {
			qrInclude = *o.Include
		}
		qInclude := qrInclude
		if qInclude != "" {
			if err := r.SetQueryParam("include", qInclude); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}