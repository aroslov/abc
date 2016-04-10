package xray

import (
	"encoding/json"
)

const (
	JOURNAL = "/journals"
	TIMESTAMP = "/timestamp"
)


type Timestamp struct {
	Txid string `json:"txid"`
}

type Journal struct {
	Id string `json:"id"`
	Timestamps []Timestamp `json:"timestamps"`
	xray_service XRayService
}

func (s XRayService) GetJournal(journal_id string) (*Journal) {
	data := s.get(JOURNAL + "/" + journal_id)
	j := Journal{xray_service: s}
	json.Unmarshal(data, &j)
	return &j
}

func (j Journal) Timestamp() {
	j.xray_service.post(JOURNAL + "/" + j.Id + TIMESTAMP, nil)
}

func (j Journal) NumeberOfTimestamps() int {
	n := 0
	for _, t := range j.Timestamps {
		if len(t.Txid) > 0 {
			n++
		}
	}
	return n
}



