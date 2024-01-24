package acl

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	ContentTypeHeader        = "Content-Type"
	ApplicationJSONMediaType = "application/json"
)

var _ = godotenv.Load()

type Headers map[string]string

var HeadersMap = Headers{
	ContentTypeHeader: ApplicationJSONMediaType,
	"Cookie":          fmt.Sprintf("csrftoken=%s", os.Getenv("CSRFToken")),
}

// User API endpoints
var (
	GET_USER_BY_ID_API      = os.Getenv("GET_USER_BY_ID_API")
	GET_USER_ROLE_BY_ID_API = os.Getenv("GET_USER_ROLE_BY_ID_API")
	CREATE_USER_API         = os.Getenv("CREATE_USER_API")
	CREATE_USER_ROLE_API    = os.Getenv("CREATE_USER_ROLE_API")
)

// Auth API endpoints
var (
	USER_LOGIN_API            = os.Getenv("USER_LOGIN_API")
	USER_SIGNUP_API           = os.Getenv("USER_SIGNUP_API")
	USER_ROLE_PERMISSIONS_API = os.Getenv("USER_ROLE_PERMISSIONS_API")
)

// Role API endpoints
var (
	CREATE_ROLE_API             = os.Getenv("CREATE_ROLE_API")
	CREATE_ROLE_PERMISSIONS_API = os.Getenv("CREATE_ROLE_PERMISSIONS_API")
)

// Permission API endpoints
var (
	CREATE_PERMISSION_API = os.Getenv("CREATE_PERMISSION_API")
)
