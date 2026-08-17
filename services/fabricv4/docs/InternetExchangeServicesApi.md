# \InternetExchangeServicesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExchangeServiceById**](InternetExchangeServicesApi.md#GetExchangeServiceById) | **Get** /fabric/v4/exchangeServices/{exchangeServiceId} | Get Internet Exchange Service
[**SearchExchangeService**](InternetExchangeServicesApi.md#SearchExchangeService) | **Post** /fabric/v4/exchangeServices/search | Search Internet Exchange Service



## GetExchangeServiceById

> ExchangeServiceResponse GetExchangeServiceById(ctx, exchangeServiceId).Execute()

Get Internet Exchange Service



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
	exchangeServiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Internet Exchange Service Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetExchangeServicesApi.GetExchangeServiceById(context.Background(), exchangeServiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetExchangeServicesApi.GetExchangeServiceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExchangeServiceById`: ExchangeServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InternetExchangeServicesApi.GetExchangeServiceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchangeServiceId** | **string** | Internet Exchange Service Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeServiceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExchangeServiceResponse**](ExchangeServiceResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchExchangeService

> ExchangeServiceSearchResponse SearchExchangeService(ctx).ExchangeServiceSearchRequest(exchangeServiceSearchRequest).Execute()

Search Internet Exchange Service



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
	exchangeServiceSearchRequest := *openapiclient.NewExchangeServiceSearchRequest() // ExchangeServiceSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetExchangeServicesApi.SearchExchangeService(context.Background()).ExchangeServiceSearchRequest(exchangeServiceSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetExchangeServicesApi.SearchExchangeService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchExchangeService`: ExchangeServiceSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `InternetExchangeServicesApi.SearchExchangeService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchExchangeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchangeServiceSearchRequest** | [**ExchangeServiceSearchRequest**](ExchangeServiceSearchRequest.md) |  | 

### Return type

[**ExchangeServiceSearchResponse**](ExchangeServiceSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

