package handlers

import "net/http"
func LoadHome(w http.ResponseWriter, r *http.Request) {

	appKey,	sessionToken := getApiCredentials(r)

	if sessionToken != "" && appKey != "" {
		http.Redirect(w, r, "/www/home.html", 302)
		return
	} else {
		Login(w,r)
	}
}


