package handlers

import (
	"net/http"
	"html/template"
)
func Login (w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("/home/doowad/git/betgo/www/Login.html")
	t.Execute(w, nil)
}

