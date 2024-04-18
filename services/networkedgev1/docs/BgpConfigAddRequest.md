# BgpConfigAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationKey** | Pointer to **string** | Provide a key value that you can use later to authenticate. | [optional] 
**ConnectionUuid** | Pointer to **string** | The unique Id of a connection between the virtual device and the cloud service provider | [optional] 
**LocalAsn** | Pointer to **int64** | Local ASN (autonomous system network). This is the ASN of the virtual device. | [optional] 
**LocalIpAddress** | Pointer to **string** | Local IP Address. This is the IP address of the virtual device in CIDR format. | [optional] 
**RemoteAsn** | Pointer to **int64** | Remote ASN (autonomous system network). This is the ASN of the cloud service provider. | [optional] 
**RemoteIpAddress** | Pointer to **string** | Remote IP Address. This is the IP address of the cloud service provider. | [optional] 

## Methods

### NewBgpConfigAddRequest

`func NewBgpConfigAddRequest() *BgpConfigAddRequest`

NewBgpConfigAddRequest instantiates a new BgpConfigAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpConfigAddRequestWithDefaults

`func NewBgpConfigAddRequestWithDefaults() *BgpConfigAddRequest`

NewBgpConfigAddRequestWithDefaults instantiates a new BgpConfigAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationKey

`func (o *BgpConfigAddRequest) GetAuthenticationKey() string`

GetAuthenticationKey returns the AuthenticationKey field if non-nil, zero value otherwise.

### GetAuthenticationKeyOk

`func (o *BgpConfigAddRequest) GetAuthenticationKeyOk() (*string, bool)`

GetAuthenticationKeyOk returns a tuple with the AuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKey

`func (o *BgpConfigAddRequest) SetAuthenticationKey(v string)`

SetAuthenticationKey sets AuthenticationKey field to given value.

### HasAuthenticationKey

`func (o *BgpConfigAddRequest) HasAuthenticationKey() bool`

HasAuthenticationKey returns a boolean if a field has been set.

### GetConnectionUuid

`func (o *BgpConfigAddRequest) GetConnectionUuid() string`

GetConnectionUuid returns the ConnectionUuid field if non-nil, zero value otherwise.

### GetConnectionUuidOk

`func (o *BgpConfigAddRequest) GetConnectionUuidOk() (*string, bool)`

GetConnectionUuidOk returns a tuple with the ConnectionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUuid

`func (o *BgpConfigAddRequest) SetConnectionUuid(v string)`

SetConnectionUuid sets ConnectionUuid field to given value.

### HasConnectionUuid

`func (o *BgpConfigAddRequest) HasConnectionUuid() bool`

HasConnectionUuid returns a boolean if a field has been set.

### GetLocalAsn

`func (o *BgpConfigAddRequest) GetLocalAsn() int64`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *BgpConfigAddRequest) GetLocalAsnOk() (*int64, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *BgpConfigAddRequest) SetLocalAsn(v int64)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *BgpConfigAddRequest) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetLocalIpAddress

`func (o *BgpConfigAddRequest) GetLocalIpAddress() string`

GetLocalIpAddress returns the LocalIpAddress field if non-nil, zero value otherwise.

### GetLocalIpAddressOk

`func (o *BgpConfigAddRequest) GetLocalIpAddressOk() (*string, bool)`

GetLocalIpAddressOk returns a tuple with the LocalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIpAddress

`func (o *BgpConfigAddRequest) SetLocalIpAddress(v string)`

SetLocalIpAddress sets LocalIpAddress field to given value.

### HasLocalIpAddress

`func (o *BgpConfigAddRequest) HasLocalIpAddress() bool`

HasLocalIpAddress returns a boolean if a field has been set.

### GetRemoteAsn

`func (o *BgpConfigAddRequest) GetRemoteAsn() int64`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *BgpConfigAddRequest) GetRemoteAsnOk() (*int64, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *BgpConfigAddRequest) SetRemoteAsn(v int64)`

SetRemoteAsn sets RemoteAsn field to given value.

### HasRemoteAsn

`func (o *BgpConfigAddRequest) HasRemoteAsn() bool`

HasRemoteAsn returns a boolean if a field has been set.

### GetRemoteIpAddress

`func (o *BgpConfigAddRequest) GetRemoteIpAddress() string`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *BgpConfigAddRequest) GetRemoteIpAddressOk() (*string, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *BgpConfigAddRequest) SetRemoteIpAddress(v string)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.

### HasRemoteIpAddress

`func (o *BgpConfigAddRequest) HasRemoteIpAddress() bool`

HasRemoteIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


