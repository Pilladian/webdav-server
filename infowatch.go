package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Send Logs to InfoWatch
func sendLogsToInfoWatch(pid string, authorized string, user string, method string, resource string) (int, error) {
	request_body, request_body_err := json.Marshal(map[string]string{
		"authorized": authorized,
		"user":       user,
		"method":     method,
		"resource":   resource,
		"time":       time.Now().Format(time.RFC3339),
	})

	if request_body_err != nil {
		return 1, request_body_err
	}

	url := os.Getenv("INFOWATCH_URL")
	rev_proxy_username := os.Getenv("INFOWATCH_REV_PROXY_USERNAME")
	rev_proxy_password := os.Getenv("INFOWATCH_REV_PROXY_PASSWORD")

	req, _ := http.NewRequest("POST", fmt.Sprintf(url, pid), bytes.NewBuffer(request_body))
	req.SetBasicAuth(rev_proxy_username, rev_proxy_password)
	re, re_err := CLIENT.Do(req)
	if re_err != nil {
		return 1, re_err
	}
	return re.StatusCode, nil
}
