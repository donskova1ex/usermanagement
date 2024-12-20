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
	"context"
	"errors"
	"net/http"
)

// UserRolesAPIService is a service that implements the logic for the UserRolesAPIServicer
// This service should implement the business logic for every endpoint for the UserRolesAPI API.
// Include any external packages or services that will be required by this service.
type UserRolesAPIService struct {
}

// NewUserRolesAPIService creates a default api service
func NewUserRolesAPIService() *UserRolesAPIService {
	return &UserRolesAPIService{}
}

// GetUserRolesByUserId - Find user-roles by userID
func (s *UserRolesAPIService) GetUserRolesByUserId(ctx context.Context, userid string) (ImplResponse, error) {
	// TODO - update GetUserRolesByUserId with the required logic for this service method.
	// Add api_user_roles_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []UserRole{}) or use other options such as http.Ok ...
	// return Response(200, []UserRole{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUserRolesByUserId method not implemented")
}

// EditUserRolesByUserID - edit user role by userID
func (s *UserRolesAPIService) EditUserRolesByUserID(ctx context.Context, userid string, userRole []UserRole) (ImplResponse, error) {
	// TODO - update EditUserRolesByUserID with the required logic for this service method.
	// Add api_user_roles_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, User{}) or use other options such as http.Ok ...
	// return Response(200, User{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("EditUserRolesByUserID method not implemented")
}
