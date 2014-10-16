package handlers

import (
	"net/http"
	"fmt"
)
func LoadHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
	//if no cookie redirect to login

	sessionToken,_ := r.Cookie("ssoid")

	fmt.Println(r.Cookies())
	fmt.Println(sessionToken)

	if sessionToken != nil {
		http.Redirect(w, r, "/www/home.html", 302)
		return
	} else {
		Login(w,r)
	}


}
