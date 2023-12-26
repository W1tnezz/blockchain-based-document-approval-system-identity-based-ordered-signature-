package main

import (

	
	"os"
	"os/signal"

)

func main() {


	go func() {

	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	// node.Stop()
}
