# \DeploymentsApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionDeployment**](DeploymentsApi.md#ActionDeployment) | **Post** /fabric/v4/deployments/{deploymentId}/actions | Deploy, Dry Run or Destroy Deployment
[**CreateTopologyDeployment**](DeploymentsApi.md#CreateTopologyDeployment) | **Post** /fabric/v4/deployments | Create a new topology deployment
[**DeleteDeployment**](DeploymentsApi.md#DeleteDeployment) | **Delete** /fabric/v4/deployments/{deploymentId} | Delete Deployment using UUID
[**GetDeployment**](DeploymentsApi.md#GetDeployment) | **Get** /fabric/v4/deployments/{deploymentId} | Retrieve Deployment details using UUID
[**SearchDeployments**](DeploymentsApi.md#SearchDeployments) | **Post** /fabric/v4/deployments/searchDeployments | Search deployments
[**SearchProviderResources**](DeploymentsApi.md#SearchProviderResources) | **Post** /fabric/v4/providerResources/search | Search provider resources



## ActionDeployment

> CloudRouterActionResponse ActionDeployment(ctx, deploymentId).DeploymentActionRequest(deploymentActionRequest).Execute()

Deploy, Dry Run or Destroy Deployment



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment UUID
	deploymentActionRequest := *openapiclient.NewDeploymentActionRequest(openapiclient.DeploymentActionType("DEPLOY"), []openapiclient.ActionRequest{openapiclient.ActionRequest{AWSPermission: openapiclient.NewAWSPermission(openapiclient.AWSPermission_type("PERMISSION"), "arn:aws:iam:::role/", "us-east-1", *openapiclient.NewTopologyProperties("13fc5d30-036d-11f0-96a3-5336ee83f43f"))}}) // DeploymentActionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsApi.ActionDeployment(context.Background(), deploymentId).DeploymentActionRequest(deploymentActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.ActionDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionDeployment`: CloudRouterActionResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.ActionDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** | Deployment UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentActionRequest** | [**DeploymentActionRequest**](DeploymentActionRequest.md) |  | 

### Return type

[**CloudRouterActionResponse**](CloudRouterActionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTopologyDeployment

> DeploymentResponse CreateTopologyDeployment(ctx).Deployment(deployment).Execute()

Create a new topology deployment



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
	deployment := *openapiclient.NewDeployment("TOPOLOGY_DEPLOYMENT", "MCN-Deployment", []openapiclient.OrchestratorProviders{openapiclient.OrchestratorProviders{AWSProvider: openapiclient.NewAWSProvider(openapiclient.AWSProvider_type("AWS_PROVIDER"), []openapiclient.AWSProviderResource{openapiclient.AWSProviderResource{AWSDirectConnect: openapiclient.NewAWSDirectConnect(openapiclient.SearchDirectConnect_type("DIRECT_CONNECT"), *openapiclient.NewTopologyProperties("13fc5d30-036d-11f0-96a3-5336ee83f43f"))}})}}, *openapiclient.NewProject("44f4c4f8-2f39-494e-838c-d8e640591be5"), *openapiclient.NewSimplifiedAccount()) // Deployment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsApi.CreateTopologyDeployment(context.Background()).Deployment(deployment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.CreateTopologyDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTopologyDeployment`: DeploymentResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.CreateTopologyDeployment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTopologyDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment** | [**Deployment**](Deployment.md) |  | 

### Return type

[**DeploymentResponse**](DeploymentResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeployment

> DeploymentResponse DeleteDeployment(ctx, deploymentId).Execute()

Delete Deployment using UUID



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsApi.DeleteDeployment(context.Background(), deploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.DeleteDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDeployment`: DeploymentResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.DeleteDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** | Deployment UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeploymentResponse**](DeploymentResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployment

> DeploymentResponse GetDeployment(ctx, deploymentId).Execute()

Retrieve Deployment details using UUID



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
	deploymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Deployment UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsApi.GetDeployment(context.Background(), deploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.GetDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployment`: DeploymentResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.GetDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** | Deployment UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeploymentResponse**](DeploymentResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDeployments

> DeploymentSearchResponse SearchDeployments(ctx).DeploymentSearchRequest(deploymentSearchRequest).Execute()

Search deployments



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
	deploymentSearchRequest := *openapiclient.NewDeploymentSearchRequest() // DeploymentSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsApi.SearchDeployments(context.Background()).DeploymentSearchRequest(deploymentSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.SearchDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDeployments`: DeploymentSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.SearchDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentSearchRequest** | [**DeploymentSearchRequest**](DeploymentSearchRequest.md) |  | 

### Return type

[**DeploymentSearchResponse**](DeploymentSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchProviderResources

> ProviderSearchResponse SearchProviderResources(ctx).ProviderSearchRequest(providerSearchRequest).Execute()

Search provider resources



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
	providerSearchRequest := *openapiclient.NewProviderSearchRequest() // ProviderSearchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsApi.SearchProviderResources(context.Background()).ProviderSearchRequest(providerSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsApi.SearchProviderResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchProviderResources`: ProviderSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsApi.SearchProviderResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProviderResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerSearchRequest** | [**ProviderSearchRequest**](ProviderSearchRequest.md) |  | 

### Return type

[**ProviderSearchResponse**](ProviderSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

