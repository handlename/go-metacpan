package main

// go run main.go -q Mojo

import (
	"flag"
	"fmt"

	"github.com/handlename/go-metacpan"
)

func main() {
	var query string
	flag.StringVar(&query, "q", "", "search query")
	flag.Parse()

	dists, err := metacpan.SearchDistribution(query)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i, dist := range dists {
		fmt.Println("---")
		fmt.Printf("%d\n", i+1)
		fmt.Printf("name\t: %s\n", dist.Name())
		fmt.Printf("url\t: %s\n", dist.URL())
	}
}
