# \PrecisionTimeApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTimeServices**](PrecisionTimeApi.md#CreateTimeServices) | **Post** /fabric/v4/timeServices | Create Time Service
[**DeleteTimeServiceById**](PrecisionTimeApi.md#DeleteTimeServiceById) | **Delete** /fabric/v4/timeServices/{serviceId} | Delete time service
[**GetTimeServicesById**](PrecisionTimeApi.md#GetTimeServicesById) | **Get** /fabric/v4/timeServices/{serviceId} | Get Time Service
[**UpdateTimeServicesById**](PrecisionTimeApi.md#UpdateTimeServicesById) | **Patch** /fabric/v4/timeServices/{serviceId} | Patch time service



## CreateTimeServices

> PrecisionTimeServiceCreateResponse CreateTimeServices(ctx).PrecisionTimeServiceRequest(precisionTimeServiceRequest).Execute()

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
	precisionTimeServiceRequest := *openapiclient.NewPrecisionTimeServiceRequest(openapiclient.precisionTimeServiceRequest_type("NTP"), "Name_example", *openapiclient.NewPrecisionTimePackageRequest(openapiclient.precisionTimePackageResponse_code("NTP_STANDARD")), []openapiclient.FabricConnectionUuid{*openapiclient.NewFabricConnectionUuid("Uuid_example")}, *openapiclient.NewIpv4()) // PrecisionTimeServiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeApi.CreateTimeServices(context.Background()).PrecisionTimeServiceRequest(precisionTimeServiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeApi.CreateTimeServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTimeServices`: PrecisionTimeServiceCreateResponse
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

[**PrecisionTimeServiceCreateResponse**](PrecisionTimeServiceCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTimeServiceById

> PrecisionTimeServiceCreateResponse DeleteTimeServiceById(ctx, serviceId).Execute()

Delete time service



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
	// response from `DeleteTimeServiceById`: PrecisionTimeServiceCreateResponse
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

[**PrecisionTimeServiceCreateResponse**](PrecisionTimeServiceCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeServicesById

> PrecisionTimeServiceCreateResponse GetTimeServicesById(ctx, serviceId).Execute()

Get Time Service



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
	// response from `GetTimeServicesById`: PrecisionTimeServiceCreateResponse
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

[**PrecisionTimeServiceCreateResponse**](PrecisionTimeServiceCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTimeServicesById

> PrecisionTimeServiceCreateResponse UpdateTimeServicesById(ctx, serviceId).PrecisionTimeChangeOperation(precisionTimeChangeOperation).Execute()

Patch time service



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
	precisionTimeChangeOperation := []openapiclient.PrecisionTimeChangeOperation{*openapiclient.NewPrecisionTimeChangeOperation(openapiclient.precisionTimeChangeOperation_op("replace"), openapiclient.precisionTimeChangeOperation_path("/ipv4"), map[string]interface{}(123))} // []PrecisionTimeChangeOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrecisionTimeApi.UpdateTimeServicesById(context.Background(), serviceId).PrecisionTimeChangeOperation(precisionTimeChangeOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimeApi.UpdateTimeServicesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTimeServicesById`: PrecisionTimeServiceCreateResponse
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

[**PrecisionTimeServiceCreateResponse**](PrecisionTimeServiceCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

