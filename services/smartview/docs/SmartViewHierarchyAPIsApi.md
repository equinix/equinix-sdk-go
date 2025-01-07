# \SmartViewHierarchyAPIsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocationHierarchy**](SmartViewHierarchyAPIsApi.md#GetLocationHierarchy) | **Get** /smartview/v1/hierarchy/location | Fetch the Location Hierarchy
[**GetPowerHierarchy**](SmartViewHierarchyAPIsApi.md#GetPowerHierarchy) | **Get** /smartview/v1/hierarchy/power | Fetch the Power Hierarchy



## GetLocationHierarchy

> []HierarchyNode GetLocationHierarchy(ctx).Authorization(authorization).AccountNo(accountNo).Authorization2(authorization2).Ibx(ibx).Execute()

Fetch the Location Hierarchy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smartview"
)

func main() {
	authorization := "authorization_example" // string | Specify the OAuth Bearer token with prefix 'Bearer '.
	accountNo := "accountNo_example" // string | Customer Account Number
	authorization2 := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	ibx := "ibx_example" // string | IBX Code (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewHierarchyAPIsApi.GetLocationHierarchy(context.Background()).Authorization(authorization).AccountNo(accountNo).Authorization2(authorization2).Ibx(ibx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewHierarchyAPIsApi.GetLocationHierarchy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocationHierarchy`: []HierarchyNode
	fmt.Fprintf(os.Stdout, "Response from `SmartViewHierarchyAPIsApi.GetLocationHierarchy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationHierarchyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the OAuth Bearer token with prefix &#39;Bearer &#39;. | 
 **accountNo** | **string** | Customer Account Number | 
 **authorization2** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **ibx** | **string** | IBX Code | 

### Return type

[**[]HierarchyNode**](HierarchyNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPowerHierarchy

> []PowerHierarchyNode GetPowerHierarchy(ctx).Authorization(authorization).AccountNo(accountNo).Authorization2(authorization2).Ibx(ibx).Execute()

Fetch the Power Hierarchy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smartview"
)

func main() {
	authorization := "authorization_example" // string | Specify the OAuth Bearer token with prefix 'Bearer '.
	accountNo := "accountNo_example" // string | Customer Account Number
	authorization2 := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	ibx := "ibx_example" // string | IBX Code (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SmartViewHierarchyAPIsApi.GetPowerHierarchy(context.Background()).Authorization(authorization).AccountNo(accountNo).Authorization2(authorization2).Ibx(ibx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartViewHierarchyAPIsApi.GetPowerHierarchy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPowerHierarchy`: []PowerHierarchyNode
	fmt.Fprintf(os.Stdout, "Response from `SmartViewHierarchyAPIsApi.GetPowerHierarchy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPowerHierarchyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the OAuth Bearer token with prefix &#39;Bearer &#39;. | 
 **accountNo** | **string** | Customer Account Number | 
 **authorization2** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **ibx** | **string** | IBX Code | 

### Return type

[**[]PowerHierarchyNode**](PowerHierarchyNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

