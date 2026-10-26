# \PrecisionTimePackagesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTimeServicesPackageByCode**](PrecisionTimePackagesApi.md#GetTimeServicesPackageByCode) | **Get** /fabric/v4/timeServicePackages/{packageCode} | Get Package By Code
[**GetTimeServicesPackages**](PrecisionTimePackagesApi.md#GetTimeServicesPackages) | **Get** /fabric/v4/timeServicePackages | Get Packages



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
	resp, r, err := apiClient.PrecisionTimePackagesApi.GetTimeServicesPackageByCode(context.Background(), packageCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimePackagesApi.GetTimeServicesPackageByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeServicesPackageByCode`: PrecisionTimePackageResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimePackagesApi.GetTimeServicesPackageByCode`: %v\n", resp)
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
	resp, r, err := apiClient.PrecisionTimePackagesApi.GetTimeServicesPackages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrecisionTimePackagesApi.GetTimeServicesPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeServicesPackages`: PrecisionTimeServicePackagesResponse
	fmt.Fprintf(os.Stdout, "Response from `PrecisionTimePackagesApi.GetTimeServicesPackages`: %v\n", resp)
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

