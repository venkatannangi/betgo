package main

import (
	"net/http"
	"fmt"
	"os"
	"com.betfair.api/betgo/handlers"
)

func login(w http.ResponseWriter, r *http.Request) {
	handlers.Login(w,r)
}

func home(w http.ResponseWriter, r *http.Request) {
	handlers.LoadHome(w,r)
}

func getSports(w http.ResponseWriter, r *http.Request) {
	handlers.GetSports(w,r)
}

var wwwDir string


func main() {
	args := os.Args
	if len(args) != 2 {
		wwwDir = "/var/www"
	} else {
		wwwDir = args[1]
	}

	fmt.Println("Using wwwDir location " + wwwDir)

	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/getSports", getSports)
	http.Handle("/www/", http.StripPrefix("/www/", http.FileServer(http.Dir(wwwDir))))
	http.ListenAndServe(":9000", nil)
}
