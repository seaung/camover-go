package logger

import "fmt"

type Logger struct{}

const (
	Reset     = "\033[0m"
	Blue      = "\033[1;34m"
	Green     = "\033[1;32m"
	Red       = "\033[1;31m"
	Yellow    = "\033[1;33m"
	Gray      = "\033[1;77m"
	ClearLine = "\033[1K\r"
)

func NewLogger() *Logger {
    return &Logger{}
}

func (l *Logger) printLoggerPrefix(prefix, msg string) {
    fmt.Print(ClearLine + prefix + msg + Reset + "\n")
}

func (l *Logger) printEmpty(msg string) {
    l.printLoggerPrefix("", msg)
}

func (l *Logger) Process(msg string) {
    l.printLoggerPrefix(Blue + "[*] " + Reset, msg)
}

func (l *Logger) Success(msg string) {
    l.printLoggerPrefix(Green + "[+] " + Reset, msg)
}

func (l *Logger) Errorw(msg string) {
    l.printLoggerPrefix(Red + "[-] " + Reset, msg)
}

func (l *Logger) Warning(msg string) {
    l.printLoggerPrefix(Yellow + "[!] " + Reset, msg)
}

func (l *Logger) Info(msg string) {
    l.printLoggerPrefix(Gray + "[i] " + Reset, msg)
}

