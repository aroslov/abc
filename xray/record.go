package xray

import (
	"encoding/json"
)

const (
	RECORD = "/records"
	JOURNAL_COMMIT = "/commit"
)

type Record struct {
	Id string `json:"id"`
	xray_service XRayService
}

func (xray_service XRayService) CreateRecord(record []byte) (*Record) {
	data := xray_service.post(RECORD, record)
	r := Record{xray_service: xray_service}
	json.Unmarshal(data, &r)
	return &r
}

func (r Record) CommitToJournal(journal_id string) {
	r.xray_service.post(JOURNAL + "/" + journal_id + JOURNAL_COMMIT + "/" + r.Id, nil)
}
