# PhysicalPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**PhysicalPortType**](PhysicalPortType.md) |  | [optional] 
**Id** | Pointer to **int32** | Equinix assigned response attribute for Physical Port Id | [optional] 
**Href** | Pointer to **string** | Equinix assigned response attribute for an absolute URL that is the subject of the link&#39;s context. | [optional] [readonly] 
**State** | Pointer to [**PortState**](PortState.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**InterfaceSpeed** | Pointer to **int32** | Physical Port Speed in Mbps | [optional] 
**InterfaceType** | Pointer to **string** | Physical Port Interface Type | [optional] 
**Tether** | Pointer to [**PortTether**](PortTether.md) |  | [optional] 
**DemarcationPoint** | Pointer to [**PortDemarcationPoint**](PortDemarcationPoint.md) |  | [optional] 
**AdditionalInfo** | Pointer to [**[]PortAdditionalInfo**](PortAdditionalInfo.md) | Physical Port additional information | [optional] 
**Order** | Pointer to [**PortOrder**](PortOrder.md) |  | [optional] 
**Operation** | Pointer to [**PortOperation**](PortOperation.md) |  | [optional] 
**Loas** | Pointer to [**[]PortLoa**](PortLoa.md) | Port Loas | [optional] 

## Methods

### NewPhysicalPort

`func NewPhysicalPort() *PhysicalPort`

NewPhysicalPort instantiates a new PhysicalPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalPortWithDefaults

`func NewPhysicalPortWithDefaults() *PhysicalPort`

NewPhysicalPortWithDefaults instantiates a new PhysicalPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PhysicalPort) GetType() PhysicalPortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PhysicalPort) GetTypeOk() (*PhysicalPortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PhysicalPort) SetType(v PhysicalPortType)`

SetType sets Type field to given value.

### HasType

`func (o *PhysicalPort) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *PhysicalPort) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhysicalPort) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhysicalPort) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PhysicalPort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *PhysicalPort) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PhysicalPort) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PhysicalPort) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PhysicalPort) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetState

`func (o *PhysicalPort) GetState() PortState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PhysicalPort) GetStateOk() (*PortState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PhysicalPort) SetState(v PortState)`

SetState sets State field to given value.

### HasState

`func (o *PhysicalPort) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAccount

`func (o *PhysicalPort) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PhysicalPort) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PhysicalPort) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PhysicalPort) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetInterfaceSpeed

`func (o *PhysicalPort) GetInterfaceSpeed() int32`

GetInterfaceSpeed returns the InterfaceSpeed field if non-nil, zero value otherwise.

### GetInterfaceSpeedOk

`func (o *PhysicalPort) GetInterfaceSpeedOk() (*int32, bool)`

GetInterfaceSpeedOk returns a tuple with the InterfaceSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceSpeed

`func (o *PhysicalPort) SetInterfaceSpeed(v int32)`

SetInterfaceSpeed sets InterfaceSpeed field to given value.

### HasInterfaceSpeed

`func (o *PhysicalPort) HasInterfaceSpeed() bool`

HasInterfaceSpeed returns a boolean if a field has been set.

### GetInterfaceType

`func (o *PhysicalPort) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *PhysicalPort) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *PhysicalPort) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *PhysicalPort) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetTether

`func (o *PhysicalPort) GetTether() PortTether`

GetTether returns the Tether field if non-nil, zero value otherwise.

### GetTetherOk

`func (o *PhysicalPort) GetTetherOk() (*PortTether, bool)`

GetTetherOk returns a tuple with the Tether field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTether

`func (o *PhysicalPort) SetTether(v PortTether)`

SetTether sets Tether field to given value.

### HasTether

`func (o *PhysicalPort) HasTether() bool`

HasTether returns a boolean if a field has been set.

### GetDemarcationPoint

`func (o *PhysicalPort) GetDemarcationPoint() PortDemarcationPoint`

GetDemarcationPoint returns the DemarcationPoint field if non-nil, zero value otherwise.

### GetDemarcationPointOk

`func (o *PhysicalPort) GetDemarcationPointOk() (*PortDemarcationPoint, bool)`

GetDemarcationPointOk returns a tuple with the DemarcationPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemarcationPoint

`func (o *PhysicalPort) SetDemarcationPoint(v PortDemarcationPoint)`

SetDemarcationPoint sets DemarcationPoint field to given value.

### HasDemarcationPoint

`func (o *PhysicalPort) HasDemarcationPoint() bool`

HasDemarcationPoint returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *PhysicalPort) GetAdditionalInfo() []PortAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *PhysicalPort) GetAdditionalInfoOk() (*[]PortAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *PhysicalPort) SetAdditionalInfo(v []PortAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *PhysicalPort) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetOrder

`func (o *PhysicalPort) GetOrder() PortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PhysicalPort) GetOrderOk() (*PortOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PhysicalPort) SetOrder(v PortOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PhysicalPort) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOperation

`func (o *PhysicalPort) GetOperation() PortOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *PhysicalPort) GetOperationOk() (*PortOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *PhysicalPort) SetOperation(v PortOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *PhysicalPort) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetLoas

`func (o *PhysicalPort) GetLoas() []PortLoa`

GetLoas returns the Loas field if non-nil, zero value otherwise.

### GetLoasOk

`func (o *PhysicalPort) GetLoasOk() (*[]PortLoa, bool)`

GetLoasOk returns a tuple with the Loas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoas

`func (o *PhysicalPort) SetLoas(v []PortLoa)`

SetLoas sets Loas field to given value.

### HasLoas

`func (o *PhysicalPort) HasLoas() bool`

HasLoas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


