package providers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetMetadata(url string, headers map[string]string) (string, error) {

	client := &http.Client{
		Timeout: time.Duration(5 * time.Second), // TODO: make configurable?
	}

	req, err := http.NewRequest("GET", url, nil)

	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Error [%v] for url [%s].", resp.StatusCode, url)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
