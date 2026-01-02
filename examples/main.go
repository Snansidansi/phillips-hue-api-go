package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/snansidansi/hueapi"
	"github.com/snansidansi/hueapi/models"
)

func main() {
	println("Hue API test.")
	println()

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error while loading .env file: %v", err)
		os.Exit(1)
	}

	bridge := models.Bridge{
		Id:       os.Getenv("HUE_BRIDGE_ID"),
		IPAdress: os.Getenv("HUE_BRIDGE_IP"),
	}

	registerInBridge(bridge)
}

func discoverBridges() {
	bridges, _ := hueapi.DiscoverBridges(http.DefaultClient)
	fmt.Printf("%+v\n", bridges)
}

func registerInBridge(bridge models.Bridge) {
	devicename := "HueApi-testing"
	generateClientKey := true
	username, clientkey, err := hueapi.RegisterBridge(http.DefaultClient, &bridge, devicename, generateClientKey)
	fmt.Printf("Username: %v\nClientkey: %v\nError: %v\n", username, clientkey, err)

	if err != nil {
		var hueErr *hueapi.HueError
		if errors.As(err, &hueErr) {
			if hueErr.Type == hueapi.ErrLinkButtonNotPressed {
				fmt.Println("Link button not pressed")
			}
		}
	}
}
