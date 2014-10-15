package handlers

import (
	"net/http"
)
func LoadHome(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w,r, "/home/doowad/git/betgo/www/Index.html");
}
