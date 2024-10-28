package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"undrakh.net/summarizer/cmd/web/app"
	"undrakh.net/summarizer/pkg/common/oapi"
	"undrakh.net/summarizer/pkg/easyOAuth2"
	"undrakh.net/summarizer/pkg/userman"
)

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func oauthLogin(oauthClient *easyOAuth2.EasyOAuthClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := oauthClient.RedirectToLogin(w, r); err != nil {
			handleOAuthError(w, r, fmt.Sprintf("%v %v %v", oauthClient.Name, "oauth2 login error:", err))
			return
		}
	}
}

func oauthCallback(oauthClient *easyOAuth2.EasyOAuthClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := oauthClient.HandleCallback(w, r)
		if err != nil {
			handleOAuthError(w, r, fmt.Sprintf("%v %v %v", oauthClient.Name, "oauth2 callback error:", err))
			return
		}

		data, err := oauthClient.GetUserInfo(token.AccessToken)
		if err != nil {
			handleOAuthError(w, r, fmt.Sprintf("%v %v %v", oauthClient.Name, "oauth2 callback error:", err))
			return
		}

		var (
			filter    *userman.User
			authTypes []string
			userData  *userman.User // Possible new user data. Less code this way.
		)
		switch oauthClient.Name {
		case "google":
			var userinfo *GoogleUserInfo
			if err := json.Unmarshal(data, &userinfo); err != nil {
				handleOAuthError(w, r, fmt.Sprintf("google unmarshal error: %v data: %v", err, string(data)))
				return
			}
			if userinfo.ID == "" {
				handleOAuthError(w, r, fmt.Sprintf("google userinfo had empty ID. Data: %v", string(data)))
				return
			}
			filter = &userman.User{Email: userinfo.Email}
			authTypes = []string{userman.AUTH_TYPE_BASIC, userman.AUTH_TYPE_GOOGLE}
			userData = &userman.User{
				AuthType:       userman.AUTH_TYPE_GOOGLE,
				GoogleID:       userinfo.ID,
				Name:           userinfo.Name,
				ProfilePicture: userinfo.Picture,
				Email:          userinfo.Email,
			}
		default:
			oapi.ServerError(w, fmt.Errorf("invalid oauth2 provider: %v", oauthClient.Name))
			return
		}

		user, err := app.Users.GetWithAuthTypes(filter, authTypes)
		if err != nil && !errors.Is(err, userman.ErrNotFound) {
			oapi.ServerError(w, err)
			return
		}

		recentlyDeleted, err := app.Users.GetRecentlyDeleted(filter, authTypes)
		if err != nil && !errors.Is(err, userman.ErrNotFound) {
			oapi.ServerError(w, err)
			return
		}

		if recentlyDeleted != nil {
			http.Redirect(w, r, "/account_deleted", http.StatusTemporaryRedirect)
			return
		}

		if user == nil {

			userData.UUID = uuid.NewString()
			userData.LastLogin = time.Now()
			userData.IsVerified = true
			c, err := app.Users.Save(userData)
			if err != nil {
				oapi.ServerError(w, err)
				return
			}
			user = c
		} else {
			user.Name = userData.Name
			user.Email = userData.Email
			user.ProfilePicture = userData.ProfilePicture
			user.LastLogin = time.Now()
			if !user.IsVerified {
				user.IsVerified = true
				user.PasswordHash = ""
				user.AuthType = userData.AuthType
			}
			if _, err := app.Users.Save(user); err != nil {
				oapi.ServerError(w, err)
				return
			}
		}

		app.Session.Put(r, "auth_user_id", user.ID)
		app.Session.Put(r, "oauth2_provider_name", oauthClient.Name)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

func handleOAuthError(w http.ResponseWriter, r *http.Request, errorStr string) {
	app.ErrorLog.Println(errorStr)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func me(w http.ResponseWriter, r *http.Request) {
	loggedUser := r.Context().Value(app.ContextKeyChosenUser).(*userman.User)
	if loggedUser == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	response, err := app.Users.GetWithRoles(loggedUser)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}
	oapi.SendResp(w, response)
}
