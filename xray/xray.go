package xray

import (
	"fmt"
	"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"encoding/json"
	"net/http"
	"bytes"
	"github.com/aroslov/abc/util"
)

type XRayService struct {
	base_url string
	access_key string
	access_secret string
}

func GetXRayService(base_url, access_key, access_secret string)(*XRayService) {
	return &XRayService{base_url, access_key, access_secret}
}

func (xray_service *XRayService) post(path string, data []byte) []byte {
	req, err := http.NewRequest("POST", xray_service.base_url + path, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(xray_service.access_key, xray_service.access_secret)

	client := &http.Client{}
	resp, err := client.Do(req)
	util.PanicOnError(err, fmt.Sprintf("Error POST'ing data to %s", path))

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		log.WithFields(log.Fields{
			"method" : "POST",
			"path" : path,
			"request": req,
			"request_body": string(data),
			"response": resp,
			"response_body": string(body),
			"response_code" : resp.StatusCode,
		}).Panic("Received erronous response code")
	}
	log.WithFields(log.Fields{
		"method" : "POST",
		"path" : path,
		"response_body": string(body),
	}).Debugf("Received API response")
	return body
}

func (xray_service *XRayService) get(path string) []byte {
	req, err := http.NewRequest("GET", xray_service.base_url + path, nil)
	req.SetBasicAuth(xray_service.access_key, xray_service.access_secret)

	client := &http.Client{}
	resp, err := client.Do(req)
	util.PanicOnError(err, fmt.Sprintf("Error GET'ing data to %s", path))

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		log.Panicf("Got %d from %s", resp.StatusCode, path)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.WithFields(log.Fields{
		"method" : "GET",
		"path" : path,
		"response_body": string(body),
	}).Debugf("Received API response")
	return body
}

func unmarshal(data []byte) map[string]interface{} {
	var x map[string]interface{}
	json.Unmarshal(data, &x)
	return x
}