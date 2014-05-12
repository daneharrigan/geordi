package main

import (
	"crypto/tls"
	"github.com/daneharrigan/geordi/logger"
	"os"
	"fmt"
	"bufio"
	"flag"
)

var (
	host = flag.String("host", "localhost", "Service Host")
	port = flag.String("port", "5000", "TLS Service Port")
)

func main() {
	flag.Parse()
	conn, err := tls.Dial("tcp", addr(), config())
	if err != nil {
		logger.Errorf("ns=client fn=Dial error=%q", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		b, err := stdin.ReadBytes('\n')
		if err != nil {
			logger.Errorf("ns=client fn=ReadBytes error=%q", err)
			continue
		}

		writer.Write(b)
		writer.WriteByte('\r')
		if err := writer.Flush(); err != nil {
			logger.Errorf("ns=client fn=Flush error=%q", err)
			os.Exit(1)
		}

		b, err = reader.ReadBytes('\r')
		if err != nil {
			logger.Errorf("ns=client fn=ReadBytes error=%q", err)
			continue
		}

		fmt.Printf("%s\n", b)
	}
}

func addr() string {
	return *host + ":" + *port
}

func config() *tls.Config {
	return &tls.Config{ InsecureSkipVerify: true }
}
