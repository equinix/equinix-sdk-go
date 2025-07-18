package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/equinix/equinix-sdk-go/extensions/equinixoauth2"
	"github.com/equinix/equinix-sdk-go/services/eiav2"
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
	// and creates an EIAv2 API client from the result
	configuration := eiav2.NewConfiguration()
	configuration.HTTPClient = standardClient
	configuration.AddDefaultHeader("X-SOURCE", "API")
	configuration.AddDefaultHeader("X-CORRELATION-ID", "asdfdkioinasdoinfiek183859573")
	client := eiav2.NewAPIClient(configuration)

	// From here to the end of the example we are defining the details of our POST /internetAccess/v2/services
	// request that will create an Equinix Internet Access service on top of the connection_uuid
	// you've passed in as a command line parameter
	eiaServiceRequest := eiav2.ServiceRequest{}
	eiaServiceRequest.SetName("EIAv2_SDK_Testing")
	eiaServiceRequest.SetType(eiav2.SERVICETYPEV2_SINGLE)
	eiaServiceRequest.SetConnections([]string{connection_uuid})

	routingProtocol := eiav2.DirectRoutingProtocolRequest{}
	routingProtocol.SetType(eiav2.ROUTINGPROTOCOLTYPE_DIRECT)
	routingProtocol.SetName("RP_EIAv2_SDK_Testing")

	ipv4 := eiav2.DirectRoutingProtocolIpv4Request{}
	ipv4CustomerRoute := eiav2.CustomerRouteIpv4Request{}
	ipv4CustomerRoute.SetPrefix("198.51.100.0/24")
	ipv4.SetCustomerRoutes([]eiav2.CustomerRouteIpv4Request{ipv4CustomerRoute})
	ipv4PeeringRoute := eiav2.DirectPeeringIpv4Request{}
	ipv4PeeringRoute.SetEquinixPeerIps([]string{"198.51.100.8"})
	ipv4.SetPeerings([]eiav2.DirectPeeringIpv4Request{ipv4PeeringRoute})

	ipv6 := eiav2.DirectRoutingProtocolIpv6Request{}
	ipv6CustomerRoute := eiav2.CustomerRouteIpv6Request{}
	ipv6CustomerRoute.SetPrefix("2001:db8::/32")
	ipv6.SetCustomerRoutes([]eiav2.CustomerRouteIpv6Request{ipv6CustomerRoute})
	ipv6PeeringRoute := eiav2.DirectPeeringIpv6Request{}
	ipv6PeeringRoute.SetEquinixPeerIps([]string{"2001:db8::34f4:0:f3dd:1"})
	ipv6.SetPeerings([]eiav2.DirectPeeringIpv6Request{ipv6PeeringRoute})

	routingProtocol.SetIpv4(ipv4)
	routingProtocol.SetIpv6(ipv6)
	serviceRequestRoutingProtocol := eiav2.DirectRoutingProtocolRequestAsServiceRequestAllOfRoutingProtocol(&routingProtocol)
	eiaServiceRequest.SetRoutingProtocol(serviceRequestRoutingProtocol)

	eiaService, resp, err := client.EIAServiceApi.CreateEquinixInternetAccessv2(ctx).ServiceRequest(eiaServiceRequest).Execute()
	if err != nil {
		log.Printf("Error creating EIA Service: %v", err)
	}

	log.Printf("Response API body from create call: %s", resp.Body)
	log.Printf("Response struct from create API call: %v", eiaService)
}
