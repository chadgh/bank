package main

import (
	"flag"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	runtests := flag.Bool("test", false, "run the tests")
	flag.Parse()
	if *runtests {
		if err := runtest(); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("yup")
	}
}
