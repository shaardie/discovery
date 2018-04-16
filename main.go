package main

import (
	"fmt"
	"os"

	"github.com/shaardie/discovery/network"
)

func main() {

	tests := network.Tests()
	for _, t := range tests {
		fmt.Printf("%v...", t.Description())
		r, err := t.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(r)
	}
}
