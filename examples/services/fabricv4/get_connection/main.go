package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/equinix/equinix-sdk-go/services/fabricv4"

	"github.com/equinix/equinix-sdk-go/extensions/equinixoauth2"
	"github.com/hashicorp/go-retryablehttp"
)

func main() {
	// Pulls the connection_uuid from the command line so that you can input it with
	// `go run main.go -connection_uuid="<your_connection_uuid>"`
	var connection_uuid string
	flag.StringVar(&connection_uuid, "connection_uuid", "", "")
	flag.Parse()

	// Initializes the Equinix oauth2 with the EQUINIX_API_CLIENTID and EQUINIX_API_CLIENTSECRET
	// details set as environment variables.
	ctx := context.Background()
	clientId := os.Getenv("EQUINIX_API_CLIENTID")
	clientSecret := os.Getenv("EQUINIX_API_CLIENTSECRET")
	baseURL := "https://api.equinix.com"
	authConfig := equinixoauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		BaseURL:      baseURL,
	}
	authTransport := authConfig.New()

	/*
		The following code is setting up configuration of the
		SDK for retries in the case that there is a network connection error or the server returns a 5XX error.
		It is using Hashicorp's retryablehttp Go package but the same could be done with an alternative like req.
		See https://req.cool/docs/tutorial/retry/ for details.
		For this example we're doing the following:
		* Adding auth to our retry client
		* Setting the minimum wait time between retries to 1 second
		* Setting the maximum time between retries to 60 seconds
		* Returning our config using the StandardClient from retryablehttp that returns http.Client with our
			config details
	*/
	retryClient := retryablehttp.NewClient()
	retryClient.HTTPClient.Transport = authTransport
	retryClient.RetryWaitMin = time.Second
	retryClient.RetryWaitMax = time.Second * 60
	standardClient := retryClient.StandardClient()

	// The following configuration allows us to add custom headers, add our retryable http.client to the SDK,
	// and creates an Fabricv4 API client from the result
	configuration := fabricv4.NewConfiguration()
	configuration.HTTPClient = standardClient
	configuration.AddDefaultHeader("X-SOURCE", "API")
	configuration.AddDefaultHeader("X-CORRELATION-ID", "asdfdkioinasdoinfiek183859573")
	client := fabricv4.NewAPIClient(configuration)

	// From here to the end of the example we are defining the details of our GET /fabric/v4/connections/{connectionId}
	// request that will retrieve a Fabric Connection from its UUID
	conn, resp, err := client.ConnectionsApi.GetConnectionByUuid(ctx, connection_uuid).Execute()
	if err != nil {
		log.Printf("Error getting Fabric Connection: %v", err)
	}

	log.Printf("Response API body from get call: %s", resp.Body)
	log.Printf("Response struct from get API call: %v", conn)
}
