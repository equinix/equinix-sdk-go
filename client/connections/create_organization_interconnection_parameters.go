// Code generated by go-swagger; DO NOT EDIT.

package connections

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

	"github.com/t0mk/gometal/models"
)

// NewCreateOrganizationInterconnectionParams creates a new CreateOrganizationInterconnectionParams object
// with the default values initialized.
func NewCreateOrganizationInterconnectionParams() *CreateOrganizationInterconnectionParams {
	var ()
	return &CreateOrganizationInterconnectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrganizationInterconnectionParamsWithTimeout creates a new CreateOrganizationInterconnectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateOrganizationInterconnectionParamsWithTimeout(timeout time.Duration) *CreateOrganizationInterconnectionParams {
	var ()
	return &CreateOrganizationInterconnectionParams{

		timeout: timeout,
	}
}

// NewCreateOrganizationInterconnectionParamsWithContext creates a new CreateOrganizationInterconnectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateOrganizationInterconnectionParamsWithContext(ctx context.Context) *CreateOrganizationInterconnectionParams {
	var ()
	return &CreateOrganizationInterconnectionParams{

		Context: ctx,
	}
}

// NewCreateOrganizationInterconnectionParamsWithHTTPClient creates a new CreateOrganizationInterconnectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateOrganizationInterconnectionParamsWithHTTPClient(client *http.Client) *CreateOrganizationInterconnectionParams {
	var ()
	return &CreateOrganizationInterconnectionParams{
		HTTPClient: client,
	}
}

/*CreateOrganizationInterconnectionParams contains all the parameters to send to the API endpoint
for the create organization interconnection operation typically these are written to a http.Request
*/
type CreateOrganizationInterconnectionParams struct {

	/*Connection
	  Connection details

	*/
	Connection *models.InterconnectionCreateInput
	/*OrganizationID
	  UUID of the organization

	*/
	OrganizationID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create organization interconnection params
func (o *CreateOrganizationInterconnectionParams) WithTimeout(timeout time.Duration) *CreateOrganizationInterconnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create organization interconnection params
func (o *CreateOrganizationInterconnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create organization interconnection params
func (o *CreateOrganizationInterconnectionParams) WithContext(ctx context.Context) *CreateOrganizationInterconnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create organization interconnection params
func (o *CreateOrganizationInterconnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create organization interconnection params
func (o *CreateOrganizationInterconnectionParams) WithHTTPClient(client *http.Client) *CreateOrganizationInterconnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create organization interconnection params
func (o *CreateOrganizationInterconnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnection adds the connection to the create organization interconnection params
func (o *CreateOrganizationInterconnectionParams) WithConnection(connection *models.InterconnectionCreateInput) *CreateOrganizationInterconnectionParams {
	o.SetConnection(connection)
	return o
}

// SetConnection adds the connection to the create organization interconnection params
func (o *CreateOrganizationInterconnectionParams) SetConnection(connection *models.InterconnectionCreateInput) {
	o.Connection = connection
}

// WithOrganizationID adds the organizationID to the create organization interconnection params
func (o *CreateOrganizationInterconnectionParams) WithOrganizationID(organizationID strfmt.UUID) *CreateOrganizationInterconnectionParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create organization interconnection params
func (o *CreateOrganizationInterconnectionParams) SetOrganizationID(organizationID strfmt.UUID) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrganizationInterconnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Connection != nil {
		if err := r.SetBodyParam(o.Connection); err != nil {
			return err
		}
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}