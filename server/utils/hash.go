package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(s string) string{
	b := []byte(s)
	res := md5.Sum(b)
	out := fmt.Sprintf("%x",res)
	return out
}