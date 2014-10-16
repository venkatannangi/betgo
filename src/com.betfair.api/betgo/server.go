package main

import (
	"net/http"
	"fmt"
	"os"
	"strings"
	"com.betfair.api/betgo/handlers"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/hello/"):] //get everything after the /hello/ part of the URL
	fmt.Fprintf(w, "Hello %s!", remPartOfURL)
}

func shouthelloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/shouthello/"):] //get everything after the /shouthello/ part of the URL
	fmt.Fprintf(w, "Hello %s!", strings.ToUpper(remPartOfURL))
}

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
		panic("Expecting localtion www dir as a parameter")
	}

	WWWDir = args[1]
	cssDir := WWWDir + "/css"

	fmt.Println(cssDir)
	http.HandleFunc("/login", login)
	http.HandleFunc("/", home)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(cssDir))))
	http.ListenAndServe(":9000", nil)
}
