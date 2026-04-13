# \AgentsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgent**](AgentsApi.md#CreateAgent) | **Post** /fabric/v4/agents | Create Agent
[**DeleteAgentByUuid**](AgentsApi.md#DeleteAgentByUuid) | **Delete** /fabric/v4/agents/{agentId} | Delete Agent by UUID
[**GetAgentActivities**](AgentsApi.md#GetAgentActivities) | **Get** /fabric/v4/agents/{agentId}/activities | Get Agent Activities
[**GetAgentByUuid**](AgentsApi.md#GetAgentByUuid) | **Get** /fabric/v4/agents/{agentId} | Get Agent by UUID
[**GetAgents**](AgentsApi.md#GetAgents) | **Get** /fabric/v4/agents | Get Agents
[**PatchAgentByUuid**](AgentsApi.md#PatchAgentByUuid) | **Patch** /fabric/v4/agents/{agentId} | Update Agent by UUID



## CreateAgent

> Agents CreateAgent(ctx).AgentPostRequest(agentPostRequest).Execute()

Create Agent



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
	agentPostRequest := *openapiclient.NewAgentPostRequest("ANO_AGENT", "Name_example", *openapiclient.NewProject("44f4c4f8-2f39-494e-838c-d8e640591be5"), *openapiclient.NewAgentTemplate()) // AgentPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsApi.CreateAgent(context.Background()).AgentPostRequest(agentPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.CreateAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAgent`: Agents
	fmt.Fprintf(os.Stdout, "Response from `AgentsApi.CreateAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentPostRequest** | [**AgentPostRequest**](AgentPostRequest.md) |  | 

### Return type

[**Agents**](Agents.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgentByUuid

> Agents DeleteAgentByUuid(ctx, agentId).Offset(offset).Limit(limit).Execute()

Delete Agent by UUID



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Agent UUID
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsApi.DeleteAgentByUuid(context.Background(), agentId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.DeleteAgentByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAgentByUuid`: Agents
	fmt.Fprintf(os.Stdout, "Response from `AgentsApi.DeleteAgentByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**Agents**](Agents.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentActivities

> AgentGetActivities GetAgentActivities(ctx, agentId).Offset(offset).Limit(limit).Execute()

Get Agent Activities



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Agent UUID
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsApi.GetAgentActivities(context.Background(), agentId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.GetAgentActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentActivities`: AgentGetActivities
	fmt.Fprintf(os.Stdout, "Response from `AgentsApi.GetAgentActivities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**AgentGetActivities**](AgentGetActivities.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentByUuid

> Agents GetAgentByUuid(ctx, agentId).Offset(offset).Limit(limit).Execute()

Get Agent by UUID



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Agent UUID
	offset := int32(1) // int32 | offset (optional)
	limit := int32(10) // int32 | number of records to fetch (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsApi.GetAgentByUuid(context.Background(), agentId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.GetAgentByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentByUuid`: Agents
	fmt.Fprintf(os.Stdout, "Response from `AgentsApi.GetAgentByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**Agents**](Agents.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgents

> AgentGetAllResponse GetAgents(ctx).Offset(offset).Limit(limit).Execute()

Get Agents



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
	resp, r, err := apiClient.AgentsApi.GetAgents(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.GetAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgents`: AgentGetAllResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentsApi.GetAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | offset | 
 **limit** | **int32** | number of records to fetch | 

### Return type

[**AgentGetAllResponse**](AgentGetAllResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAgentByUuid

> Agents PatchAgentByUuid(ctx, agentId).AgentPatchRequest(agentPatchRequest).Execute()

Update Agent by UUID



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Agent UUID
	agentPatchRequest := *openapiclient.NewAgentPatchRequest("/name", "replace", interface{}(123)) // AgentPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsApi.PatchAgentByUuid(context.Background(), agentId).AgentPatchRequest(agentPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.PatchAgentByUuid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAgentByUuid`: Agents
	fmt.Fprintf(os.Stdout, "Response from `AgentsApi.PatchAgentByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAgentByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentPatchRequest** | [**AgentPatchRequest**](AgentPatchRequest.md) |  | 

### Return type

[**Agents**](Agents.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

