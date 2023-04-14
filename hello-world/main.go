// main.go

package main

import (
	"flag"
	"fmt"
)

var (
	fSubj = flag.String("subj", "world", "")
)

func main() {
	flag.Parse()
	fmt.Printf("Hello %s!\n", *fSubj)
}
