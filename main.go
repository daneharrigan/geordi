package main

import (
	"github.com/daneharrigan/geordi/config"
	"github.com/daneharrigan/geordi/server"
)

func main() {
	c := config.New()
	server.Run(c.Port)
}
