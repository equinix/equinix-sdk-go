package main

import (
	"context"
	"github.com/equinix/equinix-sdk-go/services/fabricv4"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/equinix/oauth2-go"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
)

func main() {
	ctx := context.Background()
	clientId := os.Getenv("EQUINIX_API_CLIENTID")
	clientSecret := os.Getenv("EQUINIX_API_CLIENTSECRET")
	baseURL := "https://api.equinix.com"
	authConfig := oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		BaseURL:      baseURL,
	}
	authClient := authConfig.New(ctx)
	authClient.Transport = logging.NewTransport("Equinix", authClient.Transport)

	transport := logging.NewTransport("Equinix Fabric (fabriv4)", authClient.Transport)

	retryClient := retryablehttp.NewClient()
	retryClient.HTTPClient.Transport = transport
	retryClient.RetryWaitMin = time.Second
	retryClient.RetryWaitMax = time.Second * 60
	standardClient := retryClient.StandardClient()

	baseURLParsed, _ := url.Parse(baseURL)

	configuration := fabricv4.NewConfiguration()
	configuration.Servers = fabricv4.ServerConfigurations{
		fabricv4.ServerConfiguration{
			URL: baseURLParsed.String(),
		},
	}
	configuration.HTTPClient = standardClient
	configuration.AddDefaultHeader("X-SOURCE", "API")
	configuration.AddDefaultHeader("X-CORRELATION-ID", "asdfdkioinasdoinfiek183859573")
	// configuration.AddDefaultHeader("Content-Type", "json")
	client := fabricv4.NewAPIClient(configuration)

	conn, resp, err := client.ConnectionsApi.GetConnectionByUuid(ctx, "<connection_uuid>").Execute()

	if err != nil {
		log.Printf("Error getting Fabric Connection: %v", err)
	}

	log.Printf("Response API body from create call: %s", resp.Body)
	log.Printf("Response struct from create API call: %v", conn)

}
