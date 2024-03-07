package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"

	"github.com/bytedance/sonic"
)

var (
	DefautHeader = JwtHeader{
		Algo: "HS256",
		Type: "JWT",
	}
)

type JwtHeader struct {
	Algo string `json:"alg"` //哈希算法，默认为HMAC SHA256(写为 HS256)
	Type string `json:"typ"` //令牌(token)类型，统一写为JWT
}

type JwtPayload struct {
	ID          string         `json:"jti"` //JWT ID用于标识该JWT
	Issue       string         `json:"iss"` //发行人。比如微信
	Audience    string         `json:"aud"` //受众人。比如王者荣耀
	Subject     string         `json:"sub"` //主题
	IssueAt     int64          `json:"iat"` //发布时间,精确到秒
	NotBefore   int64          `json:"nbf"` //在此之前不可用,精确到秒
	Expiration  int64          `json:"exp"` //到期时间,精确到秒
	UserDefined map[string]any `json:"ud"`  //用户自定义的其他字段
}

func GenerateJWT(header JwtHeader, payload JwtPayload, secret string) string {
	bs1, _ := sonic.Marshal(header)
	bs2 := base64.RawURLEncoding.EncodeToString(bs1)
	ps1, _ := sonic.Marshal(payload)
	ps2 := base64.RawURLEncoding.EncodeToString(ps1)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(bs2 + "." + ps2))
	signiture := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	return bs2 + "." + ps2 + "." + signiture
}

func VertifyJWT(jwt string, secret string) (*JwtHeader, *JwtPayload, error) {
	parts := strings.Split(jwt, ".")
	header, payload, signature := parts[0], parts[1], parts[2]
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(header + "." + payload))
	if signature != base64.RawURLEncoding.EncodeToString(h.Sum(nil)) {
		return nil, nil, errors.New("token vertify failed")
	}
	var header_info JwtHeader
	var payload_info JwtPayload

	h1, err := base64.RawURLEncoding.DecodeString(header)
	if err != nil {
		return nil, nil, err
	}
	h2, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return nil, nil, err
	}
	if err := sonic.Unmarshal(h1, &header_info); err != nil {
		return nil, nil, err
	}
	if err := sonic.Unmarshal(h2, &payload_info); err != nil {
		return nil, nil, err
	}
	return &header_info, &payload_info, nil
}
