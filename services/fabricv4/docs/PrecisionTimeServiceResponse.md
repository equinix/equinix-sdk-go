# PrecisionTimeServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | Precision Time Service URI. | 
**Type** | [**PrecisionTimeServiceResponseType**](PrecisionTimeServiceResponseType.md) |  | 
**Name** | Pointer to **string** | Precision Time Service Name. | [optional] 
**Uuid** | **string** | Precision Time Service UUID. | 
**State** | [**PrecisionTimeServiceResponseState**](PrecisionTimeServiceResponseState.md) |  | 
**Package** | [**PrecisionTimePackageResponse**](PrecisionTimePackageResponse.md) |  | 
**Connections** | Pointer to [**[]VirtualConnectionTimeServiceResponse**](VirtualConnectionTimeServiceResponse.md) | Fabric Connections associated with Precision Time Service. | [optional] 
**Ipv4** | Pointer to [**Ipv4**](Ipv4.md) |  | [optional] 
**NtpAdvancedConfiguration** | Pointer to [**[]Md5**](Md5.md) | NTP Advanced configuration - MD5 Authentication. | [optional] 
**PtpAdvancedConfiguration** | Pointer to [**PtpAdvanceConfiguration**](PtpAdvanceConfiguration.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Order** | Pointer to [**PrecisionTimeOrder**](PrecisionTimeOrder.md) |  | [optional] 
**Pricing** | Pointer to [**PrecisionTimePrice**](PrecisionTimePrice.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewPrecisionTimeServiceResponse

`func NewPrecisionTimeServiceResponse(href string, type_ PrecisionTimeServiceResponseType, uuid string, state PrecisionTimeServiceResponseState, package_ PrecisionTimePackageResponse, ) *PrecisionTimeServiceResponse`

NewPrecisionTimeServiceResponse instantiates a new PrecisionTimeServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrecisionTimeServiceResponseWithDefaults

`func NewPrecisionTimeServiceResponseWithDefaults() *PrecisionTimeServiceResponse`

NewPrecisionTimeServiceResponseWithDefaults instantiates a new PrecisionTimeServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetType

`func (o *PrecisionTimeServiceResponse) GetType() PrecisionTimeServiceResponseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrecisionTimeServiceResponse) GetTypeOk() (*PrecisionTimeServiceResponseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrecisionTimeServiceResponse) SetType(v PrecisionTimeServiceResponseType)`

SetType sets Type field to given value.


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


### GetState

`func (o *PrecisionTimeServiceResponse) GetState() PrecisionTimeServiceResponseState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PrecisionTimeServiceResponse) GetStateOk() (*PrecisionTimeServiceResponseState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PrecisionTimeServiceResponse) SetState(v PrecisionTimeServiceResponseState)`

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

`func (o *PrecisionTimeServiceResponse) GetConnections() []VirtualConnectionTimeServiceResponse`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *PrecisionTimeServiceResponse) GetConnectionsOk() (*[]VirtualConnectionTimeServiceResponse, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *PrecisionTimeServiceResponse) SetConnections(v []VirtualConnectionTimeServiceResponse)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *PrecisionTimeServiceResponse) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

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

### HasIpv4

`func (o *PrecisionTimeServiceResponse) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetNtpAdvancedConfiguration

`func (o *PrecisionTimeServiceResponse) GetNtpAdvancedConfiguration() []Md5`

GetNtpAdvancedConfiguration returns the NtpAdvancedConfiguration field if non-nil, zero value otherwise.

### GetNtpAdvancedConfigurationOk

`func (o *PrecisionTimeServiceResponse) GetNtpAdvancedConfigurationOk() (*[]Md5, bool)`

GetNtpAdvancedConfigurationOk returns a tuple with the NtpAdvancedConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpAdvancedConfiguration

`func (o *PrecisionTimeServiceResponse) SetNtpAdvancedConfiguration(v []Md5)`

SetNtpAdvancedConfiguration sets NtpAdvancedConfiguration field to given value.

### HasNtpAdvancedConfiguration

`func (o *PrecisionTimeServiceResponse) HasNtpAdvancedConfiguration() bool`

HasNtpAdvancedConfiguration returns a boolean if a field has been set.

### GetPtpAdvancedConfiguration

`func (o *PrecisionTimeServiceResponse) GetPtpAdvancedConfiguration() PtpAdvanceConfiguration`

GetPtpAdvancedConfiguration returns the PtpAdvancedConfiguration field if non-nil, zero value otherwise.

### GetPtpAdvancedConfigurationOk

`func (o *PrecisionTimeServiceResponse) GetPtpAdvancedConfigurationOk() (*PtpAdvanceConfiguration, bool)`

GetPtpAdvancedConfigurationOk returns a tuple with the PtpAdvancedConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpAdvancedConfiguration

`func (o *PrecisionTimeServiceResponse) SetPtpAdvancedConfiguration(v PtpAdvanceConfiguration)`

SetPtpAdvancedConfiguration sets PtpAdvancedConfiguration field to given value.

### HasPtpAdvancedConfiguration

`func (o *PrecisionTimeServiceResponse) HasPtpAdvancedConfiguration() bool`

HasPtpAdvancedConfiguration returns a boolean if a field has been set.

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

`func (o *PrecisionTimeServiceResponse) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PrecisionTimeServiceResponse) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PrecisionTimeServiceResponse) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PrecisionTimeServiceResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOrder

`func (o *PrecisionTimeServiceResponse) GetOrder() PrecisionTimeOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PrecisionTimeServiceResponse) GetOrderOk() (*PrecisionTimeOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PrecisionTimeServiceResponse) SetOrder(v PrecisionTimeOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PrecisionTimeServiceResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPricing

`func (o *PrecisionTimeServiceResponse) GetPricing() PrecisionTimePrice`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *PrecisionTimeServiceResponse) GetPricingOk() (*PrecisionTimePrice, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *PrecisionTimeServiceResponse) SetPricing(v PrecisionTimePrice)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *PrecisionTimeServiceResponse) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetChangeLog

`func (o *PrecisionTimeServiceResponse) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *PrecisionTimeServiceResponse) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *PrecisionTimeServiceResponse) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *PrecisionTimeServiceResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


