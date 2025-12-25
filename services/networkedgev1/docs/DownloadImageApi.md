# \DownloadImageApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDownloadableImagesUsingGET**](DownloadImageApi.md#GetDownloadableImagesUsingGET) | **Get** /ne/v1/devices/{deviceType}/repositories | Get Downloadable Images
[**PostDownloadImage**](DownloadImageApi.md#PostDownloadImage) | **Post** /ne/v1/devices/{deviceType}/repositories/{version}/download | Download an Image



## GetDownloadableImagesUsingGET

> []ListOfDownloadableImages GetDownloadableImagesUsingGET(ctx, deviceType).Execute()

Get Downloadable Images



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	deviceType := "deviceType_example" // string | Device type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DownloadImageApi.GetDownloadableImagesUsingGET(context.Background(), deviceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadImageApi.GetDownloadableImagesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDownloadableImagesUsingGET`: []ListOfDownloadableImages
	fmt.Fprintf(os.Stdout, "Response from `DownloadImageApi.GetDownloadableImagesUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceType** | **string** | Device type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadableImagesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListOfDownloadableImages**](ListOfDownloadableImages.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDownloadImage

> PostDownloadImageResponse PostDownloadImage(ctx, deviceType, version).Execute()

Download an Image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	deviceType := "deviceType_example" // string | Device type
	version := "version_example" // string | Version

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DownloadImageApi.PostDownloadImage(context.Background(), deviceType, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadImageApi.PostDownloadImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDownloadImage`: PostDownloadImageResponse
	fmt.Fprintf(os.Stdout, "Response from `DownloadImageApi.PostDownloadImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceType** | **string** | Device type | 
**version** | **string** | Version | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDownloadImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostDownloadImageResponse**](PostDownloadImageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

