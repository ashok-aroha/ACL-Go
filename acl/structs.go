package acl

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
	UserID     int    `json:"user_id,omitempty"`
	UserRoleID int    `json:"id,omitempty"`
	UserRoles  []int  `json:"user_roles,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (ur *UserRole) ToMap() map[string]interface{} {
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

func (r *Role) ToMap() map[string]interface{} {
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
	RolePermissionsID int    `json:"id,omitempty"`
	RoleID            int    `json:"role_id,omitempty"`
	RolePermissions   []int  `json:"role_permissions,omitempty"`
	StatusCode        int    `json:"status_code,omitempty"`
	Message           string `json:"message,omitempty"`
}

func (rp *RolePermission) ToMap() map[string]interface{} {
	rolePermissionsMap := map[string]interface{}{
		"role_id":          rp.RoleID,
		"role_permissions": rp.RolePermissions,
	}

	if rp.RolePermissionsID != 0 {
		rolePermissionsMap["id"] = rp.RolePermissionsID
	}
	if rp.StatusCode != 0 {
		rolePermissionsMap["status_code"] = rp.StatusCode
	}
	if rp.Message != "" {
		rolePermissionsMap["message"] = rp.Message
	}

	return rolePermissionsMap
}

type Permission struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	PermissionID int    `json:"id,omitempty"`
	StatusCode   int    `json:"status_code,omitempty"`
	Message      string `json:"message,omitempty"`
}

func (p *Permission) ToMap() map[string]interface{} {
	permissionMap := map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
	}

	if p.PermissionID != 0 {
		permissionMap["id"] = p.PermissionID
	}
	if p.StatusCode != 0 {
		permissionMap["status_code"] = p.StatusCode
	}
	if p.Message != "" {
		permissionMap["message"] = p.Message
	}

	return permissionMap
}

// func main() {
// 	user := User{
// 		Email:      "example@example.com",
// 		Password:   "password123",
// 		FirstName:  "John",
// 		LastName:   "Doe",
// 		UserID:     1,
// 		StatusCode: 200,
// 		Message:    "OK",
// 	}

// 	userDict := user.ToDict()
// 	fmt.Println(userDict)

// 	role := Role{
// 		Name:        "Admin",
// 		Description: "Administrator",
// 		RoleID:      1,
// 		StatusCode:  200,
// 		Message:     "OK",
// 	}

// 	roleDict := role.GetDict()
// 	fmt.Println(roleDict)
// }
