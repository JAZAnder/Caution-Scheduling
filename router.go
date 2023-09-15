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
	r.HandleFunc("/testsql", helloHandler)

// - - - - - - Static Reasourses - - - - - - 
	//Assets
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	r.PathPrefix("/assets/").Handler(fs)

	//Pages
	wp := http.StripPrefix("/pages/", http.FileServer(http.Dir("./pages/")))
	r.PathPrefix("/pages/").Handler(wp)


	
// - - - - - - API Endpoints - - - - - - 

//Labs
	//Get List of All Labs
	r.HandleFunc("/api/labs",helloHandler).Methods("GET")
	//Adds a new lab based on Info in POST (Name) and Cookie (Session_Token)
	r.HandleFunc("/api/labs/add", helloHandler).Methods("POST")

//LabHours
	//Gets a labs hours based on LabId
	r.HandleFunc("/api/labHours/{labId}", helloHandler).Methods("GET")
	//Changes lab hours based on LabId POST(action, hours) and Cookie(Session_Token)
	r.HandleFunc("/api/labHours/add/{labId}", helloHandler).Methods("POST")
//




//Google Shit
	r.HandleFunc("/auth/{provider}", auth)
	r.HandleFunc("/auth/{provider}/callback", authCallback)

	return r
}