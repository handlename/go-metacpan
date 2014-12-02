package main

import (
	"flag"
	"fmt"

	"github.com/handlename/go-metacpan"
)

func main() {
	var query string
	flag.StringVar(&query, "q", "", "search query")
	flag.Parse()

	hits, err := metacpan.SearchDistribution(query)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i, hit := range hits {
		fmt.Println("---")
		fmt.Printf("%d\n", i)
		fmt.Printf("name\t:%s\n", hit.Name)
		fmt.Printf("url\t:%s\n", hit.URL())
	}
}
