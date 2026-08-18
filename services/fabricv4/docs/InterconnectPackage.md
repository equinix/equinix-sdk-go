# InterconnectPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Interconnect Package URI | [optional] [readonly] 
**Type** | **string** | Interconnect Package Type | 
**Uuid** | Pointer to **string** | Equinix-assigned Interconnect Package identifier | [optional] 
**Code** | **string** | Interconnect Package code (e.g. LAB, BASIC, STANDARD, PREMIUM) | 
**Description** | Pointer to **string** | Interconnect Package description | [optional] 
**RoutesMax** | Pointer to **int32** | Maximum number of routes | [optional] 
**BandwidthMax** | Pointer to **int32** | Maximum bandwidth in Mbps | [optional] 
**IsRemote** | Pointer to **bool** | Authorization to connect remotely | [optional] 

## Methods

### NewInterconnectPackage

`func NewInterconnectPackage(type_ string, code string, ) *InterconnectPackage`

NewInterconnectPackage instantiates a new InterconnectPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectPackageWithDefaults

`func NewInterconnectPackageWithDefaults() *InterconnectPackage`

NewInterconnectPackageWithDefaults instantiates a new InterconnectPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *InterconnectPackage) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *InterconnectPackage) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *InterconnectPackage) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *InterconnectPackage) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *InterconnectPackage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InterconnectPackage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InterconnectPackage) SetType(v string)`

SetType sets Type field to given value.


### GetUuid

`func (o *InterconnectPackage) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InterconnectPackage) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InterconnectPackage) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InterconnectPackage) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCode

`func (o *InterconnectPackage) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InterconnectPackage) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InterconnectPackage) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *InterconnectPackage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InterconnectPackage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InterconnectPackage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InterconnectPackage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRoutesMax

`func (o *InterconnectPackage) GetRoutesMax() int32`

GetRoutesMax returns the RoutesMax field if non-nil, zero value otherwise.

### GetRoutesMaxOk

`func (o *InterconnectPackage) GetRoutesMaxOk() (*int32, bool)`

GetRoutesMaxOk returns a tuple with the RoutesMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutesMax

`func (o *InterconnectPackage) SetRoutesMax(v int32)`

SetRoutesMax sets RoutesMax field to given value.

### HasRoutesMax

`func (o *InterconnectPackage) HasRoutesMax() bool`

HasRoutesMax returns a boolean if a field has been set.

### GetBandwidthMax

`func (o *InterconnectPackage) GetBandwidthMax() int32`

GetBandwidthMax returns the BandwidthMax field if non-nil, zero value otherwise.

### GetBandwidthMaxOk

`func (o *InterconnectPackage) GetBandwidthMaxOk() (*int32, bool)`

GetBandwidthMaxOk returns a tuple with the BandwidthMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMax

`func (o *InterconnectPackage) SetBandwidthMax(v int32)`

SetBandwidthMax sets BandwidthMax field to given value.

### HasBandwidthMax

`func (o *InterconnectPackage) HasBandwidthMax() bool`

HasBandwidthMax returns a boolean if a field has been set.

### GetIsRemote

`func (o *InterconnectPackage) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *InterconnectPackage) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *InterconnectPackage) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *InterconnectPackage) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


