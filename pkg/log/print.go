package log

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	author  = "seaung"
	version = "1.0.0"
)

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

func ShowBanner() {
	name := fmt.Sprintf("camover-go (v.%s)", version)
	banner := `
	 _________ _____ ___  ____ _   _____  _____      ____ _____ 
	 / ___/ __ '/ __ '__ \/ __ \ | / / _ \/ ___/_____/ __ '/ __ \
	/ /__/ /_/ / / / / / / /_/ / |/ /  __/ /  /_____/ /_/ / /_/ /
	\___/\__,_/_/ /_/ /_/\____/|___/\___/_/         \__, /\____/ 
						       /____/        

	`
	all_lines := strings.Split(banner, "\n")
	w := len(all_lines)
	fmt.Println(banner)
	color.Yellow(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(name))/2, name)))
	color.Cyan(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(author))/2, author)))
	fmt.Println()
}
