package main

import (
	"context"
	"fmt"
	"os"

	didactovela "github.com/axisvela/didactovela-open/sdk/go"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: version http://127.0.0.1:2056")
		os.Exit(2)
	}
	client, err := didactovela.New(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	info, err := client.Version(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%s %s (%s)\n", info.Product, info.Version, info.Stage)
}
