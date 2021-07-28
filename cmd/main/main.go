package main

import (
	"fmt"

	"github.com/verticalscope/guild-golang/internal"
)

func main() {
	fmt.Printf("I'm running on: %s", internal.SystemType)
}
