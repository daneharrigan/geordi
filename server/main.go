package main

import (
	"github.com/daneharrigan/geordi/logger"
	"github.com/daneharrigan/geordi/types"
	"github.com/daneharrigan/geordi/command"
	"github.com/daneharrigan/geordi/responder"
	"crypto/tls"
	"flag"
	"os"
	"net"
	"bufio"
)

type Operation struct {
	Command string
	Arguments []Argument
}

type Argument struct {
	Value []byte
	Type types.Type
}

var (
	port = flag.String("port", "5000", "TLS Service Port")
	pem  = flag.String("pem", "", "Path to .pem file")
	key  = flag.String("key", "", "Path to .key file")
)

func main() {
	flag.Parse()
	logger.Infof("ns=server at=start port=%s", *port)

	ln, err := tls.Listen("tcp", ":"+*port, config())
	if err != nil {
		logger.Errorf("ns=server fn=Listen error=%q", err)
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()
		logger.Infof("ns=server fn=Accept addr=%q", conn.RemoteAddr())
		if err != nil {
			logger.Errorf("ns=server fn=Accept error=%q", err)
			conn.Close()
			continue
		}

		go handle(conn)
	}
}

func config() *tls.Config {
	crt, err := tls.LoadX509KeyPair(*pem, *key)
	if err != nil {
		logger.Errorf("ns=server fn=LoadX509KeyPair error=%q", err)
		os.Exit(1)
	}

	return &tls.Config{
		ClientAuth: tls.NoClientCert,
		Certificates: []tls.Certificate{crt},
	}
}

func handle(conn net.Conn) {
	reader := bufio.NewReader(conn)
	respond := responder.New(conn)

	for {
		operation, err := reader.ReadBytes('\r')
		if err != nil {
			logger.Errorf("ns=server fn=ReadBytes error=%q", err)
			conn.Close()
			return
		}

		if err := command.Execute(operation, respond); err != nil {
			logger.Errorf("ns=server fn=Execute error=%q", err)
		}

		if err := respond.Flush(); err != nil {
			logger.Errorf("ns=server fn=Flush error=%q", err)
			conn.Close()
			return
		}
	}
}
