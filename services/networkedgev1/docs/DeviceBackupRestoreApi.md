# \DeviceBackupRestoreApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckDetailsOfBackupsUsingGET**](DeviceBackupRestoreApi.md#CheckDetailsOfBackupsUsingGET) | **Get** /ne/v1/devices/{uuid}/restoreAnalysis | Checks Feasibility of Restore
[**CreateDeviceBackupUsingPOST**](DeviceBackupRestoreApi.md#CreateDeviceBackupUsingPOST) | **Post** /ne/v1/deviceBackups | Creates Device Backup
[**DeleteDeviceBackupUsingDELETE**](DeviceBackupRestoreApi.md#DeleteDeviceBackupUsingDELETE) | **Delete** /ne/v1/deviceBackups/{uuid} | Delete Backup of Device
[**DownloadDetailsOfBackupsUsingGET**](DeviceBackupRestoreApi.md#DownloadDetailsOfBackupsUsingGET) | **Get** /ne/v1/deviceBackups/{uuid}/download | Download a Backup
[**GetDetailsOfBackupsUsingGET**](DeviceBackupRestoreApi.md#GetDetailsOfBackupsUsingGET) | **Get** /ne/v1/deviceBackups/{uuid} | Get Backups of Device {uuid}
[**GetDeviceBackupsUsingGET**](DeviceBackupRestoreApi.md#GetDeviceBackupsUsingGET) | **Get** /ne/v1/deviceBackups | Get Backups of Device
[**RestoreDeviceBackupUsingPATCH**](DeviceBackupRestoreApi.md#RestoreDeviceBackupUsingPATCH) | **Patch** /ne/v1/devices/{uuid}/restore | Restores a backup
[**UpdateDeviceBackupUsingPATCH**](DeviceBackupRestoreApi.md#UpdateDeviceBackupUsingPATCH) | **Patch** /ne/v1/deviceBackups/{uuid} | Update Device Backup



## CheckDetailsOfBackupsUsingGET

> RestoreBackupInfoVerbose CheckDetailsOfBackupsUsingGET(ctx, uuid).BackupUuid(backupUuid).Authorization(authorization).Execute()

Checks Feasibility of Restore



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	uuid := "uuid_example" // string | Unique ID of a device
	backupUuid := "backupUuid_example" // string | Unique ID of a backup
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceBackupRestoreApi.CheckDetailsOfBackupsUsingGET(context.Background(), uuid).BackupUuid(backupUuid).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceBackupRestoreApi.CheckDetailsOfBackupsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckDetailsOfBackupsUsingGET`: RestoreBackupInfoVerbose
	fmt.Fprintf(os.Stdout, "Response from `DeviceBackupRestoreApi.CheckDetailsOfBackupsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique ID of a device | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckDetailsOfBackupsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backupUuid** | **string** | Unique ID of a backup | 
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 

### Return type

[**RestoreBackupInfoVerbose**](RestoreBackupInfoVerbose.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceBackupUsingPOST

> SshUserCreateResponse CreateDeviceBackupUsingPOST(ctx).Authorization(authorization).Request(request).Execute()

Creates Device Backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	request := *openapiclient.NewDeviceBackupCreateRequest("3da0a663-20d9-4b8f-8c5d-d5cf706840c8", "My New Backup") // DeviceBackupCreateRequest | Device backup info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceBackupRestoreApi.CreateDeviceBackupUsingPOST(context.Background()).Authorization(authorization).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceBackupRestoreApi.CreateDeviceBackupUsingPOST``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceBackupUsingPOST`: SshUserCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceBackupRestoreApi.CreateDeviceBackupUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceBackupUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **request** | [**DeviceBackupCreateRequest**](DeviceBackupCreateRequest.md) | Device backup info | 

### Return type

[**SshUserCreateResponse**](SshUserCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceBackupUsingDELETE

> DeleteDeviceBackupUsingDELETE(ctx, uuid).Authorization(authorization).Execute()

Delete Backup of Device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	uuid := "uuid_example" // string | Unique Id a backup
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceBackupRestoreApi.DeleteDeviceBackupUsingDELETE(context.Background(), uuid).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceBackupRestoreApi.DeleteDeviceBackupUsingDELETE``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique Id a backup | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceBackupUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadDetailsOfBackupsUsingGET

> string DownloadDetailsOfBackupsUsingGET(ctx, uuid).Authorization(authorization).Execute()

Download a Backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	uuid := "uuid_example" // string | Unique ID of a backup
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceBackupRestoreApi.DownloadDetailsOfBackupsUsingGET(context.Background(), uuid).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceBackupRestoreApi.DownloadDetailsOfBackupsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadDetailsOfBackupsUsingGET`: string
	fmt.Fprintf(os.Stdout, "Response from `DeviceBackupRestoreApi.DownloadDetailsOfBackupsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique ID of a backup | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDetailsOfBackupsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDetailsOfBackupsUsingGET

> DeviceBackupInfoVerbose GetDetailsOfBackupsUsingGET(ctx, uuid).Authorization(authorization).Execute()

Get Backups of Device {uuid}



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	uuid := "uuid_example" // string | Unique ID of a backup
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceBackupRestoreApi.GetDetailsOfBackupsUsingGET(context.Background(), uuid).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceBackupRestoreApi.GetDetailsOfBackupsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDetailsOfBackupsUsingGET`: DeviceBackupInfoVerbose
	fmt.Fprintf(os.Stdout, "Response from `DeviceBackupRestoreApi.GetDetailsOfBackupsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique ID of a backup | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDetailsOfBackupsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 

### Return type

[**DeviceBackupInfoVerbose**](DeviceBackupInfoVerbose.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceBackupsUsingGET

> DeviceBackupPageResponse GetDeviceBackupsUsingGET(ctx).VirtualDeviceUuid(virtualDeviceUuid).Authorization(authorization).Status(status).Offset(offset).Limit(limit).Execute()

Get Backups of Device



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	virtualDeviceUuid := "virtualDeviceUuid_example" // string | Unique ID of a virtual device
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	status := []string{"Inner_example"} // []string | An array of status values (optional)
	offset := "offset_example" // string | Specifies where to start a page. It is the starting point of the collection returned from the server. (optional) (default to "0")
	limit := "limit_example" // string | Specifies the page size. (optional) (default to "20")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceBackupRestoreApi.GetDeviceBackupsUsingGET(context.Background()).VirtualDeviceUuid(virtualDeviceUuid).Authorization(authorization).Status(status).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceBackupRestoreApi.GetDeviceBackupsUsingGET``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceBackupsUsingGET`: DeviceBackupPageResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceBackupRestoreApi.GetDeviceBackupsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceBackupsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **virtualDeviceUuid** | **string** | Unique ID of a virtual device | 
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **status** | **[]string** | An array of status values | 
 **offset** | **string** | Specifies where to start a page. It is the starting point of the collection returned from the server. | [default to &quot;0&quot;]
 **limit** | **string** | Specifies the page size. | [default to &quot;20&quot;]

### Return type

[**DeviceBackupPageResponse**](DeviceBackupPageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreDeviceBackupUsingPATCH

> RestoreDeviceBackupUsingPATCH(ctx, uuid).Authorization(authorization).Request(request).Execute()

Restores a backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	uuid := "uuid_example" // string | Unique ID of a backup
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	request := *openapiclient.NewDeviceBackupUpdateRequest("My New Backup") // DeviceBackupUpdateRequest | Update device backup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceBackupRestoreApi.RestoreDeviceBackupUsingPATCH(context.Background(), uuid).Authorization(authorization).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceBackupRestoreApi.RestoreDeviceBackupUsingPATCH``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique ID of a backup | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreDeviceBackupUsingPATCHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **request** | [**DeviceBackupUpdateRequest**](DeviceBackupUpdateRequest.md) | Update device backup | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceBackupUsingPATCH

> UpdateDeviceBackupUsingPATCH(ctx, uuid).Authorization(authorization).Request(request).Execute()

Update Device Backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/networkedgev1"
)

func main() {
	uuid := "uuid_example" // string | Unique ID of a backup
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	request := *openapiclient.NewDeviceBackupUpdateRequest("My New Backup") // DeviceBackupUpdateRequest | Update device backup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeviceBackupRestoreApi.UpdateDeviceBackupUsingPATCH(context.Background(), uuid).Authorization(authorization).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceBackupRestoreApi.UpdateDeviceBackupUsingPATCH``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Unique ID of a backup | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceBackupUsingPATCHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **request** | [**DeviceBackupUpdateRequest**](DeviceBackupUpdateRequest.md) | Update device backup | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

