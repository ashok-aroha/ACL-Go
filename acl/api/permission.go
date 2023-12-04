package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	acl "github.com/ashok-aroha/ACL-Go/acl"
)

type PermissionAPI struct {
	Permission acl.Permission
}

func NewPermissionAPI(name, description string) *PermissionAPI {
	return &PermissionAPI{
		Permission: acl.Permission{
			Name:        name,
			Description: description,
		},
	}
}

func (p *PermissionAPI) CreatePermission() map[string]interface{} {
	url := acl.CreatePermissionAPI
	headers := acl.Headers

	payload, _ := json.Marshal(p.Permission)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Permission Created!!")
		return result
	}

	fmt.Println("Permission Creation Failed!!")
	return result
}
