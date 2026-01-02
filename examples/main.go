package main

import (
	"encoding/json"
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

	apiKey := os.Getenv("HUE_BRIDGE_USERNAME")
	client := hueapi.NewClient(bridge, apiKey, nil, true)

	firstLightID := getLights(client)
	getLightByID(client, firstLightID)
}

func discoverBridges() {
	bridges, _ := hueapi.DiscoverBridges(http.DefaultClient)
	fmt.Printf("%+v\n", bridges)
}

func registerInBridge(bridge *models.Bridge) {
	devicename := "HueApi-testing"
	generateClientKey := true
	username, clientkey, err := hueapi.RegisterBridge(http.DefaultClient, bridge, devicename, generateClientKey)
	fmt.Printf("Username: %v\nClientkey: %v\nError: %v\n", username, clientkey, err)

	if err != nil {
		var hueErr *models.HueError
		if errors.As(err, &hueErr) {
			if hueErr.Type == models.ErrLinkButtonNotPressed {
				fmt.Println("Link button not pressed")
			}
		}
	}
}

func getLights(client *hueapi.Client) (id string) {
	lights, hueError, err := client.Lights.GetAllLights()
	// printStructFormatted(lights)
	fmt.Print("Hue error: ")
	printStructFormatted(hueError)
	fmt.Printf("Normal error: %v\n", err)

	return lights[0].ID
}

func getLightByID(client *hueapi.Client, id string) {
	light, hueError, err := client.Lights.GetLightByID(id)
	printStructFormatted(light)
	fmt.Printf("Hue error: %v\n", hueError)
	fmt.Printf("Normal error: %v\n", err)
}

func printStructFormatted(data any) {
	bytes, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(bytes))
}
