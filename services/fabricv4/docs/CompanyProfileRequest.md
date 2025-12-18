# CompanyProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** |  | 
**Summary** | **string** |  | 
**Description** | **string** |  | 
**Notifications** | Pointer to **[]map[string]interface{}** |  | [optional] 
**WebUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewCompanyProfileRequest

`func NewCompanyProfileRequest(type_ string, name string, summary string, description string, ) *CompanyProfileRequest`

NewCompanyProfileRequest instantiates a new CompanyProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyProfileRequestWithDefaults

`func NewCompanyProfileRequestWithDefaults() *CompanyProfileRequest`

NewCompanyProfileRequestWithDefaults instantiates a new CompanyProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CompanyProfileRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompanyProfileRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompanyProfileRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CompanyProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSummary

`func (o *CompanyProfileRequest) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CompanyProfileRequest) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CompanyProfileRequest) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetDescription

`func (o *CompanyProfileRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CompanyProfileRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CompanyProfileRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNotifications

`func (o *CompanyProfileRequest) GetNotifications() []map[string]interface{}`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *CompanyProfileRequest) GetNotificationsOk() (*[]map[string]interface{}, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *CompanyProfileRequest) SetNotifications(v []map[string]interface{})`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *CompanyProfileRequest) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetWebUrl

`func (o *CompanyProfileRequest) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *CompanyProfileRequest) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *CompanyProfileRequest) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *CompanyProfileRequest) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


