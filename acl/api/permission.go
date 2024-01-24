package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	acl "github.com/ashok-aroha/ACL-Go/acl"
)

type Permission struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          int    `json:"id,omitempty"`
	StatusCode  int
	Message     string
}

func (p *Permission) ToMap() map[string]interface{} {
	permissionDict := map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
	}

	if p.ID != 0 {
		permissionDict["id"] = p.ID
	}

	if p.StatusCode != 0 {
		permissionDict["status_code"] = p.StatusCode
	}

	if p.Message != "" {
		permissionDict["message"] = p.Message
	}

	return permissionDict
}

type PermissionAPI struct {
	Headers acl.Headers
}

func (p *PermissionAPI) CreatePermission(name, description string) Permission {
	url := acl.CREATE_PERMISSION_API // Replace with your API URL

	permissionObj := Permission{Name: name, Description: description}

	payload, _ := json.Marshal(permissionObj.ToMap())

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	for key, value := range p.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return Permission{}
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		var permissionRes Permission
		json.NewDecoder(resp.Body).Decode(&permissionRes)

		permissionObj.ID = permissionRes.ID
		//TODO: API Message has writtent custom message need to be udpated into server side
		permissionObj.Message = "Permission Creation Successful!!"
	} else {
		permissionObj.Message = "Permission Creation Failed!!"
	}

	permissionObj.StatusCode = resp.StatusCode

	return permissionObj
}
