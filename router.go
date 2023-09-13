package main
import(
	"github.com/gorilla/mux"
	"net/http"
)
//Routes
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/", indexHandler)


	//Assets
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