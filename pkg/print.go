package pkg

import "fmt"

func PrintEmpty(message string) {
	fmt.Printf("\033[1K\r %s\n", message)
}

func Process(message string) {
	fmt.Printf("\033[1K\r\033[1;34m[*]\033[0m %s\n", message)
}

func Success(message string) {
	fmt.Printf("\033[1K\r\033[1;32m[+]\033[0m %s\n", message)
}

func Errorf(message string) {
	fmt.Printf("\033[1K\r\033[1;31m[-]\033[0m %s\n", message)
}

func Warning(message string) {
	fmt.Printf("\033[1K\r\033[1;33m[!]\033[0m %s\n", message)
}

func Info(message string) {
	fmt.Printf("\033[1K\r\033[1;77m[i]\033[0m %s\n", message)
}
