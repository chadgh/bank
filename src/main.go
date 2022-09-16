package main

import (
	"flag"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	runtests := flag.Bool("test", false, "run the tests")
	verbose := flag.Bool("verbose", false, "verbose output")
	flag.Parse()

	if *runtests {
		if err := runtest(*verbose); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := runserver(*verbose); err != nil {
			log.Fatal(err)
		}
	}
}
