package helpers

import (
	"encoding/json"
	"io/ioutil"
	"io"
	"com.betfair.api/betgo/to"
)

func GetMarketFilter(payload io.ReadCloser) (filter *to.MarketFilter) {

	body,_ := ioutil.ReadAll(payload)
	filter = &to.MarketFilter{}

	if err := json.Unmarshal(body, &filter); err != nil {
		panic(err)
	}

	return filter
}




