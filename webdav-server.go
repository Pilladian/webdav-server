package main

import (
	"fmt"
	"net/http"

	"github.com/Pilladian/go-helper"
	"github.com/Pilladian/logger"
	"golang.org/x/net/webdav"
)

// GLOBAL VARS
var STORAGE_PATH string = "./data"
var SERVER_PORT int = 80
var WEBDAV_SERVER *webdav.Handler
var AUTH_USERS map[string]string

// initialize environment
func initialize() {
	logger.SetLogLevel(2)
	helper.CreatePath(STORAGE_PATH)

	// add authorized users
	AUTH_USERS = make(map[string]string)
	AUTH_USERS["dav_user"] = "pass"
}

// entrypoint
func main() {
	initialize()

	WEBDAV_SERVER = &webdav.Handler{
		FileSystem: webdav.Dir(STORAGE_PATH),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				logger.Error(fmt.Sprintf("%s : %s : %s", r.Method, r.URL, err))
			}
		},
	}

	logger.Info("starting webdav server")
	http.HandleFunc("/", handleRequests)
	http.HandleFunc("/healthy", healthyRequestHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", SERVER_PORT), nil)
}
