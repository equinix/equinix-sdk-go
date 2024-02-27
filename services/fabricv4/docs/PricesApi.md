# \PricesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchPrices**](PricesApi.md#SearchPrices) | **Post** /fabric/v4/prices/search | Get Prices



## SearchPrices

> PriceSearchResponse SearchPrices(ctx).FilterBody(filterBody).Execute()

Get Prices



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
)

func main() {
	filterBody := *openapiclient.NewFilterBody() // FilterBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesApi.SearchPrices(context.Background()).FilterBody(filterBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesApi.SearchPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPrices`: PriceSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `PricesApi.SearchPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterBody** | [**FilterBody**](FilterBody.md) |  | 

### Return type

[**PriceSearchResponse**](PriceSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

