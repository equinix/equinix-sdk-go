# \PrecisionTimeServicesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTimeServices**](PrecisionTimeServicesApi.md#CreateTimeServices) | **Post** /fabric/v4/timeServices | Create Time Service
[**DeleteTimeServiceById**](PrecisionTimeServicesApi.md#DeleteTimeServiceById) | **Delete** /fabric/v4/timeServices/{serviceId} | Delete Time Service by ID
[**FulfillTimeServices**](PrecisionTimeServicesApi.md#FulfillTimeServices) | **Put** /fabric/v4/timeServices/{serviceId} | Update Time Service
[**GetTimeServicesById**](PrecisionTimeServicesApi.md#GetTimeServicesById) | **Get** /fabric/v4/timeServices/{serviceId} | Get Time Service by ID
[**GetTimeServicesConnectionsByServiceId**](PrecisionTimeServicesApi.md#GetTimeServicesConnectionsByServiceId) | **Get** /fabric/v4/timeServices/{serviceId}/connections | Get Connection Links
[**SearchTimeServices**](PrecisionTimeServicesApi.md#SearchTimeServices) | **Post** /fabric/v4/timeServices/search | Search Time Services
[**UpdateTimeServicesById**](PrecisionTimeServicesApi.md#UpdateTimeServicesById) | **Patch** /fabric/v4/timeServices/{serviceId} | Update Time Service by ID



## CreateTimeServices

> PrecisionTimeServiceResponse CreateTimeServices(ctx).PrecisionTimeServiceRequest(precisionTimeServiceRequest).Execute()

Create Time Service



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
	precisionTimeServiceRequest := *openapiclient.NewPrecisionTimeServiceRequest(openapiclient.TimeServicePrice_type("NTP"), "Name_example", *openapiclient.NewPrecisionTimePackageRequest(openapiclient.precisionTimePackageRequest_code("NTP_STANDARD")), []openapiclient.VirtualConnectionUuid{*openapiclient.NewVirtualConnectionUuid("Uuid_example")}, *openapiclient.NewIpv4("Primary_example", "Secondary_example", "NetworkMask_example")) // PrecisionTimeServiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeServicesApi.CreateTimeServices(context.Background()).PrecisionTimeServiceRequest(precisionTimeServiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeServicesApi.CreateTimeServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTimeServices`: PrecisionTimeServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeServicesApi.CreateTimeServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTimeServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **precisionTimeServiceRequest** | [**PrecisionTimeServiceRequest**](PrecisionTimeServiceRequest.md) |  | 

### Return type

[**PrecisionTimeServiceResponse**](PrecisionTimeServiceResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTimeServiceById

> PrecisionTimeServiceResponse DeleteTimeServiceById(ctx, serviceId).Execute()

Delete Time Service by ID



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
	serviceId := "557400f8-d360-11e9-bb65-2a2ae2dbcce4" // string | Service UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeServicesApi.DeleteTimeServiceById(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeServicesApi.DeleteTimeServiceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTimeServiceById`: PrecisionTimeServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeServicesApi.DeleteTimeServiceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTimeServiceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrecisionTimeServiceResponse**](PrecisionTimeServiceResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FulfillTimeServices

> PrecisionTimeServiceResponse FulfillTimeServices(ctx, serviceId).PrecisionTimeServiceRequest(precisionTimeServiceRequest).Execute()

Update Time Service



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
	serviceId := "557400f8-d360-11e9-bb65-2a2ae2dbcce4" // string | Service UUID
	precisionTimeServiceRequest := *openapiclient.NewPrecisionTimeServiceRequest(openapiclient.TimeServicePrice_type("NTP"), "Name_example", *openapiclient.NewPrecisionTimePackageRequest(openapiclient.precisionTimePackageRequest_code("NTP_STANDARD")), []openapiclient.VirtualConnectionUuid{*openapiclient.NewVirtualConnectionUuid("Uuid_example")}, *openapiclient.NewIpv4("Primary_example", "Secondary_example", "NetworkMask_example")) // PrecisionTimeServiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeServicesApi.FulfillTimeServices(context.Background(), serviceId).PrecisionTimeServiceRequest(precisionTimeServiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeServicesApi.FulfillTimeServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FulfillTimeServices`: PrecisionTimeServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeServicesApi.FulfillTimeServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFulfillTimeServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **precisionTimeServiceRequest** | [**PrecisionTimeServiceRequest**](PrecisionTimeServiceRequest.md) |  | 

### Return type

[**PrecisionTimeServiceResponse**](PrecisionTimeServiceResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeServicesById

> PrecisionTimeServiceResponse GetTimeServicesById(ctx, serviceId).Execute()

Get Time Service by ID



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
	serviceId := "557400f8-d360-11e9-bb65-2a2ae2dbcce4" // string | Service UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeServicesApi.GetTimeServicesById(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeServicesApi.GetTimeServicesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeServicesById`: PrecisionTimeServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeServicesApi.GetTimeServicesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeServicesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrecisionTimeServiceResponse**](PrecisionTimeServiceResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeServicesConnectionsByServiceId

> PrecisionTimeServiceConnectionsResponse GetTimeServicesConnectionsByServiceId(ctx, serviceId).Execute()

Get Connection Links



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
	serviceId := "557400f8-d360-11e9-bb65-2a2ae2dbcce4" // string | Service UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeServicesApi.GetTimeServicesConnectionsByServiceId(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeServicesApi.GetTimeServicesConnectionsByServiceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeServicesConnectionsByServiceId`: PrecisionTimeServiceConnectionsResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeServicesApi.GetTimeServicesConnectionsByServiceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeServicesConnectionsByServiceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrecisionTimeServiceConnectionsResponse**](PrecisionTimeServiceConnectionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTimeServices

> ServiceSearchResponse SearchTimeServices(ctx).TimeServicesSearchRequest(timeServicesSearchRequest).Execute()

Search Time Services



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
	timeServicesSearchRequest := *openapiclient.NewTimeServicesSearchRequest() // TimeServicesSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeServicesApi.SearchTimeServices(context.Background()).TimeServicesSearchRequest(timeServicesSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeServicesApi.SearchTimeServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTimeServices`: ServiceSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeServicesApi.SearchTimeServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTimeServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeServicesSearchRequest** | [**TimeServicesSearchRequest**](TimeServicesSearchRequest.md) |  | 

### Return type

[**ServiceSearchResponse**](ServiceSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTimeServicesById

> PrecisionTimeServiceResponse UpdateTimeServicesById(ctx, serviceId).PrecisionTimeChangeOperation(precisionTimeChangeOperation).Execute()

Update Time Service by ID



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
	serviceId := "557400f8-d360-11e9-bb65-2a2ae2dbcce4" // string | Service UUID
	precisionTimeChangeOperation := []openapiclient.PrecisionTimeChangeOperation{*openapiclient.NewPrecisionTimeChangeOperation(openapiclient.CloudRouterChangeOperation_op("replace"), openapiclient.precisionTimeChangeOperation_path("/name"), interface{}(123))} // []PrecisionTimeChangeOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeServicesApi.UpdateTimeServicesById(context.Background(), serviceId).PrecisionTimeChangeOperation(precisionTimeChangeOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeServicesApi.UpdateTimeServicesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTimeServicesById`: PrecisionTimeServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeServicesApi.UpdateTimeServicesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTimeServicesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **precisionTimeChangeOperation** | [**[]PrecisionTimeChangeOperation**](PrecisionTimeChangeOperation.md) |  | 

### Return type

[**PrecisionTimeServiceResponse**](PrecisionTimeServiceResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

