package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

// Entry Point
func main() {

	//Creating Google Login
	key := "SESSION_SECRET" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30    // 30 days
	isProd := false         // Set to true when serving over https
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	goth.UseProviders(
		google.New(".......apps.googleusercontent.com", "PASSWORD", "http://local.techwall.xyz/auth/google/callback", "email", "profile"),
	)

	//Ending Google Login

	err := godotenv.Load()
	if err != nil {
		log.Println("Error Loading .env file")
	}
	//Grabs Port from .env if Empty defults 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := newRouter()
	http.ListenAndServe(":"+port, r)
}
