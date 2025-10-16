// Package main demonstrates how to use the Equinix Lookupv2 API to retrieve
// available patch panel information.
//
// This example shows how to:
//   - Configure OAuth2 authentication for the Equinix API
//   - Create a Lookupv2 API client
//   - Retrieve available patch panel information
//
// Prerequisites:
//   - Set EQUINIX_API_CLIENTID and EQUINIX_API_CLIENTSECRET environment variables
//
// Usage:
//
//	go run main.go -patchpanelid="<a valid patch panel ID>"
//
// Please note that the EQUINIX_API_CLIENTID defines which cages/cabinets/patch panels
// can be accessed by this example.
package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/equinix/equinix-sdk-go/services/lookupv2"

	"github.com/equinix/equinix-sdk-go/extensions/equinixoauth2"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	CORRELATION_ID = "lookupv2-example-7d4a2b8e"
)

func main() {
	// Initialize OAuth2 configuration using environment variables
	ctx := context.Background()
	clientId := os.Getenv("EQUINIX_API_CLIENTID")
	clientSecret := os.Getenv("EQUINIX_API_CLIENTSECRET")

	if clientId == "" || clientSecret == "" {
		log.Fatal("Please set EQUINIX_API_CLIENTID and EQUINIX_API_CLIENTSECRET environment variables")
	}

	// Parse command line flags
	var patchPanelId string
	flag.StringVar(&patchPanelId, "patchpanelid", "", "A patch panel ID like e.g. PP:0000:31231")
	flag.Parse()

	if patchPanelId == "" {
		log.Fatal("Please provide a patch panel ID using -patchpanelid flag")
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

	// Configure the Lookup v2 client
	configuration := lookupv2.NewConfiguration()
	configuration.HTTPClient = standardClient
	configuration.AddDefaultHeader("X-SOURCE", "API")
	configuration.AddDefaultHeader("X-CORRELATION-ID", CORRELATION_ID)
	client := lookupv2.NewAPIClient(configuration)

	// Get available patch panel details
	log.Printf("Retrieving details for patch panel [%s]...\n", patchPanelId)

	ppDetail, resp, err := client.LookupApi.RetrievePatchPanelDetails(ctx, patchPanelId).Execute()
	if err != nil {
		log.Printf("Error retrieving patch panel details: %v", err)
		if resp != nil {
			log.Printf("Response status: %s", resp.Status)
			if respBody, err := io.ReadAll(resp.Body); err == nil {
				log.Printf("Response body: %s", string(respBody))
			}
		}
		os.Exit(1)
	}

	// Display response
	log.Printf("Successfully retrieved patch panel information")
	log.Printf("Response Status: %s", resp.Status)
	log.Printf("Response Code: %d", resp.StatusCode)
	log.Printf("Response: %+v", ppDetail)

	log.Printf("Patch panel has %d available ports and %d used ports\n",
		len(ppDetail.AvailablePorts), len(ppDetail.UsedPortsDetails))
}
