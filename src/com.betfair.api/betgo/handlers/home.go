package handlers

import (
	"net/http"
	"fmt"
)
func LoadHome(w http.ResponseWriter, r *http.Request, wwwDir string) {
	fmt.Println("home")
	//if no cookie redirect to login

	sessionToken,_ := r.Cookie("ssoid")

	fmt.Println(sessionToken)

	if sessionToken != nil {
		homePage := wwwDir + "/Index.html"
		http.ServeFile(w,r, homePage);
	} else {
		Login(w,r, wwwDir)
	}


}
