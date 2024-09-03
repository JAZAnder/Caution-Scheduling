package main

import (
	"log"
	//"net/http"
	"os"
	//"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	//"github.com/markbates/goth"
	//"github.com/markbates/goth/providers/google"
	//"github.com/gin-gonic/gin"
	. "github.com/JAZAnder/Caution-Scheduling/internal/app"
)

// Entry Point
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error Loading .env file")
	}
	
	a := App{}

	a.Initialize()
		

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	a.Run(port)
}
