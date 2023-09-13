package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth/providers/google"
)

//Entry Point
func main()  {

	//Creating Google Login
	key := "SESSION_SECRET"  // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30  // 30 days
	isProd := false       // Set to true when serving over https
	store := sessions.NewCookieStore([]byte(key))
  	store.MaxAge(maxAge)
  	store.Options.Path = "/"
  	store.Options.HttpOnly = true   // HttpOnly should always be enabled
  	store.Options.Secure = isProd

	goth.UseProviders(
		google.New("239635319851-g0dhkjdo7t00run1j0fb3rld3115434h.apps.googleusercontent.com", "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", "http://local.techwall.xyz/auth/google/callback", "email", "profile"),
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

//Routes
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/", indexHandler)

	
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	r.PathPrefix("/assets/").Handler(fs)

	//Pages
	wp := http.StripPrefix("/pages/", http.FileServer(http.Dir("./pages/")))
	r.PathPrefix("/pages/").Handler(wp)

	//Google Signin
	r.HandleFunc("/auth/{provider}", auth)
	r.HandleFunc("/auth/{provider}/callback", authCallback)

	return r
}

//handlers
func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World!")
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("./index.html"))
	tmpl.Execute(w, nil)

}
func auth(w http.ResponseWriter, r *http.Request)  {
	gothic.BeginAuthHandler(w, r)
	//fmt.Fprintf(w, "Hi")
}
func authCallback(w http.ResponseWriter, r *http.Request)  {
	user, err := gothic.CompleteUserAuth(w, r)
    if err != nil {
      fmt.Fprintln(w, err)
      return
    }
    t, _ := template.ParseFiles("templates/success.html")
    t.Execute(w, user)
}
