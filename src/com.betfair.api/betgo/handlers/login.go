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
		http.Redirect(w, r, "/www/login.html", 302)
		return;
	} else {
		//process user input, call ISS for authentication

		//set ssoid in cookie and app key
		ssoidCookie := &http.Cookie{
			Name: "ssoid",
			Value: "foobar",
		}
		appKeyCookie := &http.Cookie{
			Name: "appKey",
			Value: "foobarKey",
		}
		http.SetCookie(w, ssoidCookie)
		http.SetCookie(w, appKeyCookie)
		r.AddCookie(ssoidCookie)
		r.AddCookie(appKeyCookie)
		LoadHome(w,r)
	}


}

