package handlers

import "net/http"

func Login (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r, "/home/doowad/git/betgo/www/Login.html");
}

