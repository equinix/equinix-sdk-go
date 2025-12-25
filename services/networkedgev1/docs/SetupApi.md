# \SetupApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountsWithStatusUsingGET**](SetupApi.md#GetAccountsWithStatusUsingGET) | **Get** /ne/v1/accounts/{metro} | Get Accounts {metro}
[**GetAgreementStatusUsingGET**](SetupApi.md#GetAgreementStatusUsingGET) | **Get** /ne/v1/agreements/accounts | Get Agreement Status.
[**GetAllowedInterfacesUsingGET**](SetupApi.md#GetAllowedInterfacesUsingGET) | **Get** /ne/v1/deviceTypes/{deviceType}/interfaces | Get Allowed Interfaces
[**GetMetrosUsingGET**](SetupApi.md#GetMetrosUsingGET) | **Get** /ne/v1/metros | Get Available Metros
[**GetNotificationsUsingGET**](SetupApi.md#GetNotificationsUsingGET) | **Get** /ne/v1/notifications | Get Downtime Notifications
[**GetOrderSummaryUsingGET**](SetupApi.md#GetOrderSummaryUsingGET) | **Get** /ne/v1/orderSummaries | Get Order Summary
[**GetOrderTermsUsingGET**](SetupApi.md#GetOrderTermsUsingGET) | **Get** /ne/v1/agreements/orders | Get Order Terms
[**GetPublicKeysUsingGET**](SetupApi.md#GetPublicKeysUsingGET) | **Get** /ne/v1/publicKeys | Get Public Keys
[**GetVendorTermsUsingGET**](SetupApi.md#GetVendorTermsUsingGET) | **Get** /ne/v1/agreements/vendors | Get Vendor Terms
[**GetVirtualDevicesUsingGET**](SetupApi.md#GetVirtualDevicesUsingGET) | **Get** /ne/v1/deviceTypes | Get Device Types
[**PostPublicKeyUsingPOST**](SetupApi.md#PostPublicKeyUsingPOST) | **Post** /ne/v1/publicKeys | Create Public Key
[**RetrievePriceUsingGET**](SetupApi.md#RetrievePriceUsingGET) | **Get** /ne/v1/prices | Get the Price
[**SendAgreementUsingPOST1**](SetupApi.md#SendAgreementUsingPOST1) | **Post** /ne/v1/agreements/accounts | Create an agreement
[**UploadFileUsingPOST**](SetupApi.md#UploadFileUsingPOST) | **Post** /ne/v1/files | Upload File (Post)



## GetAccountsWithStatusUsingGET

> PageResponseDtoMetroAccountResponse GetAccountsWithStatusUsingGET(ctx, metro).AccountUcmId(accountUcmId).Execute()

Get Accounts {metro}



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
	metro := "metro_example" // string | Metro region for which you want to check your account status
	accountUcmId := "accountUcmId_example" // string | Unique ID of an account (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetupApi.GetAccountsWithStatusUsingGET(context.Background(), metro).AccountUcmId(accountUcmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetupApi.GetAccountsWithStatusUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountsWithStatusUsingGET`: PageResponseDtoMetroAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `SetupApi.GetAccountsWithStatusUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metro** | **string** | Metro region for which you want to check your account status | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsWithStatusUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountUcmId** | **string** | Unique ID of an account | 

### Return type

[**PageResponseDtoMetroAccountResponse**](PageResponseDtoMetroAccountResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgreementStatusUsingGET

> AgreementStatusResponse GetAgreementStatusUsingGET(ctx).AccountNumber(accountNumber).Execute()

Get Agreement Status.



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
	accountNumber := "accountNumber_example" // string | account_number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetupApi.GetAgreementStatusUsingGET(context.Background()).AccountNumber(accountNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetupApi.GetAgreementStatusUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgreementStatusUsingGET`: AgreementStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `SetupApi.GetAgreementStatusUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgreementStatusUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountNumber** | **string** | account_number | 

### Return type

[**AgreementStatusResponse**](AgreementStatusResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllowedInterfacesUsingGET

> AllowedInterfaceResponse GetAllowedInterfacesUsingGET(ctx, deviceType).DeviceManagementType(deviceManagementType).Core(core).Mode(mode).Cluster(cluster).Sdwan(sdwan).Connectivity(connectivity).Memory(memory).Unit(unit).Flavor(flavor).Version(version).SoftwarePkg(softwarePkg).Execute()

Get Allowed Interfaces



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
	deviceType := "deviceType_example" // string | Device type code. PA-VM.
	deviceManagementType := "deviceManagementType_example" // string | Device management type. SELF-CONFIGURED
	core := int32(56) // int32 | The desired number of cores.
	mode := "mode_example" // string | License mode, either Subscription or BYOL. (optional)
	cluster := true // bool | Whether you want a cluster device. (optional)
	sdwan := true // bool | Whether you want an SD-WAN device. (optional)
	connectivity := "connectivity_example" // string | Type of connectivity you want. INTERNET-ACCESS, PRIVATE, or INTERNET-ACCESS-WITH-PRVT-MGMT. PRIVATE devices do not have ACLs or bandwidth. (optional)
	memory := int32(56) // int32 | Desired memory. (optional)
	unit := "unit_example" // string | Unit of memory. GB or MB. (optional)
	flavor := "flavor_example" // string | Flavor of device. (optional)
	version := "version_example" // string | Version. (optional)
	softwarePkg := "softwarePkg_example" // string | Software package. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetupApi.GetAllowedInterfacesUsingGET(context.Background(), deviceType).DeviceManagementType(deviceManagementType).Core(core).Mode(mode).Cluster(cluster).Sdwan(sdwan).Connectivity(connectivity).Memory(memory).Unit(unit).Flavor(flavor).Version(version).SoftwarePkg(softwarePkg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetupApi.GetAllowedInterfacesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllowedInterfacesUsingGET`: AllowedInterfaceResponse
	fmt.Fprintf(os.Stdout, "Response from `SetupApi.GetAllowedInterfacesUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceType** | **string** | Device type code. PA-VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllowedInterfacesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceManagementType** | **string** | Device management type. SELF-CONFIGURED | 
 **core** | **int32** | The desired number of cores. | 
 **mode** | **string** | License mode, either Subscription or BYOL. | 
 **cluster** | **bool** | Whether you want a cluster device. | 
 **sdwan** | **bool** | Whether you want an SD-WAN device. | 
 **connectivity** | **string** | Type of connectivity you want. INTERNET-ACCESS, PRIVATE, or INTERNET-ACCESS-WITH-PRVT-MGMT. PRIVATE devices do not have ACLs or bandwidth. | 
 **memory** | **int32** | Desired memory. | 
 **unit** | **string** | Unit of memory. GB or MB. | 
 **flavor** | **string** | Flavor of device. | 
 **version** | **string** | Version. | 
 **softwarePkg** | **string** | Software package. | 

### Return type

[**AllowedInterfaceResponse**](AllowedInterfaceResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetrosUsingGET

> PageResponseDtoMetroResponse GetMetrosUsingGET(ctx).Region(region).Offset(offset).Limit(limit).Execute()

Get Available Metros



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
	region := "region_example" // string | Name of the region for which you want metros (e.g., AMER) (optional)
	offset := "offset_example" // string | Specifies where to start a page. It is the starting point of the collection returned from the server. (optional) (default to "0")
	limit := "limit_example" // string | Specifies the page size. (optional) (default to "20")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetupApi.GetMetrosUsingGET(context.Background()).Region(region).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetupApi.GetMetrosUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetrosUsingGET`: PageResponseDtoMetroResponse
	fmt.Fprintf(os.Stdout, "Response from `SetupApi.GetMetrosUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetrosUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** | Name of the region for which you want metros (e.g., AMER) | 
 **offset** | **string** | Specifies where to start a page. It is the starting point of the collection returned from the server. | [default to &quot;0&quot;]
 **limit** | **string** | Specifies the page size. | [default to &quot;20&quot;]

### Return type

[**PageResponseDtoMetroResponse**](PageResponseDtoMetroResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsUsingGET

> DowntimeNotification GetNotificationsUsingGET(ctx).Execute()

Get Downtime Notifications



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetupApi.GetNotificationsUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetupApi.GetNotificationsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationsUsingGET`: DowntimeNotification
	fmt.Fprintf(os.Stdout, "Response from `SetupApi.GetNotificationsUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsUsingGETRequest struct via the builder pattern


### Return type

[**DowntimeNotification**](DowntimeNotification.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderSummaryUsingGET

> GetOrderSummaryUsingGET(ctx).AccountNumber(accountNumber).Metro(metro).VendorPackage(vendorPackage).LicenseType(licenseType).SoftwarePackage(softwarePackage).Throughput(throughput).ThroughputUnit(throughputUnit).TermLength(termLength).AdditionalBandwidth(additionalBandwidth).VirtualDeviceUuid(virtualDeviceUuid).DeviceManagementType(deviceManagementType).Core(core).SecondaryAccountNumber(secondaryAccountNumber).SecondaryMetro(secondaryMetro).SecondaryAdditionalBandwidth(secondaryAdditionalBandwidth).AccountUcmId(accountUcmId).OrderingContact(orderingContact).Execute()

Get Order Summary



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
	accountNumber := int32(56) // int32 | Account number (optional)
	metro := "metro_example" // string | Metro (optional)
	vendorPackage := "vendorPackage_example" // string | Vendor package (optional)
	licenseType := "licenseType_example" // string | License type (optional)
	softwarePackage := "softwarePackage_example" // string | Software package (optional)
	throughput := int32(56) // int32 | Throughput (optional)
	throughputUnit := "throughputUnit_example" // string | Throughput unit (optional)
	termLength := "termLength_example" // string | Term length (in months) (optional)
	additionalBandwidth := int32(56) // int32 | Additional bandwidth (in Mbps) (optional)
	virtualDeviceUuid := "virtualDeviceUuid_example" // string | Virtual device unique Id (only required if existing device is being modified) (optional)
	deviceManagementType := "deviceManagementType_example" // string | The device management type (optional)
	core := int32(56) // int32 | The number of cores (optional)
	secondaryAccountNumber := int32(56) // int32 | Secondary account number (in case you have a device pair) (optional)
	secondaryMetro := "secondaryMetro_example" // string | Secondary metro (in case you have a device pair) (optional)
	secondaryAdditionalBandwidth := int32(56) // int32 | Secondary additional bandwidth (in Mbps) (optional)
	accountUcmId := "accountUcmId_example" // string | Account unique ID (optional)
	orderingContact := "orderingContact_example" // string | Reseller customer username (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SetupApi.GetOrderSummaryUsingGET(context.Background()).AccountNumber(accountNumber).Metro(metro).VendorPackage(vendorPackage).LicenseType(licenseType).SoftwarePackage(softwarePackage).Throughput(throughput).ThroughputUnit(throughputUnit).TermLength(termLength).AdditionalBandwidth(additionalBandwidth).VirtualDeviceUuid(virtualDeviceUuid).DeviceManagementType(deviceManagementType).Core(core).SecondaryAccountNumber(secondaryAccountNumber).SecondaryMetro(secondaryMetro).SecondaryAdditionalBandwidth(secondaryAdditionalBandwidth).AccountUcmId(accountUcmId).OrderingContact(orderingContact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetupApi.GetOrderSummaryUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderSummaryUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountNumber** | **int32** | Account number | 
 **metro** | **string** | Metro | 
 **vendorPackage** | **string** | Vendor package | 
 **licenseType** | **string** | License type | 
 **softwarePackage** | **string** | Software package | 
 **throughput** | **int32** | Throughput | 
 **throughputUnit** | **string** | Throughput unit | 
 **termLength** | **string** | Term length (in months) | 
 **additionalBandwidth** | **int32** | Additional bandwidth (in Mbps) | 
 **virtualDeviceUuid** | **string** | Virtual device unique Id (only required if existing device is being modified) | 
 **deviceManagementType** | **string** | The device management type | 
 **core** | **int32** | The number of cores | 
 **secondaryAccountNumber** | **int32** | Secondary account number (in case you have a device pair) | 
 **secondaryMetro** | **string** | Secondary metro (in case you have a device pair) | 
 **secondaryAdditionalBandwidth** | **int32** | Secondary additional bandwidth (in Mbps) | 
 **accountUcmId** | **string** | Account unique ID | 
 **orderingContact** | **string** | Reseller customer username | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderTermsUsingGET

> OrderTermsResponse GetOrderTermsUsingGET(ctx).Execute()

Get Order Terms



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetupApi.GetOrderTermsUsingGET(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetupApi.GetOrderTermsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderTermsUsingGET`: OrderTermsResponse
	fmt.Fprintf(os.Stdout, "Response from `SetupApi.GetOrderTermsUsingGET`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderTermsUsingGETRequest struct via the builder pattern


### Return type

[**OrderTermsResponse**](OrderTermsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicKeysUsingGET

> []PageResponsePublicKeys GetPublicKeysUsingGET(ctx).AccountUcmId(accountUcmId).Execute()

Get Public Keys



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
	accountUcmId := "accountUcmId_example" // string | This field is for resellers. Please pass the accountUcmId of your customer to get the public keys. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetupApi.GetPublicKeysUsingGET(context.Background()).AccountUcmId(accountUcmId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetupApi.GetPublicKeysUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicKeysUsingGET`: []PageResponsePublicKeys
	fmt.Fprintf(os.Stdout, "Response from `SetupApi.GetPublicKeysUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicKeysUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountUcmId** | **string** | This field is for resellers. Please pass the accountUcmId of your customer to get the public keys. | 

### Return type

[**[]PageResponsePublicKeys**](PageResponsePublicKeys.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVendorTermsUsingGET

> VendorTermsResponse GetVendorTermsUsingGET(ctx).VendorPackage(vendorPackage).LicenseType(licenseType).Execute()

Get Vendor Terms



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
	vendorPackage := "vendorPackage_example" // string | vendorPackage
	licenseType := "licenseType_example" // string | licenseType

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetupApi.GetVendorTermsUsingGET(context.Background()).VendorPackage(vendorPackage).LicenseType(licenseType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetupApi.GetVendorTermsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVendorTermsUsingGET`: VendorTermsResponse
	fmt.Fprintf(os.Stdout, "Response from `SetupApi.GetVendorTermsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVendorTermsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vendorPackage** | **string** | vendorPackage | 
 **licenseType** | **string** | licenseType | 

### Return type

[**VendorTermsResponse**](VendorTermsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualDevicesUsingGET

> PageResponseDtoVirtualDeviceType GetVirtualDevicesUsingGET(ctx).DeviceTypeCode(deviceTypeCode).Category(category).Offset(offset).Limit(limit).Execute()

Get Device Types



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
	deviceTypeCode := "deviceTypeCode_example" // string | Device type code (e.g., C8000V) (optional)
	category := "category_example" // string | Category. One of FIREWALL, ROUTER or SD-WAN (optional)
	offset := "offset_example" // string | Specifies where to start a page. It is the starting point of the collection returned from the server. (optional) (default to "0")
	limit := "limit_example" // string | Specifies the page size. (optional) (default to "20")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetupApi.GetVirtualDevicesUsingGET(context.Background()).DeviceTypeCode(deviceTypeCode).Category(category).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetupApi.GetVirtualDevicesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualDevicesUsingGET`: PageResponseDtoVirtualDeviceType
	fmt.Fprintf(os.Stdout, "Response from `SetupApi.GetVirtualDevicesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualDevicesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceTypeCode** | **string** | Device type code (e.g., C8000V) | 
 **category** | **string** | Category. One of FIREWALL, ROUTER or SD-WAN | 
 **offset** | **string** | Specifies where to start a page. It is the starting point of the collection returned from the server. | [default to &quot;0&quot;]
 **limit** | **string** | Specifies the page size. | [default to &quot;20&quot;]

### Return type

[**PageResponseDtoVirtualDeviceType**](PageResponseDtoVirtualDeviceType.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPublicKeyUsingPOST

> PostPublicKeyUsingPOST(ctx).PublicKeyRequest(publicKeyRequest).Execute()

Create Public Key



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
	publicKeyRequest := *openapiclient.NewPublicKeyRequest("myKeyName", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQC5kHcagDZ7utPan4DHWUvoJxwz/DISRFwZdpMhslhZRI+6dGOC8mJn42SlSUAUtkt8Qyl4HipPK7Xh6oGj70Iba1a9pDcURYTYcqWFBEhcdDsMnH1CICmvVdsILehFtiS3X0J1JhwmWQI/7ll3QOk8fLgWCz3idlYJqtMs8Gz/6Q== noname") // PublicKeyRequest | keyName, keyValue, and keyType

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SetupApi.PostPublicKeyUsingPOST(context.Background()).PublicKeyRequest(publicKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetupApi.PostPublicKeyUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPublicKeyUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicKeyRequest** | [**PublicKeyRequest**](PublicKeyRequest.md) | keyName, keyValue, and keyType | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePriceUsingGET

> CompositePriceResponse RetrievePriceUsingGET(ctx).AccountNumber(accountNumber).Metro(metro).VendorPackage(vendorPackage).LicenseType(licenseType).SoftwarePackage(softwarePackage).Throughput(throughput).ThroughputUnit(throughputUnit).TermLength(termLength).AdditionalBandwidth(additionalBandwidth).VirtualDeviceUuid(virtualDeviceUuid).DeviceManagementType(deviceManagementType).Core(core).SecondaryAccountNumber(secondaryAccountNumber).SecondaryMetro(secondaryMetro).SecondaryAdditionalBandwidth(secondaryAdditionalBandwidth).AccountUcmId(accountUcmId).OrderingContact(orderingContact).Execute()

Get the Price



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
	accountNumber := int32(56) // int32 | Account number (optional)
	metro := "metro_example" // string | Metro (optional)
	vendorPackage := "vendorPackage_example" // string | Vendor package (optional)
	licenseType := "licenseType_example" // string | License type (optional)
	softwarePackage := "softwarePackage_example" // string | Software package (optional)
	throughput := int32(56) // int32 | Throughput (optional)
	throughputUnit := "throughputUnit_example" // string | Throughput unit (optional)
	termLength := "termLength_example" // string | Term length (in months) (optional)
	additionalBandwidth := int32(56) // int32 | Additional bandwidth (in Mbps) (optional)
	virtualDeviceUuid := "virtualDeviceUuid_example" // string | Virtual device unique Id (only required if existing device is being modified) (optional)
	deviceManagementType := "deviceManagementType_example" // string | The device management type (optional)
	core := int32(56) // int32 | The number of cores (optional)
	secondaryAccountNumber := int32(56) // int32 | The secondary account number (for HA) (optional)
	secondaryMetro := "secondaryMetro_example" // string | Secondary metro (for HA) (optional)
	secondaryAdditionalBandwidth := int32(56) // int32 | Secondary additional bandwidth (in Mbps for HA) (optional)
	accountUcmId := "accountUcmId_example" // string | Account unique ID (optional)
	orderingContact := "orderingContact_example" // string | Reseller customer username (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetupApi.RetrievePriceUsingGET(context.Background()).AccountNumber(accountNumber).Metro(metro).VendorPackage(vendorPackage).LicenseType(licenseType).SoftwarePackage(softwarePackage).Throughput(throughput).ThroughputUnit(throughputUnit).TermLength(termLength).AdditionalBandwidth(additionalBandwidth).VirtualDeviceUuid(virtualDeviceUuid).DeviceManagementType(deviceManagementType).Core(core).SecondaryAccountNumber(secondaryAccountNumber).SecondaryMetro(secondaryMetro).SecondaryAdditionalBandwidth(secondaryAdditionalBandwidth).AccountUcmId(accountUcmId).OrderingContact(orderingContact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetupApi.RetrievePriceUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrievePriceUsingGET`: CompositePriceResponse
	fmt.Fprintf(os.Stdout, "Response from `SetupApi.RetrievePriceUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePriceUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountNumber** | **int32** | Account number | 
 **metro** | **string** | Metro | 
 **vendorPackage** | **string** | Vendor package | 
 **licenseType** | **string** | License type | 
 **softwarePackage** | **string** | Software package | 
 **throughput** | **int32** | Throughput | 
 **throughputUnit** | **string** | Throughput unit | 
 **termLength** | **string** | Term length (in months) | 
 **additionalBandwidth** | **int32** | Additional bandwidth (in Mbps) | 
 **virtualDeviceUuid** | **string** | Virtual device unique Id (only required if existing device is being modified) | 
 **deviceManagementType** | **string** | The device management type | 
 **core** | **int32** | The number of cores | 
 **secondaryAccountNumber** | **int32** | The secondary account number (for HA) | 
 **secondaryMetro** | **string** | Secondary metro (for HA) | 
 **secondaryAdditionalBandwidth** | **int32** | Secondary additional bandwidth (in Mbps for HA) | 
 **accountUcmId** | **string** | Account unique ID | 
 **orderingContact** | **string** | Reseller customer username | 

### Return type

[**CompositePriceResponse**](CompositePriceResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendAgreementUsingPOST1

> AgreementAcceptResponse SendAgreementUsingPOST1(ctx).AgreementAcceptRequest(agreementAcceptRequest).Execute()

Create an agreement



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
	agreementAcceptRequest := *openapiclient.NewAgreementAcceptRequest() // AgreementAcceptRequest | agreementAcceptRequest

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetupApi.SendAgreementUsingPOST1(context.Background()).AgreementAcceptRequest(agreementAcceptRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetupApi.SendAgreementUsingPOST1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendAgreementUsingPOST1`: AgreementAcceptResponse
	fmt.Fprintf(os.Stdout, "Response from `SetupApi.SendAgreementUsingPOST1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendAgreementUsingPOST1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agreementAcceptRequest** | [**AgreementAcceptRequest**](AgreementAcceptRequest.md) | agreementAcceptRequest | 

### Return type

[**AgreementAcceptResponse**](AgreementAcceptResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileUsingPOST

> FileUploadResponse UploadFileUsingPOST(ctx).File(file).MetroCode(metroCode).DeviceTypeCode(deviceTypeCode).ProcessType(processType).DeviceManagementType(deviceManagementType).LicenseType(licenseType).Execute()

Upload File (Post)



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
	file := os.NewFile(1234, "some_file") // *os.File | A license or a cloud_init file. For example, for Aviatrix, this is your bootsrap config file generated from the Aviatrix Controller portal.
	metroCode := "metroCode_example" // string | Two-letter metro code.
	deviceTypeCode := "deviceTypeCode_example" // string | Device type code, e.g., AVIATRIX_EDGE
	processType := "processType_example" // string | Whether you are uploading a license or a cloud_init file. LICENSE or CLOUD_INIT
	deviceManagementType := "deviceManagementType_example" // string | Device management type, whether SELF-CONFIGURED or not (optional)
	licenseType := "licenseType_example" // string | Type of license (BYOL-Bring Your Own License) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SetupApi.UploadFileUsingPOST(context.Background()).File(file).MetroCode(metroCode).DeviceTypeCode(deviceTypeCode).ProcessType(processType).DeviceManagementType(deviceManagementType).LicenseType(licenseType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SetupApi.UploadFileUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFileUsingPOST`: FileUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `SetupApi.UploadFileUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | A license or a cloud_init file. For example, for Aviatrix, this is your bootsrap config file generated from the Aviatrix Controller portal. | 
 **metroCode** | **string** | Two-letter metro code. | 
 **deviceTypeCode** | **string** | Device type code, e.g., AVIATRIX_EDGE | 
 **processType** | **string** | Whether you are uploading a license or a cloud_init file. LICENSE or CLOUD_INIT | 
 **deviceManagementType** | **string** | Device management type, whether SELF-CONFIGURED or not | 
 **licenseType** | **string** | Type of license (BYOL-Bring Your Own License) | 

### Return type

[**FileUploadResponse**](FileUploadResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

