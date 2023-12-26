package main

import (
	"generator/pkg/generator"
	"log"
	"os"
	"os/signal"
	"go.dedis.ch/kyber/v3/pairing"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	g := generator.NewGenerator(
		pairing.NewSuiteBn256(),
		"9999",
	)

	go func() {
		g.LaunchGrpcServer()
	}()
	
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	g.Close()
}
