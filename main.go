package main

import (
	"github.com/NextTourPlan/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Println(err)
	}
}
