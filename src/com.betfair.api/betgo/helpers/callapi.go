package helpers

import (
	"net/http"
	"io/ioutil"
	"log"
	"strings"
	"encoding/json"
)

func SendRequest (url string, appKey string, ssoid string, payload string, v interface{}) error {
	if req , err := http.NewRequest("POST", url, strings.NewReader(payload)); err == nil {
		req.Header.Add("X-Application", appKey)
		req.Header.Add("X-Authentication", ssoid)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-type", "application/json")

		client := &http.Client{}
		if response, err := client.Do(req); err == nil {
			defer response.Body.Close()
			if contents, err := ioutil.ReadAll(response.Body); err == nil {
				if err := json.Unmarshal(contents, v); err != nil {
					log.Fatal(err)
					return err
				}
			} else {
				log.Fatal(err)
				return err
			}
		} else {
			log.Fatal(err)
			return err
		}
		return nil;

	}else {
		log.Fatal(err)
		return err
	}
}
