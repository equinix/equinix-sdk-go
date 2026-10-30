# \InternetExchangeServicesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExchangeService**](InternetExchangeServicesApi.md#CreateExchangeService) | **Post** /fabric/v4/exchangeServices | Create Internet Exchange Service
[**CreateExchangeServiceAction**](InternetExchangeServicesApi.md#CreateExchangeServiceAction) | **Post** /fabric/v4/exchangeServices/{exchangeServiceId}/actions | Create Internet Exchange Service Action
[**GetExchangeServiceById**](InternetExchangeServicesApi.md#GetExchangeServiceById) | **Get** /fabric/v4/exchangeServices/{exchangeServiceId} | Get Internet Exchange Service
[**SearchExchangeService**](InternetExchangeServicesApi.md#SearchExchangeService) | **Post** /fabric/v4/exchangeServices/search | Search Internet Exchange Services
[**UpdateExchangeServiceById**](InternetExchangeServicesApi.md#UpdateExchangeServiceById) | **Patch** /fabric/v4/exchangeServices/{exchangeServiceId} | Update Internet Exchange Service



## CreateExchangeService

> ExchangeServiceResponse CreateExchangeService(ctx).ExchangeServiceCreateRequest(exchangeServiceCreateRequest).Execute()

Create Internet Exchange Service



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
	exchangeServiceCreateRequest := *openapiclient.NewExchangeServiceCreateRequest(openapiclient.ExchangeServiceCreateRequest_type("IX"), "ix_exchange_service", *openapiclient.NewPublicPeeringConnectionRequest(), *openapiclient.NewRoutingProtocolRequest()) // ExchangeServiceCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetExchangeServicesApi.CreateExchangeService(context.Background()).ExchangeServiceCreateRequest(exchangeServiceCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetExchangeServicesApi.CreateExchangeService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExchangeService`: ExchangeServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InternetExchangeServicesApi.CreateExchangeService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExchangeServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchangeServiceCreateRequest** | [**ExchangeServiceCreateRequest**](ExchangeServiceCreateRequest.md) |  | 

### Return type

[**ExchangeServiceResponse**](ExchangeServiceResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExchangeServiceAction

> ExchangeServiceActionResponse CreateExchangeServiceAction(ctx, exchangeServiceId).ExchangeServiceActionRequest(exchangeServiceActionRequest).Execute()

Create Internet Exchange Service Action



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
	exchangeServiceId := "a8ba52de-faae-43b5-b0b1-6904d37ee011" // string | Internet Exchange Service Id
	exchangeServiceActionRequest := *openapiclient.NewExchangeServiceActionRequest(openapiclient.ExchangeServiceActionRequest_type("EXCHANGE_SERVICE_ACTIVATE")) // ExchangeServiceActionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetExchangeServicesApi.CreateExchangeServiceAction(context.Background(), exchangeServiceId).ExchangeServiceActionRequest(exchangeServiceActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetExchangeServicesApi.CreateExchangeServiceAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExchangeServiceAction`: ExchangeServiceActionResponse
	fmt.Fprintf(os.Stdout, "Response from `InternetExchangeServicesApi.CreateExchangeServiceAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchangeServiceId** | **string** | Internet Exchange Service Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExchangeServiceActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exchangeServiceActionRequest** | [**ExchangeServiceActionRequest**](ExchangeServiceActionRequest.md) |  | 

### Return type

[**ExchangeServiceActionResponse**](ExchangeServiceActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	exchangeServiceId := "a8ba52de-faae-43b5-b0b1-6904d37ee011" // string | Internet Exchange Service Id

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

Search Internet Exchange Services



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


## UpdateExchangeServiceById

> ExchangeServiceResponse UpdateExchangeServiceById(ctx, exchangeServiceId).ExchangeServiceReplaceOperation(exchangeServiceReplaceOperation).Execute()

Update Internet Exchange Service



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
	exchangeServiceId := "a8ba52de-faae-43b5-b0b1-6904d37ee011" // string | Internet Exchange Service Id
	exchangeServiceReplaceOperation := []openapiclient.ExchangeServiceReplaceOperation{*openapiclient.NewExchangeServiceReplaceOperation(openapiclient.ExchangeServiceOpEnum("replace"), "Path_example", interface{}(123))} // []ExchangeServiceReplaceOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetExchangeServicesApi.UpdateExchangeServiceById(context.Background(), exchangeServiceId).ExchangeServiceReplaceOperation(exchangeServiceReplaceOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetExchangeServicesApi.UpdateExchangeServiceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExchangeServiceById`: ExchangeServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InternetExchangeServicesApi.UpdateExchangeServiceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exchangeServiceId** | **string** | Internet Exchange Service Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExchangeServiceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exchangeServiceReplaceOperation** | [**[]ExchangeServiceReplaceOperation**](ExchangeServiceReplaceOperation.md) |  | 

### Return type

[**ExchangeServiceResponse**](ExchangeServiceResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

