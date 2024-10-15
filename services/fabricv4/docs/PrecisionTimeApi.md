# \PrecisionTimeApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTimeServices**](PrecisionTimeApi.md#CreateTimeServices) | **Post** /fabric/v4/timeServices | Create Time Service
[**DeleteTimeServiceById**](PrecisionTimeApi.md#DeleteTimeServiceById) | **Delete** /fabric/v4/timeServices/{serviceId} | Delete by ID.
[**FulfillTimeServices**](PrecisionTimeApi.md#FulfillTimeServices) | **Put** /fabric/v4/timeServices/{serviceId} | Configure Service.
[**GetTimeServicesById**](PrecisionTimeApi.md#GetTimeServicesById) | **Get** /fabric/v4/timeServices/{serviceId} | Get Service By ID.
[**GetTimeServicesConnectionsByServiceId**](PrecisionTimeApi.md#GetTimeServicesConnectionsByServiceId) | **Get** /fabric/v4/timeServices/{serviceId}/connections | Get Connection Links
[**GetTimeServicesPackageByCode**](PrecisionTimeApi.md#GetTimeServicesPackageByCode) | **Get** /fabric/v4/timeServicePackages/{packageCode} | Get Package By Code
[**GetTimeServicesPackages**](PrecisionTimeApi.md#GetTimeServicesPackages) | **Get** /fabric/v4/timeServicePackages | Get Packages
[**SearchTimeServices**](PrecisionTimeApi.md#SearchTimeServices) | **Post** /fabric/v4/timeServices/search | Search Time Services
[**UpdateTimeServicesById**](PrecisionTimeApi.md#UpdateTimeServicesById) | **Patch** /fabric/v4/timeServices/{serviceId} | Update By ID.



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
	precisionTimeServiceRequest := *openapiclient.NewPrecisionTimeServiceRequest(openapiclient.precisionTimeServiceRequest_type("NTP"), "Name_example", *openapiclient.NewPrecisionTimePackageRequest(openapiclient.precisionTimePackageRequest_code("NTP_STANDARD")), []openapiclient.VirtualConnectionUuid{*openapiclient.NewVirtualConnectionUuid("Uuid_example")}, *openapiclient.NewIpv4("Primary_example", "Secondary_example", "NetworkMask_example")) // PrecisionTimeServiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeApi.CreateTimeServices(context.Background()).PrecisionTimeServiceRequest(precisionTimeServiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeApi.CreateTimeServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTimeServices`: PrecisionTimeServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeApi.CreateTimeServices`: %v\n", resp)
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

Delete by ID.



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
	serviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeApi.DeleteTimeServiceById(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeApi.DeleteTimeServiceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTimeServiceById`: PrecisionTimeServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeApi.DeleteTimeServiceById`: %v\n", resp)
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

Configure Service.



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
	serviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service UUID
	precisionTimeServiceRequest := *openapiclient.NewPrecisionTimeServiceRequest(openapiclient.precisionTimeServiceRequest_type("NTP"), "Name_example", *openapiclient.NewPrecisionTimePackageRequest(openapiclient.precisionTimePackageRequest_code("NTP_STANDARD")), []openapiclient.VirtualConnectionUuid{*openapiclient.NewVirtualConnectionUuid("Uuid_example")}, *openapiclient.NewIpv4("Primary_example", "Secondary_example", "NetworkMask_example")) // PrecisionTimeServiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeApi.FulfillTimeServices(context.Background(), serviceId).PrecisionTimeServiceRequest(precisionTimeServiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeApi.FulfillTimeServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FulfillTimeServices`: PrecisionTimeServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeApi.FulfillTimeServices`: %v\n", resp)
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

Get Service By ID.



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
	serviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeApi.GetTimeServicesById(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeApi.GetTimeServicesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeServicesById`: PrecisionTimeServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeApi.GetTimeServicesById`: %v\n", resp)
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
	serviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeApi.GetTimeServicesConnectionsByServiceId(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeApi.GetTimeServicesConnectionsByServiceId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeServicesConnectionsByServiceId`: PrecisionTimeServiceConnectionsResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeApi.GetTimeServicesConnectionsByServiceId`: %v\n", resp)
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


## GetTimeServicesPackageByCode

> PrecisionTimePackageResponse GetTimeServicesPackageByCode(ctx, packageCode).Execute()

Get Package By Code



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
	packageCode := openapiclient.getTimeServicesPackageByCode_packageCode_parameter("NTP_STANDARD") // GetTimeServicesPackageByCodePackageCodeParameter | Package Code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeApi.GetTimeServicesPackageByCode(context.Background(), packageCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeApi.GetTimeServicesPackageByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeServicesPackageByCode`: PrecisionTimePackageResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeApi.GetTimeServicesPackageByCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageCode** | [**GetTimeServicesPackageByCodePackageCodeParameter**](.md) | Package Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeServicesPackageByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrecisionTimePackageResponse**](PrecisionTimePackageResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeServicesPackages

> PrecisionTimeServicePackagesResponse GetTimeServicesPackages(ctx).Execute()

Get Packages



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeApi.GetTimeServicesPackages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeApi.GetTimeServicesPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeServicesPackages`: PrecisionTimeServicePackagesResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeApi.GetTimeServicesPackages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeServicesPackagesRequest struct via the builder pattern


### Return type

[**PrecisionTimeServicePackagesResponse**](PrecisionTimeServicePackagesResponse.md)

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
	resp, r, err := apiClient.PrecisionTimeApi.SearchTimeServices(context.Background()).TimeServicesSearchRequest(timeServicesSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeApi.SearchTimeServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTimeServices`: ServiceSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeApi.SearchTimeServices`: %v\n", resp)
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

Update By ID.



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
	serviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service UUID
	precisionTimeChangeOperation := []openapiclient.PrecisionTimeChangeOperation{*openapiclient.NewPrecisionTimeChangeOperation(openapiclient.precisionTimeChangeOperation_op("replace"), openapiclient.precisionTimeChangeOperation_path("/name"), interface{}(123))} // []PrecisionTimeChangeOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeApi.UpdateTimeServicesById(context.Background(), serviceId).PrecisionTimeChangeOperation(precisionTimeChangeOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeApi.UpdateTimeServicesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTimeServicesById`: PrecisionTimeServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimeApi.UpdateTimeServicesById`: %v\n", resp)
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

