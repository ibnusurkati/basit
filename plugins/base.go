package plugins

import "github.com/ibnusurkati/basit/context"

// Plugin parse user defined request
type Setup interface {

	// Apply the plugin to http request
	Apply(*context.Context)

	// Valid is the plugin need run?
	Valid() bool
}
