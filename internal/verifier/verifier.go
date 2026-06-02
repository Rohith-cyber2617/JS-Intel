package verifier

import (
	"net/http"
	"time"
)

func VerifyURL(target string) (bool, int) {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(target)
	if err != nil {
		return false, 0
	}
	defer resp.Body.Close()

	return true, resp.StatusCode
}

func VerifyGraphQL(target string) bool {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(target)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true
	}

	return false
}
