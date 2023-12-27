package main

import (

	"generator/pkg/generator"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	g := generator.NewGenerator(
		
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
