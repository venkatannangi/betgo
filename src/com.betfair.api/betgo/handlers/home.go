package handlers

import (
	"net/http"
	"fmt"
)
func LoadHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
	//if no cookie
	http.ServeFile(w,r, "/home/doowad/git/betgo/www/Index.html");
}
