# ServiceProfileActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Service Profile Action URI | [optional] [readonly] 
**Type** | Pointer to **string** | Action type. Example values: PROFILE_UPDATE_ACCEPTANCE, PROFILE_UPDATE_REJECTION | [optional] 
**Uuid** | Pointer to **string** | Equinix-assigned action identifier | [optional] 
**Comments** | Pointer to **string** | Action comments | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewServiceProfileActionResponse

`func NewServiceProfileActionResponse() *ServiceProfileActionResponse`

NewServiceProfileActionResponse instantiates a new ServiceProfileActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileActionResponseWithDefaults

`func NewServiceProfileActionResponseWithDefaults() *ServiceProfileActionResponse`

NewServiceProfileActionResponseWithDefaults instantiates a new ServiceProfileActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ServiceProfileActionResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ServiceProfileActionResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ServiceProfileActionResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ServiceProfileActionResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *ServiceProfileActionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceProfileActionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceProfileActionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceProfileActionResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *ServiceProfileActionResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceProfileActionResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceProfileActionResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServiceProfileActionResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetComments

`func (o *ServiceProfileActionResponse) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ServiceProfileActionResponse) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ServiceProfileActionResponse) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ServiceProfileActionResponse) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetChangeLog

`func (o *ServiceProfileActionResponse) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *ServiceProfileActionResponse) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *ServiceProfileActionResponse) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *ServiceProfileActionResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


