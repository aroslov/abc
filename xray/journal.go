package xray

import (
	"encoding/json"
)

type Timestamp struct {
	Txid string `json:"txid"`
}

type Journal struct {
	Id string `json:"id"`
	Timestamps []Timestamp `json:"timestamps"`
}

func (s XRayService) GetJournal(journal_id string) (*Journal) {
	data := get(s.base_url + JOURNAL + "/" + journal_id)
	j := Journal{}
	json.Unmarshal(data, &j)
	return &j
}

func (j Journal) HasTimestamp() bool {
	return (len(j.Timestamps)>0) && (len(j.Timestamps[0].Txid)>0)
}



