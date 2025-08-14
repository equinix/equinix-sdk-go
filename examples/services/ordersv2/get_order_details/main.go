// Package main demonstrates how to use the Equinix Orders API v2 to retrieve
// order details and manage order operations.
//
// This example shows how to:
//   - Configure OAuth2 authentication for the Equinix API
//   - Create an Orders v2 API client
//   - Retrieve order details by order ID
//   - Display order information including status, line items, and pricing
//
// Prerequisites:
//   - Set EQUINIX_API_CLIENTID and EQUINIX_API_CLIENTSECRET environment variables
//   - Have a valid Equinix order ID
//
// Usage:
//
//	go run main.go -order="<your_order_id>"
package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/equinix/equinix-sdk-go/services/ordersv2"

	"github.com/equinix/equinix-sdk-go/extensions/equinixoauth2"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	CORRELATION_ID = "ordersv2-example-8e5b9c6f"
)

func main() {
	// Parse command line flags
	var orderID string
	flag.StringVar(&orderID, "order", "", "Order ID to retrieve")
	flag.Parse()

	if orderID == "" {
		log.Fatal("Please provide an order ID using -order flag")
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

	// Configure the Orders v2 client
	configuration := ordersv2.NewConfiguration()
	configuration.HTTPClient = standardClient
	configuration.AddDefaultHeader("X-SOURCE", "API")
	configuration.AddDefaultHeader("X-CORRELATION-ID", CORRELATION_ID)
	client := ordersv2.NewAPIClient(configuration)

	// Call the GET order details endpoint
	log.Printf("Retrieving order details for order ID: %s", orderID)

	order, resp, err := client.OrdersApi.GETOrderDetails(ctx, orderID).Execute()
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

	// Print basic order information
	log.Printf("Order ID: %s", order.GetOrderId())
	log.Printf("Status: %s", order.GetStatus())
	if accountNumber, ok := order.GetAccountNumberOk(); ok && accountNumber != nil {
		log.Printf("Account Number: %s", *accountNumber)
	}

	// For debugging, show the full raw response body
	if respBody, err := io.ReadAll(resp.Body); err == nil {
		log.Printf("Response Body:\n%s", string(respBody))
	}
}
