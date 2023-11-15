package main

import (
	"context"
	"fmt"
	"os"

	metal "github.com/equinix-labs/metal-go/metal/v1"
	"sigs.k8s.io/yaml"
)

func main() {
	// TODO Because the swagger2 spec can not declare that facilities can be
	// returned with plan results, including available_in results in:
	// available_in:
	// - href: ""
	include := []string{}
	exclude := []string{"available_in"}
	configuration := metal.NewConfiguration()
	configuration.Debug = true
	configuration.AddDefaultHeader("X-Auth-Token", os.Getenv("METAL_AUTH_TOKEN"))
	api_client := metal.NewAPIClient(configuration)
	resp, r, err := api_client.PlansApi.FindPlans(context.Background()).Include(include).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.FindPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// response from `FindPlans`: PlanList
	s, err := yaml.Marshal(resp.Plans)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "Response from `PlansApi.FindPlans`: %v\n", string(s))
}
