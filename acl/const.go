package acl

import (
	"fmt"
)

const (
	BaseURL                  = "http://127.0.0.1:8000"
	CSRFToken                = "OVwpp5rokWAkZ6OHFBXSKxoctE0uLFPl"
	ContentTypeHeader        = "Content-Type"
	ApplicationJSONMediaType = "application/json"
)

type Headers map[string]string;

var HeadersMap = Headers{
	ContentTypeHeader: ApplicationJSONMediaType,
	"Cookie":          fmt.Sprintf("csrftoken=%s", CSRFToken),
}

// User API endpoints
var (
	GetUserByIDAPI    = fmt.Sprintf("%s/get-user-by-id", BaseURL)
	GetUserRoleByID   = fmt.Sprintf("%s/get-user-roles-by-user-id", BaseURL)
	CreateUserAPI     = fmt.Sprintf("%s/create-user", BaseURL)
	CreateUserRoleAPI = fmt.Sprintf("%s/create-user-role", BaseURL)
)

// Auth API endpoints
var (
	UserLoginAPI           = fmt.Sprintf("%s/user-login", BaseURL)
	UserSignupAPI          = fmt.Sprintf("%s/create-user-role", BaseURL)
	UserRolePermissionsAPI = fmt.Sprintf("%s/get-user-with-role-permissions", BaseURL)
)

// Role API endpoints
var (
	CreateRoleAPI         = fmt.Sprintf("%s/create-role", BaseURL)
	CreateRolePermissions = fmt.Sprintf("%s/create-role-permissions", BaseURL)
)

// Permission API endpoints
var (
	CreatePermissionAPI = fmt.Sprintf("%s/create-permission", BaseURL)
)
