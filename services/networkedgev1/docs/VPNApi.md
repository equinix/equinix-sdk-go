# \VPNApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVpnUsingPOST**](VPNApi.md#CreateVpnUsingPOST) | **Post** /ne/v1/vpn | Create VPN Configuration
[**GetVpnByUuidUsingGET**](VPNApi.md#GetVpnByUuidUsingGET) | **Get** /ne/v1/vpn/{uuid} | Get VPN Configuration {uuid}
[**GetVpnsUsingGET**](VPNApi.md#GetVpnsUsingGET) | **Get** /ne/v1/vpn | Get VPN Configurations
[**RemoveVpnConfigurationUsingDELETE**](VPNApi.md#RemoveVpnConfigurationUsingDELETE) | **Delete** /ne/v1/vpn/{uuid} | Delete VPN Configuration
[**UpdateVpnConfigurationUsingPut**](VPNApi.md#UpdateVpnConfigurationUsingPut) | **Put** /ne/v1/vpn/{uuid} | Update VPN Configuration



## CreateVpnUsingPOST

> CreateVpnUsingPOST(ctx).Request(request).Execute()

Create VPN Configuration



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
	request := *openapiclient.NewVpn("Chicago", "f79eead8-b837-41d3-9095-9b15c2c4996d", "110.11.12.222", "5bb2424e888bd", int64(65413), "100.210.1.31", "pass123", "192.168.7.2/30") // Vpn | VPN info (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VPNApi.CreateVpnUsingPOST(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNApi.CreateVpnUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpnUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**Vpn**](Vpn.md) | VPN info | 

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


## GetVpnByUuidUsingGET

> VpnResponse GetVpnByUuidUsingGET(ctx, uuid).Execute()

Get VPN Configuration {uuid}



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
	resp, r, err := apiClient.VPNApi.GetVpnByUuidUsingGET(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNApi.GetVpnByUuidUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpnByUuidUsingGET`: VpnResponse
	fmt.Fprintf(os.Stdout, "Response from `VPNApi.GetVpnByUuidUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpnByUuidUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VpnResponse**](VpnResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpnsUsingGET

> VpnResponseDto GetVpnsUsingGET(ctx).StatusList(statusList).VirtualDeviceUuid(virtualDeviceUuid).Offset(offset).Limit(limit).Execute()

Get VPN Configurations



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
	statusList := []openapiclient.GetVpnsUsingGETStatusListParameterInner{openapiclient.getVpnsUsingGET_statusList_parameter_inner("PROVISIONED")} // []GetVpnsUsingGETStatusListParameterInner | One or more desired status (optional)
	virtualDeviceUuid := "virtualDeviceUuid_example" // string | Unique Id of a virtual device (optional)
	offset := "offset_example" // string | Specifies where to start a page. It is the starting point of the collection returned from the server. (optional) (default to "0")
	limit := "limit_example" // string | Specifies the page size. (optional) (default to "20")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPNApi.GetVpnsUsingGET(context.Background()).StatusList(statusList).VirtualDeviceUuid(virtualDeviceUuid).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNApi.GetVpnsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpnsUsingGET`: VpnResponseDto
	fmt.Fprintf(os.Stdout, "Response from `VPNApi.GetVpnsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVpnsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statusList** | [**[]GetVpnsUsingGETStatusListParameterInner**](GetVpnsUsingGETStatusListParameterInner.md) | One or more desired status | 
 **virtualDeviceUuid** | **string** | Unique Id of a virtual device | 
 **offset** | **string** | Specifies where to start a page. It is the starting point of the collection returned from the server. | [default to &quot;0&quot;]
 **limit** | **string** | Specifies the page size. | [default to &quot;20&quot;]

### Return type

[**VpnResponseDto**](VpnResponseDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveVpnConfigurationUsingDELETE

> RemoveVpnConfigurationUsingDELETE(ctx, uuid).Execute()

Delete VPN Configuration



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
	r, err := apiClient.VPNApi.RemoveVpnConfigurationUsingDELETE(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNApi.RemoveVpnConfigurationUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveVpnConfigurationUsingDELETERequest struct via the builder pattern


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


## UpdateVpnConfigurationUsingPut

> UpdateVpnConfigurationUsingPut(ctx, uuid).Request(request).Execute()

Update VPN Configuration



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
	request := *openapiclient.NewVpn("Chicago", "f79eead8-b837-41d3-9095-9b15c2c4996d", "110.11.12.222", "5bb2424e888bd", int64(65413), "100.210.1.31", "pass123", "192.168.7.2/30") // Vpn | VPN info (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VPNApi.UpdateVpnConfigurationUsingPut(context.Background(), uuid).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPNApi.UpdateVpnConfigurationUsingPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVpnConfigurationUsingPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**Vpn**](Vpn.md) | VPN info | 

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

