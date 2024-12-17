// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Swagger user management service - OpenAPI 3.0
 *
 * This is a sample User management server based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.0
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// PersonsAPIController binds http requests to an api service and writes the service results to the http response
type PersonsAPIController struct {
	service      PersonsAPIServicer
	errorHandler ErrorHandler
}

// PersonsAPIOption for how the controller is set up.
type PersonsAPIOption func(*PersonsAPIController)

// WithPersonsAPIErrorHandler inject ErrorHandler into controller
func WithPersonsAPIErrorHandler(h ErrorHandler) PersonsAPIOption {
	return func(c *PersonsAPIController) {
		c.errorHandler = h
	}
}

// NewPersonsAPIController creates a default api controller
func NewPersonsAPIController(s PersonsAPIServicer, opts ...PersonsAPIOption) *PersonsAPIController {
	controller := &PersonsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PersonsAPIController
func (c *PersonsAPIController) Routes() Routes {
	return Routes{
		"GetAllPersons": Route{
			strings.ToUpper("Get"),
			"/api/v1/persons",
			c.GetAllPersons,
		},
		"GetPersonById": Route{
			strings.ToUpper("Get"),
			"/api/v1/persons/{id}",
			c.GetPersonById,
		},
		"EditPersonByID": Route{
			strings.ToUpper("Put"),
			"/api/v1/persons/{id}",
			c.EditPersonByID,
		},
		"DeletePersonByID": Route{
			strings.ToUpper("Delete"),
			"/api/v1/persons/{id}",
			c.DeletePersonByID,
		},
	}
}

// GetAllPersons - Provides a list of persons
func (c *PersonsAPIController) GetAllPersons(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetAllPersons(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetPersonById - Find person by ID
func (c *PersonsAPIController) GetPersonById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &RequiredError{"id"}, nil)
		return
	}
	result, err := c.service.GetPersonById(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// EditPersonByID - edit person by ID
func (c *PersonsAPIController) EditPersonByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &RequiredError{"id"}, nil)
		return
	}
	personParam := Person{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&personParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPersonRequired(personParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPersonConstraints(personParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.EditPersonByID(r.Context(), idParam, personParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeletePersonByID - delete person by ID
func (c *PersonsAPIController) DeletePersonByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &RequiredError{"id"}, nil)
		return
	}
	result, err := c.service.DeletePersonByID(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}