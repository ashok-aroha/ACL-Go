package main

import (
	"fmt"

	// acl "github.com/ashok-aroha/ACL-Go/acl"
	api "github.com/ashok-aroha/ACL-Go/acl/api"
)

func main() {
	// Permission api
	permissionAPI := api.NewPermissionAPI("Test Permission", "Description of Test Permission")
	createdPermission := permissionAPI.CreatePermission()
	fmt.Println(createdPermission)
}

// func main() {
// 	log := acl.Log{}

// 	// Example usage: Creating a Role
// 	newRole := api.CreateRole{Name: "Admin", Description: "Admin Role"}
// 	log.PrintLog(newRole)

// 	// Example usage: Creating Role Permissions
// 	permissions := []int{1, 2, 3} // Example permissions
// 	newRolePermission := acl.RolePermission{RoleID: 123, RolePermissionsIds: permissions}
// 	log.PrintLog(newRolePermission)
// }
