package main

import (
	"net/http"
	"fmt"
	"os"
	"com.betfair.api/betgo/handlers"
)

func login(w http.ResponseWriter, r *http.Request) {
	handlers.Login(w,r, WWWDir)
}

func home(w http.ResponseWriter, r *http.Request) {
	handlers.LoadHome(w,r, WWWDir)
}

var WWWDir string

func main() {
	args := os.Args
	if len(args) != 2 {
		WWWDir = "/var/www"
	} else {
		WWWDir = args[1]
	}

	fmt.Printf("Using wwwDir location %s", WWWDir)
	cssDir := WWWDir + "/css"

	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(cssDir))))
	http.ListenAndServe(":9000", nil)
}
