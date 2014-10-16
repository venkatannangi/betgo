package handlers

import (
	"net/http"
	"fmt"
)
func LoadHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
	//if no cookie redirect to login

	appKey,	sessionToken := getApiCredentials(r)

	if sessionToken != "" && appKey != "" {
		http.Redirect(w, r, "/www/home.html", 302)
		return
	} else {
		Login(w,r)
	}
}



func GetMenu(w http.ResponseWriter, r *http.Request) {

}

func getApiCredentials(r *http.Request) (appKey string, sessionToken string) {
	appKeyCookie,_ := r.Cookie("appKey")
	if appKeyCookie != nil {
		appKey = appKeyCookie.Value
	}


	sessionTokenCookie,_ := r.Cookie("ssoid")
	if sessionTokenCookie != nil {
		sessionToken = sessionTokenCookie.Value
	}

	return appKey, sessionToken
}
