package test

import (
	"blog/database"
	"fmt"
	"testing"
)

func TestSetToken(t *testing.T) {
	database.SetToken("123456","78900")
}

func TestGetToken(t *testing.T) {
	res,err := database.GetToken("1256")
	fmt.Printf("res: %v\n", res)
	fmt.Printf("err: %v\n", err)
}
