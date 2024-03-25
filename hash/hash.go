package hash

import (
	"crypto/md5"
	"fmt"
)

func Md5str(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	re := h.Sum(nil)
	md5Str := fmt.Sprintf("%x", re)

	return md5Str
}
