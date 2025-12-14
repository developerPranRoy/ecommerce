package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub       int    `json:"sub"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Isowner   string `json:"is_owner"`
}

func CreateJWT(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	byteArray, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerB64 := base64UrlEncode(byteArray)
	byteArraydata, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	payload64 := base64UrlEncode(byteArraydata)
	message := headerB64 + "." + payload64

	bytArraySecret := []byte(secret)
	bytArrayMessage := []byte(message)

	has := hmac.New(sha256.New, bytArraySecret)
	has.Write(bytArrayMessage)

	signature := has.Sum(nil)
	signatureB64 := base64UrlEncode(signature)
	jwt := headerB64 + "." + payload64 + "." + signatureB64
	return jwt, nil
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
