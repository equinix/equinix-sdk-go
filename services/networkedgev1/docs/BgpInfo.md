# BgpInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationKey** | Pointer to **string** |  | [optional] 
**ConnectionUuid** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedByEmail** | Pointer to **string** |  | [optional] 
**CreatedByFullName** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **string** |  | [optional] 
**DeletedBy** | Pointer to **string** |  | [optional] 
**DeletedByEmail** | Pointer to **string** |  | [optional] 
**DeletedByFullName** | Pointer to **string** |  | [optional] 
**DeletedDate** | Pointer to **string** |  | [optional] 
**LastUpdatedBy** | Pointer to **string** |  | [optional] 
**LastUpdatedByEmail** | Pointer to **string** |  | [optional] 
**LastUpdatedByFullName** | Pointer to **string** |  | [optional] 
**LastUpdatedDate** | Pointer to **string** |  | [optional] 
**LocalAsn** | Pointer to **int64** |  | [optional] 
**LocalIpAddress** | Pointer to **string** |  | [optional] 
**ProvisioningStatus** | Pointer to **string** |  | [optional] 
**RemoteAsn** | Pointer to **int64** |  | [optional] 
**RemoteIpAddress** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**VirtualDeviceUuid** | Pointer to **string** |  | [optional] 

## Methods

### NewBgpInfo

`func NewBgpInfo() *BgpInfo`

NewBgpInfo instantiates a new BgpInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpInfoWithDefaults

`func NewBgpInfoWithDefaults() *BgpInfo`

NewBgpInfoWithDefaults instantiates a new BgpInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationKey

`func (o *BgpInfo) GetAuthenticationKey() string`

GetAuthenticationKey returns the AuthenticationKey field if non-nil, zero value otherwise.

### GetAuthenticationKeyOk

`func (o *BgpInfo) GetAuthenticationKeyOk() (*string, bool)`

GetAuthenticationKeyOk returns a tuple with the AuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKey

`func (o *BgpInfo) SetAuthenticationKey(v string)`

SetAuthenticationKey sets AuthenticationKey field to given value.

### HasAuthenticationKey

`func (o *BgpInfo) HasAuthenticationKey() bool`

HasAuthenticationKey returns a boolean if a field has been set.

### GetConnectionUuid

`func (o *BgpInfo) GetConnectionUuid() string`

GetConnectionUuid returns the ConnectionUuid field if non-nil, zero value otherwise.

### GetConnectionUuidOk

`func (o *BgpInfo) GetConnectionUuidOk() (*string, bool)`

GetConnectionUuidOk returns a tuple with the ConnectionUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUuid

`func (o *BgpInfo) SetConnectionUuid(v string)`

SetConnectionUuid sets ConnectionUuid field to given value.

### HasConnectionUuid

`func (o *BgpInfo) HasConnectionUuid() bool`

HasConnectionUuid returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BgpInfo) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BgpInfo) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BgpInfo) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BgpInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByEmail

`func (o *BgpInfo) GetCreatedByEmail() string`

GetCreatedByEmail returns the CreatedByEmail field if non-nil, zero value otherwise.

### GetCreatedByEmailOk

`func (o *BgpInfo) GetCreatedByEmailOk() (*string, bool)`

GetCreatedByEmailOk returns a tuple with the CreatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByEmail

`func (o *BgpInfo) SetCreatedByEmail(v string)`

SetCreatedByEmail sets CreatedByEmail field to given value.

### HasCreatedByEmail

`func (o *BgpInfo) HasCreatedByEmail() bool`

HasCreatedByEmail returns a boolean if a field has been set.

### GetCreatedByFullName

`func (o *BgpInfo) GetCreatedByFullName() string`

GetCreatedByFullName returns the CreatedByFullName field if non-nil, zero value otherwise.

### GetCreatedByFullNameOk

`func (o *BgpInfo) GetCreatedByFullNameOk() (*string, bool)`

GetCreatedByFullNameOk returns a tuple with the CreatedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByFullName

`func (o *BgpInfo) SetCreatedByFullName(v string)`

SetCreatedByFullName sets CreatedByFullName field to given value.

### HasCreatedByFullName

`func (o *BgpInfo) HasCreatedByFullName() bool`

HasCreatedByFullName returns a boolean if a field has been set.

### GetCreatedDate

`func (o *BgpInfo) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *BgpInfo) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *BgpInfo) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *BgpInfo) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDeletedBy

`func (o *BgpInfo) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *BgpInfo) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *BgpInfo) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *BgpInfo) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedByEmail

`func (o *BgpInfo) GetDeletedByEmail() string`

GetDeletedByEmail returns the DeletedByEmail field if non-nil, zero value otherwise.

### GetDeletedByEmailOk

`func (o *BgpInfo) GetDeletedByEmailOk() (*string, bool)`

GetDeletedByEmailOk returns a tuple with the DeletedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByEmail

`func (o *BgpInfo) SetDeletedByEmail(v string)`

SetDeletedByEmail sets DeletedByEmail field to given value.

### HasDeletedByEmail

`func (o *BgpInfo) HasDeletedByEmail() bool`

HasDeletedByEmail returns a boolean if a field has been set.

### GetDeletedByFullName

`func (o *BgpInfo) GetDeletedByFullName() string`

GetDeletedByFullName returns the DeletedByFullName field if non-nil, zero value otherwise.

### GetDeletedByFullNameOk

`func (o *BgpInfo) GetDeletedByFullNameOk() (*string, bool)`

GetDeletedByFullNameOk returns a tuple with the DeletedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByFullName

`func (o *BgpInfo) SetDeletedByFullName(v string)`

SetDeletedByFullName sets DeletedByFullName field to given value.

### HasDeletedByFullName

`func (o *BgpInfo) HasDeletedByFullName() bool`

HasDeletedByFullName returns a boolean if a field has been set.

### GetDeletedDate

`func (o *BgpInfo) GetDeletedDate() string`

GetDeletedDate returns the DeletedDate field if non-nil, zero value otherwise.

### GetDeletedDateOk

`func (o *BgpInfo) GetDeletedDateOk() (*string, bool)`

GetDeletedDateOk returns a tuple with the DeletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDate

`func (o *BgpInfo) SetDeletedDate(v string)`

SetDeletedDate sets DeletedDate field to given value.

### HasDeletedDate

`func (o *BgpInfo) HasDeletedDate() bool`

HasDeletedDate returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *BgpInfo) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *BgpInfo) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *BgpInfo) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *BgpInfo) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetLastUpdatedByEmail

`func (o *BgpInfo) GetLastUpdatedByEmail() string`

GetLastUpdatedByEmail returns the LastUpdatedByEmail field if non-nil, zero value otherwise.

### GetLastUpdatedByEmailOk

`func (o *BgpInfo) GetLastUpdatedByEmailOk() (*string, bool)`

GetLastUpdatedByEmailOk returns a tuple with the LastUpdatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedByEmail

`func (o *BgpInfo) SetLastUpdatedByEmail(v string)`

SetLastUpdatedByEmail sets LastUpdatedByEmail field to given value.

### HasLastUpdatedByEmail

`func (o *BgpInfo) HasLastUpdatedByEmail() bool`

HasLastUpdatedByEmail returns a boolean if a field has been set.

### GetLastUpdatedByFullName

`func (o *BgpInfo) GetLastUpdatedByFullName() string`

GetLastUpdatedByFullName returns the LastUpdatedByFullName field if non-nil, zero value otherwise.

### GetLastUpdatedByFullNameOk

`func (o *BgpInfo) GetLastUpdatedByFullNameOk() (*string, bool)`

GetLastUpdatedByFullNameOk returns a tuple with the LastUpdatedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedByFullName

`func (o *BgpInfo) SetLastUpdatedByFullName(v string)`

SetLastUpdatedByFullName sets LastUpdatedByFullName field to given value.

### HasLastUpdatedByFullName

`func (o *BgpInfo) HasLastUpdatedByFullName() bool`

HasLastUpdatedByFullName returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *BgpInfo) GetLastUpdatedDate() string`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *BgpInfo) GetLastUpdatedDateOk() (*string, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *BgpInfo) SetLastUpdatedDate(v string)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *BgpInfo) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetLocalAsn

`func (o *BgpInfo) GetLocalAsn() int64`

GetLocalAsn returns the LocalAsn field if non-nil, zero value otherwise.

### GetLocalAsnOk

`func (o *BgpInfo) GetLocalAsnOk() (*int64, bool)`

GetLocalAsnOk returns a tuple with the LocalAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsn

`func (o *BgpInfo) SetLocalAsn(v int64)`

SetLocalAsn sets LocalAsn field to given value.

### HasLocalAsn

`func (o *BgpInfo) HasLocalAsn() bool`

HasLocalAsn returns a boolean if a field has been set.

### GetLocalIpAddress

`func (o *BgpInfo) GetLocalIpAddress() string`

GetLocalIpAddress returns the LocalIpAddress field if non-nil, zero value otherwise.

### GetLocalIpAddressOk

`func (o *BgpInfo) GetLocalIpAddressOk() (*string, bool)`

GetLocalIpAddressOk returns a tuple with the LocalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIpAddress

`func (o *BgpInfo) SetLocalIpAddress(v string)`

SetLocalIpAddress sets LocalIpAddress field to given value.

### HasLocalIpAddress

`func (o *BgpInfo) HasLocalIpAddress() bool`

HasLocalIpAddress returns a boolean if a field has been set.

### GetProvisioningStatus

`func (o *BgpInfo) GetProvisioningStatus() string`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BgpInfo) GetProvisioningStatusOk() (*string, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BgpInfo) SetProvisioningStatus(v string)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BgpInfo) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### GetRemoteAsn

`func (o *BgpInfo) GetRemoteAsn() int64`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *BgpInfo) GetRemoteAsnOk() (*int64, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *BgpInfo) SetRemoteAsn(v int64)`

SetRemoteAsn sets RemoteAsn field to given value.

### HasRemoteAsn

`func (o *BgpInfo) HasRemoteAsn() bool`

HasRemoteAsn returns a boolean if a field has been set.

### GetRemoteIpAddress

`func (o *BgpInfo) GetRemoteIpAddress() string`

GetRemoteIpAddress returns the RemoteIpAddress field if non-nil, zero value otherwise.

### GetRemoteIpAddressOk

`func (o *BgpInfo) GetRemoteIpAddressOk() (*string, bool)`

GetRemoteIpAddressOk returns a tuple with the RemoteIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpAddress

`func (o *BgpInfo) SetRemoteIpAddress(v string)`

SetRemoteIpAddress sets RemoteIpAddress field to given value.

### HasRemoteIpAddress

`func (o *BgpInfo) HasRemoteIpAddress() bool`

HasRemoteIpAddress returns a boolean if a field has been set.

### GetState

`func (o *BgpInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BgpInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BgpInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BgpInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *BgpInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *BgpInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *BgpInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *BgpInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVirtualDeviceUuid

`func (o *BgpInfo) GetVirtualDeviceUuid() string`

GetVirtualDeviceUuid returns the VirtualDeviceUuid field if non-nil, zero value otherwise.

### GetVirtualDeviceUuidOk

`func (o *BgpInfo) GetVirtualDeviceUuidOk() (*string, bool)`

GetVirtualDeviceUuidOk returns a tuple with the VirtualDeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceUuid

`func (o *BgpInfo) SetVirtualDeviceUuid(v string)`

SetVirtualDeviceUuid sets VirtualDeviceUuid field to given value.

### HasVirtualDeviceUuid

`func (o *BgpInfo) HasVirtualDeviceUuid() bool`

HasVirtualDeviceUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


