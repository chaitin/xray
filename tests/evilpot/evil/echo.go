package evil

import (
	"io"
	"log"
	"net"
)

func ServeEchoServer(addr string) error {
	server, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			_, _ = io.Copy(conn, conn)
		}()
	}
}
