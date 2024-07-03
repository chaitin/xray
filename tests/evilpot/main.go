package main

import (
	"flag"
	"log"

	"github.com/chaitin/xray/tests/evilpot/evil"
)

func main() {
	evilHardAddr := flag.String("evil-hard", ":8887", "evil server 困难模式 监听地址")
	evilAddr := flag.String("evil", ":8888", "evil server 监听地址")
	echoAddr := flag.String("echo", ":8889", "echo server 监听地址")
	flag.Parse()
	go func() { log.Fatalln(evil.ServeEvilServer(*evilHardAddr, true)) }()
	go func() { log.Fatalln(evil.ServeEvilServer(*evilAddr, false)) }()
	go func() { log.Fatalln(evil.ServeEchoServer(*echoAddr)) }()
	select {}
}
