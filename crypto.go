package gotil

import "github.com/gotilty/gotil/internal"

//MD5 returns the MD5 checksum of the data.
//Just works with all primitive types.
//If given parameter is different type instead of string, firstly, it will convert the given parameter to string.
func Md5(a interface{}) (string, error) {
	return internal.Md5(a)
}
