package acl

import (
	"fmt"
)

type User struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	UserID     int    `json:"id,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (u *User) ToMap() map[string]interface{} {
	userMap := map[string]interface{}{
		"email":    u.Email,
		"password": u.Password,
	}

	if u.UserID != 0 {
		userMap["id"] = u.UserID
	}
	if u.FirstName != "" {
		userMap["first_name"] = u.FirstName
	}
	if u.LastName != "" {
		userMap["last_name"] = u.LastName
	}
	if u.StatusCode != 0 {
		userMap["status_code"] = u.StatusCode
	}
	if u.Message != "" {
		userMap["message"] = u.Message
	}

	return userMap
}

type UserRole struct {
	UserID     int    `json:"user_id"`
	UserRoleID int    `json:"id,omitempty"`
	UserRoles  []int  `json:"user_roles,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (ur *UserRole) GetMap() map[string]interface{} {
	userRoleMap := map[string]interface{}{
		"user_id":    ur.UserID,
		"user_roles": ur.UserRoles,
	}

	if ur.UserRoleID != 0 {
		userRoleMap["id"] = ur.UserRoleID
	}
	if ur.StatusCode != 0 {
		userRoleMap["status_code"] = ur.StatusCode
	}
	if ur.Message != "" {
		userRoleMap["message"] = ur.Message
	}

	return userRoleMap
}

type Role struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	RoleID      int    `json:"id,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	Message     string `json:"message,omitempty"`
}

func (r *Role) GetMap() map[string]interface{} {
	roleMap := map[string]interface{}{
		"name":        r.Name,
		"description": r.Description,
	}

	if r.RoleID != 0 {
		roleMap["id"] = r.RoleID
	}
	if r.StatusCode != 0 {
		roleMap["status_code"] = r.StatusCode
	}
	if r.Message != "" {
		roleMap["message"] = r.Message
	}

	return roleMap
}

type RolePermission struct {
	ID              int    `json:"id,omitempty"`
	RoleID          int    `json:"role_id"`
	RolePermissions []int  `json:"role_permissions"`
	StatusCode      int    `json:"status_code,omitempty"`
	Message         string `json:"message,omitempty"`
}

func (rp *RolePermission) GetMap() map[string]interface{} {
	rolePermissionMap := map[string]interface{}{
		"role_id":          rp.RoleID,
		"role_permissions": rp.RolePermissions,
	}

	if rp.ID != 0 {
		rolePermissionMap["id"] = rp.ID
	}
	if rp.StatusCode != 0 {
		rolePermissionMap["status_code"] = rp.StatusCode
	}
	if rp.Message != "" {
		rolePermissionMap["message"] = rp.Message
	}

	return rolePermissionMap
}

type Permission struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (p *Permission) GetMap() map[string]interface{} {
	return map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
	}
}

func main() {
	// Usage example:
	user := User{
		Email:     "test@example.com",
		Password:  "password",
		FirstName: "John",
		LastName:  "Doe",
		UserID:    1,
	}
	userMap := user.ToMap()
	fmt.Println(userMap)

	role := Role{
		Name:        "Admin",
		Description: "Administrator role",
		RoleID:      123,
	}
	roleMap := role.GetMap()
	fmt.Println(roleMap)
}
