# \ClientInterfacesApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTerraformTemplates**](ClientInterfacesApi.md#CreateTerraformTemplates) | **Post** /fabric/v4/deployments/{deploymentId}/download | Generate Terraform Deployment Templates



## CreateTerraformTemplates

> *os.File CreateTerraformTemplates(ctx, deploymentId).ClientInterfaces(clientInterfaces).Execute()

Generate Terraform Deployment Templates



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
	clientInterfaces := *openapiclient.NewClientInterfaces("TERRAFORM_DOWNLOAD", "terraform-deployment-download", "Deployment Terraform Templates in DC") // ClientInterfaces | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientInterfacesApi.CreateTerraformTemplates(context.Background(), deploymentId).ClientInterfaces(clientInterfaces).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInterfacesApi.CreateTerraformTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTerraformTemplates`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ClientInterfacesApi.CreateTerraformTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** | Deployment UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTerraformTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientInterfaces** | [**ClientInterfaces**](ClientInterfaces.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

