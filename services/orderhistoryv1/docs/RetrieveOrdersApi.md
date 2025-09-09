# \RetrieveOrdersApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETRetrieveOrdersLocations**](RetrieveOrdersApi.md#GETRetrieveOrdersLocations) | **Get** /v1/retrieve-orders/locations | Retrieve order permissible IBX locations
[**POSTOrdersHistory**](RetrieveOrdersApi.md#POSTOrdersHistory) | **Post** /v1/retrieve-orders | Search Orders History



## GETRetrieveOrdersLocations

> []PermissibleLocation GETRetrieveOrdersLocations(ctx).Authorization(authorization).Execute()

Retrieve order permissible IBX locations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/orderhistoryv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrieveOrdersApi.GETRetrieveOrdersLocations(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrieveOrdersApi.GETRetrieveOrdersLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GETRetrieveOrdersLocations`: []PermissibleLocation
	fmt.Fprintf(os.Stdout, "Response from `RetrieveOrdersApi.GETRetrieveOrdersLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGETRetrieveOrdersLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 

### Return type

[**[]PermissibleLocation**](PermissibleLocation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTOrdersHistory

> OrderHistoryResponse POSTOrdersHistory(ctx).Authorization(authorization).Body(body).Execute()

Search Orders History



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/orderhistoryv1"
)

func main() {
	authorization := "authorization_example" // string | Specify the Access token with prefix 'Bearer'.
	body := *openapiclient.NewOrderhistoryapirequest() // Orderhistoryapirequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RetrieveOrdersApi.POSTOrdersHistory(context.Background()).Authorization(authorization).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RetrieveOrdersApi.POSTOrdersHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `POSTOrdersHistory`: OrderHistoryResponse
	fmt.Fprintf(os.Stdout, "Response from `RetrieveOrdersApi.POSTOrdersHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTOrdersHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Specify the Access token with prefix &#39;Bearer&#39;. | 
 **body** | [**Orderhistoryapirequest**](Orderhistoryapirequest.md) |  | 

### Return type

[**OrderHistoryResponse**](OrderHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

