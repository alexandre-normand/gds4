package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexandre-normand/gds4"
	"github.com/alexandre-normand/gds4/bluetooth"
)

func main() {
	log.Println("Connection to DS4")
	bt, err := bluetooth.NewBT("30:0E:D5:8E:7A:FC")
	if err != nil {
		log.Panic(err)
		return
	}
	ds4, err := gds4.NewDS4(bt)
	if err != nil {
		log.Panic(err)
		return
	}

	ossingal := make(chan os.Signal, syscall.SIGTERM)
	signal.Notify(ossingal, os.Interrupt)

	for {
		select {
		case <-ossingal:
			return
		default:
			log.Printf("%+v\n", ds4.Status)
		}
	}
}
