# \DeviceLinkApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLinkGroupUsingPOST**](DeviceLinkApi.md#CreateLinkGroupUsingPOST) | **Post** /ne/v1/links | Create Device Link
[**DeleteLinkGroupUsingDELETE**](DeviceLinkApi.md#DeleteLinkGroupUsingDELETE) | **Delete** /ne/v1/links/{uuid} | Delete Device Link
[**GetLinkGroupByUUIDUsingGET**](DeviceLinkApi.md#GetLinkGroupByUUIDUsingGET) | **Get** /ne/v1/links/{uuid} | Get Device Link {uuid}
[**GetLinkGroupsUsingGET1**](DeviceLinkApi.md#GetLinkGroupsUsingGET1) | **Get** /ne/v1/links | Get Device Links.
[**UpdateLinkGroupUsingPATCH**](DeviceLinkApi.md#UpdateLinkGroupUsingPATCH) | **Patch** /ne/v1/links/{uuid} | Update Device Link



## CreateLinkGroupUsingPOST

> DeviceLinkGroupResponse CreateLinkGroupUsingPOST(ctx).Authorization(authorization).DeviceLinkGroup(deviceLinkGroup).Execute()

Create Device Link



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
	deviceLinkGroup := *openapiclient.NewDeviceLinkRequest("linkGroup", "192.164.0.0/29") // DeviceLinkRequest | New Device Link Group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceLinkApi.CreateLinkGroupUsingPOST(context.Background()).Authorization(authorization).DeviceLinkGroup(deviceLinkGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceLinkApi.CreateLinkGroupUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLinkGroupUsingPOST`: DeviceLinkGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceLinkApi.CreateLinkGroupUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLinkGroupUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **deviceLinkGroup** | [**DeviceLinkRequest**](DeviceLinkRequest.md) | New Device Link Group | 

### Return type

[**DeviceLinkGroupResponse**](DeviceLinkGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLinkGroupUsingDELETE

> DeleteLinkGroupUsingDELETE(ctx, uuid).Authorization(authorization).Execute()

Delete Device Link



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
	uuid := "uuid_example" // string | Unique Id of a device link group.
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceLinkApi.DeleteLinkGroupUsingDELETE(context.Background(), uuid).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceLinkApi.DeleteLinkGroupUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique Id of a device link group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinkGroupUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkGroupByUUIDUsingGET

> DeviceLinkGroupDto GetLinkGroupByUUIDUsingGET(ctx, uuid).Authorization(authorization).Execute()

Get Device Link {uuid}



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
	uuid := "uuid_example" // string | Unique Id of a device link group.
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceLinkApi.GetLinkGroupByUUIDUsingGET(context.Background(), uuid).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceLinkApi.GetLinkGroupByUUIDUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkGroupByUUIDUsingGET`: DeviceLinkGroupDto
	fmt.Fprintf(os.Stdout, "Response from `DeviceLinkApi.GetLinkGroupByUUIDUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique Id of a device link group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkGroupByUUIDUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 

### Return type

[**DeviceLinkGroupDto**](DeviceLinkGroupDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkGroupsUsingGET1

> []LinksPageResponse GetLinkGroupsUsingGET1(ctx).Authorization(authorization).Metro(metro).VirtualDeviceUuid(virtualDeviceUuid).AccountUcmId(accountUcmId).GroupUuid(groupUuid).GroupName(groupName).Offset(offset).Limit(limit).Execute()

Get Device Links.



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
	metro := "metro_example" // string | Metro Code (optional)
	virtualDeviceUuid := "virtualDeviceUuid_example" // string | Unique Id of a virtual device. (optional)
	accountUcmId := "accountUcmId_example" // string | Unique Id of the account. A reseller querying for a customer's link groups can pass the accountUcmId of the customer's account. To get the accountUcmId of your customer's account, please check the Equinix account creation portal (ECP) or call Get account API. (optional)
	groupUuid := "groupUuid_example" // string | Unique Id of a link group (optional)
	groupName := "groupName_example" // string | The name of a link group (optional)
	offset := "offset_example" // string | Specifies where to start a page. It is the starting point of the collection returned from the server. (optional) (default to "0")
	limit := "limit_example" // string | Specifies the page size. (optional) (default to "20")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceLinkApi.GetLinkGroupsUsingGET1(context.Background()).Authorization(authorization).Metro(metro).VirtualDeviceUuid(virtualDeviceUuid).AccountUcmId(accountUcmId).GroupUuid(groupUuid).GroupName(groupName).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceLinkApi.GetLinkGroupsUsingGET1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkGroupsUsingGET1`: []LinksPageResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceLinkApi.GetLinkGroupsUsingGET1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkGroupsUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **metro** | **string** | Metro Code | 
 **virtualDeviceUuid** | **string** | Unique Id of a virtual device. | 
 **accountUcmId** | **string** | Unique Id of the account. A reseller querying for a customer&#39;s link groups can pass the accountUcmId of the customer&#39;s account. To get the accountUcmId of your customer&#39;s account, please check the Equinix account creation portal (ECP) or call Get account API. | 
 **groupUuid** | **string** | Unique Id of a link group | 
 **groupName** | **string** | The name of a link group | 
 **offset** | **string** | Specifies where to start a page. It is the starting point of the collection returned from the server. | [default to &quot;0&quot;]
 **limit** | **string** | Specifies the page size. | [default to &quot;20&quot;]

### Return type

[**[]LinksPageResponse**](LinksPageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLinkGroupUsingPATCH

> UpdateLinkGroupUsingPATCH(ctx, uuid).Authorization(authorization).DeviceLinkGroup(deviceLinkGroup).Execute()

Update Device Link



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
	uuid := "uuid_example" // string | Unique Id of a device link group.
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	deviceLinkGroup := *openapiclient.NewDeviceLinkRequest("linkGroup", "192.164.0.0/29") // DeviceLinkRequest | Device Link Group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceLinkApi.UpdateLinkGroupUsingPATCH(context.Background(), uuid).Authorization(authorization).DeviceLinkGroup(deviceLinkGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceLinkApi.UpdateLinkGroupUsingPATCH``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique Id of a device link group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLinkGroupUsingPATCHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **deviceLinkGroup** | [**DeviceLinkRequest**](DeviceLinkRequest.md) | Device Link Group | 

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

