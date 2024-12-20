// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Swagger user management service - OpenAPI 3.0
 *
 * This is a sample User management server based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.0
 */

package openapi

type User struct {
	Id string `json:"Id,omitempty"`

	PersonId string `json:"personId,omitempty"`

	UserName string `json:"userName,omitempty"`

	UserGroupId string `json:"userGroupId,omitempty"`

	Active bool `json:"active,omitempty"`

	Deleted bool `json:"deleted,omitempty"`
}

// AssertUserRequired checks if the required fields are not zero-ed
func AssertUserRequired(obj User) error {
	return nil
}

// AssertUserConstraints checks if the values respects the defined constraints
func AssertUserConstraints(obj User) error {
	return nil
}
