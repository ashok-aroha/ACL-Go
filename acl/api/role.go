package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	acl "github.com/ashok-aroha/ACL-GO/acl"
)


type RoleAPI struct {
	Role acl.Role
}

func NewRoleAPI(name string, description string) *RoleAPI {
	return &RoleAPI{
		Role: acl.Role{
			Name:        name,
			Description: description,
		},
	}
}