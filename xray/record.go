package xray

import (
	"github.com/aroslov/abc/util"
	"encoding/json"
)

type Record struct {
	Id string `json:"id"`
	xray_service XRayService;
}

func (s XRayService) CreateRecord(record []byte) (*Record) {
	data := post(s.base_url + RECORD, record)
	j := Journal{s: s}
	json.Unmarshal(data, &j)
	return j
}

func (r Record) CommitToJournal(journal_id string) {
	post(r.xray_service.base_url + JOURNAL + "/" + journal_id + JOURNAL_COMMIT + "/" + r.Id, nil)
}

func (r Record) CommitJournalRecord(journal_id string, record_id string) {
	post(r.xray_service.base_url + JOURNAL + "/" + journal_id + JOURNAL_COMMIT + "/" + r.Id, nil)
}