# PrecisionTimeServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PrecisionTimeServiceCreateResponseType**](PrecisionTimeServiceCreateResponseType.md) |  | 
**Href** | **string** |  | 
**Uuid** | **string** | uuid of the ept service | 
**Name** | Pointer to **string** | name of the ept service | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**State** | [**PrecisionTimeServiceCreateResponseState**](PrecisionTimeServiceCreateResponseState.md) |  | 
**Package** | [**PrecisionTimePackageResponse**](PrecisionTimePackageResponse.md) |  | 
**Connections** | Pointer to [**[]FabricConnectionUuid**](FabricConnectionUuid.md) | fabric l2 connections used for the ept service | [optional] 
**Order** | [**Order**](Order.md) |  | 
**Ipv4** | [**Ipv4**](Ipv4.md) |  | 
**AdvanceConfiguration** | Pointer to [**AdvanceConfiguration**](AdvanceConfiguration.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Account** | Pointer to [**Account**](Account.md) |  | [optional] 

## Methods

### NewPrecisionTimeServiceResponse

`func NewPrecisionTimeServiceResponse(type_ PrecisionTimeServiceCreateResponseType, href string, uuid string, state PrecisionTimeServiceCreateResponseState, package_ PrecisionTimePackageResponse, order Order, ipv4 Ipv4, ) *PrecisionTimeServiceResponse`

NewPrecisionTimeServiceResponse instantiates a new PrecisionTimeServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrecisionTimeServiceResponseWithDefaults

`func NewPrecisionTimeServiceResponseWithDefaults() *PrecisionTimeServiceResponse`

NewPrecisionTimeServiceResponseWithDefaults instantiates a new PrecisionTimeServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PrecisionTimeServiceResponse) GetType() PrecisionTimeServiceCreateResponseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrecisionTimeServiceResponse) GetTypeOk() (*PrecisionTimeServiceCreateResponseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrecisionTimeServiceResponse) SetType(v PrecisionTimeServiceCreateResponseType)`

SetType sets Type field to given value.


### GetHref

`func (o *PrecisionTimeServiceResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PrecisionTimeServiceResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PrecisionTimeServiceResponse) SetHref(v string)`

SetHref sets Href field to given value.


### GetUuid

`func (o *PrecisionTimeServiceResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PrecisionTimeServiceResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PrecisionTimeServiceResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *PrecisionTimeServiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrecisionTimeServiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrecisionTimeServiceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrecisionTimeServiceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PrecisionTimeServiceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrecisionTimeServiceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrecisionTimeServiceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrecisionTimeServiceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *PrecisionTimeServiceResponse) GetState() PrecisionTimeServiceCreateResponseState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PrecisionTimeServiceResponse) GetStateOk() (*PrecisionTimeServiceCreateResponseState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PrecisionTimeServiceResponse) SetState(v PrecisionTimeServiceCreateResponseState)`

SetState sets State field to given value.


### GetPackage

`func (o *PrecisionTimeServiceResponse) GetPackage() PrecisionTimePackageResponse`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *PrecisionTimeServiceResponse) GetPackageOk() (*PrecisionTimePackageResponse, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *PrecisionTimeServiceResponse) SetPackage(v PrecisionTimePackageResponse)`

SetPackage sets Package field to given value.


### GetConnections

`func (o *PrecisionTimeServiceResponse) GetConnections() []FabricConnectionUuid`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *PrecisionTimeServiceResponse) GetConnectionsOk() (*[]FabricConnectionUuid, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *PrecisionTimeServiceResponse) SetConnections(v []FabricConnectionUuid)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *PrecisionTimeServiceResponse) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetOrder

`func (o *PrecisionTimeServiceResponse) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PrecisionTimeServiceResponse) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PrecisionTimeServiceResponse) SetOrder(v Order)`

SetOrder sets Order field to given value.


### GetIpv4

`func (o *PrecisionTimeServiceResponse) GetIpv4() Ipv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *PrecisionTimeServiceResponse) GetIpv4Ok() (*Ipv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *PrecisionTimeServiceResponse) SetIpv4(v Ipv4)`

SetIpv4 sets Ipv4 field to given value.


### GetAdvanceConfiguration

`func (o *PrecisionTimeServiceResponse) GetAdvanceConfiguration() AdvanceConfiguration`

GetAdvanceConfiguration returns the AdvanceConfiguration field if non-nil, zero value otherwise.

### GetAdvanceConfigurationOk

`func (o *PrecisionTimeServiceResponse) GetAdvanceConfigurationOk() (*AdvanceConfiguration, bool)`

GetAdvanceConfigurationOk returns a tuple with the AdvanceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanceConfiguration

`func (o *PrecisionTimeServiceResponse) SetAdvanceConfiguration(v AdvanceConfiguration)`

SetAdvanceConfiguration sets AdvanceConfiguration field to given value.

### HasAdvanceConfiguration

`func (o *PrecisionTimeServiceResponse) HasAdvanceConfiguration() bool`

HasAdvanceConfiguration returns a boolean if a field has been set.

### GetProject

`func (o *PrecisionTimeServiceResponse) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PrecisionTimeServiceResponse) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PrecisionTimeServiceResponse) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *PrecisionTimeServiceResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAccount

`func (o *PrecisionTimeServiceResponse) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PrecisionTimeServiceResponse) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PrecisionTimeServiceResponse) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PrecisionTimeServiceResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


