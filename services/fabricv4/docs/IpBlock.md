# IpBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Unique identifier for the Ip Block | 
**Href** | **string** | Resource URL path for the Ip Block | 
**Type** | [**TypeOfIpBlockProduct**](TypeOfIpBlockProduct.md) |  | 
**State** | [**IpBlockState**](IpBlockState.md) |  | 
**Ownership** | [**IpBlockOwnership**](IpBlockOwnership.md) |  | 
**Location** | Pointer to [**IpBlockLocation**](IpBlockLocation.md) |  | [optional] 
**PrefixLength** | **int32** | IpBlockPrefix length | 
**Prefix** | Pointer to **string** | CIDR prefix | [optional] 
**Order** | Pointer to [**IpBlockOrderResponse**](IpBlockOrderResponse.md) |  | [optional] 
**Account** | Pointer to [**IpBlockAccount**](IpBlockAccount.md) |  | [optional] 
**Project** | [**IpBlockProject**](IpBlockProject.md) |  | 
**Regulations** | Pointer to [**IpBlockRegulations**](IpBlockRegulations.md) |  | [optional] 
**Assets** | Pointer to [**[]IpBlockAsset**](IpBlockAsset.md) | Products using this Ip Block | [optional] 
**Change** | Pointer to [**IpBlockChange**](IpBlockChange.md) |  | [optional] 
**ChangeLog** | [**IpBlockChangeLog**](IpBlockChangeLog.md) |  | 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewIpBlock

`func NewIpBlock(uuid string, href string, type_ TypeOfIpBlockProduct, state IpBlockState, ownership IpBlockOwnership, prefixLength int32, project IpBlockProject, changeLog IpBlockChangeLog, ) *IpBlock`

NewIpBlock instantiates a new IpBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlockWithDefaults

`func NewIpBlockWithDefaults() *IpBlock`

NewIpBlockWithDefaults instantiates a new IpBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *IpBlock) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *IpBlock) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *IpBlock) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetHref

`func (o *IpBlock) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IpBlock) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IpBlock) SetHref(v string)`

SetHref sets Href field to given value.


### GetType

`func (o *IpBlock) GetType() TypeOfIpBlockProduct`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpBlock) GetTypeOk() (*TypeOfIpBlockProduct, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpBlock) SetType(v TypeOfIpBlockProduct)`

SetType sets Type field to given value.


### GetState

`func (o *IpBlock) GetState() IpBlockState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IpBlock) GetStateOk() (*IpBlockState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IpBlock) SetState(v IpBlockState)`

SetState sets State field to given value.


### GetOwnership

`func (o *IpBlock) GetOwnership() IpBlockOwnership`

GetOwnership returns the Ownership field if non-nil, zero value otherwise.

### GetOwnershipOk

`func (o *IpBlock) GetOwnershipOk() (*IpBlockOwnership, bool)`

GetOwnershipOk returns a tuple with the Ownership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnership

`func (o *IpBlock) SetOwnership(v IpBlockOwnership)`

SetOwnership sets Ownership field to given value.


### GetLocation

`func (o *IpBlock) GetLocation() IpBlockLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IpBlock) GetLocationOk() (*IpBlockLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IpBlock) SetLocation(v IpBlockLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *IpBlock) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPrefixLength

`func (o *IpBlock) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *IpBlock) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *IpBlock) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.


### GetPrefix

`func (o *IpBlock) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *IpBlock) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *IpBlock) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *IpBlock) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetOrder

`func (o *IpBlock) GetOrder() IpBlockOrderResponse`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *IpBlock) GetOrderOk() (*IpBlockOrderResponse, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *IpBlock) SetOrder(v IpBlockOrderResponse)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *IpBlock) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetAccount

`func (o *IpBlock) GetAccount() IpBlockAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IpBlock) GetAccountOk() (*IpBlockAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IpBlock) SetAccount(v IpBlockAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IpBlock) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetProject

`func (o *IpBlock) GetProject() IpBlockProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *IpBlock) GetProjectOk() (*IpBlockProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *IpBlock) SetProject(v IpBlockProject)`

SetProject sets Project field to given value.


### GetRegulations

`func (o *IpBlock) GetRegulations() IpBlockRegulations`

GetRegulations returns the Regulations field if non-nil, zero value otherwise.

### GetRegulationsOk

`func (o *IpBlock) GetRegulationsOk() (*IpBlockRegulations, bool)`

GetRegulationsOk returns a tuple with the Regulations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulations

`func (o *IpBlock) SetRegulations(v IpBlockRegulations)`

SetRegulations sets Regulations field to given value.

### HasRegulations

`func (o *IpBlock) HasRegulations() bool`

HasRegulations returns a boolean if a field has been set.

### GetAssets

`func (o *IpBlock) GetAssets() []IpBlockAsset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *IpBlock) GetAssetsOk() (*[]IpBlockAsset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *IpBlock) SetAssets(v []IpBlockAsset)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *IpBlock) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetChange

`func (o *IpBlock) GetChange() IpBlockChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *IpBlock) GetChangeOk() (*IpBlockChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *IpBlock) SetChange(v IpBlockChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *IpBlock) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetChangeLog

`func (o *IpBlock) GetChangeLog() IpBlockChangeLog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *IpBlock) GetChangeLogOk() (*IpBlockChangeLog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *IpBlock) SetChangeLog(v IpBlockChangeLog)`

SetChangeLog sets ChangeLog field to given value.


### GetError

`func (o *IpBlock) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IpBlock) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IpBlock) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *IpBlock) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


