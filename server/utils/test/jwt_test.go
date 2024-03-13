package util_test

import (
	"blog/utils"
	"fmt"
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	var secret = "12345678"
	var header = utils.JwtHeader{
		Algo: "sha256",
		Type: "JWT",
	}
	var payload = utils.JwtPayload{
		ID:          123,
		Issue:       "微信",
		Audience:    "群众",
		Subject:     "没啥事",
		IssueAt:     time.Now().Unix(),
		Expiration:  time.Now().Add(7 * time.Hour).Unix(),
		UserDefined: map[string]any{"name": "rurenyi", "age": 23},
	}
	jwt := utils.GenerateJWT(header, payload, secret)
	fmt.Printf("jwt: %v\n", jwt)

	var decode_header *utils.JwtHeader
	var decode_payload *utils.JwtPayload
	decode_header, decode_payload, err := utils.VertifyJWT(jwt, secret)
	fmt.Printf("decode_header: %v\n", decode_header)
	fmt.Printf("decode_payload: %v\n", decode_payload)
	fmt.Printf("err: %v\n", err)
}

func TestSecret(t *testing.T) {
	fmt.Printf("utils.Secret: %v\n", utils.Secret)
}
