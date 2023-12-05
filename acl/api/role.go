package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	acl "github.com/ashok-aroha/ACL-Go/acl"
)

type Role struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          int    `json:"id,omitempty"`
	StatusCode  int
	Message     string
}

func (r *Role) ToMap() map[string]interface{} {
	roleDict := map[string]interface{}{
		"name":        r.Name,
		"description": r.Description,
	}

	if r.ID != 0 {
		roleDict["id"] = r.ID
	}

	if r.StatusCode != 0 {
		roleDict["status_code"] = r.StatusCode
	}

	if r.Message != "" {
		roleDict["message"] = r.Message
	}

	return roleDict
}

type RolePermission struct {
	RoleID          int   `json:"role_id"`
	RolePermissions []int `json:"role_permissions"`
	ID              int   `json:"id,omitempty"`
	StatusCode      int
	Message         string
}

func (rp *RolePermission) ToMap() map[string]interface{} {
	rolePermDict := map[string]interface{}{
		"role_id":          rp.RoleID,
		"role_permissions": rp.RolePermissions,
	}

	if rp.ID != 0 {
		rolePermDict["id"] = rp.ID
	}

	if rp.StatusCode != 0 {
		rolePermDict["status_code"] = rp.StatusCode
	}

	if rp.Message != "" {
		rolePermDict["message"] = rp.Message
	}

	return rolePermDict
}

type RoleAPI struct {
	Headers acl.Headers
}

func (ra *RoleAPI) CreateRole(name, description string) Role {
	url := acl.CREATE_ROLE_API

	roleObj := Role{Name: name, Description: description}

	payload, _ := json.Marshal(roleObj.ToMap())

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	for key, value := range ra.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, _ := client.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Role Creation Successful!!")

		var roleRes Role
		json.NewDecoder(resp.Body).Decode(&roleRes)

		roleObj.ID = roleRes.ID
		roleObj.Message = "Role Created!!"
	} else {
		roleObj.Message = "Role Creation Failed!!"
	}

	roleObj.StatusCode = resp.StatusCode

	return roleObj
}

func (ra *RoleAPI) CreateRolePermissions(roleID int, rolePermissions []int) RolePermission {

	url := acl.CREATE_ROLE_PERMISSIONS_API

	rolePermObj := RolePermission{RoleID: roleID, RolePermissions: rolePermissions}

	payload, _ := json.Marshal(rolePermObj.ToMap())

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	for key, value := range ra.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, _ := client.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		var rolePermData RolePermission
		json.NewDecoder(resp.Body).Decode(&rolePermData)

		rolePermObj.ID = rolePermData.ID
		rolePermObj.Message = "Role Permission Created!!"
	} else {
		rolePermObj.Message = "Role Permission Creation Failed!!"
	}

	rolePermObj.StatusCode = resp.StatusCode

	return rolePermObj
}
