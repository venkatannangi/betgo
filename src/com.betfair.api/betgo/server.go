package main

import (
	"net/http"
	"fmt"
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
	handlers.Login(w,r)
}

func home(w http.ResponseWriter, r *http.Request) {
	handlers.LoadHome(w,r)
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/", home)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("www/css"))))
	http.ListenAndServe(":9000", nil)
}
