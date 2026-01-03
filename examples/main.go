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

	lights, err := getLights(client)
	if err != nil {
		fmt.Printf("Error getting lights: %v\n", err)
		os.Exit(1)
	}
	// fmt.Println("All Lights:")
	// printStructFormatted(lights)

	firstLightID := ""
	if len(lights) > 0 {
		firstLightID = lights[0].ID
	} else {
		fmt.Println("No lights found.")
		os.Exit(0)
	}

	light, err := getLightByID(client, firstLightID)
	if err != nil {
		fmt.Printf("Error getting light by ID: %v\n", err)
		os.Exit(1)
	}
	// fmt.Printf("Light by ID %s:\n", firstLightID)
	// printStructFormatted(light)

	SetLight(client, &light)

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

func getLights(client *hueapi.Client) ([]models.Light, error) {
	hueResp, err := client.Lights.GetAllLights()

	if err != nil {
		return nil, err
	}

	if len(hueResp.Errors) > 0 {
		return nil, fmt.Errorf("Hue API error: %+v", hueResp.Errors)
	}

	return hueResp.Data, nil
}

func getLightByID(client *hueapi.Client, id string) (models.Light, error) {
	hueResp, err := client.Lights.GetLightByID(id)

	if err != nil {
		return models.Light{}, err
	}

	if len(hueResp.Errors) > 0 {
		return models.Light{}, fmt.Errorf("Hue API error: %+v", hueResp.Errors)
	}

	if len(hueResp.Data) > 0 {
		return hueResp.Data[0], nil
	}

	fmt.Println("Fatal error no example lights found")
	os.Exit(1)
	return models.Light{}, nil
}

func SetLight(client *hueapi.Client, light *models.Light) {
	// lightEnabled := !light.On.On
	// brightness := rand.Intn(100) + 1
	// lightBuilder := builders.NewLightBuilder()
	// lightBuilder.SetOnOff(lightEnabled)
	// lightBuilder.Brightness(float64(brightness))
	// update := lightBuilder.Build()
	//
	// hueResp, err := client.Lights.SetLightState(light.ID, update)

	hueResp, err := client.Lights.On(light.ID)
	fmt.Print("Hue error: ")
	printStructFormatted(hueResp)
	fmt.Printf("Error: %v\n", err)

	hueResp, err = client.Lights.Identify(light.ID, 2000)
	fmt.Print("Hue error: ")
	printStructFormatted(hueResp)
	fmt.Printf("Error: %v\n", err)
}

func printStructFormatted(data any) {
	bytes, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(bytes))
}
