# \AvailabilityApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProductsAvailability**](AvailabilityApi.md#GetProductsAvailability) | **Get** /securecabinet/v1/availability/{accountNumber} | Secure Cabinet availability.



## GetProductsAvailability

> []ProductsAvailability GetProductsAvailability(ctx, accountNumber).Execute()

Secure Cabinet availability.



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
	accountNumber := "accountNumber_example" // string | Billing Account Number.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvailabilityApi.GetProductsAvailability(context.Background(), accountNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityApi.GetProductsAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductsAvailability`: []ProductsAvailability
	fmt.Fprintf(os.Stdout, "Response from `AvailabilityApi.GetProductsAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountNumber** | **string** | Billing Account Number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductsAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProductsAvailability**](ProductsAvailability.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

