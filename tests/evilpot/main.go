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

	log.Println("Starting servers...")

	go func() {
		log.Printf("Starting evil server in hard mode on %s...\n", *evilHardAddr)
		if err := evil.ServeEvilServer(*evilHardAddr, true); err != nil {
			log.Fatalf("Evil server hard mode failed: %v\n", err)
		}
	}()

	go func() {
		log.Printf("Starting evil server on %s...\n", *evilAddr)
		if err := evil.ServeEvilServer(*evilAddr, false); err != nil {
			log.Fatalf("Evil server failed: %v\n", err)
		}
	}()

	go func() {
		log.Printf("Starting echo server on %s...\n", *echoAddr)
		if err := evil.ServeEchoServer(*echoAddr); err != nil {
			log.Fatalf("Echo server failed: %v\n", err)
		}
	}()

	select {}
}
