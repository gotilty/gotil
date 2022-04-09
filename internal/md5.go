package internal

import (
	"crypto/md5"
	"fmt"

	"github.com/gotilty/gotil/internal/errs"
)

//MD5 returns the MD5 checksum of the data.
//Just works with all primitive types.
//If given parameter is different type instead of string, firstly, it will convert the given parameter to string.
func Md5(a interface{}) (string, error) {
	if s, err := ToString(a); err == nil {
		return md5String(s)
	} else {
		return "", err
	}
}

func md5String(s string) (string, error) {
	if s == "" {
		return "", errs.NilReferenceTypeError()
	}
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data)), nil
}
