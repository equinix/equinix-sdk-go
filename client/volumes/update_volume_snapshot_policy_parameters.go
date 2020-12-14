// Code generated by go-swagger; DO NOT EDIT.

package volumes

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
	"github.com/go-openapi/swag"
)

// NewUpdateVolumeSnapshotPolicyParams creates a new UpdateVolumeSnapshotPolicyParams object
// with the default values initialized.
func NewUpdateVolumeSnapshotPolicyParams() *UpdateVolumeSnapshotPolicyParams {
	var ()
	return &UpdateVolumeSnapshotPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVolumeSnapshotPolicyParamsWithTimeout creates a new UpdateVolumeSnapshotPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateVolumeSnapshotPolicyParamsWithTimeout(timeout time.Duration) *UpdateVolumeSnapshotPolicyParams {
	var ()
	return &UpdateVolumeSnapshotPolicyParams{

		timeout: timeout,
	}
}

// NewUpdateVolumeSnapshotPolicyParamsWithContext creates a new UpdateVolumeSnapshotPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateVolumeSnapshotPolicyParamsWithContext(ctx context.Context) *UpdateVolumeSnapshotPolicyParams {
	var ()
	return &UpdateVolumeSnapshotPolicyParams{

		Context: ctx,
	}
}

// NewUpdateVolumeSnapshotPolicyParamsWithHTTPClient creates a new UpdateVolumeSnapshotPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateVolumeSnapshotPolicyParamsWithHTTPClient(client *http.Client) *UpdateVolumeSnapshotPolicyParams {
	var ()
	return &UpdateVolumeSnapshotPolicyParams{
		HTTPClient: client,
	}
}

/*UpdateVolumeSnapshotPolicyParams contains all the parameters to send to the API endpoint
for the update volume snapshot policy operation typically these are written to a http.Request
*/
type UpdateVolumeSnapshotPolicyParams struct {

	/*ID
	  Snapshot Policy UUID

	*/
	ID strfmt.UUID
	/*SnapshotCount
	  Snapshot count

	*/
	SnapshotCount *int64
	/*SnapshotFrequency
	  Snapshot frequency

	*/
	SnapshotFrequency string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update volume snapshot policy params
func (o *UpdateVolumeSnapshotPolicyParams) WithTimeout(timeout time.Duration) *UpdateVolumeSnapshotPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update volume snapshot policy params
func (o *UpdateVolumeSnapshotPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update volume snapshot policy params
func (o *UpdateVolumeSnapshotPolicyParams) WithContext(ctx context.Context) *UpdateVolumeSnapshotPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update volume snapshot policy params
func (o *UpdateVolumeSnapshotPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update volume snapshot policy params
func (o *UpdateVolumeSnapshotPolicyParams) WithHTTPClient(client *http.Client) *UpdateVolumeSnapshotPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update volume snapshot policy params
func (o *UpdateVolumeSnapshotPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update volume snapshot policy params
func (o *UpdateVolumeSnapshotPolicyParams) WithID(id strfmt.UUID) *UpdateVolumeSnapshotPolicyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update volume snapshot policy params
func (o *UpdateVolumeSnapshotPolicyParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithSnapshotCount adds the snapshotCount to the update volume snapshot policy params
func (o *UpdateVolumeSnapshotPolicyParams) WithSnapshotCount(snapshotCount *int64) *UpdateVolumeSnapshotPolicyParams {
	o.SetSnapshotCount(snapshotCount)
	return o
}

// SetSnapshotCount adds the snapshotCount to the update volume snapshot policy params
func (o *UpdateVolumeSnapshotPolicyParams) SetSnapshotCount(snapshotCount *int64) {
	o.SnapshotCount = snapshotCount
}

// WithSnapshotFrequency adds the snapshotFrequency to the update volume snapshot policy params
func (o *UpdateVolumeSnapshotPolicyParams) WithSnapshotFrequency(snapshotFrequency string) *UpdateVolumeSnapshotPolicyParams {
	o.SetSnapshotFrequency(snapshotFrequency)
	return o
}

// SetSnapshotFrequency adds the snapshotFrequency to the update volume snapshot policy params
func (o *UpdateVolumeSnapshotPolicyParams) SetSnapshotFrequency(snapshotFrequency string) {
	o.SnapshotFrequency = snapshotFrequency
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVolumeSnapshotPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.SnapshotCount != nil {

		// query param snapshot_count
		var qrSnapshotCount int64
		if o.SnapshotCount != nil {
			qrSnapshotCount = *o.SnapshotCount
		}
		qSnapshotCount := swag.FormatInt64(qrSnapshotCount)
		if qSnapshotCount != "" {
			if err := r.SetQueryParam("snapshot_count", qSnapshotCount); err != nil {
				return err
			}
		}

	}

	// query param snapshot_frequency
	qrSnapshotFrequency := o.SnapshotFrequency
	qSnapshotFrequency := qrSnapshotFrequency
	if qSnapshotFrequency != "" {
		if err := r.SetQueryParam("snapshot_frequency", qSnapshotFrequency); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}