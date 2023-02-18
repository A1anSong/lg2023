package lg

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateMD5String(s []byte) string {
	h := md5.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}
