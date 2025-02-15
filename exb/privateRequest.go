package exb

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (ex *Keys) PrivateRequest(u string, params map[string]string, query string, get ...string) interface{} {
	var req *http.Request
	var err error
	nonce := strconv.FormatInt(time.Now().Unix()*1000, 10)

	key := []byte(ex.PrivateKey)
	message := nonce + ex.PublicKey

	sig := hmac.New(sha256.New, key)
	sig.Write([]byte(message))

	signhash := hex.EncodeToString(sig.Sum(nil))

	if get != nil {
		req, err = http.NewRequest("GET", u+query, nil)
		if err != nil {
			log.Printf("Error: %s", err.Error())
		}
		q := req.URL.Query()
		for papam, val := range params {
			q.Add(papam, val)
		}

		req.URL.RawQuery = q.Encode()
	} else {
		postBody, err := json.Marshal(params)
		if err != nil {
			log.Printf("Error: %s", err.Error())
		}
		req, err = http.NewRequest("POST", u+query, bytes.NewBuffer(postBody))
		if err != nil {
			log.Printf("Error: %s", err.Error())
		}
	}
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}
	req.Header.Set("X-Auth-Apikey", ex.PublicKey)
	req.Header.Set("X-Auth-Nonce", nonce)
	req.Header.Set("X-Auth-Signature", signhash)
	req.Header.Set("content-Type", "application/json")

	//log.Println(req.URL.String())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	//log.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	//log.Println("response Body:", string(body))

	if resp.StatusCode >= http.StatusBadRequest {
		var m interface{}
		e := json.Unmarshal(body, &m)
		if e != nil {
			log.Printf("failed to unmarshal json: %s", e)
		}
		targets := m.(map[string]interface{})
		for i, t := range targets {
			if i == "errors" {
				log.Fatalf("Ошибка API: %s", t)
			}
		}
	}

	//res = make([]*Order, 0)
	//err = json.Unmarshal(body, &res)
	//if err != nil {
	//	return []*Order{}, err
	//}
	return body
}
