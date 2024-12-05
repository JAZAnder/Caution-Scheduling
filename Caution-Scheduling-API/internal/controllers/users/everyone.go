package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
)

type GoogleTokenPayload struct {
	Email      string `json:"email"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Sub        string `json:"sub"`
}

func googleLoginUser(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		responses.RespondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	googleUser, err := verifyGoogleToken(request.Token)
	if err != nil {
		responses.RespondWithError(w, http.StatusUnauthorized, "Invalid Google token")
		return
	}

	if !strings.HasSuffix(googleUser.Email, "@selu.edu") {
		responses.RespondWithError(w, http.StatusForbidden, "Access restricted to @selu.edu email addresses")
		return
	}

	var u user.LocalUser
	u.Email = googleUser.Email
	err = u.GetUserByEmail(database)

	if err != nil && err.Error() == "user not found" {
		u.FirstName = googleUser.GivenName
		u.LastName = googleUser.FamilyName
		u.UserName = strings.ToLower(googleUser.Email)
		u.Role = 1
		u.FullName = u.FirstName + " " + u.LastName
		u.GoogleId = googleUser.Sub
		if err = u.SignUp(database); err != nil {
			responses.RespondWithError(w, http.StatusInternalServerError, "Failed to Sign In")
			return
		}
	} else if err != nil {
		responses.RespondWithError(w, http.StatusInternalServerError, "Database error")
		return
	}

	var c user.SessionCookie
	c.UserName = u.UserName
	c.CreateSession(database)
	http.SetCookie(w, &http.Cookie{
		Name:     "key",
		Value:    c.Cookie,
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	userDto, _ := u.ToSelfViewInformation()
	responses.RespondWithJSON(w, http.StatusOK, userDto)
}

func verifyGoogleToken(token string) (*GoogleTokenPayload, error) {
	resp, err := http.Get("https://oauth2.googleapis.com/tokeninfo?id_token=" + token)
	if err != nil {
		return nil, fmt.Errorf("error verifying Google token: %v", err)
	}
	defer resp.Body.Close()

	var googlePayload GoogleTokenPayload
	if err := json.NewDecoder(resp.Body).Decode(&googlePayload); err != nil {
		return nil, fmt.Errorf("error decoding token response: %v", err)
	}
	return &googlePayload, nil
}

func loginLocalUser(w http.ResponseWriter, r *http.Request) {
	var u user.LocalUser
	var c user.SessionCookie

	u.UserName = r.PostFormValue("userName")
	u.Password = r.PostFormValue("password")
	err := u.Login(database)
	u.Password = "REDACTED"

	if err != nil {
		responses.RespondWithError(w, http.StatusUnauthorized, "Username or Password Incorrect")
		return
	}

	c.UserName = u.UserName
	c.CreateSession(database)
	cookie := http.Cookie{
		Name:     "key",
		Value:    c.Cookie,
		MaxAge:   3600,
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)

	userDto, _ := u.ToSelfViewInformation()
	responses.RespondWithJSON(w, http.StatusOK, userDto)

}
