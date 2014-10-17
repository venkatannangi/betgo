package handlers

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"time"
	"encoding/json"
)

const menuEndpoint = "https://api.betfair.com/exchange/betting/rest/v1/en/navigation/menu.json"


type MarketFilter struct {
	ExchangeIds []string `json:"exchangeIds"`
/*TextQuery string `json:"textQuery"`
	EventTypeIds []string `json:"eventTypeIds"`
	EventIds []string `json:"eventIds"`
	CompetitionIds []string `json:"competitionIds"`
	Venues []string `json:"venues"`
	MarketBettingTypes []string `json:"marketBettingTypes"`
	MarketCountries []string `json:"marketCountries"`
	MarketTypeCodes []string `json:"marketTypeCodes"`
	WithOrders []string `json:"withOrders"`
	BspOnly bool `json:"bspOnly"`
	TurnInPlayEnabled bool `json:"turnInPlayEnabled"`
	InPlayOnly bool `json:"inPlayOnly"`
	MarketStartTime timerange `json:"marketStartTime"`*/
}

type timerange struct {
	From time.Time `json:"from"`
	To time.Time `json:"to"`
}

func GetMenu(w http.ResponseWriter, r *http.Request) {
	//appKey,	sessionToken := getApiCredentials(r)
	client := &http.Client{}
	req,_ := http.NewRequest("GET", menuEndpoint, nil)

	addCredentialToHeader(req, "", "")
	fmt.Println(req)
	fmt.Println(req.Header)
	resp,_ := client.Do(req)

	fmt.Println(resp)
	fmt.Println(resp.Header.Get("Content-Encoding"))

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))

}

func GetSports(w http.ResponseWriter, r *http.Request) {

	marketFilter := getMarketFilter(r)

	//client := &http.Client{}
	//req,_ := http.NewRequest("GET", menuEndpoint, nil)
	fmt.Println(marketFilter)
	fmt.Println(marketFilter.ExchangeIds)

}

func getMarketFilter(r *http.Request) (filter *MarketFilter) {

	body,err := ioutil.ReadAll(r.Body)
	filter = &MarketFilter{}

	if err != nil {
		if err := json.Unmarshal(body, &filter); err != nil {
			panic(err)
		}
	}

	fmt.Println("request Body:", string(body))

	return filter
}


