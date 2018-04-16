package main

import (
	"fmt"

	"github.com/shaardie/discovery/base"
	"github.com/shaardie/discovery/network"
	"github.com/shaardie/discovery/utils"
)

func test(topic string, tests []utils.Test) {
	fmt.Printf("\n%v:\n\n", topic)
	for _, t := range tests {
		fmt.Printf("%v...", t.Description())
		r, err := t.Run()
		if err != nil {
			fmt.Printf("failed, %v", err)
			continue
		}
		fmt.Println(r)
	}
}

func main() {
	test("BASE", base.Tests())
	test("NETWORK", network.Tests())
}
