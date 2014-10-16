package handlers

import (
	"net/http"
	"fmt"
)

func Login (w http.ResponseWriter, r *http.Request, wwwDir string) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	formData := r.Form;
	fmt.Println(formData)

	if len(formData) == 0 {
		loginPage := wwwDir + "/Login.html"
		http.ServeFile(w,r, loginPage);
	} else {
		//do login

		//set ssoid in cookie and app key
		//cookie := http.Cookie
	}

	// call iss

	//set cookie
}

