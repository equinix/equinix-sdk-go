# \DeviceRMAApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostDeviceRMAUsingPOST**](DeviceRMAApi.md#PostDeviceRMAUsingPOST) | **Post** /ne/v1/devices/{virtualDeviceUuid}/rma | Trigger Device RMA



## PostDeviceRMAUsingPOST

> PostDeviceRMAUsingPOST(ctx, virtualDeviceUuid).Authorization(authorization).Request(request).Execute()

Trigger Device RMA



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
	virtualDeviceUuid := "virtualDeviceUuid_example" // string | Unique ID of a device
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	request := *openapiclient.NewDeviceRMAPostRequest("17.09.01a") // DeviceRMAPostRequest | Post RMA request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceRMAApi.PostDeviceRMAUsingPOST(context.Background(), virtualDeviceUuid).Authorization(authorization).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceRMAApi.PostDeviceRMAUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualDeviceUuid** | **string** | Unique ID of a device | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeviceRMAUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **request** | [**DeviceRMAPostRequest**](DeviceRMAPostRequest.md) | Post RMA request | 

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

