# SimplifiedNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Network URI | [optional] [readonly] 
**Uuid** | **string** | Equinix-assigned network identifier | 
**Name** | Pointer to **string** | Customer-assigned network name | [optional] 
**State** | Pointer to [**NetworkState**](NetworkState.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Change** | Pointer to [**SimplifiedNetworkChange**](SimplifiedNetworkChange.md) |  | [optional] 
**Operation** | Pointer to [**NetworkOperation**](NetworkOperation.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | Network sub-resources links | [optional] [readonly] 
**Type** | Pointer to [**NetworkType**](NetworkType.md) |  | [optional] 
**Scope** | Pointer to [**NetworkScope**](NetworkScope.md) |  | [optional] 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 

## Methods

### NewSimplifiedNetwork

`func NewSimplifiedNetwork(uuid string, ) *SimplifiedNetwork`

NewSimplifiedNetwork instantiates a new SimplifiedNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedNetworkWithDefaults

`func NewSimplifiedNetworkWithDefaults() *SimplifiedNetwork`

NewSimplifiedNetworkWithDefaults instantiates a new SimplifiedNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *SimplifiedNetwork) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SimplifiedNetwork) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SimplifiedNetwork) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SimplifiedNetwork) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *SimplifiedNetwork) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SimplifiedNetwork) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SimplifiedNetwork) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *SimplifiedNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimplifiedNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimplifiedNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimplifiedNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *SimplifiedNetwork) GetState() NetworkState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SimplifiedNetwork) GetStateOk() (*NetworkState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SimplifiedNetwork) SetState(v NetworkState)`

SetState sets State field to given value.

### HasState

`func (o *SimplifiedNetwork) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAccount

`func (o *SimplifiedNetwork) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SimplifiedNetwork) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SimplifiedNetwork) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SimplifiedNetwork) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetChange

`func (o *SimplifiedNetwork) GetChange() SimplifiedNetworkChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *SimplifiedNetwork) GetChangeOk() (*SimplifiedNetworkChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *SimplifiedNetwork) SetChange(v SimplifiedNetworkChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *SimplifiedNetwork) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetOperation

`func (o *SimplifiedNetwork) GetOperation() NetworkOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *SimplifiedNetwork) GetOperationOk() (*NetworkOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *SimplifiedNetwork) SetOperation(v NetworkOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *SimplifiedNetwork) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetChangeLog

`func (o *SimplifiedNetwork) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *SimplifiedNetwork) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *SimplifiedNetwork) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *SimplifiedNetwork) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetLinks

`func (o *SimplifiedNetwork) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SimplifiedNetwork) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SimplifiedNetwork) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SimplifiedNetwork) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetType

`func (o *SimplifiedNetwork) GetType() NetworkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedNetwork) GetTypeOk() (*NetworkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedNetwork) SetType(v NetworkType)`

SetType sets Type field to given value.

### HasType

`func (o *SimplifiedNetwork) HasType() bool`

HasType returns a boolean if a field has been set.

### GetScope

`func (o *SimplifiedNetwork) GetScope() NetworkScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SimplifiedNetwork) GetScopeOk() (*NetworkScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SimplifiedNetwork) SetScope(v NetworkScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SimplifiedNetwork) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetLocation

`func (o *SimplifiedNetwork) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SimplifiedNetwork) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SimplifiedNetwork) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SimplifiedNetwork) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


