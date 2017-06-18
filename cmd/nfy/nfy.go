package main

import (
	"flag"
	"log"
	"os"

	"github.com/imishinist/nfy"
)

func main() {
	title := flag.String("t", "", "title")
	message := flag.String("m", "", "message")
	flag.Parse()

	if len(flag.Args()) <= 0 {
		flag.Usage()
		os.Exit(0)
	}

	if err := nfy.Notify(*title, *message); err != nil {
		log.Fatal(err)
	}
}
