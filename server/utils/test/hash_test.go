package util_test

import (
	"blog/utils"
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {
	a := utils.MD5("123456")
	fmt.Printf("a: %v\n", a)
}