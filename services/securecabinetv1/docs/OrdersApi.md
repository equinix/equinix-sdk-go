# \OrdersApi

All URIs are relative to *https://virtserver.swaggerhub.com/equinix-api/secure-cab-express-external-api/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrder**](OrdersApi.md#CreateOrder) | **Post** /securecabinet/v1/orders | Order a new Secure Cabinet deployment



## CreateOrder

> OrderCreateResponse CreateOrder(ctx).OrderCreateRequest(orderCreateRequest).Execute()

Order a new Secure Cabinet deployment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/securecabinetv1"
)

func main() {
	orderCreateRequest := *openapiclient.NewOrderCreateRequest("132028", "CH1", openapiclient.ContractTerm("TERM_24_MONTHS"), *openapiclient.NewOrderItem(float64(5), true, int32(1), *openapiclient.NewDimensions(*openapiclient.NewDimension(int32(482), openapiclient.Dimension_unit("MILLIMETER")), *openapiclient.NewDimension(int32(482), openapiclient.Dimension_unit("MILLIMETER")), *openapiclient.NewDimension(int32(482), openapiclient.Dimension_unit("MILLIMETER"))), true)) // OrderCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersApi.CreateOrder(context.Background()).OrderCreateRequest(orderCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrder`: OrderCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderCreateRequest** | [**OrderCreateRequest**](OrderCreateRequest.md) |  | 

### Return type

[**OrderCreateResponse**](OrderCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

