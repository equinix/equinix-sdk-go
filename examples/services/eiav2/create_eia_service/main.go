package main

import (
	"context"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"

	"github.com/equinix/equinix-sdk-go/services/eiav2"
	"github.com/equinix/oauth2-go"
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

	transport := logging.NewTransport("Equinix Fabric (eiav2)", authClient.Transport)

	retryClient := retryablehttp.NewClient()
	retryClient.HTTPClient.Transport = transport
	retryClient.RetryWaitMin = time.Second
	retryClient.RetryWaitMax = time.Second * 60
	standardClient := retryClient.StandardClient()

	baseURLParsed, _ := url.Parse(baseURL)

	configuration := eiav2.NewConfiguration()
	configuration.Servers = eiav2.ServerConfigurations{
		eiav2.ServerConfiguration{
			URL: baseURLParsed.String(),
		},
	}
	configuration.HTTPClient = standardClient
	configuration.AddDefaultHeader("X-SOURCE", "API")
	configuration.AddDefaultHeader("X-CORRELATION-ID", "asdfdkioinasdoinfiek183859573")
	client := eiav2.NewAPIClient(configuration)

	eiaServiceRequest := eiav2.ServiceRequest{}
	eiaServiceRequest.SetName("EIAv2_SDK_Testing")
	eiaServiceRequest.SetType(eiav2.SERVICETYPEV2_SINGLE)
	eiaServiceRequest.SetConnections([]string{"<connection_uuid>"})

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
