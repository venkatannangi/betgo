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

func loginCall(w http.ResponseWriter, r *http.Request) {
	handlers.Login(w,r)

}

func main() {
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/shouthello/", shouthelloHandler)
	http.HandleFunc("/login", loginCall)
	http.ListenAndServe("localhost:9000", nil)
}
