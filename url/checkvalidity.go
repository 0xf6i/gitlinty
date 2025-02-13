package url

import (
	"errors"
	"net/http"
)

func CheckValidity(urlString string) (bool, error) {
	// in courtesy of
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return false, errors.New("could not create a new HTTP request")
	}
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0")

	resp, err := client.Do(req)
	if err != nil {
		return false, errors.New("could not send the HTTP request to Github")
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 && resp.Header.Get("Server") == "GitHub.com" {
		return true, nil
	}

	return false, errors.New("link is not a Github URL")
}
