package main

import (
	"fmt"
	"net/http"
	"os"

	"haservice/config"
	"haservice/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config := config.GetConfig()
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	routes.RegisterRoutes(r)

	serverUrl := os.Getenv("SERVER_URL")
	url := fmt.Sprintf("%s:%d", serverUrl, config.Port)
	fmt.Println("Server started on " + url)
	http.ListenAndServe(url, r)
}
