// Package main demonstrates how to use the Equinix Smart Hands API v1 to retrieve
// available Smart Hands service types.
//
// This example shows how to:
//   - Configure OAuth2 authentication for the Equinix API
//   - Create a Smart Hands v1 API client
//   - Retrieve available Smart Hands service types
//
// Prerequisites:
//   - Set EQUINIX_API_CLIENTID and EQUINIX_API_CLIENTSECRET environment variables
//   - Have Smart Hands permissions on your Equinix account
//
// Usage:
//
//	go run main.go
package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/equinix/equinix-sdk-go/services/smarthandsv1"

	"github.com/equinix/equinix-sdk-go/extensions/equinixoauth2"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	CORRELATION_ID = "smarthandsv1-example-7d4a2b8e"
)

func main() {
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

	// Configure the Smart Hands v1 client
	configuration := smarthandsv1.NewConfiguration()
	configuration.HTTPClient = standardClient
	configuration.AddDefaultHeader("X-SOURCE", "API")
	configuration.AddDefaultHeader("X-CORRELATION-ID", CORRELATION_ID)
	client := smarthandsv1.NewAPIClient(configuration)

	// Get available Smart Hands service types
	log.Println("Retrieving available Smart Hands service types...")

	typesResp, resp, err := client.SmarthandsApi.SmartHandTypes(ctx).Execute()
	if err != nil {
		log.Printf("Error retrieving Smart Hands types: %v", err)
		if resp != nil {
			log.Printf("Response status: %s", resp.Status)
			if respBody, err := io.ReadAll(resp.Body); err == nil {
				log.Printf("Response body: %s", string(respBody))
			}
		}
		os.Exit(1)
	}

	// Display response
	log.Printf("Successfully retrieved Smart Hands service types")
	log.Printf("Response Status: %s", resp.Status)
	log.Printf("Response Code: %d", resp.StatusCode)
	log.Printf("Response: %+v", typesResp)
}
