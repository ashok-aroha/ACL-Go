package main

import (
	"fmt"

	acl "github.com/ashok-aroha/ACL-Go/acl"
	api "github.com/ashok-aroha/ACL-Go/acl/api"
)

// To test create permission
// func main() {
// 	log := acl.Log()
// 	permissionAPI := api.PermissionAPI{}
// 	permissionAPI.CreatePermission("example_four_permission", "Example description")
// 	log.PrintLog(permissionAPI)

// }

// To test create role and role permissions
// func main() {

// 	log := acl.Log{}
// 	roleAPI := api.RoleAPI{}

// 	//Creating a Role
// 	createdRole := roleAPI.CreateRole("Example Role", "This is an example role.")
// 	log.PrintLog(createdRole)

// 	// Creating Role Permissions
// 	permissions := []int{1, 2, 3}                                           // Replace with your permissions
// 	createdRolePermissions := roleAPI.CreateRolePermissions(4, permissions) // Assuming role ID 1
// 	log.PrintLog(createdRolePermissions)

// }

// func main() {

// 	log := acl.Log{}
// 	userAPI := api.UserAPI{}

// 	// TODO: Testing Pending of this API (New task created by Ashok: https://github.com/ashok-aroha/ACL-Go/issues/1)
// 	// Example usage of methods
// 	userByID := userAPI.GetUserByID(1) // Replace 123 with the actual user ID
// 	log.PrintLog(userByID)
// 	fmt.Println("User by ID:", userByID)

// 	userRoleByID := userAPI.GetUserRoleByID(1) // Replace 123 with the actual user ID
// 	log.PrintLog(userRoleByID)
// 	fmt.Println("User Role by ID:", userRoleByID.ToMap())

// 	// TODO: API is working but need handle proper response (https://trello.com/c/xzq3FFsT/60-api-create-user-api)
// 	createdUser := userAPI.CreateUser("example_one@example.com", "password", "John", "Doe")
// 	log.PrintLog(createdUser)

// 	createdUserRole := userAPI.CreateUserRole(1, []int{6}) // Replace 123 with the actual user ID and roles accordingly
// 	log.PrintLog(createdUserRole)
// }

func main() {

	log := acl.Log{} // Initialize your logging mechanism here

	authAPI := api.AuthAPI{}

	// Example: Sign up a user
	// authAPI.Signup("example@example.com", "password123", "John", "Doe")

	// Example: Login and check authorization
	auth := authAPI.Login("ashok@gmail.com", "ashok")
	log.PrintLog(auth)
	fmt.Println("User ID:", auth.UserID)

	isAuthorized := authAPI.IsAuthorized([]string{"Project Manager"}, []string{"write", "read"}, auth.UserID)

	if isAuthorized {
		println("User is authorized!")
	} else {
		println("User is not authorized.")
	}
}
