// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// docs endpoints
//
// Command:
// $ goa gen github.com/arduino/arduino-create-agent/design

package docs

import (
	goa "goa.design/goa"
)

// Endpoints wraps the "docs" service endpoints.
type Endpoints struct {
}

// NewEndpoints wraps the methods of the "docs" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{}
}

// Use applies the given middleware to all the "docs" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
}