package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Pilladian/logger"
)

func sendLogsToElastic(pid string, authorized bool, user string, method string, resource string) (int, error) {
	request_body, request_body_err := json.Marshal(map[string]interface{}{
		"authorized": authorized,
		"user":       user,
		"method":     method,
		"resource":   resource,
		"time":       time.Now().Format(time.RFC3339),
	})

	if request_body_err != nil {
		logger.Error(request_body_err.Error())
		return 1, request_body_err
	}

	url := os.Getenv("ELASTIC_SEARCH_URL")
	rev_proxy_username := os.Getenv("ELASTIC_SEARCH_USERNAME")
	rev_proxy_password := os.Getenv("ELASTIC_SEARCH_PASSWORD")

	req, _ := http.NewRequest("POST", fmt.Sprintf(url+"/%s/_doc", pid), bytes.NewBuffer(request_body))
	req.SetBasicAuth(rev_proxy_username, rev_proxy_password)
	req.Header.Set("Content-Type", "application/json")
	re, re_err := CLIENT.Do(req)
	if re_err != nil {
		return 1, re_err
	}
	body, _ := io.ReadAll(re.Body)
	logger.Info("Elastic Response: " + string(body))
	return re.StatusCode, nil
}
