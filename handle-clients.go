package main

import (
	"fmt"
	"net/http"

	"github.com/Pilladian/logger"
)

func handleRequests(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()

	if ok {

		logger.Info(fmt.Sprintf("user \"%s\" - method %s -  %s ", username, r.Method, r.RequestURI))
		authorized, authentication_err := authenticate(username, password)
		if authentication_err != nil {
			logger.Error(authentication_err.Error())
		}

		if authorized {
			w.Header().Set("Timeout", "99999999")
			WEBDAV_SERVER.ServeHTTP(w, r)
			return
		}
	}

	logger.Error(fmt.Sprintf("user \"%s\" tried accessing %s", username, r.URL.String()))
	w.Header().Set("WWW-Authenticate", `Basic realm="BASIC WebDAV REALM"`)
	w.WriteHeader(401)
	w.Write([]byte("401 Unauthorized\n"))
}

func healthyRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "running")
}
