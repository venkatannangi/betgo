package handlers

import (
	"net/http"
	"fmt"
	"com.betfair.api/betgo/helpers"
	"com.betfair.api/betgo/to"
)

const menuEndpoint = "https://api.betfair.com/exchange/betting/rest/v1/en/navigation/menu.json"



func GetMenu(w http.ResponseWriter, r *http.Request) {
	//appKey,	sessionToken := getApiCredentials(r)
	/*client := &http.Client{}
	req,_ := http.NewRequest("GET", menuEndpoint, nil)

	//addCredentialToHeader(req, "", "")
	fmt.Println(req)
	fmt.Println(req.Header)
	resp,_ := client.Do(req)

	fmt.Println(resp)
	fmt.Println(resp.Header.Get("Content-Encoding"))

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))*/

}

func GetSports(w http.ResponseWriter, r *http.Request) {
	marketFilter := &to.MarketFilter{}
	appKey,	sessionToken := getApiCredentials(r)

	helpers.Parse(r.Body, marketFilter)
	fmt.Println(marketFilter)
	fmt.Println(marketFilter)

	eventTypeRequest := to.EventTypeRequest{
		*marketFilter,
		"",
	}
	eventTypeResponse := &to.EventTypeResponse{}

	eventTypeUrl := api_betting_endpoint + "/listEventTypes/"
	helpers.SendRequest(eventTypeUrl, appKey, sessionToken, eventTypeRequest, eventTypeResponse)
	fmt.Println(eventTypeResponse)
}




