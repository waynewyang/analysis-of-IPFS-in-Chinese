package main

import (
	"flag"
	"fmt"
	isd "gx/ipfs/QmZmmuAXgX73UQmX1jRKjTGmjzq24Jinqkq8vzkBtno4uX/go-is-domain"
	"os"
)

func main() {
	domain := flag.String("d", "", "input a domain")
	flag.Parse()
	if *domain == "" {
		fmt.Println("Please Input -d your domain!")
		os.Exit(0)
	}

	if isd.IsDomain(*domain) {
		fmt.Println("its a domain")
	} else {
		fmt.Println("not a domain")
	}
}
