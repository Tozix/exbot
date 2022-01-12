package exb

import (
	"io/ioutil"
	"log"
	"net/http"
)

func PublicRequest(u string, query string) interface{} {

	req, err := http.NewRequest("GET", u+query, nil)
	if err != nil {
		log.Println(err.Error())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	return body
}
