package handlers

import (
	"net/http"
	"fmt"
)

func Login (w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	formData := r.Form;
	fmt.Println(formData)

	if len(formData) == 0 {
		//ask for user input
		http.Redirect(w, r, "/www/Login.html", 302)
		return;
	} else {
		//process user input, call ISS for authentication

		//set ssoid in cookie and app key
		cookie := &http.Cookie{
			Name: "ssoid",
			Value: "foobar",
		}
		http.SetCookie(w, cookie)
		r.AddCookie(cookie)
		LoadHome(w,r)
	}


}

