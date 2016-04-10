package util

import (
	"github.com/nu7hatch/gouuid"
	"regexp"
	"encoding/hex"
	"io/ioutil"
	"strings"
	"strconv"
	crand "crypto/rand"
	mrand "math/rand"
	"time"
)

func LoadTemplate(message_template string) []byte {
	b, err := ioutil.ReadFile(message_template)
	PanicOnError(err, "Unable to read message template file")
	return b
}

const (
	MACRO_UUID = "${uuid}"
	MACRO_RANDOM_HEX = "\\$\\{random_hex_([0-9]+)_(u|l)\\}"
	MACRO_RANDOM_UINT = "\\$\\{random_uint_([0-9]+)\\}"
	MACRO_TIMESTAMP = "${timestamp}"
	TIMESTAMP_FORMAT = "2006-01-02T15:04:05.999Z"
)

func SubstituteTemplate(template []byte) []byte {
	s := string(template[:])
	for strings.Contains(s, MACRO_UUID) {
		u, _ := uuid.NewV4()
		s = strings.Replace(s, MACRO_UUID, u.String(), 1)
	}
	for strings.Contains(s, MACRO_TIMESTAMP) {
		ts := time.Now().Format(TIMESTAMP_FORMAT)
		s = strings.Replace(s, MACRO_TIMESTAMP, ts, 1)
	}
	re := regexp.MustCompile(MACRO_RANDOM_UINT)
	for re.MatchString(s) {
		sub := re.FindStringSubmatch(s)
		n, _ := strconv.Atoi(sub[1])
		s = strings.Replace(s, sub[0], strconv.Itoa(mrand.Intn(n)), 1)
	}
	re = regexp.MustCompile(MACRO_RANDOM_HEX)
	for re.MatchString(s) {
		sub := re.FindStringSubmatch(s)
		n, _ := strconv.Atoi(sub[1])
		len := n / 2
		r := make([]byte, len, len)
		crand.Read(r)
		h := hex.EncodeToString(r)
		if sub[2] == "u" {
			h = strings.ToUpper(h)
		}
		s = strings.Replace(s, sub[0], h, 1)
	}
	return []byte(s)
}