package xray

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/aroslov/abc/util"
	"encoding/json"
)

type Fingerprint struct {
	Metadata string `json:"metadata"`
	MetadataContentType string `json:"metadataContentType"`
	MetadataHash string `json:"metadataHash"`
	Nonce string `json:"nonce"`

	xray_service XRayService;
}

func (xray_service XRayService) NewFingerprint(message []byte) (*Fingerprint) {
	nonce := util.Nonce(13)
	tmp := []byte{}
	tmp = append(tmp, nonce...)
	tmp = append(tmp, message...)
	hash := sha256.Sum256(tmp)
	b64 := base64.StdEncoding
	fp := Fingerprint{b64.EncodeToString(message), "application/json;enc=v1",
		hex.EncodeToString(hash[:]), b64.EncodeToString(nonce), xray_service}
	return &fp
}

func (fp Fingerprint) json() ([]byte) {
	fp_json,err := json.Marshal(fp)
	util.FailOnError(err, "Unable to create fingerprint")
	return fp_json
}

func (fp Fingerprint) CreateRecordFingerprint (record_id string) {
	post(fp.xray_service.base_url + RECORD + "/" + record_id + RECORD_FINGERPRINT, fp.json())
}
