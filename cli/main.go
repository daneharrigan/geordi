package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/daneharrigan/geordi/logger"
	"os"
)

var (
	host = flag.String("host", "localhost", "Service Host")
	port = flag.String("port", "5000", "TLS Service Port")
)

func main() {
	flag.Parse()
	for {
		conn, err := tls.Dial("tcp", addr(), config())
		if err != nil {
			logger.Errorf("ns=client fn=Dial error=%q", err)
			os.Exit(1)
		}

		accept(conn)
		logger.Infof("ns=client at=reconnect")
	}
}

func accept(conn *tls.Conn) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		b, err := stdin.ReadBytes('\n')
		if err != nil {
			logger.Errorf("ns=client fn=ReadBytes error=%q", err)
			return
		}

		writer.Write(b)
		writer.WriteByte('\r')
		if err := writer.Flush(); err != nil {
			logger.Errorf("ns=client fn=Flush error=%q", err)
			return
		}

		b, err = reader.ReadBytes('\r')
		if err != nil {
			logger.Errorf("ns=client fn=ReadBytes error=%q", err)
			return
		}

		fmt.Printf("%s\n", b)
	}
}

func addr() string {
	return *host + ":" + *port
}

func config() *tls.Config {
	return &tls.Config{InsecureSkipVerify: true}
}
