# \OrdersApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNotesToAnOrder**](OrdersApi.md#AddNotesToAnOrder) | **Post** /colocations/v2/orders/{orderId}/notes | Add notes to an order
[**CancelAnOrder**](OrdersApi.md#CancelAnOrder) | **Post** /colocations/v2/orders/{orderId}/cancel | Cancel an order
[**GETOrderDetails**](OrdersApi.md#GETOrderDetails) | **Get** /colocations/v2/orders/{orderId} | Retrieve an order
[**ReplyToAnOrderNegotiation**](OrdersApi.md#ReplyToAnOrderNegotiation) | **Post** /colocations/v2/orders/{orderId}/negotiations | Reply to an order negotiation
[**RetrieveOrderNegotiations**](OrdersApi.md#RetrieveOrderNegotiations) | **Get** /colocations/v2/orders/{orderId}/negotiations | Retrieve order negotiations



## AddNotesToAnOrder

> AddNotesToAnOrder(ctx, orderId).NoteRequest(noteRequest).Execute()

Add notes to an order



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/ordersv2"
)

func main() {
	orderId := "{orderId}" // string | Identifier of the order
	noteRequest := *openapiclient.NewNoteRequest("The text of the note") // NoteRequest | The Order Id to be updated (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrdersApi.AddNotesToAnOrder(context.Background(), orderId).NoteRequest(noteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.AddNotesToAnOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | Identifier of the order | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddNotesToAnOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noteRequest** | [**NoteRequest**](NoteRequest.md) | The Order Id to be updated | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelAnOrder

> CancelAnOrder(ctx, orderId).CancelRequest(cancelRequest).Execute()

Cancel an order



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/ordersv2"
)

func main() {
	orderId := "orderId_example" // string | Identifier of the order
	cancelRequest := *openapiclient.NewCancelRequest("Reason_example") // CancelRequest | The Orders to be cancelled

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrdersApi.CancelAnOrder(context.Background(), orderId).CancelRequest(cancelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CancelAnOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | Identifier of the order | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelAnOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelRequest** | [**CancelRequest**](CancelRequest.md) | The Orders to be cancelled | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderDetails

> Orders GETOrderDetails(ctx, orderId).Ibxs(ibxs).Execute()

Retrieve an order



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/ordersv2"
)

func main() {
	orderId := "orderId_example" // string | Identifier of the Order
	ibxs := []string{"Inner_example"} // []string | List of IBXs to be filtered (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersApi.GETOrderDetails(context.Background(), orderId).Ibxs(ibxs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GETOrderDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GETOrderDetails`: Orders
	fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GETOrderDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | Identifier of the Order | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ibxs** | **[]string** | List of IBXs to be filtered | 

### Return type

[**Orders**](Orders.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplyToAnOrderNegotiation

> ReplyToAnOrderNegotiation(ctx, orderId).NegotiationsRequest(negotiationsRequest).Execute()

Reply to an order negotiation



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/ordersv2"
)

func main() {
	orderId := "{orderId}" // string | Identifier of the order
	negotiationsRequest := *openapiclient.NewNegotiationsRequest("4-9091830", openapiclient.Negotiations_request_action("APPROVE")) // NegotiationsRequest | The action to be taken for an order negotiation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrdersApi.ReplyToAnOrderNegotiation(context.Background(), orderId).NegotiationsRequest(negotiationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ReplyToAnOrderNegotiation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | Identifier of the order | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplyToAnOrderNegotiationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **negotiationsRequest** | [**NegotiationsRequest**](NegotiationsRequest.md) | The action to be taken for an order negotiation | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveOrderNegotiations

> []Negotiations RetrieveOrderNegotiations(ctx, orderId).Execute()

Retrieve order negotiations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/ordersv2"
)

func main() {
	orderId := "{orderId}" // string | Identifier of the order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersApi.RetrieveOrderNegotiations(context.Background(), orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.RetrieveOrderNegotiations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveOrderNegotiations`: []Negotiations
	fmt.Fprintf(os.Stdout, "Response from `OrdersApi.RetrieveOrderNegotiations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | Identifier of the order | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveOrderNegotiationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Negotiations**](Negotiations.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

