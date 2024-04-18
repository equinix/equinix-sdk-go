# \LicensingApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateLicenseUsingPOST**](LicensingApi.md#UpdateLicenseUsingPOST) | **Put** /ne/v1/devices/{uuid}/licenseTokens | Update License Token/ID/Code
[**UploadLicenseForDeviceUsingPOST**](LicensingApi.md#UploadLicenseForDeviceUsingPOST) | **Post** /ne/v1/devices/licenseFiles/{uuid} | Post License {uuid}
[**UploadLicenseUsingPOST**](LicensingApi.md#UploadLicenseUsingPOST) | **Post** /ne/v1/devices/licenseFiles | Post License File



## UpdateLicenseUsingPOST

> LicenseUploadResponse UpdateLicenseUsingPOST(ctx, uuid).Authorization(authorization).Request(request).Execute()

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
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	request := *openapiclient.NewLicenseUpdateRequest() // LicenseUpdateRequest | License token

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensingApi.UpdateLicenseUsingPOST(context.Background(), uuid).Authorization(authorization).Request(request).Execute()
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

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **request** | [**LicenseUpdateRequest**](LicenseUpdateRequest.md) | License token | 

### Return type

[**LicenseUploadResponse**](LicenseUploadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLicenseForDeviceUsingPOST

> LicenseUploadResponse UploadLicenseForDeviceUsingPOST(ctx, uuid).Authorization(authorization).File(file).Execute()

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
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	file := os.NewFile(1234, "some_file") // *os.File | License file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensingApi.UploadLicenseForDeviceUsingPOST(context.Background(), uuid).Authorization(authorization).File(file).Execute()
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

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **file** | ***os.File** | License file | 

### Return type

[**LicenseUploadResponse**](LicenseUploadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLicenseUsingPOST

> LicenseUploadResponse UploadLicenseUsingPOST(ctx).MetroCode(metroCode).DeviceTypeCode(deviceTypeCode).LicenseType(licenseType).Authorization(authorization).File(file).Execute()

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
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	file := os.NewFile(1234, "some_file") // *os.File | file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensingApi.UploadLicenseUsingPOST(context.Background()).MetroCode(metroCode).DeviceTypeCode(deviceTypeCode).LicenseType(licenseType).Authorization(authorization).File(file).Execute()
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
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **file** | ***os.File** | file | 

### Return type

[**LicenseUploadResponse**](LicenseUploadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

