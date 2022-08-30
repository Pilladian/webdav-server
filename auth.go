package main

import (
	"errors"
	"fmt"
)

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
