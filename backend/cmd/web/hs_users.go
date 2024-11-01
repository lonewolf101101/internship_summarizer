package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"undrakh.net/summarizer/cmd/web/app"
	"undrakh.net/summarizer/cmd/web/validators"
	"undrakh.net/summarizer/pkg/common/oapi"
	"undrakh.net/summarizer/pkg/roleman"
	"undrakh.net/summarizer/pkg/userman"
)

type CreateRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AddRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type AddUserRoleRequest struct {
	User userman.User `json:"user"`
	Role roleman.Role `json:"role"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	size, _ := strconv.Atoi(q.Get("size"))

	if page <= 0 {
		page = 1
	}

	if size <= 0 || size > 100 {
		size = 25
	}

	filter := new(userman.Filter)
	filter.Role = q.Get("role")
	filter.Keyword = q.Get("keyword")

	users, total, err := app.Users.GetAll(filter, page, size)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}

	oapi.SendResp(w, map[string]interface{}{
		"items": users,
		"total": total,
	})
}

func getUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(app.ContextKeyChosenUser).(*userman.User)
	oapi.SendResp(w, user)
}

func editUser(w http.ResponseWriter, r *http.Request) {
	var data *userman.User
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		oapi.CustomError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidateUser(data); err != nil {
		oapi.CustomError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, _ := r.Context().Value(app.ContextKeyChosenUser).(*userman.User)

	user.Name = data.Name
	user.PhoneNumber = data.PhoneNumber

	savedUser, err := app.Users.Save(user)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}

	oapi.SendResp(w, savedUser)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(app.ContextKeyChosenUser).(*userman.User)

	if err := app.Users.Delete(user.ID); err != nil {
		oapi.ServerError(w, err)
		return
	}

	oapi.SendResp(w, user)
}

func updateUserInfo(w http.ResponseWriter, r *http.Request) {
	loggedUser := r.Context().Value(app.ContextKeyAuthUser).(*userman.User)

	var data *userman.User
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		oapi.CustomError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidateUser(data); err != nil {
		oapi.CustomError(w, http.StatusBadRequest, err.Error())
		return
	}

	loggedUser.Name = data.Name
	loggedUser.PhoneNumber = data.PhoneNumber

	savedUser, err := app.Users.Save(loggedUser)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}

	oapi.SendResp(w, savedUser)
}

func AddRoleHandler(w http.ResponseWriter, r *http.Request) {
	var input AddRoleRequest

	// Read and parse the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		oapi.CustomError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := json.Unmarshal(body, &input); err != nil {
		oapi.CustomError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Create a new Role instance
	role := &roleman.Role{
		Name:        input.Name,
		Description: input.Description,
	}

	// Save the role to the database
	role, err = app.Roles.Save(role)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}

	// Respond with the created role
	oapi.SendResp(w, role)
}

func addUserRole(w http.ResponseWriter, r *http.Request) {
	var data *AddUserRoleRequest

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		oapi.CustomError(w, http.StatusBadRequest, err.Error())
		return
	}
	var user = data.User
	var role = data.Role
	err := app.Users.AddRole(&user, &role)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}
	resp, err := app.Users.GetWithRoles(&user)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}

	oapi.SendResp(w, resp)
}
