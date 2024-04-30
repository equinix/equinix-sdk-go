# ServiceV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**Type** | Pointer to [**ServiceTypeV2**](ServiceTypeV2.md) |  | [optional] 
**Bandwidth** | Pointer to **int32** | Service bandwidth in Mbps | [optional] 
**Account** | Pointer to [**CustomerBillingAccount**](CustomerBillingAccount.md) |  | [optional] 
**ChangeLog** | Pointer to [**ServiceChangeLog**](ServiceChangeLog.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 
**Order** | Pointer to [**ServiceOrderReference**](ServiceOrderReference.md) |  | [optional] 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**State** | [**ServiceState**](ServiceState.md) |  | 

## Methods

### NewServiceV2

`func NewServiceV2(uuid string, state ServiceState, ) *ServiceV2`

NewServiceV2 instantiates a new ServiceV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceV2WithDefaults

`func NewServiceV2WithDefaults() *ServiceV2`

NewServiceV2WithDefaults instantiates a new ServiceV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ServiceV2) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceV2) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceV2) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *ServiceV2) GetType() ServiceTypeV2`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceV2) GetTypeOk() (*ServiceTypeV2, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceV2) SetType(v ServiceTypeV2)`

SetType sets Type field to given value.

### HasType

`func (o *ServiceV2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBandwidth

`func (o *ServiceV2) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *ServiceV2) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *ServiceV2) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *ServiceV2) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetAccount

`func (o *ServiceV2) GetAccount() CustomerBillingAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ServiceV2) GetAccountOk() (*CustomerBillingAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ServiceV2) SetAccount(v CustomerBillingAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ServiceV2) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetChangeLog

`func (o *ServiceV2) GetChangeLog() ServiceChangeLog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *ServiceV2) GetChangeLogOk() (*ServiceChangeLog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *ServiceV2) SetChangeLog(v ServiceChangeLog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *ServiceV2) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetLinks

`func (o *ServiceV2) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceV2) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceV2) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceV2) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetOrder

`func (o *ServiceV2) GetOrder() ServiceOrderReference`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ServiceV2) GetOrderOk() (*ServiceOrderReference, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ServiceV2) SetOrder(v ServiceOrderReference)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ServiceV2) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetProject

`func (o *ServiceV2) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ServiceV2) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ServiceV2) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *ServiceV2) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *ServiceV2) GetState() ServiceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ServiceV2) GetStateOk() (*ServiceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ServiceV2) SetState(v ServiceState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


