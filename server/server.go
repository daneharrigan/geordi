package server

import (
	"github.com/daneharrigan/geordi/command"
	"github.com/daneharrigan/geordi/logger"
	"github.com/daneharrigan/geordi/respond"
	"net"
	"os"
)

func Run(port string) {
	logger.Infof("at=start port=%s", port)
	listen(port)
}

func listen(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Errorf("fn=listen error=%q", err)
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()
		logger.Infof("fn=Accept addr=%q", conn.RemoteAddr())

		if err != nil {
			logger.Infof("fn=Accept error=%q", err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	c, err := command.Parse(conn)
	if err != nil {
		logger.Infof("fn=Parse error=q", err)
		respond.Errorf(conn, "%s", err)
		return
	}

	b, err := c.Run()
	if err != nil {
		logger.Infof("fn=Parse error=q", err)
		respond.Errorf(conn, "%s", err)
		return
	}

	respond.Sendf(conn, "%s", b)
}
