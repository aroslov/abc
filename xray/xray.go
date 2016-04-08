package xray

const (
	JOURNAL = "/journals"
	JOURNAL_COMMIT = "/commit"
	RECORD = "/records"
	RECORD_FINGERPRINT = "/fingerprints"
)

type XRayService struct {
	base_url string
}

func GetJournalService(base_url string)(*XRayService) {
	return &XRayService{base_url}
}
