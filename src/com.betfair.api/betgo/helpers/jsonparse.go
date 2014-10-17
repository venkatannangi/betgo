package helpers

import (
	"encoding/json"
	"io/ioutil"
	"io"
	"com.betfair.api/betgo/to"
	"log"
)



func GetMarketFilter(payload io.ReadCloser) (filter *to.MarketFilter) {

	if body,err := ioutil.ReadAll(payload); err == nil{
		filter = &to.MarketFilter{}
		if err := json.Unmarshal(body, &filter); err == nil {
			return filter
		} else {
			log.Fatal(err)
			return nil
		}

	}else {
		log.Fatal(err)
		return nil
	}

}
