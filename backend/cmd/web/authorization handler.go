package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"undrakh.net/summarizer/cmd/web/app"
	"undrakh.net/summarizer/cmd/web/validators"
	"undrakh.net/summarizer/pkg/common/oapi"
	"undrakh.net/summarizer/pkg/roleman"
	"undrakh.net/summarizer/pkg/userman"
)

// LoginRequest and RegisterRequest structures for parsing JSON input
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PhoneNumber    string `json:"phone_number"`
	ProfilePicture string `json:"profile_picture,omitempty"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	var (
		input  RegisterRequest
		filter *userman.Filter
	)

	err = json.Unmarshal(body, &input)
	if err != nil {
		oapi.CustomError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	var user *userman.User
	hashedPassword, err := userman.HashPassword(input.Password)
	if err != nil {
		oapi.CustomError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user = &userman.User{
		Name:           input.Name,
		Email:          input.Email,
		PasswordHash:   "",
		PhoneNumber:    input.PhoneNumber,
		ProfilePicture: input.ProfilePicture,
	}

	if err := validators.ValidateUser(user); err != nil {
		oapi.CustomError(w, http.StatusBadRequest, err.Error())
		return
	}

	filter = &userman.Filter{Email: input.Email}
	_, size, err := app.Users.GetAll(filter, 5, 2)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}
	if size > 0 {
		oapi.CustomError(w, http.StatusConflict, "Email already registered")
		return
	}

	user = &userman.User{
		Name:           input.Name,
		Email:          input.Email,
		PhoneNumber:    input.PhoneNumber,
		ProfilePicture: input.ProfilePicture,
		UUID:           uuid.NewString(),
		AuthType:       userman.AUTH_TYPE_BASIC,
		PasswordHash:   hashedPassword,
		IsVerified:     false,
		LastLogin:      time.Now(),
	}
	fmt.Print(user)
	user, err = app.Users.Save(user)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}
	var role *roleman.Role

	role = &roleman.Role{
		RID: roleman.ROLE_BASIC,
	}
	role, err = app.Roles.Get(role)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}
	err = app.Users.AddRole(user, role)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}
	resp, err := app.Users.GetWithRoles(user)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}

	oapi.SendResp(w, resp)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input LoginRequest

	body, err := io.ReadAll(r.Body)
	if err != nil || json.Unmarshal(body, &input) != nil {
		oapi.CustomError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	filter := &userman.User{Email: input.Email}
	user, err := app.Users.GetWithAuthTypes(filter, []string{userman.AUTH_TYPE_BASIC})
	if err != nil {
		if errors.Is(err, userman.ErrNotFound) {
			oapi.CustomError(w, http.StatusUnauthorized, "Invalid email or password")
			return
		}
		oapi.ServerError(w, err)
		return
	}

	if !userman.VerifyPassword(input.Password, user.PasswordHash) {
		oapi.CustomError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	user.LastLogin = time.Now()
	if _, err := app.Users.Save(user); err != nil {
		oapi.ServerError(w, err)
		return
	}

	app.Session.Put(r, "auth_user_id", user.ID)

	http.Redirect(w, r, "https://localhost:3000/profile", http.StatusTemporaryRedirect)
}
