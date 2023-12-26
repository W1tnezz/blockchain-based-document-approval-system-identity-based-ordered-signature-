package main

import (
	"generator/pkg/generator"
	"os"
	"os/signal"

	"go.dedis.ch/kyber/v3/pairing"
)

func main() {

	g := generator.NewGenerator(
		pairing.NewSuiteBn256().Suite,
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
