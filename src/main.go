package main

import (
	"flag"
	"log"

	"github.com/apex/gateway"
	_ "github.com/lib/pq"
)

func main() {
	runtests := flag.Bool("test", false, "run the tests")
	rundevserver := flag.Bool("server", false, "run the dev server")
	verbose := flag.Bool("verbose", false, "verbose output")
	flag.Parse()

	if *runtests {
		if err := runtest(*verbose); err != nil {
			log.Fatal(err)
		}
	} else if *rundevserver {
		if err := runserver(*verbose); err != nil {
			log.Fatal(err)
		}
	} else {
		server := NewServer(false)
		gateway.ListenAndServe(":8080", server.router)
	}
}
