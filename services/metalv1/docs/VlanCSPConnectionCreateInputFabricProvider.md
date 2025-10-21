# VlanCSPConnectionCreateInputFabricProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AWSFabricProviderType**](AWSFabricProviderType.md) |  | 
**AccountId** | **string** | AWS Account ID | 
**Location** | Pointer to **string** |  | [optional] 

## Methods

### NewVlanCSPConnectionCreateInputFabricProvider

`func NewVlanCSPConnectionCreateInputFabricProvider(type_ AWSFabricProviderType, accountId string, ) *VlanCSPConnectionCreateInputFabricProvider`

NewVlanCSPConnectionCreateInputFabricProvider instantiates a new VlanCSPConnectionCreateInputFabricProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanCSPConnectionCreateInputFabricProviderWithDefaults

`func NewVlanCSPConnectionCreateInputFabricProviderWithDefaults() *VlanCSPConnectionCreateInputFabricProvider`

NewVlanCSPConnectionCreateInputFabricProviderWithDefaults instantiates a new VlanCSPConnectionCreateInputFabricProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VlanCSPConnectionCreateInputFabricProvider) GetType() AWSFabricProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VlanCSPConnectionCreateInputFabricProvider) GetTypeOk() (*AWSFabricProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VlanCSPConnectionCreateInputFabricProvider) SetType(v AWSFabricProviderType)`

SetType sets Type field to given value.


### GetAccountId

`func (o *VlanCSPConnectionCreateInputFabricProvider) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VlanCSPConnectionCreateInputFabricProvider) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VlanCSPConnectionCreateInputFabricProvider) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetLocation

`func (o *VlanCSPConnectionCreateInputFabricProvider) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VlanCSPConnectionCreateInputFabricProvider) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VlanCSPConnectionCreateInputFabricProvider) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *VlanCSPConnectionCreateInputFabricProvider) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


