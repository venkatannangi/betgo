package handlers

import (
	"net/http"
	"fmt"
	"com.betfair.api/betgo/helpers"
	"com.betfair.api/betgo/to"
)

func Login(w http.ResponseWriter, r *http.Request) {
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
		username := r.FormValue("username")
		applicationkey := r.FormValue("appKey")
		password := r.FormValue("password")

	    url := "https://identitysso.betfair.com/api/login?username="+username+"&password="+password

		req, err := http.NewRequest("POST", url, nil)
		if (err != nil) {
			fmt.Println(err)
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Application", applicationkey)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		response := &to.SessionInfo{}
		helpers.Parse(resp.Body, response)

		fmt.Println("session token"+response.Token)
		//set ssoid in cookie and app key
		ssoidCookie := &http.Cookie{
			Name: "ssoid",
			Value: response.Token,
		}
		appKeyCookie := &http.Cookie{
			Name: "appKey",
			Value: applicationkey,
		}
		http.SetCookie(w, ssoidCookie)
		http.SetCookie(w, appKeyCookie)
		r.AddCookie(ssoidCookie)
		r.AddCookie(appKeyCookie)
		LoadHome(w, r)
	}
}

