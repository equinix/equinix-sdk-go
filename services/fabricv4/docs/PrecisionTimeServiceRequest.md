# PrecisionTimeServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PrecisionTimeServiceRequestType**](PrecisionTimeServiceRequestType.md) |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Package** | [**PrecisionTimePackageRequest**](PrecisionTimePackageRequest.md) |  | 
**Connections** | [**[]FabricConnectionUuid**](FabricConnectionUuid.md) |  | 
**Ipv4** | [**Ipv4**](Ipv4.md) |  | 
**AdvanceConfiguration** | Pointer to [**AdvanceConfiguration**](AdvanceConfiguration.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 

## Methods

### NewPrecisionTimeServiceRequest

`func NewPrecisionTimeServiceRequest(type_ PrecisionTimeServiceRequestType, name string, package_ PrecisionTimePackageRequest, connections []FabricConnectionUuid, ipv4 Ipv4, ) *PrecisionTimeServiceRequest`

NewPrecisionTimeServiceRequest instantiates a new PrecisionTimeServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrecisionTimeServiceRequestWithDefaults

`func NewPrecisionTimeServiceRequestWithDefaults() *PrecisionTimeServiceRequest`

NewPrecisionTimeServiceRequestWithDefaults instantiates a new PrecisionTimeServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PrecisionTimeServiceRequest) GetType() PrecisionTimeServiceRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrecisionTimeServiceRequest) GetTypeOk() (*PrecisionTimeServiceRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrecisionTimeServiceRequest) SetType(v PrecisionTimeServiceRequestType)`

SetType sets Type field to given value.


### GetName

`func (o *PrecisionTimeServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrecisionTimeServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrecisionTimeServiceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PrecisionTimeServiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrecisionTimeServiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrecisionTimeServiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrecisionTimeServiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPackage

`func (o *PrecisionTimeServiceRequest) GetPackage() PrecisionTimePackageRequest`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *PrecisionTimeServiceRequest) GetPackageOk() (*PrecisionTimePackageRequest, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *PrecisionTimeServiceRequest) SetPackage(v PrecisionTimePackageRequest)`

SetPackage sets Package field to given value.


### GetConnections

`func (o *PrecisionTimeServiceRequest) GetConnections() []FabricConnectionUuid`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *PrecisionTimeServiceRequest) GetConnectionsOk() (*[]FabricConnectionUuid, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *PrecisionTimeServiceRequest) SetConnections(v []FabricConnectionUuid)`

SetConnections sets Connections field to given value.


### GetIpv4

`func (o *PrecisionTimeServiceRequest) GetIpv4() Ipv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *PrecisionTimeServiceRequest) GetIpv4Ok() (*Ipv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *PrecisionTimeServiceRequest) SetIpv4(v Ipv4)`

SetIpv4 sets Ipv4 field to given value.


### GetAdvanceConfiguration

`func (o *PrecisionTimeServiceRequest) GetAdvanceConfiguration() AdvanceConfiguration`

GetAdvanceConfiguration returns the AdvanceConfiguration field if non-nil, zero value otherwise.

### GetAdvanceConfigurationOk

`func (o *PrecisionTimeServiceRequest) GetAdvanceConfigurationOk() (*AdvanceConfiguration, bool)`

GetAdvanceConfigurationOk returns a tuple with the AdvanceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanceConfiguration

`func (o *PrecisionTimeServiceRequest) SetAdvanceConfiguration(v AdvanceConfiguration)`

SetAdvanceConfiguration sets AdvanceConfiguration field to given value.

### HasAdvanceConfiguration

`func (o *PrecisionTimeServiceRequest) HasAdvanceConfiguration() bool`

HasAdvanceConfiguration returns a boolean if a field has been set.

### GetProject

`func (o *PrecisionTimeServiceRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PrecisionTimeServiceRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PrecisionTimeServiceRequest) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *PrecisionTimeServiceRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


