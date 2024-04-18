# \BGPApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBgpConfigurationUsingPOST**](BGPApi.md#AddBgpConfigurationUsingPOST) | **Post** /ne/v1/bgp | Create BGP Peering
[**GetBgpConfigurationUsingGET**](BGPApi.md#GetBgpConfigurationUsingGET) | **Get** /ne/v1/bgp/{uuid} | Get BGP Peering {uuid}
[**GetBgpConfigurationsUsingGET**](BGPApi.md#GetBgpConfigurationsUsingGET) | **Get** /ne/v1/bgp | Get BGP Peering
[**UpdateBgpConfigurationUsingPUT**](BGPApi.md#UpdateBgpConfigurationUsingPUT) | **Put** /ne/v1/bgp/{uuid} | Update BGP Peering



## AddBgpConfigurationUsingPOST

> BgpAsyncResponse AddBgpConfigurationUsingPOST(ctx).Authorization(authorization).Request(request).Execute()

Create BGP Peering



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
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	request := *openapiclient.NewBgpConfigAddRequest() // BgpConfigAddRequest | BGP configuration details (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BGPApi.AddBgpConfigurationUsingPOST(context.Background()).Authorization(authorization).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.AddBgpConfigurationUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddBgpConfigurationUsingPOST`: BgpAsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BGPApi.AddBgpConfigurationUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddBgpConfigurationUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **request** | [**BgpConfigAddRequest**](BgpConfigAddRequest.md) | BGP configuration details | 

### Return type

[**BgpAsyncResponse**](BgpAsyncResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBgpConfigurationUsingGET

> BgpInfo GetBgpConfigurationUsingGET(ctx, uuid).Authorization(authorization).Execute()

Get BGP Peering {uuid}



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
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BGPApi.GetBgpConfigurationUsingGET(context.Background(), uuid).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.GetBgpConfigurationUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBgpConfigurationUsingGET`: BgpInfo
	fmt.Fprintf(os.Stdout, "Response from `BGPApi.GetBgpConfigurationUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBgpConfigurationUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 

### Return type

[**BgpInfo**](BgpInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBgpConfigurationsUsingGET

> BgpInfo GetBgpConfigurationsUsingGET(ctx).Authorization(authorization).VirtualDeviceUuid(virtualDeviceUuid).ConnectionUuid(connectionUuid).Status(status).AccountUcmId(accountUcmId).Offset(offset).Limit(limit).Execute()

Get BGP Peering



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
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	virtualDeviceUuid := "virtualDeviceUuid_example" // string | Unique Id of a virtual device (optional)
	connectionUuid := "connectionUuid_example" // string | Unique Id of a connection (optional)
	status := "status_example" // string | Provisioning status of BGP Peering (optional)
	accountUcmId := "accountUcmId_example" // string | Unique Id of an account. A reseller querying for a customer's devices can input the accountUcmId of the customer's account. (optional)
	offset := "offset_example" // string | Specifies where to start a page. It is the starting point of the collection returned from the server. (optional) (default to "0")
	limit := "limit_example" // string | Specifies the page size. (optional) (default to "20")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BGPApi.GetBgpConfigurationsUsingGET(context.Background()).Authorization(authorization).VirtualDeviceUuid(virtualDeviceUuid).ConnectionUuid(connectionUuid).Status(status).AccountUcmId(accountUcmId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.GetBgpConfigurationsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBgpConfigurationsUsingGET`: BgpInfo
	fmt.Fprintf(os.Stdout, "Response from `BGPApi.GetBgpConfigurationsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBgpConfigurationsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **virtualDeviceUuid** | **string** | Unique Id of a virtual device | 
 **connectionUuid** | **string** | Unique Id of a connection | 
 **status** | **string** | Provisioning status of BGP Peering | 
 **accountUcmId** | **string** | Unique Id of an account. A reseller querying for a customer&#39;s devices can input the accountUcmId of the customer&#39;s account. | 
 **offset** | **string** | Specifies where to start a page. It is the starting point of the collection returned from the server. | [default to &quot;0&quot;]
 **limit** | **string** | Specifies the page size. | [default to &quot;20&quot;]

### Return type

[**BgpInfo**](BgpInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBgpConfigurationUsingPUT

> BgpAsyncResponse UpdateBgpConfigurationUsingPUT(ctx, uuid).Authorization(authorization).Request(request).Execute()

Update BGP Peering



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
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	request := *openapiclient.NewBgpUpdateRequest() // BgpUpdateRequest | BGP config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BGPApi.UpdateBgpConfigurationUsingPUT(context.Background(), uuid).Authorization(authorization).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.UpdateBgpConfigurationUsingPUT``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBgpConfigurationUsingPUT`: BgpAsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `BGPApi.UpdateBgpConfigurationUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBgpConfigurationUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **request** | [**BgpUpdateRequest**](BgpUpdateRequest.md) | BGP config | 

### Return type

[**BgpAsyncResponse**](BgpAsyncResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

