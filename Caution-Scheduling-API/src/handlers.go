package main
import(
	"fmt"
	"html/template"
	"github.com/markbates/goth/gothic"
	"net/http"
)
	
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