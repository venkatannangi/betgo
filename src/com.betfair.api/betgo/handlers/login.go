package handlers

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"com.betfair.api/betgo/helpers"
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
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))

		var response helpers.SessionInfo
		response = helpers.ParseSSOResponse(string(body))
		/*if ( response.Error != "" && response.Token == "" ) {
			http.Redirect(w, r, "/www/home.html", 401)
			return;
		}*/
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

