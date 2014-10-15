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
		http.ServeFile(w,r, "www/Login.html");
	} else {
		//do login

		//set ssoid in cookie and app key
		//cookie := http.Cookie
	}

	// call iss

	//set cookie
}

