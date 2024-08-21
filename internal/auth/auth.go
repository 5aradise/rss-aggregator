package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(h http.Header) (string, error) {
	authH := h.Get("Authorization")
	if authH == "" {
		return "", errors.New("authorization header not found")
	}

	authVals := strings.Split(authH, " ")
	if len(authVals) != 2 {
		return "", errors.New("invalid authorization header")
	}
	if authVals[0] != "ApiKey" {
		return "", errors.New("invalid authorization header")
	}

	return authVals[1], nil
}
