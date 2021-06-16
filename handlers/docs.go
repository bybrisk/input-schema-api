// Package classification of Input-Schema API
//
// Documentation for Input-Schema API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta

package handlers
import "github.com/bybrisk/input-schema-api/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// Success message on creating a bot schema for a business
// swagger:response createSchemaResponse
type createSchemaResponseWrapper struct {
	// Success message on creating a response schema
	// in: body
	Body data.CreateSchemaResponse
}

// swagger:parameters createSchema
type createSchemaParmsWrapper struct {
	// Data structure to create a bot schema for a business.
	// in: body
	// required: true
	Body data.CreateSchemaRequest
}