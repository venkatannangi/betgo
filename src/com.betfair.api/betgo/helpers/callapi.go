package helpers

import (
	"net/http"
	"io/ioutil"
	"log"
	"strings"
	"encoding/json"
)

func SendRequest (url string, appKey string, ssoid string, payload string, v interface{}) error {

	client := &http.Client{}
	req , err := http.NewRequest("POST", url, strings.NewReader(payload))
	req.Header.Add("X-Application", appKey)
	req.Header.Add("X-Authentication", ssoid)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")

	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		if err := json.Unmarshal(contents, v); err != nil {
			log.Fatal(err)
			return err
		}

		if err != nil {
			log.Fatal(err)
			return err
		}

		return nil

	}
}
