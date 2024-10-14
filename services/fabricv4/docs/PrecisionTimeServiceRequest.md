# PrecisionTimeServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PrecisionTimeServiceRequestType**](PrecisionTimeServiceRequestType.md) |  | 
**Name** | **string** | Precision Time Service name. | 
**Package** | [**PrecisionTimePackageRequest**](PrecisionTimePackageRequest.md) |  | 
**Connections** | [**[]VirtualConnectionUuid**](VirtualConnectionUuid.md) |  | 
**Ipv4** | [**Ipv4**](Ipv4.md) |  | 
**NtpAdvancedConfiguration** | Pointer to [**[]Md5**](Md5.md) | NTP Advanced configuration - MD5 Authentication. | [optional] 
**PtpAdvancedConfiguration** | Pointer to [**PtpAdvanceConfiguration**](PtpAdvanceConfiguration.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Order** | Pointer to [**PrecisionTimeOrder**](PrecisionTimeOrder.md) |  | [optional] 

## Methods

### NewPrecisionTimeServiceRequest

`func NewPrecisionTimeServiceRequest(type_ PrecisionTimeServiceRequestType, name string, package_ PrecisionTimePackageRequest, connections []VirtualConnectionUuid, ipv4 Ipv4, ) *PrecisionTimeServiceRequest`

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

`func (o *PrecisionTimeServiceRequest) GetConnections() []VirtualConnectionUuid`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *PrecisionTimeServiceRequest) GetConnectionsOk() (*[]VirtualConnectionUuid, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *PrecisionTimeServiceRequest) SetConnections(v []VirtualConnectionUuid)`

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


### GetNtpAdvancedConfiguration

`func (o *PrecisionTimeServiceRequest) GetNtpAdvancedConfiguration() []Md5`

GetNtpAdvancedConfiguration returns the NtpAdvancedConfiguration field if non-nil, zero value otherwise.

### GetNtpAdvancedConfigurationOk

`func (o *PrecisionTimeServiceRequest) GetNtpAdvancedConfigurationOk() (*[]Md5, bool)`

GetNtpAdvancedConfigurationOk returns a tuple with the NtpAdvancedConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpAdvancedConfiguration

`func (o *PrecisionTimeServiceRequest) SetNtpAdvancedConfiguration(v []Md5)`

SetNtpAdvancedConfiguration sets NtpAdvancedConfiguration field to given value.

### HasNtpAdvancedConfiguration

`func (o *PrecisionTimeServiceRequest) HasNtpAdvancedConfiguration() bool`

HasNtpAdvancedConfiguration returns a boolean if a field has been set.

### GetPtpAdvancedConfiguration

`func (o *PrecisionTimeServiceRequest) GetPtpAdvancedConfiguration() PtpAdvanceConfiguration`

GetPtpAdvancedConfiguration returns the PtpAdvancedConfiguration field if non-nil, zero value otherwise.

### GetPtpAdvancedConfigurationOk

`func (o *PrecisionTimeServiceRequest) GetPtpAdvancedConfigurationOk() (*PtpAdvanceConfiguration, bool)`

GetPtpAdvancedConfigurationOk returns a tuple with the PtpAdvancedConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpAdvancedConfiguration

`func (o *PrecisionTimeServiceRequest) SetPtpAdvancedConfiguration(v PtpAdvanceConfiguration)`

SetPtpAdvancedConfiguration sets PtpAdvancedConfiguration field to given value.

### HasPtpAdvancedConfiguration

`func (o *PrecisionTimeServiceRequest) HasPtpAdvancedConfiguration() bool`

HasPtpAdvancedConfiguration returns a boolean if a field has been set.

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

### GetOrder

`func (o *PrecisionTimeServiceRequest) GetOrder() PrecisionTimeOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PrecisionTimeServiceRequest) GetOrderOk() (*PrecisionTimeOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PrecisionTimeServiceRequest) SetOrder(v PrecisionTimeOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PrecisionTimeServiceRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


