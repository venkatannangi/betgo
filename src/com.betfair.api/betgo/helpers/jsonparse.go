package helpers

import (
	"encoding/json"
	"io/ioutil"
	"io"
	"log"
)

func Parse(payload io.ReadCloser, v interface {}){
	if body,err := ioutil.ReadAll(payload); err == nil{
		if err := json.Unmarshal(body, v); err == nil {
			return
		}else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}

}




