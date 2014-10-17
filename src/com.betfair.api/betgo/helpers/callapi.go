package helpers

import (
	"net/http"
	"io/ioutil"
	"log"
	"bytes"
	"encoding/json"
)

func SendRequest (url string, appKey string, ssoid string, payload interface {}, v interface{}) error {
	if body, err := json.Marshal(payload); err == nil {
		if req , err := http.NewRequest("POST", url, bytes.NewReader(body)); err == nil {
			req.Header.Add("X-Application", appKey)
			req.Header.Add("X-Authentication", ssoid)
			req.Header.Add("Accept", "application/json")
			req.Header.Add("Content-type", "application/json")
			client := &http.Client{}
			if response, err := client.Do(req); err == nil {
				defer response.Body.Close()
				if contents, err := ioutil.ReadAll(response.Body); err == nil {
					if err := json.Unmarshal(contents, v); err == nil {
						return nil
					} else {
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
		}else {
			log.Fatal(err)
			return err
		}
	}else {
		log.Fatal(err)
		return err
	}

}
