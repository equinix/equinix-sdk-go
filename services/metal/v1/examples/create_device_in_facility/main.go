package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	metal "github.com/equinix-labs/metal-go/metal/v1"
	"sigs.k8s.io/yaml"
)

func main() {
	var projectId string
	flag.StringVar(&projectId, "projectId", "", "")
	flag.Parse()

	configuration := metal.NewConfiguration()
	configuration.Debug = true
	configuration.AddDefaultHeader("X-Auth-Token", os.Getenv("METAL_AUTH_TOKEN"))
	api_client := metal.NewAPIClient(configuration)
	deviceRequest := metal.CreateDeviceRequest{
		DeviceCreateInFacilityInput: &metal.DeviceCreateInFacilityInput{
			Facility:        []string{"fr2"},
			Plan:            "c3.medium.x86",
			OperatingSystem: "ubuntu_20_04",
		},
	}
	device, r, err := api_client.DevicesApi.CreateDevice(context.Background(), projectId).CreateDeviceRequest(deviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDevice`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// response from `CreateDevice`: Device
	s, err := yaml.Marshal(device)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "Response from `PlansApi.FindPlans`: %v\n", string(s))
}
