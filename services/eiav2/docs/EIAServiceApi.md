# \EIAServiceApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEquinixInternetAccessv2**](EIAServiceApi.md#CreateEquinixInternetAccessv2) | **Post** /internetAccess/v2/services | Creates Equinix Internet Access Service



## CreateEquinixInternetAccessv2

> ServiceV2 CreateEquinixInternetAccessv2(ctx).ServiceRequest(serviceRequest).Execute()

Creates Equinix Internet Access Service



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/eiav2"
)

func main() {
	serviceRequest := *openapiclient.NewServiceRequest(openapiclient.ServiceTypeV2("SINGLE"), []string{"Connections_example"}, *openapiclient.NewRoutingProtocolRequest(openapiclient.RoutingProtocolType("DIRECT"))) // ServiceRequest | Options for creating Equinix Internet Access Service product 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EIAServiceApi.CreateEquinixInternetAccessv2(context.Background()).ServiceRequest(serviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EIAServiceApi.CreateEquinixInternetAccessv2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEquinixInternetAccessv2`: ServiceV2
	fmt.Fprintf(os.Stdout, "Response from `EIAServiceApi.CreateEquinixInternetAccessv2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEquinixInternetAccessv2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceRequest** | [**ServiceRequest**](ServiceRequest.md) | Options for creating Equinix Internet Access Service product  | 

### Return type

[**ServiceV2**](ServiceV2.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

