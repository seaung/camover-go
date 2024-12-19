### camover-go

camover-go是[EntySec/CamOver](https://www.github.com/EntySec/CamOver)Golang版本的实现.
目前前只是单纯的将python代码翻译成Golang代码,喜欢的就点个Start再走吧!

### 如何编译?

```bash
git clone https://github.com/seaung/camover-go.git

cd camover-go/
go mod tidy

cd camover-go/cmd/camover/

go build .
```

### 如何使用
```bash
➜  camover-go git:(main) ✗ ./cmd/camover/camover --help
NAME:
   CamOver - A camera exploit tool to extract credentials.

USAGE:
   CamOver [global options] command [command options]

COMMANDS:
   exploit  Exploit vulnerabilities in network cameras.
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
➜  camover-go git:(main) ✗ ./cmd/camover/camover exploit --help
NAME:
   CamOver exploit - Exploit vulnerabilities in network cameras.

USAGE:
   CamOver exploit [command options]

OPTIONS:
   --threads        Use threads for faster processing. (default: false)
   --output value   Output file to save results.
   --input value    Input file with list of addresses.
   --address value  Single address to test.
   --shodan value   Shodan API Key.
   --zoomeye value  Zoomeye API Key.
   --pages value    Number of Zoomeye pages to fetch. (default: 100)
   --delay value    Delay between threads (default: 100ms)
   --help, -h       show help

➜  camover-go git:(main) ✗ ./cmd/camover/camover exploit --address 192.168.10.110

```

---
that's all

