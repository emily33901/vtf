package main

import (
	"log"

	"github.com/emily33901/vtf"
)

func main() {
	_, err := vtf.ReadFromFile("samples/read/test.vtf")
	if err != nil {
		log.Fatal(err)
	}
}
