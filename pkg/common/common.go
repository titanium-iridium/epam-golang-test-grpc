package common

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const DateTimeFormat = "2006-01-02 15:04:05"

type Config struct {
	Address string
}

// gRPC communication message structure
type Message struct {
	Time time.Time
	Text string
}

// Get app config
func GetConfig() *Config {
	var (
		host string
		port string
	)

	flag.StringVar(&host, "host", "127.0.0.1", "Server host")
	flag.StringVar(&port, "port", "9999", "Server port")
	flag.Parse()

	host = strings.TrimSuffix(host, ":")
	port = strings.TrimPrefix(port, ":")

	return &Config{Address: host + ":" + port}
}

// Endlessly read user input from console and put it to message channel
func ConsoleInput(channel chan string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		//fmt.Printf("[%s] Input message: ", time.Now().Format(DateTimeFormat))
		fmt.Print("Input message: ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSuffix(msg, "\n")
		channel <- msg
	}
}

// Endlessly print messages from message channel to console
func ConsoleOutput(channel chan *Message) {
	for msg := range channel {
		fmt.Printf("[%s] %s\n", msg.Time.Format(DateTimeFormat), msg.Text)
	}
}

// Output message with error description
func LogError(message string, err error) {
	txt := fmt.Sprintf("%s: %s", message, err)
	// marker line with same length as message text
	line := strings.Repeat("=", len(txt))
	fmt.Println("\n" + line + "\n" + txt + "\n" + line)
}
