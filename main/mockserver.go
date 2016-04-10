package main

import (
	"net/http"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"github.com/aroslov/abc/util"
)

func handler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("config/mockresponse.json")
	util.PanicOnError(err, "Unable to read mock file")
	w.Write(b)
	body, _ := ioutil.ReadAll(r.Body)
	log.Infof("%s %s %d %s", r.Method, r.URL.Path, r.ContentLength, body)
}

func main() {
	util.SetupLogging("info", "")
	http.HandleFunc("/", handler)
	log.Infof("Listening on http://localhost:8888/")
	http.ListenAndServe(":8888", nil)
}
