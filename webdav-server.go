package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/Pilladian/go-helper"
	"github.com/Pilladian/logger"
	"golang.org/x/net/webdav"
)

// GLOBAL VARS
var STORAGE_PATH string = "./data"
var SERVER_PORT int = 8080
var WEBDAV_SERVER *webdav.Handler
var AUTH_USERS map[string]string
var CLIENT *http.Client
var INFOWATCH_PID = os.Getenv("INFOWATCH_PID")

// initialize environment
func initialize() {
	logger.SetLogLevel(2)
	if create_path_err := helper.CreatePath(STORAGE_PATH); create_path_err != nil {
		logger.Error(create_path_err.Error())
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	CLIENT = &http.Client{}

	// add authorized users
	AUTH_USERS = make(map[string]string)
	AUTH_USERS[os.Getenv("USERNAME")] = os.Getenv("PASSWORD")
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
	http_server_err := http.ListenAndServe(fmt.Sprintf(":%d", SERVER_PORT), nil)
	if http_server_err != nil {
		logger.Error(http_server_err.Error())
	}
}
