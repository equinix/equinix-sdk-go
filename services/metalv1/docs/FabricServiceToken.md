# FabricServiceToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | The expiration date and time of the Fabric service token. Once a service token is expired, it is no longer redeemable. | [optional] 
**Id** | Pointer to **string** | The UUID that can be used on the Fabric Portal to redeem either an A-Side or Z-Side Service Token. For Fabric VCs (Metal Billed), this UUID will represent an A-Side Service Token, which will allow interconnections to be made from Equinix Metal to other Service Providers on Fabric. For Fabric VCs (Fabric Billed), this UUID will represent a Z-Side Service Token, which will allow interconnections to be made to connect an owned Fabric Port or  Virtual Device to Equinix Metal. | [optional] 
**MaxAllowedSpeed** | Pointer to **int64** | The maximum speed that can be selected on the Fabric Portal when configuring a interconnection with either  an A-Side or Z-Side Service Token. For Fabric VCs (Metal Billed), this is what the billing is based off of, and can be one of the following options, &#39;50mbps&#39;, &#39;200mbps&#39;, &#39;500mbps&#39;, &#39;1gbps&#39;, &#39;2gbps&#39;, &#39;5gbps&#39; or &#39;10gbps&#39;. For Fabric VCs (Fabric Billed), this will default to 10Gbps. | [optional] 
**Role** | Pointer to [**FabricServiceTokenRole**](FabricServiceTokenRole.md) |  | [optional] 
**ServiceTokenType** | Pointer to [**FabricServiceTokenServiceTokenType**](FabricServiceTokenServiceTokenType.md) |  | [optional] 
**State** | Pointer to [**FabricServiceTokenState**](FabricServiceTokenState.md) |  | [optional] 

## Methods

### NewFabricServiceToken

`func NewFabricServiceToken() *FabricServiceToken`

NewFabricServiceToken instantiates a new FabricServiceToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricServiceTokenWithDefaults

`func NewFabricServiceTokenWithDefaults() *FabricServiceToken`

NewFabricServiceTokenWithDefaults instantiates a new FabricServiceToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *FabricServiceToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *FabricServiceToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *FabricServiceToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *FabricServiceToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *FabricServiceToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FabricServiceToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FabricServiceToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FabricServiceToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxAllowedSpeed

`func (o *FabricServiceToken) GetMaxAllowedSpeed() int64`

GetMaxAllowedSpeed returns the MaxAllowedSpeed field if non-nil, zero value otherwise.

### GetMaxAllowedSpeedOk

`func (o *FabricServiceToken) GetMaxAllowedSpeedOk() (*int64, bool)`

GetMaxAllowedSpeedOk returns a tuple with the MaxAllowedSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedSpeed

`func (o *FabricServiceToken) SetMaxAllowedSpeed(v int64)`

SetMaxAllowedSpeed sets MaxAllowedSpeed field to given value.

### HasMaxAllowedSpeed

`func (o *FabricServiceToken) HasMaxAllowedSpeed() bool`

HasMaxAllowedSpeed returns a boolean if a field has been set.

### GetRole

`func (o *FabricServiceToken) GetRole() FabricServiceTokenRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *FabricServiceToken) GetRoleOk() (*FabricServiceTokenRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *FabricServiceToken) SetRole(v FabricServiceTokenRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *FabricServiceToken) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetServiceTokenType

`func (o *FabricServiceToken) GetServiceTokenType() FabricServiceTokenServiceTokenType`

GetServiceTokenType returns the ServiceTokenType field if non-nil, zero value otherwise.

### GetServiceTokenTypeOk

`func (o *FabricServiceToken) GetServiceTokenTypeOk() (*FabricServiceTokenServiceTokenType, bool)`

GetServiceTokenTypeOk returns a tuple with the ServiceTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenType

`func (o *FabricServiceToken) SetServiceTokenType(v FabricServiceTokenServiceTokenType)`

SetServiceTokenType sets ServiceTokenType field to given value.

### HasServiceTokenType

`func (o *FabricServiceToken) HasServiceTokenType() bool`

HasServiceTokenType returns a boolean if a field has been set.

### GetState

`func (o *FabricServiceToken) GetState() FabricServiceTokenState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FabricServiceToken) GetStateOk() (*FabricServiceTokenState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FabricServiceToken) SetState(v FabricServiceTokenState)`

SetState sets State field to given value.

### HasState

`func (o *FabricServiceToken) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


