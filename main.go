package main

import (
	"log"

	"github.com/miltian/homie/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute failed: %v", err)
	}
}
