# SrvListClientRegistrationsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageToken** | Pointer to **string** | When paging through results, the PageToken is an opaque indicator that identifies a page. | [optional] 
**PageSize** | **int32** | When paging through results, this is an integer indicating the maximum number of results to return. Note, it is   possible for fewer results to be returned even when the end of the result set has not been reached. | 
**ServiceId** | Pointer to **string** | Fully qualified, universally unique id of a service. Starts with the NamespaceId. Formatted like   \&quot;service:&amp;lt;namespace&amp;gt;/&amp;lt;service&amp;gt;\&quot;. | [optional] 
**ServiceErn** | Pointer to **string** | Equinix resource name, a universally unique identifier for a resource across all clouds, regions, and   services. Formatted as   \&quot;ern:&amp;lt;cloudId&amp;gt;:&amp;lt;serviceId&amp;gt;:&amp;lt;regionId&amp;gt;:&amp;lt;projectId&amp;gt;:&amp;lt;resourceType&amp;gt;:&amp;lt;resourceId&amp;gt;\&quot;. | [optional] 

## Methods

### NewSrvListClientRegistrationsPostRequest

`func NewSrvListClientRegistrationsPostRequest(pageSize int32, ) *SrvListClientRegistrationsPostRequest`

NewSrvListClientRegistrationsPostRequest instantiates a new SrvListClientRegistrationsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrvListClientRegistrationsPostRequestWithDefaults

`func NewSrvListClientRegistrationsPostRequestWithDefaults() *SrvListClientRegistrationsPostRequest`

NewSrvListClientRegistrationsPostRequestWithDefaults instantiates a new SrvListClientRegistrationsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageToken

`func (o *SrvListClientRegistrationsPostRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *SrvListClientRegistrationsPostRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *SrvListClientRegistrationsPostRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *SrvListClientRegistrationsPostRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.

### GetPageSize

`func (o *SrvListClientRegistrationsPostRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *SrvListClientRegistrationsPostRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *SrvListClientRegistrationsPostRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetServiceId

`func (o *SrvListClientRegistrationsPostRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SrvListClientRegistrationsPostRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SrvListClientRegistrationsPostRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *SrvListClientRegistrationsPostRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceErn

`func (o *SrvListClientRegistrationsPostRequest) GetServiceErn() string`

GetServiceErn returns the ServiceErn field if non-nil, zero value otherwise.

### GetServiceErnOk

`func (o *SrvListClientRegistrationsPostRequest) GetServiceErnOk() (*string, bool)`

GetServiceErnOk returns a tuple with the ServiceErn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceErn

`func (o *SrvListClientRegistrationsPostRequest) SetServiceErn(v string)`

SetServiceErn sets ServiceErn field to given value.

### HasServiceErn

`func (o *SrvListClientRegistrationsPostRequest) HasServiceErn() bool`

HasServiceErn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


