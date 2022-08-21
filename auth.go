package main

import (
	"encoding/base64"
	"errors"
	"fmt"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func authenticate(u string, p string) (bool, error) {
	if u == "" || p == "" {
		return true, nil
	}

	if AUTH_USERS[u] == p {
		return true, nil
	} else {
		return false, errors.New(fmt.Sprintf("Invalid credentials \"%s:%s\"", u, p))
	}
}
