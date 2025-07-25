// Package main demonstrates how to use the Equinix Secure Cabinet API to check
// availability of secure cabinets across different IBX locations.
//
// This example shows how to:
//   - Configure OAuth2 authentication for the Equinix API
//   - Create a Secure Cabinet API client
//   - Query the availability endpoint for a specific billing account
//   - Display the results including power capacity, cabinet dimensions, and fabric port availability
//
// Prerequisites:
//   - Set EQUINIX_API_CLIENTID and EQUINIX_API_CLIENTSECRET environment variables
//   - Have a valid Equinix billing account number
//
// Usage:
//
//	go run main.go -account="<your_account_number>"
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/equinix/equinix-sdk-go/services/securecabinetv1"

	"github.com/equinix/equinix-sdk-go/extensions/equinixoauth2"
	"github.com/hashicorp/go-retryablehttp"
)

const CORRELATION_ID = "securecab-example-7d4a8c5e"

func main() {
	// Parse command line flags
	var accountNumber string
	flag.StringVar(&accountNumber, "account", "", "Billing account number")
	flag.Parse()

	if accountNumber == "" {
		log.Fatal("Please provide an account number using -account flag")
	}

	// Initialize OAuth2 configuration using environment variables
	ctx := context.Background()
	clientId := os.Getenv("EQUINIX_API_CLIENTID")
	clientSecret := os.Getenv("EQUINIX_API_CLIENTSECRET")

	if clientId == "" || clientSecret == "" {
		log.Fatal("Please set EQUINIX_API_CLIENTID and EQUINIX_API_CLIENTSECRET environment variables")
	}

	baseURL := "https://api.equinix.com"
	authConfig := equinixoauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		BaseURL:      baseURL,
	}
	authTransport := authConfig.New()

	// Configure retry client
	retryClient := retryablehttp.NewClient()
	retryClient.HTTPClient.Transport = authTransport
	retryClient.RetryWaitMin = time.Second
	retryClient.RetryWaitMax = time.Second * 60
	retryClient.Logger = log.New(os.Stdout, "", log.LstdFlags)
	retryClient.RequestLogHook = func(logger retryablehttp.Logger, req *http.Request, retry int) {
		if retry > 0 {
			log.Printf("Retrying request (attempt %d) with wait time between %v and %v", retry+1, retryClient.RetryWaitMin, retryClient.RetryWaitMax)
		}
	}
	standardClient := retryClient.StandardClient()

	// Configure the Secure Cabinet client
	configuration := securecabinetv1.NewConfiguration()
	configuration.HTTPClient = standardClient
	configuration.AddDefaultHeader("X-SOURCE", "API")
	configuration.AddDefaultHeader("X-CORRELATION-ID", CORRELATION_ID)
	client := securecabinetv1.NewAPIClient(configuration)

	// Call the availability endpoint
	log.Printf("Checking Secure Cabinet availability for account: %s", accountNumber)

	availability, resp, err := client.AvailabilityApi.GetProductsAvailability(ctx, accountNumber).Execute()
	if err != nil {
		log.Printf("Error checking availability: %v", err)
		if resp != nil {
			log.Printf("Response status: %s", resp.Status)
			log.Printf("Response body: %s", resp.Body)
		}
		os.Exit(1)
	}

	// Print the results
	log.Printf("Found %d IBX locations with Secure Cabinet availability", len(availability))

	// Collect formatted output
	var output []string
	for _, ibxAvailability := range availability {
		output = append(output, fmt.Sprintf("IBX: %s", ibxAvailability.GetIbx()))
		output = append(output, fmt.Sprintf("  Max cabinets per order: %d", ibxAvailability.GetMaximumNumberOfCabinetsToOrder()))
		output = append(output, fmt.Sprintf("  Min power draw per cabinet: %.2f kW", ibxAvailability.GetMinimumDrawCapacityPerCabinet()))
		output = append(output, fmt.Sprintf("  Max power draw per cabinet: %.2f kW", ibxAvailability.GetMaximumDrawCapacityPerCabinet()))

		dimensions := ibxAvailability.GetCabinetDimensions()
		output = append(output, "  Cabinet dimensions:")
		output = append(output, fmt.Sprintf("    Width: %d %s", dimensions.Width.GetValue(), dimensions.Width.GetUnit()))
		output = append(output, fmt.Sprintf("    Depth: %d %s", dimensions.Depth.GetValue(), dimensions.Depth.GetUnit()))
		output = append(output, fmt.Sprintf("    Height: %d %s", dimensions.Height.GetValue(), dimensions.Height.GetUnit()))

		if pdu, ok := ibxAvailability.GetPduConfigurationOk(); ok && pdu != nil {
			output = append(output, "  PDU available: Yes")
		} else {
			output = append(output, "  PDU available: No")
		}

		if fabric, ok := ibxAvailability.GetFabricPortSpeedOk(); ok && fabric != nil {
			output = append(output, "  Fabric port available: Yes")
		} else {
			output = append(output, "  Fabric port available: No")
		}

		output = append(output, "") // Empty line between IBX entries
	}

	// Limit output to first 30 lines to avoid flooding logs with availability data
	maxLines := 30
	if len(output) > maxLines {
		log.Printf("Showing first %d lines of availability data (total lines: %d):", maxLines, len(output))
		log.Println(strings.Join(output[:maxLines], "\n"))
		log.Printf("... truncated %d lines to keep logs readable", len(output)-maxLines)
	} else {
		log.Println("Availability data:")
		log.Println(strings.Join(output, "\n"))
	}
}
