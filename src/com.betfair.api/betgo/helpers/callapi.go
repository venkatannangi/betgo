package helpers

import (
	"net/http"
	"io/ioutil"
	"log"
	"strings"
)

func SendRequest (url string, appKey string, ssoid string, payload string) string {

	client := &http.Client{}
	req ,err := http.NewRequest("POST", url , strings.NewReader(payload))
	req.Header.Add("X-Application",appKey)
	req.Header.Add("X-Authentication",ssoid)
	req.Header.Add("Accept","application/json")
	req.Header.Add("Content-type","application/json")

	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		return string(contents)
		/*fmt.Println("The calculated length is:", len(string(contents)), "for the url:", url)
		fmt.Println("The content is:", string(contents), "for the url:", url)
		fmt.Println("   ", response.StatusCode)
		hdr := response.Header
		for key, value := range hdr {
			fmt.Println("   ", key, ":", value)
		} */
	}
	return ""
}
