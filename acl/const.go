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

type Headers map[string]string

var HeadersMap = Headers{
	ContentTypeHeader: ApplicationJSONMediaType,
	"Cookie":          fmt.Sprintf("csrftoken=%s", CSRFToken),
}

// User API endpoints
var (
	GET_USER_BY_ID_API      = fmt.Sprintf("%s/get-user-by-id", BaseURL)
	GET_USER_ROLE_BY_ID_API = fmt.Sprintf("%s/get-user-roles-by-user-id", BaseURL)
	CREATE_USER_API         = fmt.Sprintf("%s/create-user", BaseURL)
	CREATE_USER_ROLE_API    = fmt.Sprintf("%s/create-user-role", BaseURL)
)

// Auth API endpoints
var (
	USER_LOGIN_API            = fmt.Sprintf("%s/user-login", BaseURL)
	USER_SIGNUP_API           = fmt.Sprintf("%s/create-user-role", BaseURL)
	USER_ROLE_PERMISSIONS_API = fmt.Sprintf("%s/get-user-with-role-permissions", BaseURL)
)

// Role API endpoints
var (
	CREATE_ROLE_API             = fmt.Sprintf("%s/create-role", BaseURL)
	CREATE_ROLE_PERMISSIONS_API = fmt.Sprintf("%s/create-role-permissions", BaseURL)
)

// Permission API endpoints
var (
	CREATE_PERMISSION_API = fmt.Sprintf("%s/create-permission", BaseURL)
)
