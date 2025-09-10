// Package main demonstrates how to use the Equinix OrderHistory API v1 to retrieve
// order details.
//
// This example shows how to:
//   - Configure OAuth2 authentication for the Equinix API
//   - Create an OrdersHistoryv1 API client
//   - Search Orders History
//   - Display order information including status, line items, and pricing
//
// Prerequisites:
//   - Set EQUINIX_API_CLIENTID and EQUINIX_API_CLIENTSECRET environment variables
//
// Usage:
//
//	go run main.go
package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/equinix/equinix-sdk-go/services/orderhistoryv1"

	"github.com/equinix/equinix-sdk-go/extensions/equinixoauth2"
	"github.com/hashicorp/go-retryablehttp"
)

const CORRELATION_ID = "orderhistory-example-7d4a8c5e"

func main() {
	// Initialize OAuth2 configuration using environment variables
	// ctx := context.Background()
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

	// Configure the Orderhistoryv1 client
	configuration := orderhistoryv1.NewConfiguration()
	configuration.HTTPClient = standardClient
	configuration.Debug = true
	configuration.AddDefaultHeader("X-SOURCE", "API")
	configuration.AddDefaultHeader("X-CORRELATION-ID", CORRELATION_ID)
	client := orderhistoryv1.NewAPIClient(configuration)

	// Call the GET order details endpoint
	log.Printf("Retrieving all orders for the account associated with EQUINIX_API_CLIENTID")

	req := orderhistoryv1.ApiPOSTOrdersHistoryRequest{
		ApiService: client.RetrieveOrdersApi,
	}.Authorization("Bearer 123").Body(
		orderhistoryv1.Orderhistoryapirequest{
			// You can filter orders by date range, status, order ID, etc.
			// For this example, we're retrieving all orders without filters.
			// See https://docs.equinix.com/api-catalog/orderhistoryv1 for details.
			Filters: orderhistoryv1.NewFilters(),
		},
	)

	order, resp, err := client.RetrieveOrdersApi.POSTOrdersHistoryExecute(req)
	if err != nil {
		log.Printf("Error retrieving order: %v", err)
		if resp != nil {
			log.Printf("Response status: %s", resp.Status)
			log.Printf("Response body: %s", resp.Body)
		}
		os.Exit(1)
	}

	// Success! Print basic info and raw response
	log.Printf("Successfully retrieved order details")
	log.Printf("Response Status: %s", resp.Status)
	log.Printf("Response Code: %d", resp.StatusCode)

	_ = order

	/*
		// Print basic order information
		log.Printf("Order ID: %s", order.GetOrderId())
		log.Printf("Status: %s", order.GetStatus())
		if accountNumber, ok := order.GetAccountNumberOk(); ok && accountNumber != nil {
			log.Printf("Account Number: %s", *accountNumber)
		}
	*/
	// For debugging, show the full raw response body
	if respBody, err := io.ReadAll(resp.Body); err == nil {
		log.Printf("Response Body:\n%s", string(respBody))
	}
}
