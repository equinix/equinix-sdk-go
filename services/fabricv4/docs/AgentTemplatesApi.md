# \AgentTemplatesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgentTemplateByUuid**](AgentTemplatesApi.md#GetAgentTemplateByUuid) | **Get** /fabric/v4/agentTemplates/{agentTemplateId} | Get Agent Template by UUID
[**GetAgentTemplates**](AgentTemplatesApi.md#GetAgentTemplates) | **Get** /fabric/v4/agentTemplates | Get Agent Templates



## GetAgentTemplateByUuid

> AgentTemplates GetAgentTemplateByUuid(ctx, agentTemplateId).Offset(offset).Limit(limit).Execute()

Get Agent Template by UUID



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
	agentTemplateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Agent Template UUID
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTemplatesApi.GetAgentTemplateByUuid(context.Background(), agentTemplateId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTemplatesApi.GetAgentTemplateByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentTemplateByUuid`: AgentTemplates
	fmt.Fprintf(os.Stdout, "Response from `AgentTemplatesApi.GetAgentTemplateByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentTemplateId** | **string** | Agent Template UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentTemplateByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**AgentTemplates**](AgentTemplates.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentTemplates

> AgentTemplateGetAllResponse GetAgentTemplates(ctx).Offset(offset).Limit(limit).Execute()

Get Agent Templates



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
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTemplatesApi.GetAgentTemplates(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTemplatesApi.GetAgentTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentTemplates`: AgentTemplateGetAllResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentTemplatesApi.GetAgentTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**AgentTemplateGetAllResponse**](AgentTemplateGetAllResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

