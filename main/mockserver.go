package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"github.com/aroslov/abc/util"
)

func handler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("config/mockresponse.json")
	util.FailOnError(err, "Unable to read mock file")
	w.Write(b)
	body, _ := ioutil.ReadAll(r.Body)
	log.Printf("%s %s %d %s", r.Method, r.URL.Path, r.ContentLength, body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
