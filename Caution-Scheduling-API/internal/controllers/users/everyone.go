package users

import (
	"net/http"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/responses"
	"github.com/JAZAnder/Caution-Scheduling/internal/objects/user"
)

//User Login Method
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