# \AlertLogApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlertActivityLog**](AlertLogApi.md#GetAlertActivityLog) | **Get** /smartview/v1/alerts/getAlertActivityLog | get alert activity logs



## GetAlertActivityLog

> AlertActivityLog GetAlertActivityLog(ctx).Authorization(authorization).PageNum(pageNum).Limit(limit).Ibx(ibx).Category(category).EventType(eventType).IsAcknowledged(isAcknowledged).AccountNo(accountNo).OrderBy(orderBy).SortBy(sortBy).Execute()

get alert activity logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/smartview"
)

func main() {
	authorization := "authorization_example" // string | The OAuth Bearer token. Please add the prefix 'Bearer ' before the token.
	pageNum := "pageNum_example" // string | Page Number
	limit := "limit_example" // string | Limit
	ibx := "ibx_example" // string | IBX (optional)
	category := "category_example" // string | Category (optional)
	eventType := "eventType_example" // string | Event Type (optional)
	isAcknowledged := "isAcknowledged_example" // string | Is Acknowledged (optional)
	accountNo := "accountNo_example" // string | Account Number (optional)
	orderBy := "orderBy_example" // string | order by ascending or descending  (optional)
	sortBy := "sortBy_example" // string | sortBy value  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertLogApi.GetAlertActivityLog(context.Background()).Authorization(authorization).PageNum(pageNum).Limit(limit).Ibx(ibx).Category(category).EventType(eventType).IsAcknowledged(isAcknowledged).AccountNo(accountNo).OrderBy(orderBy).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertLogApi.GetAlertActivityLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertActivityLog`: AlertActivityLog
	fmt.Fprintf(os.Stdout, "Response from `AlertLogApi.GetAlertActivityLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertActivityLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | The OAuth Bearer token. Please add the prefix &#39;Bearer &#39; before the token. | 
 **pageNum** | **string** | Page Number | 
 **limit** | **string** | Limit | 
 **ibx** | **string** | IBX | 
 **category** | **string** | Category | 
 **eventType** | **string** | Event Type | 
 **isAcknowledged** | **string** | Is Acknowledged | 
 **accountNo** | **string** | Account Number | 
 **orderBy** | **string** | order by ascending or descending  | 
 **sortBy** | **string** | sortBy value  | 

### Return type

[**AlertActivityLog**](AlertActivityLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

