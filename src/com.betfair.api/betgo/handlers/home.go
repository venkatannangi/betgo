package handlers

import (
	"net/http"
	"fmt"
)
func LoadHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
	//if no cookie redirect to login

	sessionToken,_ := r.Cookie("ssoid")

	fmt.Println(sessionToken)

	if sessionToken != nil {
		http.ServeFile(w,r, "/home/doowad/git/betgo/www/Index.html");
	} else {
		Login(w,r)
	}


}
