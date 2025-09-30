# \VirtualDeviceApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVirtualDeviceUsingPOST**](VirtualDeviceApi.md#CreateVirtualDeviceUsingPOST) | **Post** /ne/v1/devices | Create Virtual Device
[**DeleteVRouterUsingDELETE**](VirtualDeviceApi.md#DeleteVRouterUsingDELETE) | **Delete** /ne/v1/devices/{uuid} | Delete Virtual Devices
[**GetDeviceReloadUsingGET1**](VirtualDeviceApi.md#GetDeviceReloadUsingGET1) | **Get** /ne/v1/devices/{virtualDeviceUUID}/softReboot | Device Reload History
[**GetDeviceUpgradeUsingGET1**](VirtualDeviceApi.md#GetDeviceUpgradeUsingGET1) | **Get** /ne/v1/devices/{virtualDeviceUuid}/resourceUpgrade | Get Device Upgrade History
[**GetInterfaceStatisticsUsingGET**](VirtualDeviceApi.md#GetInterfaceStatisticsUsingGET) | **Get** /ne/v1/devices/{virtualDeviceUuid}/interfaces/{interfaceId}/stats | Get Interface Statistics
[**GetVirtualDeviceInterfacesUsingGET**](VirtualDeviceApi.md#GetVirtualDeviceInterfacesUsingGET) | **Get** /ne/v1/devices/{uuid}/interfaces | Get Device Interfaces
[**GetVirtualDeviceUsingGET**](VirtualDeviceApi.md#GetVirtualDeviceUsingGET) | **Get** /ne/v1/devices/{uuid} | Get Virtual Device {uuid}
[**GetVirtualDevicesUsingGET1**](VirtualDeviceApi.md#GetVirtualDevicesUsingGET1) | **Get** /ne/v1/devices | Get Virtual Devices
[**PingDeviceUsingGET**](VirtualDeviceApi.md#PingDeviceUsingGET) | **Get** /ne/v1/devices/{uuid}/ping | Ping Virtual Device
[**PostDeviceReloadUsingPOST1**](VirtualDeviceApi.md#PostDeviceReloadUsingPOST1) | **Post** /ne/v1/devices/{virtualDeviceUUID}/softReboot | Trigger Soft Reboot
[**UpdateAdditionalBandwidth**](VirtualDeviceApi.md#UpdateAdditionalBandwidth) | **Put** /ne/v1/devices/{uuid}/additionalBandwidths | Update Additional Bandwidth
[**UpdateVirtualDeviceUsingPATCH1**](VirtualDeviceApi.md#UpdateVirtualDeviceUsingPATCH1) | **Patch** /ne/v1/devices/{uuid} | Update Virtual Device
[**UpdateVirtualDeviceUsingPUT**](VirtualDeviceApi.md#UpdateVirtualDeviceUsingPUT) | **Put** /ne/v1/devices/{uuid} | Update Device Draft



## CreateVirtualDeviceUsingPOST

> VirtualDeviceCreateResponse CreateVirtualDeviceUsingPOST(ctx).VirtualDevice(virtualDevice).Draft(draft).DraftUuid(draftUuid).Execute()

Create Virtual Device



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
	virtualDevice := *openapiclient.NewVirtualDeviceRequest("16.09.03", "C8000V", true, "SUB", "SV", "Router1-c8000v", []string{"Notifications_example"}, "EQUINIX-CONFIGURED", int32(123)) // VirtualDeviceRequest | Create a virtual device (e.g., a router or a firewall)
	draft := true // bool | draft (optional) (default to false)
	draftUuid := "draftUuid_example" // string | draftUuid (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualDeviceApi.CreateVirtualDeviceUsingPOST(context.Background()).VirtualDevice(virtualDevice).Draft(draft).DraftUuid(draftUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualDeviceApi.CreateVirtualDeviceUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVirtualDeviceUsingPOST`: VirtualDeviceCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualDeviceApi.CreateVirtualDeviceUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualDeviceUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualDevice** | [**VirtualDeviceRequest**](VirtualDeviceRequest.md) | Create a virtual device (e.g., a router or a firewall) | 
 **draft** | **bool** | draft | [default to false]
 **draftUuid** | **string** | draftUuid | 

### Return type

[**VirtualDeviceCreateResponse**](VirtualDeviceCreateResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVRouterUsingDELETE

> DeleteVRouterUsingDELETE(ctx, uuid).DeleteRedundantDevice(deleteRedundantDevice).DeletionInfo(deletionInfo).Execute()

Delete Virtual Devices



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
	uuid := "uuid_example" // string | Unique Id of the virtual device.
	deleteRedundantDevice := true // bool | Optional parameter in case you have a secondary device. As both primary and secondary devices are deleted simultaneously, this field must be marked True so we delete both the devices simultaneously. (optional) (default to false)
	deletionInfo := *openapiclient.NewVirtualDeviceDeleteRequest() // VirtualDeviceDeleteRequest | deletionInfo (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualDeviceApi.DeleteVRouterUsingDELETE(context.Background(), uuid).DeleteRedundantDevice(deleteRedundantDevice).DeletionInfo(deletionInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualDeviceApi.DeleteVRouterUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique Id of the virtual device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVRouterUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteRedundantDevice** | **bool** | Optional parameter in case you have a secondary device. As both primary and secondary devices are deleted simultaneously, this field must be marked True so we delete both the devices simultaneously. | [default to false]
 **deletionInfo** | [**VirtualDeviceDeleteRequest**](VirtualDeviceDeleteRequest.md) | deletionInfo | 

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


## GetDeviceReloadUsingGET1

> DeviceRebootPageResponse GetDeviceReloadUsingGET1(ctx, virtualDeviceUUID).Offset(offset).Limit(limit).Execute()

Device Reload History



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
	virtualDeviceUUID := "virtualDeviceUUID_example" // string | Unique ID of a virtual device.
	offset := "offset_example" // string | Specifies where to start a page. It is the starting point of the collection returned from the server. (optional) (default to "0")
	limit := "limit_example" // string | Specifies the page size. (optional) (default to "20")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualDeviceApi.GetDeviceReloadUsingGET1(context.Background(), virtualDeviceUUID).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualDeviceApi.GetDeviceReloadUsingGET1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceReloadUsingGET1`: DeviceRebootPageResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualDeviceApi.GetDeviceReloadUsingGET1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualDeviceUUID** | **string** | Unique ID of a virtual device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceReloadUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **string** | Specifies where to start a page. It is the starting point of the collection returned from the server. | [default to &quot;0&quot;]
 **limit** | **string** | Specifies the page size. | [default to &quot;20&quot;]

### Return type

[**DeviceRebootPageResponse**](DeviceRebootPageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceUpgradeUsingGET1

> DeviceUpgradePageResponse GetDeviceUpgradeUsingGET1(ctx, virtualDeviceUuid).Offset(offset).Limit(limit).Execute()

Get Device Upgrade History



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
	virtualDeviceUuid := "virtualDeviceUuid_example" // string | Unique ID of a virtual device.
	offset := "offset_example" // string | Specifies where to start a page. It is the starting point of the collection returned from the server. (optional) (default to "0")
	limit := "limit_example" // string | Specifies the page size. (optional) (default to "20")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualDeviceApi.GetDeviceUpgradeUsingGET1(context.Background(), virtualDeviceUuid).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualDeviceApi.GetDeviceUpgradeUsingGET1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceUpgradeUsingGET1`: DeviceUpgradePageResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualDeviceApi.GetDeviceUpgradeUsingGET1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualDeviceUuid** | **string** | Unique ID of a virtual device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceUpgradeUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **string** | Specifies where to start a page. It is the starting point of the collection returned from the server. | [default to &quot;0&quot;]
 **limit** | **string** | Specifies the page size. | [default to &quot;20&quot;]

### Return type

[**DeviceUpgradePageResponse**](DeviceUpgradePageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterfaceStatisticsUsingGET

> InterfaceStatsObject GetInterfaceStatisticsUsingGET(ctx, virtualDeviceUuid, interfaceId).StartDateTime(startDateTime).EndDateTime(endDateTime).Execute()

Get Interface Statistics



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
	virtualDeviceUuid := "virtualDeviceUuid_example" // string | Unique Id of a device
	interfaceId := "interfaceId_example" // string | Interface Id
	startDateTime := "startDateTime_example" // string | Start time of the duration for which you want stats. (YYYY-MM-DDThh:mm:ssZ)
	endDateTime := "endDateTime_example" // string | End time of the duration for which you want stats. (YYYY-MM-DDThh:mm:ssZ)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualDeviceApi.GetInterfaceStatisticsUsingGET(context.Background(), virtualDeviceUuid, interfaceId).StartDateTime(startDateTime).EndDateTime(endDateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualDeviceApi.GetInterfaceStatisticsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInterfaceStatisticsUsingGET`: InterfaceStatsObject
	fmt.Fprintf(os.Stdout, "Response from `VirtualDeviceApi.GetInterfaceStatisticsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualDeviceUuid** | **string** | Unique Id of a device | 
**interfaceId** | **string** | Interface Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterfaceStatisticsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDateTime** | **string** | Start time of the duration for which you want stats. (YYYY-MM-DDThh:mm:ssZ) | 
 **endDateTime** | **string** | End time of the duration for which you want stats. (YYYY-MM-DDThh:mm:ssZ) | 

### Return type

[**InterfaceStatsObject**](InterfaceStatsObject.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualDeviceInterfacesUsingGET

> []InterfaceBasicInfoResponse GetVirtualDeviceInterfacesUsingGET(ctx, uuid).Execute()

Get Device Interfaces



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
	uuid := "uuid_example" // string | uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualDeviceApi.GetVirtualDeviceInterfacesUsingGET(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualDeviceApi.GetVirtualDeviceInterfacesUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualDeviceInterfacesUsingGET`: []InterfaceBasicInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualDeviceApi.GetVirtualDeviceInterfacesUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualDeviceInterfacesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InterfaceBasicInfoResponse**](InterfaceBasicInfoResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualDeviceUsingGET

> VirtualDeviceDetailsResponse GetVirtualDeviceUsingGET(ctx, uuid).Execute()

Get Virtual Device {uuid}



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
	uuid := "uuid_example" // string | uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualDeviceApi.GetVirtualDeviceUsingGET(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualDeviceApi.GetVirtualDeviceUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualDeviceUsingGET`: VirtualDeviceDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualDeviceApi.GetVirtualDeviceUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualDeviceUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualDeviceDetailsResponse**](VirtualDeviceDetailsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualDevicesUsingGET1

> VirtualDevicePageResponse GetVirtualDevicesUsingGET1(ctx).Offset(offset).Limit(limit).MetroCode(metroCode).Status(status).ShowOnlySubCustomerDevices(showOnlySubCustomerDevices).AccountUcmId(accountUcmId).SearchText(searchText).Sort(sort).Execute()

Get Virtual Devices



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
	offset := "offset_example" // string | Specifies where to start a page. It is the starting point of the collection returned from the server. (optional) (default to "0")
	limit := "limit_example" // string | Specifies the page size. (optional) (default to "20")
	metroCode := "metroCode_example" // string | metroCode (optional)
	status := "status_example" // string | status (optional)
	showOnlySubCustomerDevices := true // bool | Resellers may mark this Yes to see sub customer devices. (optional)
	accountUcmId := "accountUcmId_example" // string | Unique ID of the account. (optional)
	searchText := "searchText_example" // string | Enter text to fetch only matching device names (optional)
	sort := []string{"Inner_example"} // []string | Sorts the output based on field names. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualDeviceApi.GetVirtualDevicesUsingGET1(context.Background()).Offset(offset).Limit(limit).MetroCode(metroCode).Status(status).ShowOnlySubCustomerDevices(showOnlySubCustomerDevices).AccountUcmId(accountUcmId).SearchText(searchText).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualDeviceApi.GetVirtualDevicesUsingGET1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualDevicesUsingGET1`: VirtualDevicePageResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualDeviceApi.GetVirtualDevicesUsingGET1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualDevicesUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **string** | Specifies where to start a page. It is the starting point of the collection returned from the server. | [default to &quot;0&quot;]
 **limit** | **string** | Specifies the page size. | [default to &quot;20&quot;]
 **metroCode** | **string** | metroCode | 
 **status** | **string** | status | 
 **showOnlySubCustomerDevices** | **bool** | Resellers may mark this Yes to see sub customer devices. | 
 **accountUcmId** | **string** | Unique ID of the account. | 
 **searchText** | **string** | Enter text to fetch only matching device names | 
 **sort** | **[]string** | Sorts the output based on field names. | 

### Return type

[**VirtualDevicePageResponse**](VirtualDevicePageResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingDeviceUsingGET

> PingDeviceUsingGET(ctx, uuid).Execute()

Ping Virtual Device



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
	uuid := "uuid_example" // string | Virtual Device unique Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualDeviceApi.PingDeviceUsingGET(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualDeviceApi.PingDeviceUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Virtual Device unique Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPingDeviceUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PostDeviceReloadUsingPOST1

> PostDeviceReloadUsingPOST1(ctx, virtualDeviceUUID).Execute()

Trigger Soft Reboot



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
	virtualDeviceUUID := "virtualDeviceUUID_example" // string | Unique ID of a virtual device.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualDeviceApi.PostDeviceReloadUsingPOST1(context.Background(), virtualDeviceUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualDeviceApi.PostDeviceReloadUsingPOST1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualDeviceUUID** | **string** | Unique ID of a virtual device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeviceReloadUsingPOST1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UpdateAdditionalBandwidth

> UpdateAdditionalBandwidth(ctx, uuid).Request(request).Execute()

Update Additional Bandwidth



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
	uuid := "uuid_example" // string | The unique Id of a virtual device
	request := *openapiclient.NewAdditionalBandwidthRequest() // AdditionalBandwidthRequest | Additional Bandwidth

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualDeviceApi.UpdateAdditionalBandwidth(context.Background(), uuid).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualDeviceApi.UpdateAdditionalBandwidth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | The unique Id of a virtual device | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdditionalBandwidthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**AdditionalBandwidthRequest**](AdditionalBandwidthRequest.md) | Additional Bandwidth | 

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


## UpdateVirtualDeviceUsingPATCH1

> UpdateVirtualDeviceUsingPATCH1(ctx, uuid).VirtualDeviceUpdateRequestDto(virtualDeviceUpdateRequestDto).Execute()

Update Virtual Device



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
	uuid := "uuid_example" // string | The unique Id of the device.
	virtualDeviceUpdateRequestDto := *openapiclient.NewVirtualDeviceInternalPatchRequestDto() // VirtualDeviceInternalPatchRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualDeviceApi.UpdateVirtualDeviceUsingPATCH1(context.Background(), uuid).VirtualDeviceUpdateRequestDto(virtualDeviceUpdateRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualDeviceApi.UpdateVirtualDeviceUsingPATCH1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | The unique Id of the device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualDeviceUsingPATCH1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **virtualDeviceUpdateRequestDto** | [**VirtualDeviceInternalPatchRequestDto**](VirtualDeviceInternalPatchRequestDto.md) |  | 

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


## UpdateVirtualDeviceUsingPUT

> UpdateVirtualDeviceUsingPUT(ctx, uuid).Draft(draft).VirtualDevice(virtualDevice).Execute()

Update Device Draft



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
	draft := true // bool | draft (default to true)
	uuid := "uuid_example" // string | Unique Id of a Virtual Device
	virtualDevice := *openapiclient.NewVirtualDeviceRequest("16.09.03", "C8000V", true, "SUB", "SV", "Router1-c8000v", []string{"Notifications_example"}, "EQUINIX-CONFIGURED", int32(123)) // VirtualDeviceRequest | Update virtual device details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VirtualDeviceApi.UpdateVirtualDeviceUsingPUT(context.Background(), uuid).Draft(draft).VirtualDevice(virtualDevice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualDeviceApi.UpdateVirtualDeviceUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique Id of a Virtual Device | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualDeviceUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **draft** | **bool** | draft | [default to true]

 **virtualDevice** | [**VirtualDeviceRequest**](VirtualDeviceRequest.md) | Update virtual device details | 

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

