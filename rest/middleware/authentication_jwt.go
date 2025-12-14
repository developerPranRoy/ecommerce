package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/utils"
	"encoding/base64"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticationJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			utils.SenError(w, http.StatusBadRequest, "Unauthorized")
			return
		}

		headerArray := strings.Split(header, " ")
		if len(headerArray) != 2 {
			utils.SenError(w, http.StatusBadRequest, "Unauthorized")
			return
		}

		accesToken := headerArray[1]
		tokenParts := strings.Split(accesToken, ".")
		if len(tokenParts) != 3 {
			utils.SenError(w, http.StatusBadRequest, "Unauthorized")
			return
		}
		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		// cnf := config.GetConfig()

		byteArraySecret := []byte(m.cnf.JWTSecret)
		byteArrayMessage := []byte(message)

		hm := hmac.New(sha256.New, byteArraySecret)
		hm.Write(byteArrayMessage)

		hash := hm.Sum(nil)
		newSignature := base64UrlEncode(hash)
		if newSignature != signature {
			utils.SenError(w, http.StatusBadRequest, "Unauthorized, U changed Data")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
