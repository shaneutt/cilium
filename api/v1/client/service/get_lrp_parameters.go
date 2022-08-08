// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package service

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

// NewGetLrpParams creates a new GetLrpParams object
// with the default values initialized.
func NewGetLrpParams() *GetLrpParams {

	return &GetLrpParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLrpParamsWithTimeout creates a new GetLrpParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLrpParamsWithTimeout(timeout time.Duration) *GetLrpParams {

	return &GetLrpParams{

		timeout: timeout,
	}
}

// NewGetLrpParamsWithContext creates a new GetLrpParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLrpParamsWithContext(ctx context.Context) *GetLrpParams {

	return &GetLrpParams{

		Context: ctx,
	}
}

// NewGetLrpParamsWithHTTPClient creates a new GetLrpParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLrpParamsWithHTTPClient(client *http.Client) *GetLrpParams {

	return &GetLrpParams{
		HTTPClient: client,
	}
}

/*
GetLrpParams contains all the parameters to send to the API endpoint
for the get lrp operation typically these are written to a http.Request
*/
type GetLrpParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lrp params
func (o *GetLrpParams) WithTimeout(timeout time.Duration) *GetLrpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lrp params
func (o *GetLrpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lrp params
func (o *GetLrpParams) WithContext(ctx context.Context) *GetLrpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lrp params
func (o *GetLrpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lrp params
func (o *GetLrpParams) WithHTTPClient(client *http.Client) *GetLrpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lrp params
func (o *GetLrpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLrpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
