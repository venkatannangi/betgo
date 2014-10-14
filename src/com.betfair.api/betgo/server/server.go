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
	if (strings.Contains(r.URL.Path, ".")) {
		chttp.ServeHTTP(w, r)
	} else {
		//fmt.Fprintf(w, "HomeHandler")
		handlers.Login(w,r)
	}

}
var chttp = http.NewServeMux()
func main() {
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/shouthello/", shouthelloHandler)
	http.HandleFunc("/login", loginCall)
	chttp.Handle("/login", http.FileServer(http.Dir("/home/doowad/git/betgo/www/")))
	http.ListenAndServe("localhost:9000", nil)
}
