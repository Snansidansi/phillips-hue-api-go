package main

import (
	"fmt"
	"net/http"

	"github.com/snansidansi/hueapi"
)

func main() {
	println("Hue API test.")
	println()
	discoverBridges()
}

func discoverBridges() {
	bridges, _ := hueapi.DiscoverBridges(http.DefaultClient)
	fmt.Printf("%+v\n", bridges)
}
