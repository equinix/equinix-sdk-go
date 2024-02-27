# PrecisionTimeServiceCreateResponse

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
**Ipv4** | [**Ipv4**](Ipv4.md) |  | 
**Account** | Pointer to [**Account**](Account.md) |  | [optional] 
**AdvanceConfiguration** | Pointer to [**AdvanceConfiguration**](AdvanceConfiguration.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 

## Methods

### NewPrecisionTimeServiceCreateResponse

`func NewPrecisionTimeServiceCreateResponse(type_ PrecisionTimeServiceCreateResponseType, href string, uuid string, state PrecisionTimeServiceCreateResponseState, package_ PrecisionTimePackageResponse, ipv4 Ipv4, ) *PrecisionTimeServiceCreateResponse`

NewPrecisionTimeServiceCreateResponse instantiates a new PrecisionTimeServiceCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrecisionTimeServiceCreateResponseWithDefaults

`func NewPrecisionTimeServiceCreateResponseWithDefaults() *PrecisionTimeServiceCreateResponse`

NewPrecisionTimeServiceCreateResponseWithDefaults instantiates a new PrecisionTimeServiceCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PrecisionTimeServiceCreateResponse) GetType() PrecisionTimeServiceCreateResponseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrecisionTimeServiceCreateResponse) GetTypeOk() (*PrecisionTimeServiceCreateResponseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrecisionTimeServiceCreateResponse) SetType(v PrecisionTimeServiceCreateResponseType)`

SetType sets Type field to given value.


### GetHref

`func (o *PrecisionTimeServiceCreateResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PrecisionTimeServiceCreateResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PrecisionTimeServiceCreateResponse) SetHref(v string)`

SetHref sets Href field to given value.


### GetUuid

`func (o *PrecisionTimeServiceCreateResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PrecisionTimeServiceCreateResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PrecisionTimeServiceCreateResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *PrecisionTimeServiceCreateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrecisionTimeServiceCreateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrecisionTimeServiceCreateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrecisionTimeServiceCreateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PrecisionTimeServiceCreateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrecisionTimeServiceCreateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrecisionTimeServiceCreateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrecisionTimeServiceCreateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *PrecisionTimeServiceCreateResponse) GetState() PrecisionTimeServiceCreateResponseState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PrecisionTimeServiceCreateResponse) GetStateOk() (*PrecisionTimeServiceCreateResponseState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PrecisionTimeServiceCreateResponse) SetState(v PrecisionTimeServiceCreateResponseState)`

SetState sets State field to given value.


### GetPackage

`func (o *PrecisionTimeServiceCreateResponse) GetPackage() PrecisionTimePackageResponse`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *PrecisionTimeServiceCreateResponse) GetPackageOk() (*PrecisionTimePackageResponse, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *PrecisionTimeServiceCreateResponse) SetPackage(v PrecisionTimePackageResponse)`

SetPackage sets Package field to given value.


### GetConnections

`func (o *PrecisionTimeServiceCreateResponse) GetConnections() []FabricConnectionUuid`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *PrecisionTimeServiceCreateResponse) GetConnectionsOk() (*[]FabricConnectionUuid, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *PrecisionTimeServiceCreateResponse) SetConnections(v []FabricConnectionUuid)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *PrecisionTimeServiceCreateResponse) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetIpv4

`func (o *PrecisionTimeServiceCreateResponse) GetIpv4() Ipv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *PrecisionTimeServiceCreateResponse) GetIpv4Ok() (*Ipv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *PrecisionTimeServiceCreateResponse) SetIpv4(v Ipv4)`

SetIpv4 sets Ipv4 field to given value.


### GetAccount

`func (o *PrecisionTimeServiceCreateResponse) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PrecisionTimeServiceCreateResponse) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PrecisionTimeServiceCreateResponse) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PrecisionTimeServiceCreateResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAdvanceConfiguration

`func (o *PrecisionTimeServiceCreateResponse) GetAdvanceConfiguration() AdvanceConfiguration`

GetAdvanceConfiguration returns the AdvanceConfiguration field if non-nil, zero value otherwise.

### GetAdvanceConfigurationOk

`func (o *PrecisionTimeServiceCreateResponse) GetAdvanceConfigurationOk() (*AdvanceConfiguration, bool)`

GetAdvanceConfigurationOk returns a tuple with the AdvanceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanceConfiguration

`func (o *PrecisionTimeServiceCreateResponse) SetAdvanceConfiguration(v AdvanceConfiguration)`

SetAdvanceConfiguration sets AdvanceConfiguration field to given value.

### HasAdvanceConfiguration

`func (o *PrecisionTimeServiceCreateResponse) HasAdvanceConfiguration() bool`

HasAdvanceConfiguration returns a boolean if a field has been set.

### GetProject

`func (o *PrecisionTimeServiceCreateResponse) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PrecisionTimeServiceCreateResponse) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PrecisionTimeServiceCreateResponse) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *PrecisionTimeServiceCreateResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


