package main

import (
	api "github.com/ashok-aroha/ACL-Go/acl/api"
)

func main() {
	permissionAPI := api.PermissionAPI{}
	permissionAPI.CreatePermission("example_four_permission", "Example description")
}
