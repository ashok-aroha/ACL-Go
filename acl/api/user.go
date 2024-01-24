package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	// "io/ioutil"
	"net/http"

	acl "github.com/ashok-aroha/ACL-Go/acl"
)

type User struct {
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	UserID     int    `json:"user_id,omitempty"`
	StatusCode int
	Message    string
}

func (u *User) ToMap() map[string]interface{} {
	userDict := map[string]interface{}{
		"email":      u.Email,
		"password":   u.Password,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
	}

	if u.UserID != 0 {
		userDict["id"] = u.UserID
	}

	if u.StatusCode != 0 {
		userDict["status_code"] = u.StatusCode
	}

	if u.Message != "" {
		userDict["message"] = u.Message
	}

	return userDict
}

type UserRole struct {
	UserID     int   `json:"user_id,omitempty"`
	UserRoles  []int `json:"user_roles,omitempty"`
	ID         int   `json:"id,omitempty"`
	StatusCode int
	Message    string
}

func (ur *UserRole) ToMap() map[string]interface{} {
	userRoleDict := map[string]interface{}{
		"user_id":    ur.UserID,
		"user_roles": ur.UserRoles,
	}

	if ur.ID != 0 {
		userRoleDict["id"] = ur.ID
	}

	if ur.StatusCode != 0 {
		userRoleDict["status_code"] = ur.StatusCode
	}

	if ur.Message != "" {
		userRoleDict["message"] = ur.Message
	}

	return userRoleDict
}

type UserAPI struct {
	Headers acl.Headers
}

func (ua *UserAPI) GetUserByID(userID int) User {

	url := acl.GET_USER_BY_ID_API

	payload, _ := json.Marshal(map[string]interface{}{"id": userID})

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	for key, value := range ua.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return User{}
	}

	defer resp.Body.Close()

	var userObj User

	if resp.StatusCode == http.StatusOK {
		fmt.Println("User Found!!")

		json.NewDecoder(resp.Body).Decode(&userObj)

		userObj.Message = "User Found!!"
	} else {
		userObj.Message = "User Not Found!!"
	}

	userObj.StatusCode = resp.StatusCode

	return userObj
}

func (ua *UserAPI) GetUserRoleByID(userID int) UserRole {

	url := acl.GET_USER_ROLE_BY_ID_API

	payload, _ := json.Marshal(map[string]interface{}{"id": userID})

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	for key, value := range ua.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return UserRole{}
	}

	defer resp.Body.Close()

	var userRoleObj UserRole

	if resp.StatusCode == http.StatusOK {
		fmt.Println("User Role Found!!")

		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(string(body))

		json.NewDecoder(resp.Body).Decode(&userRoleObj)

		userRoleObj.Message = "User Role Found!!"
	} else {
		userRoleObj.Message = "User Role Not Found!!"
	}

	userRoleObj.StatusCode = resp.StatusCode

	return userRoleObj
}

func (ua *UserAPI) CreateUser(email, password, firstName, lastName string) User {
	userObj := User{
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}

	url := acl.CREATE_USER_API

	payload, _ := json.Marshal(userObj.ToMap())

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	for key, value := range ua.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return User{}
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("User Created!!")

		userData := map[string]interface{}{}
		json.NewDecoder(resp.Body).Decode(&userData)

		userObj.UserID = int(userData["id"].(float64))
		userObj.Message = "User Created!!"
	} else {
		userObj.Message = "User Creation Failed!!"
	}

	userObj.StatusCode = resp.StatusCode

	return userObj
}

func (ua *UserAPI) CreateUserRole(userID int, userRoles []int) UserRole {
	userRoleObj := UserRole{
		UserID:    userID,
		UserRoles: userRoles,
	}

	url := acl.CREATE_USER_ROLE_API // Replace with your API URL

	payload, _ := json.Marshal(userRoleObj.ToMap())

	fmt.Println(userRoleObj.ToMap())
	fmt.Println(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	for key, value := range ua.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return UserRole{}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	if resp.StatusCode == http.StatusOK {

		userRoleData := map[string]interface{}{}
		json.NewDecoder(resp.Body).Decode(&userRoleData)

		userRoleObj.ID = int(userRoleData["id"].(float64))
		userRoleObj.Message = "User Role Created!!"
	} else {
		userRoleObj.Message = "User Role Creation Failed!!"
	}

	userRoleObj.StatusCode = resp.StatusCode

	return userRoleObj
}
