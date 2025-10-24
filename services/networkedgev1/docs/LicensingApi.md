# \LicensingApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateLicenseUsingPOST**](LicensingApi.md#UpdateLicenseUsingPOST) | **Put** /ne/v1/devices/{uuid}/licenseTokens | Update License Token/ID/Code
[**UploadLicenseForDeviceUsingPOST**](LicensingApi.md#UploadLicenseForDeviceUsingPOST) | **Post** /ne/v1/devices/licenseFiles/{uuid} | Post License {uuid}
[**UploadLicenseUsingPOST**](LicensingApi.md#UploadLicenseUsingPOST) | **Post** /ne/v1/devices/licenseFiles | Post License File



## UpdateLicenseUsingPOST

> LicenseUploadResponse UpdateLicenseUsingPOST(ctx, uuid).Request(request).Execute()

Update License Token/ID/Code



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
	uuid := "uuid_example" // string | Unique Id of a virtual device
	request := *openapiclient.NewLicenseUpdateRequest() // LicenseUpdateRequest | License token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensingApi.UpdateLicenseUsingPOST(context.Background(), uuid).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensingApi.UpdateLicenseUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLicenseUsingPOST`: LicenseUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `LicensingApi.UpdateLicenseUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique Id of a virtual device | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLicenseUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**LicenseUpdateRequest**](LicenseUpdateRequest.md) | License token | 

### Return type

[**LicenseUploadResponse**](LicenseUploadResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLicenseForDeviceUsingPOST

> LicenseUploadResponse UploadLicenseForDeviceUsingPOST(ctx, uuid).File(file).Execute()

Post License {uuid}



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
	uuid := "uuid_example" // string | Unique Id of a virtual device
	file := os.NewFile(1234, "some_file") // *os.File | License file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensingApi.UploadLicenseForDeviceUsingPOST(context.Background(), uuid).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensingApi.UploadLicenseForDeviceUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadLicenseForDeviceUsingPOST`: LicenseUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `LicensingApi.UploadLicenseForDeviceUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique Id of a virtual device | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadLicenseForDeviceUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | License file | 

### Return type

[**LicenseUploadResponse**](LicenseUploadResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLicenseUsingPOST

> LicenseUploadResponse UploadLicenseUsingPOST(ctx).MetroCode(metroCode).DeviceTypeCode(deviceTypeCode).LicenseType(licenseType).File(file).Execute()

Post License File



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
	metroCode := "metroCode_example" // string | metroCode
	deviceTypeCode := "deviceTypeCode_example" // string | deviceTypeCode
	licenseType := "licenseType_example" // string | licenseType
	file := os.NewFile(1234, "some_file") // *os.File | file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensingApi.UploadLicenseUsingPOST(context.Background()).MetroCode(metroCode).DeviceTypeCode(deviceTypeCode).LicenseType(licenseType).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensingApi.UploadLicenseUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadLicenseUsingPOST`: LicenseUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `LicensingApi.UploadLicenseUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadLicenseUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metroCode** | **string** | metroCode | 
 **deviceTypeCode** | **string** | deviceTypeCode | 
 **licenseType** | **string** | licenseType | 
 **file** | ***os.File** | file | 

### Return type

[**LicenseUploadResponse**](LicenseUploadResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

