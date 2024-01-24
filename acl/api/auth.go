package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	acl "github.com/ashok-aroha/ACL-Go/acl"
)

type Auth struct {
	UserID     string `json:"user_id,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	StatusCode int
	Message    string
}

func (a *Auth) GetDict() map[string]interface{} {
	authDict := map[string]interface{}{
		"email":    a.Email,
		"password": a.Password,
	}

	if a.UserID != "" {
		authDict["id"] = a.UserID
	}

	if a.FirstName != "" {
		authDict["first_name"] = a.FirstName
	}

	if a.LastName != "" {
		authDict["last_name"] = a.LastName
	}

	if a.StatusCode != 0 {
		authDict["status_code"] = a.StatusCode
	}

	if a.Message != "" {
		authDict["message"] = a.Message
	}

	return authDict
}

type Log struct{}

func (l *Log) PrintLog(obj interface{}) {
	// Implementation for logging goes here
}

type AuthAPI struct {
	Headers acl.Headers
}

func (aa *AuthAPI) Signup(email, password, firstName, lastName string) Auth {
	// Implementation for signup method goes here
	return Auth{}
}

func (aa *AuthAPI) Login(email, password string) Auth {
	authObj := Auth{
		Email:    email,
		Password: password,
	}

	url := acl.USER_LOGIN_API // Replace with your API URL

	payload, _ := json.Marshal(authObj.GetDict())

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	for key, value := range aa.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return Auth{}
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("User Found!!")

		var authData map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&authData); err != nil {
			fmt.Println("Error decoding response:", err)
			return Auth{}
		}

		id, idExists := authData["id"]
		if !idExists {
			fmt.Println("ID field missing in response")
			return Auth{}
		}

		var userID string
		switch id := id.(type) {
		case string:
			userID = id
		case float64:
			userID = fmt.Sprintf("%.0f", id)
		default:
			fmt.Println("ID field is not a recognized type")
			return Auth{}
		}

		authObj.UserID = userID
		authObj.Email, _ = authData["email"].(string)
		authObj.FirstName, _ = authData["first_name"].(string)
		authObj.LastName, _ = authData["last_name"].(string)
		authObj.StatusCode = resp.StatusCode
		authObj.Message = "User Found!!"

		return authObj
	}

	authObj.Message = "User Not Found!!"
	authObj.StatusCode = resp.StatusCode

	return authObj
}

func (aa *AuthAPI) IsAuthorized(roles, permissions []string, userID string) bool {
	url := acl.USER_ROLE_PERMISSIONS_API // Replace with your API URL

	payload, _ := json.Marshal(map[string]interface{}{"id": userID})

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	for key, value := range aa.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("_is_authorized: User Role Permissions Found!!")

		authData := map[string]interface{}{}
		json.NewDecoder(resp.Body).Decode(&authData)

		userRoles := authData["user_roles"].([]interface{})
		rolePermissions := authData["role_permissions"].(map[string]interface{})

		for _, role := range roles {
			if contains(userRoles, role) {
				permissionsForRole := rolePermissions[role].([]interface{})
				if hasAnyPermission(permissionsForRole, permissions) {
					return true
				}
			}
		}
	}

	return false
}

func contains(arr []interface{}, val interface{}) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func hasAnyPermission(arr []interface{}, permissions []string) bool {
	for _, permission := range permissions {
		if contains(arr, permission) {
			return true
		}
	}
	return false
}
