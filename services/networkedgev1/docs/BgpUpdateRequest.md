# BgpUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationKey** | Pointer to **string** | Authentication Key | [optional] 
**LocalAsn** | Pointer to **int64** | Local ASN | [optional] 
**LocalIpAddress** | Pointer to **string** | Local IP Address with subnet | [optional] 
**RemoteAsn** | Pointer to **int64** | Remote ASN | [optional] 
**RemoteIpAddress** | Pointer to **string** | Remote IP Address | [optional] 

## Methods

### NewBgpUpdateRequest

`func NewBgpUpdateRequest() *BgpUpdateRequest`

NewBgpUpdateRequest instantiates a new BgpUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpUpdateRequestWithDefaults

`func NewBgpUpdateRequestWithDefaults() *BgpUpdateRequest`

NewBgpUpdateRequestWithDefaults instantiates a new BgpUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationKey

`func (o *BgpUpdateRequest) GetAuthenticationKey() string`

GetAuthenticationKey returns the AuthenticationKey field if non-nil, zero value otherwise.

### GetAuthenticationKeyOk

`func (o *BgpUpdateRequest) GetAuthenticationKeyOk() (*string, bool)`

GetAuthenticationKeyOk returns a tuple with the AuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKey

`func (o *BgpUpdateRequest) SetAuthenticationKey(v string)`

SetAuthenticationKey sets AuthenticationKey field to given value.

### HasAuthenticationKey

`func (o *BgpUpdateRequest) HasAuthenticationKey() bool`

HasAuthenticationKey returns a boolean if a field has been set.

### GetLocalAsn

`func (o *BgpUpdateRequest) GetLocalAsn() int64`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *BgpUpdateRequest) GetLocalAsnOk() (*int64, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *BgpUpdateRequest) SetLocalAsn(v int64)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *BgpUpdateRequest) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetLocalIpAddress

`func (o *BgpUpdateRequest) GetLocalIpAddress() string`

GetLocalIpAddress returns the LocalIpAddress field if non-nil, zero value otherwise.

### GetLocalIpAddressOk

`func (o *BgpUpdateRequest) GetLocalIpAddressOk() (*string, bool)`

GetLocalIpAddressOk returns a tuple with the LocalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIpAddress

`func (o *BgpUpdateRequest) SetLocalIpAddress(v string)`

SetLocalIpAddress sets LocalIpAddress field to given value.

### HasLocalIpAddress

`func (o *BgpUpdateRequest) HasLocalIpAddress() bool`

HasLocalIpAddress returns a boolean if a field has been set.

### GetRemoteAsn

`func (o *BgpUpdateRequest) GetRemoteAsn() int64`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *BgpUpdateRequest) GetRemoteAsnOk() (*int64, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *BgpUpdateRequest) SetRemoteAsn(v int64)`

SetRemoteAsn sets RemoteAsn field to given value.

### HasRemoteAsn

`func (o *BgpUpdateRequest) HasRemoteAsn() bool`

HasRemoteAsn returns a boolean if a field has been set.

### GetRemoteIpAddress

`func (o *BgpUpdateRequest) GetRemoteIpAddress() string`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *BgpUpdateRequest) GetRemoteIpAddressOk() (*string, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *BgpUpdateRequest) SetRemoteIpAddress(v string)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.

### HasRemoteIpAddress

`func (o *BgpUpdateRequest) HasRemoteIpAddress() bool`

HasRemoteIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


