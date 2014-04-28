package respond

import (
	"fmt"
	"net"
)

func Errorf(conn net.Conn, format string, args ...interface{}) {
	fmt.Fprintf(conn, "-"+format+"\r\n", args...)
}

func Sendf(conn net.Conn, format string, args ...interface{}) {
	fmt.Fprintf(conn, "+"+format+"\r\n", args...)
}
