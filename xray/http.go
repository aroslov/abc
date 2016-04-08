package xray

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/aroslov/abc/util"
	"encoding/json"
	"bytes"
	"log"
)

// TODO: escalate errors instead of failing
func post(url string, data []byte) []byte {
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	util.FailOnError(err, fmt.Sprintf("Error POST'ing data to %s", url))

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		util.Fail(fmt.Sprintf("Got %d from %s", resp.StatusCode, url))
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("Received from %s: %s", url, body)
	return body
}

func get(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	util.FailOnError(err, fmt.Sprintf("Error GET'ing data to %s", url))

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		util.Fail(fmt.Sprintf("Got %d from %s", resp.StatusCode, url))
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("Received from %s: %s", url, body)
	return body
}

func unmarshal(data []byte) map[string]interface{} {
	var x map[string]interface{}
	json.Unmarshal(data, &x)
	return x
}