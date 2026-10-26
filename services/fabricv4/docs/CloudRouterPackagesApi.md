# \CloudRouterPackagesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudRouterPackageByCode**](CloudRouterPackagesApi.md#GetCloudRouterPackageByCode) | **Get** /fabric/v4/routerPackages/{routerPackageCode} | Get Package Details
[**GetCloudRouterPackages**](CloudRouterPackagesApi.md#GetCloudRouterPackages) | **Get** /fabric/v4/routerPackages | List Packages



## GetCloudRouterPackageByCode

> CloudRouterPackage GetCloudRouterPackageByCode(ctx, routerPackageCode).Execute()

Get Package Details



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
	routerPackageCode := openapiclient.RouterPackageCode("LAB") // RouterPackageCode | Equinix-assigned Cloud Router package identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRouterPackagesApi.GetCloudRouterPackageByCode(context.Background(), routerPackageCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRouterPackagesApi.GetCloudRouterPackageByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudRouterPackageByCode`: CloudRouterPackage
	fmt.Fprintf(os.Stdout, "Response from `CloudRouterPackagesApi.GetCloudRouterPackageByCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routerPackageCode** | [**RouterPackageCode**](.md) | Equinix-assigned Cloud Router package identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudRouterPackageByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudRouterPackage**](CloudRouterPackage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudRouterPackages

> PackageResponse GetCloudRouterPackages(ctx).Offset(offset).Limit(limit).Execute()

List Packages



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
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudRouterPackagesApi.GetCloudRouterPackages(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudRouterPackagesApi.GetCloudRouterPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudRouterPackages`: PackageResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudRouterPackagesApi.GetCloudRouterPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudRouterPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**PackageResponse**](PackageResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

