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

// UserGroupsAPIService is a service that implements the logic for the UserGroupsAPIServicer
// This service should implement the business logic for every endpoint for the UserGroupsAPI API.
// Include any external packages or services that will be required by this service.
type UserGroupsAPIService struct {
}

// NewUserGroupsAPIService creates a default api service
func NewUserGroupsAPIService() *UserGroupsAPIService {
	return &UserGroupsAPIService{}
}

// GetAllUserGroups - Provides a list of user groups
func (s *UserGroupsAPIService) GetAllUserGroups(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetAllUserGroups with the required logic for this service method.
	// Add api_user_groups_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []UserGroups{}) or use other options such as http.Ok ...
	// return Response(200, []UserGroups{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetAllUserGroups method not implemented")
}

// GetUserGroupsByUserId - Find user-groups by userID
func (s *UserGroupsAPIService) GetUserGroupsByUserId(ctx context.Context, userid string) (ImplResponse, error) {
	// TODO - update GetUserGroupsByUserId with the required logic for this service method.
	// Add api_user_groups_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []UserGroups{}) or use other options such as http.Ok ...
	// return Response(200, []UserGroups{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUserGroupsByUserId method not implemented")
}

// EditUserGroupByUserID - edit user group by userID
func (s *UserGroupsAPIService) EditUserGroupByUserID(ctx context.Context, userid string, userGroups []UserGroups) (ImplResponse, error) {
	// TODO - update EditUserGroupByUserID with the required logic for this service method.
	// Add api_user_groups_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []UserGroups{}) or use other options such as http.Ok ...
	// return Response(200, []UserGroups{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("EditUserGroupByUserID method not implemented")
}
