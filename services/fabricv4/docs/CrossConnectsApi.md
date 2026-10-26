# \CrossConnectsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCrossConnect**](CrossConnectsApi.md#CreateCrossConnect) | **Post** /fabric/v4/crossConnects | Create Cross Connect
[**CreateDiverseCrossConnect**](CrossConnectsApi.md#CreateDiverseCrossConnect) | **Post** /fabric/v4/crossConnects/bulk | Create Diverse Cross Connect
[**GetCrossConnectActionByUuid**](CrossConnectsApi.md#GetCrossConnectActionByUuid) | **Get** /fabric/v4/crossConnects/{crossConnectId}/actions/{actionId} | Get Cross Connect Action by Action ID
[**GetCrossConnectAttachmentsByUuid**](CrossConnectsApi.md#GetCrossConnectAttachmentsByUuid) | **Get** /fabric/v4/crossConnects/{crossConnectId}/attachments | Get Cross Connect Attachments
[**GetCrossConnectByUuid**](CrossConnectsApi.md#GetCrossConnectByUuid) | **Get** /fabric/v4/crossConnects/{crossConnectId} | Get Cross Connect
[**GetCrossConnectNotesByUuid**](CrossConnectsApi.md#GetCrossConnectNotesByUuid) | **Get** /fabric/v4/crossConnects/{crossConnectId}/notes | Get Cross Connect Notes
[**GetCrossConnectSchedulesByUuid**](CrossConnectsApi.md#GetCrossConnectSchedulesByUuid) | **Get** /fabric/v4/crossConnects/{crossConnectId}/schedules | Get Cross Connect Schedules
[**GetCrossConnectServicesByUuid**](CrossConnectsApi.md#GetCrossConnectServicesByUuid) | **Get** /fabric/v4/crossConnects/{crossConnectId}/services | Get Cross Connect Services
[**PerformCrossConnectAction**](CrossConnectsApi.md#PerformCrossConnectAction) | **Post** /fabric/v4/crossConnects/{crossConnectId}/actions | Cross Connect Actions
[**SearchCrossConnectAction**](CrossConnectsApi.md#SearchCrossConnectAction) | **Post** /fabric/v4/crossConnects/{crossConnectId}/actions/search | Search Cross Connect Actions
[**SearchCrossConnectSchedules**](CrossConnectsApi.md#SearchCrossConnectSchedules) | **Post** /fabric/v4/crossConnectSchedules/search | Search Cross Connect Schedules
[**SearchCrossConnects**](CrossConnectsApi.md#SearchCrossConnects) | **Post** /fabric/v4/crossConnects/search | Search Cross Connects



## CreateCrossConnect

> CrossConnectResponse CreateCrossConnect(ctx).CreateCrossConnect(createCrossConnect).Execute()

Create Cross Connect



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
	createCrossConnect := *openapiclient.NewCreateCrossConnect(openapiclient.CrossConnectType("XC"), "My-CC-service", *openapiclient.NewCrossConnectASideRequest(), *openapiclient.NewCrossConnectZSideRequest()) // CreateCrossConnect | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrossConnectsApi.CreateCrossConnect(context.Background()).CreateCrossConnect(createCrossConnect).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrossConnectsApi.CreateCrossConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCrossConnect`: CrossConnectResponse
	fmt.Fprintf(os.Stdout, "Response from `CrossConnectsApi.CreateCrossConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCrossConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCrossConnect** | [**CreateCrossConnect**](CreateCrossConnect.md) |  | 

### Return type

[**CrossConnectResponse**](CrossConnectResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDiverseCrossConnect

> CrossConnectDiverse CreateDiverseCrossConnect(ctx).DiverseCrossConnectRequest(diverseCrossConnectRequest).Execute()

Create Diverse Cross Connect



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
	diverseCrossConnectRequest := *openapiclient.NewDiverseCrossConnectRequest([]openapiclient.CreateCrossConnect{*openapiclient.NewCreateCrossConnect(openapiclient.CrossConnectType("XC"), "My-CC-service", *openapiclient.NewCrossConnectASideRequest(), *openapiclient.NewCrossConnectZSideRequest())}) // DiverseCrossConnectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrossConnectsApi.CreateDiverseCrossConnect(context.Background()).DiverseCrossConnectRequest(diverseCrossConnectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrossConnectsApi.CreateDiverseCrossConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDiverseCrossConnect`: CrossConnectDiverse
	fmt.Fprintf(os.Stdout, "Response from `CrossConnectsApi.CreateDiverseCrossConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDiverseCrossConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diverseCrossConnectRequest** | [**DiverseCrossConnectRequest**](DiverseCrossConnectRequest.md) |  | 

### Return type

[**CrossConnectDiverse**](CrossConnectDiverse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrossConnectActionByUuid

> CrossConnectActionResponse GetCrossConnectActionByUuid(ctx, crossConnectId, actionId).Execute()

Get Cross Connect Action by Action ID



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
	crossConnectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cross Connect UUID.
	actionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cross Connect Action UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrossConnectsApi.GetCrossConnectActionByUuid(context.Background(), crossConnectId, actionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrossConnectsApi.GetCrossConnectActionByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrossConnectActionByUuid`: CrossConnectActionResponse
	fmt.Fprintf(os.Stdout, "Response from `CrossConnectsApi.GetCrossConnectActionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crossConnectId** | **string** | Cross Connect UUID. | 
**actionId** | **string** | Cross Connect Action UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrossConnectActionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CrossConnectActionResponse**](CrossConnectActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrossConnectAttachmentsByUuid

> CrossConnectAttachmentsResponse GetCrossConnectAttachmentsByUuid(ctx, crossConnectId).Execute()

Get Cross Connect Attachments



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
	crossConnectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cross Connect UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrossConnectsApi.GetCrossConnectAttachmentsByUuid(context.Background(), crossConnectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrossConnectsApi.GetCrossConnectAttachmentsByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrossConnectAttachmentsByUuid`: CrossConnectAttachmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `CrossConnectsApi.GetCrossConnectAttachmentsByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crossConnectId** | **string** | Cross Connect UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrossConnectAttachmentsByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CrossConnectAttachmentsResponse**](CrossConnectAttachmentsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrossConnectByUuid

> CrossConnectResponse GetCrossConnectByUuid(ctx, crossConnectId).Execute()

Get Cross Connect



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
	crossConnectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cross Connect UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrossConnectsApi.GetCrossConnectByUuid(context.Background(), crossConnectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrossConnectsApi.GetCrossConnectByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrossConnectByUuid`: CrossConnectResponse
	fmt.Fprintf(os.Stdout, "Response from `CrossConnectsApi.GetCrossConnectByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crossConnectId** | **string** | Cross Connect UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrossConnectByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CrossConnectResponse**](CrossConnectResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrossConnectNotesByUuid

> CrossConnectNotesResponse GetCrossConnectNotesByUuid(ctx, crossConnectId).Offset(offset).Limit(limit).Execute()

Get Cross Connect Notes



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
	crossConnectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cross Connect UUID.
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrossConnectsApi.GetCrossConnectNotesByUuid(context.Background(), crossConnectId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrossConnectsApi.GetCrossConnectNotesByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrossConnectNotesByUuid`: CrossConnectNotesResponse
	fmt.Fprintf(os.Stdout, "Response from `CrossConnectsApi.GetCrossConnectNotesByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crossConnectId** | **string** | Cross Connect UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrossConnectNotesByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**CrossConnectNotesResponse**](CrossConnectNotesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrossConnectSchedulesByUuid

> SchedulesResponse GetCrossConnectSchedulesByUuid(ctx, crossConnectId).Execute()

Get Cross Connect Schedules



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
	crossConnectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cross Connect UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrossConnectsApi.GetCrossConnectSchedulesByUuid(context.Background(), crossConnectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrossConnectsApi.GetCrossConnectSchedulesByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrossConnectSchedulesByUuid`: SchedulesResponse
	fmt.Fprintf(os.Stdout, "Response from `CrossConnectsApi.GetCrossConnectSchedulesByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crossConnectId** | **string** | Cross Connect UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrossConnectSchedulesByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SchedulesResponse**](SchedulesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrossConnectServicesByUuid

> CrossConnectServicesResponse GetCrossConnectServicesByUuid(ctx, crossConnectId).Execute()

Get Cross Connect Services



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
	crossConnectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cross Connect UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrossConnectsApi.GetCrossConnectServicesByUuid(context.Background(), crossConnectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrossConnectsApi.GetCrossConnectServicesByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrossConnectServicesByUuid`: CrossConnectServicesResponse
	fmt.Fprintf(os.Stdout, "Response from `CrossConnectsApi.GetCrossConnectServicesByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crossConnectId** | **string** | Cross Connect UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrossConnectServicesByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CrossConnectServicesResponse**](CrossConnectServicesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PerformCrossConnectAction

> CrossConnectActionResponse PerformCrossConnectAction(ctx, crossConnectId).CrossConnectActionRequest(crossConnectActionRequest).Execute()

Cross Connect Actions



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
	crossConnectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cross Connect UUID.
	crossConnectActionRequest := *openapiclient.NewCrossConnectActionRequest(openapiclient.CrossConnectActionType("XC_SIGNATURE_SUBMISSION")) // CrossConnectActionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrossConnectsApi.PerformCrossConnectAction(context.Background(), crossConnectId).CrossConnectActionRequest(crossConnectActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrossConnectsApi.PerformCrossConnectAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PerformCrossConnectAction`: CrossConnectActionResponse
	fmt.Fprintf(os.Stdout, "Response from `CrossConnectsApi.PerformCrossConnectAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crossConnectId** | **string** | Cross Connect UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPerformCrossConnectActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **crossConnectActionRequest** | [**CrossConnectActionRequest**](CrossConnectActionRequest.md) |  | 

### Return type

[**CrossConnectActionResponse**](CrossConnectActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCrossConnectAction

> CrossConnectActionSearchResponse SearchCrossConnectAction(ctx, crossConnectId).CrossConnectActionSearchRequest(crossConnectActionSearchRequest).Execute()

Search Cross Connect Actions



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
	crossConnectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cross Connect UUID.
	crossConnectActionSearchRequest := *openapiclient.NewCrossConnectActionSearchRequest() // CrossConnectActionSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrossConnectsApi.SearchCrossConnectAction(context.Background(), crossConnectId).CrossConnectActionSearchRequest(crossConnectActionSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrossConnectsApi.SearchCrossConnectAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCrossConnectAction`: CrossConnectActionSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `CrossConnectsApi.SearchCrossConnectAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**crossConnectId** | **string** | Cross Connect UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCrossConnectActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **crossConnectActionSearchRequest** | [**CrossConnectActionSearchRequest**](CrossConnectActionSearchRequest.md) |  | 

### Return type

[**CrossConnectActionSearchResponse**](CrossConnectActionSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCrossConnectSchedules

> SearchSchedulesResponse SearchCrossConnectSchedules(ctx).CrossConnectScheduleSearchRequest(crossConnectScheduleSearchRequest).Execute()

Search Cross Connect Schedules



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
	crossConnectScheduleSearchRequest := *openapiclient.NewCrossConnectScheduleSearchRequest() // CrossConnectScheduleSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrossConnectsApi.SearchCrossConnectSchedules(context.Background()).CrossConnectScheduleSearchRequest(crossConnectScheduleSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrossConnectsApi.SearchCrossConnectSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCrossConnectSchedules`: SearchSchedulesResponse
	fmt.Fprintf(os.Stdout, "Response from `CrossConnectsApi.SearchCrossConnectSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCrossConnectSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crossConnectScheduleSearchRequest** | [**CrossConnectScheduleSearchRequest**](CrossConnectScheduleSearchRequest.md) |  | 

### Return type

[**SearchSchedulesResponse**](SearchSchedulesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCrossConnects

> CrossConnectSearchResponse SearchCrossConnects(ctx).CrossConnectSearchRequest(crossConnectSearchRequest).Execute()

Search Cross Connects



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
	crossConnectSearchRequest := *openapiclient.NewCrossConnectSearchRequest() // CrossConnectSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrossConnectsApi.SearchCrossConnects(context.Background()).CrossConnectSearchRequest(crossConnectSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrossConnectsApi.SearchCrossConnects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCrossConnects`: CrossConnectSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `CrossConnectsApi.SearchCrossConnects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCrossConnectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crossConnectSearchRequest** | [**CrossConnectSearchRequest**](CrossConnectSearchRequest.md) |  | 

### Return type

[**CrossConnectSearchResponse**](CrossConnectSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

