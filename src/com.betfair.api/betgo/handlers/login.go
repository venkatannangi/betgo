package handlers

import (
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
)

type login_request struct{
	username string
	password string
}

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

		var jsonStr = []byte(`{"username":"` + username + `","password":"` + password + `"}`)

		/* decoder := json.NewDecoder(r.Body);
		 var lr login_request
		 err := decoder.Decode(&lr)
		 fmt.Println(lr)
		if err != nil {
			fmt.Println(err)
		}*/

		url := "https://identitysso.betfair.com/api/login"
		fmt.Println(bytes.NewBuffer(jsonStr))
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))

		//set ssoid in cookie and app key
		ssoidCookie := &http.Cookie{
			Name: "ssoid",
			Value: "foobar",
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

