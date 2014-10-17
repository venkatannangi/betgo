package handlers

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

const menuEndpoint = "https://api.betfair.com/exchange/betting/rest/v1/en/navigation/menu.json"

func GetMenu(w http.ResponseWriter, r *http.Request) {
	//appKey,	sessionToken := getApiCredentials(r)
	client := &http.Client{}
	req,_ := http.NewRequest("GET", menuEndpoint, nil)

	addCredentialToHeader(req, "appkey", "sessionToken")
	fmt.Println(req)
	fmt.Println(req.Header)
	resp,_ := client.Do(req)

	fmt.Println(resp)
	fmt.Println(resp.Header.Get("Content-Encoding"))

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))




}

