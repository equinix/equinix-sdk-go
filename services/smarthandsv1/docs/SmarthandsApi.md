# \SmarthandsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocation**](SmarthandsApi.md#GetLocation) | **Get** /v1/orders/smarthands/locations | Get locations(ibx, cage) details for placing smart hands based on smart hands permission for user
[**HandleSmartHandCableRequestOrder**](SmarthandsApi.md#HandleSmartHandCableRequestOrder) | **Post** /v1/orders/smarthands/cableRequest | This API is used to request cables per your specifications for your equipment.
[**HandleSmartHandCageCleanupOrder**](SmarthandsApi.md#HandleSmartHandCageCleanupOrder) | **Post** /v1/orders/smarthands/cageCleanup | Request trash removal or cage cleanup
[**HandleSmartHandCageEscortOrder**](SmarthandsApi.md#HandleSmartHandCageEscortOrder) | **Post** /v1/orders/smarthands/cageEscort | Request a patch cable to be moved between devices
[**HandleSmartHandLocatePackageOrder**](SmarthandsApi.md#HandleSmartHandLocatePackageOrder) | **Post** /v1/orders/smarthands/locatePackage | Get the location of your package(s) at an IBX.
[**HandleSmartHandMoveJumperCableOrder**](SmarthandsApi.md#HandleSmartHandMoveJumperCableOrder) | **Post** /v1/orders/smarthands/moveJumperCable | This API is used to request a patch cable to be moved between devices.
[**HandleSmartHandOrder**](SmarthandsApi.md#HandleSmartHandOrder) | **Post** /v1/orders/smarthands/equipmentInstall | Request equipment installation per your specifications by an IBX Technician
[**HandleSmartHandOthersOrder**](SmarthandsApi.md#HandleSmartHandOthersOrder) | **Post** /v1/orders/smarthands/other | Request a Smart Hands order not listed in catalogue
[**HandleSmartHandPatchCableInstallOrder**](SmarthandsApi.md#HandleSmartHandPatchCableInstallOrder) | **Post** /v1/orders/smarthands/patchCableInstall | Request installation of a cross connect patch cable by an IBX Technician.
[**HandleSmartHandPatchCableRemovalOrder**](SmarthandsApi.md#HandleSmartHandPatchCableRemovalOrder) | **Post** /v1/orders/smarthands/patchCableRemoval | Request removal of a cross connect cable by an IBX Technician.
[**HandleSmartHandPicturesDocumentOrder**](SmarthandsApi.md#HandleSmartHandPicturesDocumentOrder) | **Post** /v1/orders/smarthands/picturesDocument | Request an IBX Technician to provide cage-related pictures or documentation.
[**HandleSmartHandRunJumperCableOrder**](SmarthandsApi.md#HandleSmartHandRunJumperCableOrder) | **Post** /v1/orders/smarthands/runJumperCable | This API is used to request a jumper cable to be ran between devices.
[**HandleSmartHandShipmentUnpackOrder**](SmarthandsApi.md#HandleSmartHandShipmentUnpackOrder) | **Post** /v1/orders/smarthands/shipmentUnpack | Request inbound shipment unpacking and packaging disposal.
[**SmartHandTypes**](SmarthandsApi.md#SmartHandTypes) | **Get** /v1/orders/smarthands/types | All supported smart hands types API.



## GetLocation

> GetLocation200Response GetLocation(ctx).Authorization(authorization).Detail(detail).Ibxs(ibxs).Cages(cages).Execute()

Get locations(ibx, cage) details for placing smart hands based on smart hands permission for user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smarthandsv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.
	detail := true // bool | When enable this flag returns detailed permission with Cage & Cabinets. (optional) (default to false)
	ibxs := "ibxs_example" // string | Example: AM1,AM2 (optional)
	cages := "cages_example" // string | Example: AM1:02:002MC1 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmarthandsApi.GetLocation(context.Background()).Authorization(authorization).Detail(detail).Ibxs(ibxs).Cages(cages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmarthandsApi.GetLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocation`: GetLocation200Response
	fmt.Fprintf(os.Stdout, "Response from `SmarthandsApi.GetLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 
 **detail** | **bool** | When enable this flag returns detailed permission with Cage &amp; Cabinets. | [default to false]
 **ibxs** | **string** | Example: AM1,AM2 | 
 **cages** | **string** | Example: AM1:02:002MC1 | 

### Return type

[**GetLocation200Response**](GetLocation200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleSmartHandCableRequestOrder

> HandleSmartHandCableRequestOrder(ctx).Authorization(authorization).Body(body).Execute()

This API is used to request cables per your specifications for your equipment.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smarthandsv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.
	body := *openapiclient.NewCableRequestRequest(*openapiclient.NewIbxLocation("AM1", []openapiclient.IbxLocationCagesInner{*openapiclient.NewIbxLocationCagesInner("AM1:01:001MC3", "12345")}), []openapiclient.ContactInfo{*openapiclient.NewContactInfo(openapiclient.contactInfo_contactType("TECHNICAL"))}, *openapiclient.NewScheduleInfo(openapiclient.scheduleInfo_scheduleType("STANDARD")), *openapiclient.NewCableRequestRequestServiceDetails(openapiclient.cableRequestRequest_serviceDetails_quantity("1"), "Scope of work")) // CableRequestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmarthandsApi.HandleSmartHandCableRequestOrder(context.Background()).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmarthandsApi.HandleSmartHandCableRequestOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleSmartHandCableRequestOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 
 **body** | [**CableRequestRequest**](CableRequestRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleSmartHandCageCleanupOrder

> HandleSmartHandCageCleanupOrder(ctx).Authorization(authorization).Body(body).Execute()

Request trash removal or cage cleanup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smarthandsv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.
	body := *openapiclient.NewCageCleanupRequest(*openapiclient.NewIbxLocation("AM1", []openapiclient.IbxLocationCagesInner{*openapiclient.NewIbxLocationCagesInner("AM1:01:001MC3", "12345")}), []openapiclient.ContactInfo{*openapiclient.NewContactInfo(openapiclient.contactInfo_contactType("TECHNICAL"))}, *openapiclient.NewCageCleanupRequestServiceDetails(false, true, "Scope of work")) // CageCleanupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmarthandsApi.HandleSmartHandCageCleanupOrder(context.Background()).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmarthandsApi.HandleSmartHandCageCleanupOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleSmartHandCageCleanupOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 
 **body** | [**CageCleanupRequest**](CageCleanupRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleSmartHandCageEscortOrder

> HandleSmartHandCageEscortOrder(ctx).Authorization(authorization).Body(body).Execute()

Request a patch cable to be moved between devices



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smarthandsv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.
	body := *openapiclient.NewCageEscortRequest(*openapiclient.NewIbxLocation("AM1", []openapiclient.IbxLocationCagesInner{*openapiclient.NewIbxLocationCagesInner("AM1:01:001MC3", "12345")}), []openapiclient.ContactInfo{*openapiclient.NewContactInfo(openapiclient.contactInfo_contactType("TECHNICAL"))}, *openapiclient.NewScheduleInfo(openapiclient.scheduleInfo_scheduleType("STANDARD")), *openapiclient.NewCageEscortRequestServiceDetails(openapiclient.cageEscortRequest_serviceDetails_durationVisit("30 Minutes"), false, "Scope of work", false, "1-108050984499")) // CageEscortRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmarthandsApi.HandleSmartHandCageEscortOrder(context.Background()).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmarthandsApi.HandleSmartHandCageEscortOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleSmartHandCageEscortOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 
 **body** | [**CageEscortRequest**](CageEscortRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleSmartHandLocatePackageOrder

> HandleSmartHandLocatePackageOrder(ctx).Authorization(authorization).Body(body).Execute()

Get the location of your package(s) at an IBX.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smarthandsv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.
	body := *openapiclient.NewLocatePackageRequest(*openapiclient.NewIbxLocation("AM1", []openapiclient.IbxLocationCagesInner{*openapiclient.NewIbxLocationCagesInner("AM1:01:001MC3", "12345")}), []openapiclient.ContactInfo{*openapiclient.NewContactInfo(openapiclient.contactInfo_contactType("TECHNICAL"))}, *openapiclient.NewScheduleInfo(openapiclient.scheduleInfo_scheduleType("STANDARD")), *openapiclient.NewLocatePackageRequestServiceDetails("1-1111111111", "323-12121", "Near cabinet 1010", "It's small box with 1kg weight.", "Scope of work")) // LocatePackageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmarthandsApi.HandleSmartHandLocatePackageOrder(context.Background()).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmarthandsApi.HandleSmartHandLocatePackageOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleSmartHandLocatePackageOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 
 **body** | [**LocatePackageRequest**](LocatePackageRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleSmartHandMoveJumperCableOrder

> HandleSmartHandMoveJumperCableOrder(ctx).Authorization(authorization).Body(body).Execute()

This API is used to request a patch cable to be moved between devices.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smarthandsv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.
	body := *openapiclient.NewMoveJumperCableRequest(*openapiclient.NewIbxLocation("AM1", []openapiclient.IbxLocationCagesInner{*openapiclient.NewIbxLocationCagesInner("AM1:01:001MC3", "12345")}), []openapiclient.ContactInfo{*openapiclient.NewContactInfo(openapiclient.contactInfo_contactType("TECHNICAL"))}, *openapiclient.NewScheduleInfo(openapiclient.scheduleInfo_scheduleType("STANDARD")), *openapiclient.NewMoveJumperCableRequestServiceDetails(openapiclient.moveJumperCableRequest_serviceDetails_quantity("1"), "Scope of work")) // MoveJumperCableRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmarthandsApi.HandleSmartHandMoveJumperCableOrder(context.Background()).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmarthandsApi.HandleSmartHandMoveJumperCableOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleSmartHandMoveJumperCableOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 
 **body** | [**MoveJumperCableRequest**](MoveJumperCableRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleSmartHandOrder

> HandleSmartHandOrder(ctx).Authorization(authorization).Body(body).Execute()

Request equipment installation per your specifications by an IBX Technician



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smarthandsv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.
	body := *openapiclient.NewEquipmentInstallRequest(*openapiclient.NewIbxLocation("AM1", []openapiclient.IbxLocationCagesInner{*openapiclient.NewIbxLocationCagesInner("AM1:01:001MC3", "12345")}), []openapiclient.ContactInfo{*openapiclient.NewContactInfo(openapiclient.contactInfo_contactType("TECHNICAL"))}, *openapiclient.NewScheduleInfo(openapiclient.scheduleInfo_scheduleType("STANDARD")), *openapiclient.NewEquipmentInstallRequestServiceDetails("abc location", false, "abc", false, true, true, true, "Scope of work")) // EquipmentInstallRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmarthandsApi.HandleSmartHandOrder(context.Background()).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmarthandsApi.HandleSmartHandOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleSmartHandOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 
 **body** | [**EquipmentInstallRequest**](EquipmentInstallRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleSmartHandOthersOrder

> HandleSmartHandOthersOrder(ctx).Authorization(authorization).Body(body).Execute()

Request a Smart Hands order not listed in catalogue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smarthandsv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.
	body := *openapiclient.NewOtherRequest(*openapiclient.NewIbxLocation("AM1", []openapiclient.IbxLocationCagesInner{*openapiclient.NewIbxLocationCagesInner("AM1:01:001MC3", "12345")}), []openapiclient.ContactInfo{*openapiclient.NewContactInfo(openapiclient.contactInfo_contactType("TECHNICAL"))}, *openapiclient.NewScheduleInfo(openapiclient.scheduleInfo_scheduleType("STANDARD")), *openapiclient.NewOtherRequestServiceDetails("Scope of work")) // OtherRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmarthandsApi.HandleSmartHandOthersOrder(context.Background()).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmarthandsApi.HandleSmartHandOthersOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleSmartHandOthersOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 
 **body** | [**OtherRequest**](OtherRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleSmartHandPatchCableInstallOrder

> HandleSmartHandPatchCableInstallOrder(ctx).Authorization(authorization).Body(body).Execute()

Request installation of a cross connect patch cable by an IBX Technician.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smarthandsv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.
	body := *openapiclient.NewPatchCableInstallRequest(*openapiclient.NewIbxLocation("AM1", []openapiclient.IbxLocationCagesInner{*openapiclient.NewIbxLocationCagesInner("AM1:01:001MC3", "12345")}), []openapiclient.ContactInfo{*openapiclient.NewContactInfo(openapiclient.contactInfo_contactType("TECHNICAL"))}, *openapiclient.NewScheduleInfo(openapiclient.scheduleInfo_scheduleType("STANDARD")), *openapiclient.NewPatchCableInstallRequestServiceDetails([]openapiclient.CrossConnectInstall{*openapiclient.NewCrossConnectInstall("SerialNumber_example", "DeviceCabinet_example", "DeviceConnectorType_example", "DeviceDetails_example", "DevicePort_example")})) // PatchCableInstallRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmarthandsApi.HandleSmartHandPatchCableInstallOrder(context.Background()).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmarthandsApi.HandleSmartHandPatchCableInstallOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleSmartHandPatchCableInstallOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 
 **body** | [**PatchCableInstallRequest**](PatchCableInstallRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleSmartHandPatchCableRemovalOrder

> HandleSmartHandPatchCableRemovalOrder(ctx).Authorization(authorization).Body(body).Execute()

Request removal of a cross connect cable by an IBX Technician.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smarthandsv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.
	body := *openapiclient.NewPatchCableRemovalRequest(*openapiclient.NewIbxLocation("AM1", []openapiclient.IbxLocationCagesInner{*openapiclient.NewIbxLocationCagesInner("AM1:01:001MC3", "12345")}), []openapiclient.ContactInfo{*openapiclient.NewContactInfo(openapiclient.contactInfo_contactType("TECHNICAL"))}, *openapiclient.NewScheduleInfo(openapiclient.scheduleInfo_scheduleType("STANDARD")), *openapiclient.NewPatchCableRemovalRequestServiceDetails([]openapiclient.CrossConnectRemoval{*openapiclient.NewCrossConnectRemoval("SerialNumber_example", "DeviceCabinet_example", "DeviceConnectorType_example", "DeviceDetails_example", "DevicePort_example")})) // PatchCableRemovalRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmarthandsApi.HandleSmartHandPatchCableRemovalOrder(context.Background()).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmarthandsApi.HandleSmartHandPatchCableRemovalOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleSmartHandPatchCableRemovalOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 
 **body** | [**PatchCableRemovalRequest**](PatchCableRemovalRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleSmartHandPicturesDocumentOrder

> HandleSmartHandPicturesDocumentOrder(ctx).Authorization(authorization).Body(body).Execute()

Request an IBX Technician to provide cage-related pictures or documentation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smarthandsv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.
	body := *openapiclient.NewPicturesDocumentRequest(*openapiclient.NewIbxLocation("AM1", []openapiclient.IbxLocationCagesInner{*openapiclient.NewIbxLocationCagesInner("AM1:01:001MC3", "12345")}), []openapiclient.ContactInfo{*openapiclient.NewContactInfo(openapiclient.contactInfo_contactType("TECHNICAL"))}, *openapiclient.NewScheduleInfo(openapiclient.scheduleInfo_scheduleType("STANDARD")), *openapiclient.NewPicturesDocumentRequestServiceDetails(true, "Scope of work")) // PicturesDocumentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmarthandsApi.HandleSmartHandPicturesDocumentOrder(context.Background()).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmarthandsApi.HandleSmartHandPicturesDocumentOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleSmartHandPicturesDocumentOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 
 **body** | [**PicturesDocumentRequest**](PicturesDocumentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleSmartHandRunJumperCableOrder

> HandleSmartHandRunJumperCableOrder(ctx).Authorization(authorization).Body(body).Execute()

This API is used to request a jumper cable to be ran between devices.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smarthandsv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.
	body := *openapiclient.NewRunJumperCableRequest(*openapiclient.NewIbxLocation("AM1", []openapiclient.IbxLocationCagesInner{*openapiclient.NewIbxLocationCagesInner("AM1:01:001MC3", "12345")}), []openapiclient.ContactInfo{*openapiclient.NewContactInfo(openapiclient.contactInfo_contactType("TECHNICAL"))}, *openapiclient.NewScheduleInfo(openapiclient.scheduleInfo_scheduleType("STANDARD")), *openapiclient.NewRunJumperCableRequestServiceDetails(openapiclient.moveJumperCableRequest_serviceDetails_quantity("1"), "Scope of work")) // RunJumperCableRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmarthandsApi.HandleSmartHandRunJumperCableOrder(context.Background()).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmarthandsApi.HandleSmartHandRunJumperCableOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleSmartHandRunJumperCableOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 
 **body** | [**RunJumperCableRequest**](RunJumperCableRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleSmartHandShipmentUnpackOrder

> HandleSmartHandShipmentUnpackOrder(ctx).Authorization(authorization).Body(body).Execute()

Request inbound shipment unpacking and packaging disposal.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smarthandsv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.
	body := *openapiclient.NewShipmentUnpackRequest(*openapiclient.NewIbxLocation("AM1", []openapiclient.IbxLocationCagesInner{*openapiclient.NewIbxLocationCagesInner("AM1:01:001MC3", "12345")}), []openapiclient.ContactInfo{*openapiclient.NewContactInfo(openapiclient.contactInfo_contactType("TECHNICAL"))}, *openapiclient.NewScheduleInfo(openapiclient.scheduleInfo_scheduleType("STANDARD")), *openapiclient.NewShipmentUnpackRequestServiceDetails("1-12122121", false, "Scope of work", false)) // ShipmentUnpackRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SmarthandsApi.HandleSmartHandShipmentUnpackOrder(context.Background()).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmarthandsApi.HandleSmartHandShipmentUnpackOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleSmartHandShipmentUnpackOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 
 **body** | [**ShipmentUnpackRequest**](ShipmentUnpackRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartHandTypes

> SmartHandTypes200Response SmartHandTypes(ctx).Authorization(authorization).Execute()

All supported smart hands types API.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smarthandsv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmarthandsApi.SmartHandTypes(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmarthandsApi.SmartHandTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmartHandTypes`: SmartHandTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `SmarthandsApi.SmartHandTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmartHandTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 

### Return type

[**SmartHandTypes200Response**](SmartHandTypes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

