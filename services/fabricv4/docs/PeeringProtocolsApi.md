# \PeeringProtocolsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectionPeeringProtocol**](PeeringProtocolsApi.md#CreateConnectionPeeringProtocol) | **Post** /fabric/v4/connections/{uuid}/peeringProtocols | Create Peering Protocol



## CreateConnectionPeeringProtocol

> PeeringProtocolData CreateConnectionPeeringProtocol(ctx, uuid).ConnectionPeeringProtocolPostRequest(connectionPeeringProtocolPostRequest).Execute()

Create Peering Protocol



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | uuid
	connectionPeeringProtocolPostRequest := *openapiclient.NewConnectionPeeringProtocolPostRequest(openapiclient.ConnectionPeeringProtocolPostRequest_type("BGP"), "ix_bgp", "ix_bgp", int64(1100), "00:11:22:33:44:55", *openapiclient.NewPeeringConnectionIpv4(), *openapiclient.NewPeeringConnectionIpv6()) // ConnectionPeeringProtocolPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeeringProtocolsApi.CreateConnectionPeeringProtocol(context.Background(), uuid).ConnectionPeeringProtocolPostRequest(connectionPeeringProtocolPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeeringProtocolsApi.CreateConnectionPeeringProtocol``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectionPeeringProtocol`: PeeringProtocolData
	fmt.Fprintf(os.Stdout, "Response from `PeeringProtocolsApi.CreateConnectionPeeringProtocol`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionPeeringProtocolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionPeeringProtocolPostRequest** | [**ConnectionPeeringProtocolPostRequest**](ConnectionPeeringProtocolPostRequest.md) |  | 

### Return type

[**PeeringProtocolData**](PeeringProtocolData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

