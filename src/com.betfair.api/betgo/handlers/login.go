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
		http.ServeFile(w,r, "/home/doowad/git/betgo/www/Login.html");
	}

	// call iss

	//set cookie
}

