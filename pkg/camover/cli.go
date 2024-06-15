package camover

import (
	"bufio"
	"flag"
	"os"
	"sync"

	"github.com/seaung/camover-go/pkg/logger"
)

type CamOverCli struct {
	logger.Logger
	thread     bool
	output     string
	input      string
	address    string
	shodanKey  string
	zoomeyeKey string
	pages      int
}

type CamOverResult struct {
    Username string
    Password string
}

func (cli *CamOverCli)initOptions() {
    flag.BoolVar(&cli.thread, "t", false, "Use threads for fastes work.")
    flag.StringVar(&cli.output, "o", "", "Ouput result to file.")
    flag.StringVar(&cli.input, "i", "", "Input file of file.")
    flag.StringVar(&cli.address, "a", "", "Single address.")
    flag.StringVar(&cli.shodanKey, "shodan", "", "Shodan API key for exploiting devices over Inernet.")
    flag.StringVar(&cli.zoomeyeKey, "zoomeye", "", "Zoome API key for exploiting devices over Internet.")
    flag.IntVar(&cli.pages, "p", 100, "Number of pages you want to get from ZoomEye")
    flag.Parse()
}

func (cli *CamOverCli) crackAddresses(addresses []string) {
    var wg sync.WaitGroup

    for _, addr := range addresses {
        wg.Add(1)

        go func(addr string) {
            defer wg.Done()
        }(addr)
    }

    wg.Wait()
}

func (cli *CamOverCli) readAddressesFromFile() ([]string, error) {
    fs, err := os.Open(cli.input)
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

func (cli *CamOverCli) exploit(address string) (*CamOverResult, error) {
    return nil, nil
}

func (cli *CamOverCli) start() {}
