// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kata-containers/kata-containers/src/runtime/virtcontainers/pkg/firecracker/client/models"
)

// NewPatchBalloonParams creates a new PatchBalloonParams object
// with the default values initialized.
func NewPatchBalloonParams() *PatchBalloonParams {
	var ()
	return &PatchBalloonParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchBalloonParamsWithTimeout creates a new PatchBalloonParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchBalloonParamsWithTimeout(timeout time.Duration) *PatchBalloonParams {
	var ()
	return &PatchBalloonParams{

		timeout: timeout,
	}
}

// NewPatchBalloonParamsWithContext creates a new PatchBalloonParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchBalloonParamsWithContext(ctx context.Context) *PatchBalloonParams {
	var ()
	return &PatchBalloonParams{

		Context: ctx,
	}
}

// NewPatchBalloonParamsWithHTTPClient creates a new PatchBalloonParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchBalloonParamsWithHTTPClient(client *http.Client) *PatchBalloonParams {
	var ()
	return &PatchBalloonParams{
		HTTPClient: client,
	}
}

/*PatchBalloonParams contains all the parameters to send to the API endpoint
for the patch balloon operation typically these are written to a http.Request
*/
type PatchBalloonParams struct {

	/*Body
	  Balloon properties

	*/
	Body *models.BalloonUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch balloon params
func (o *PatchBalloonParams) WithTimeout(timeout time.Duration) *PatchBalloonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch balloon params
func (o *PatchBalloonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch balloon params
func (o *PatchBalloonParams) WithContext(ctx context.Context) *PatchBalloonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch balloon params
func (o *PatchBalloonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch balloon params
func (o *PatchBalloonParams) WithHTTPClient(client *http.Client) *PatchBalloonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch balloon params
func (o *PatchBalloonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch balloon params
func (o *PatchBalloonParams) WithBody(body *models.BalloonUpdate) *PatchBalloonParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch balloon params
func (o *PatchBalloonParams) SetBody(body *models.BalloonUpdate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchBalloonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
