package config

import "flag"

type Config struct {
	Port string
}

var port *string

func init() {
	port = flag.String("p", "5000", "TCP Service Port")
	flag.Parse()
}

func New() Config {
	return Config{*port}
}
