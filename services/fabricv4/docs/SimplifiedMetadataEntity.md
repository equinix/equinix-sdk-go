# SimplifiedMetadataEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | url to entity | [optional] 
**Uuid** | Pointer to **string** | Equinix assigned Identifier | [optional] 
**Type** | Pointer to **string** | Type of Port | [optional] 
**CvpId** | Pointer to **int32** | Customer virtual port Id | [optional] 
**Bandwidth** | Pointer to **float32** | Port Bandwidth | [optional] 
**PortName** | Pointer to **string** | Port Name | [optional] 
**EncapsulationProtocolType** | Pointer to **string** | Port Encapsulation | [optional] 
**AccountName** | Pointer to **string** | Account Name | [optional] 
**Priority** | Pointer to **string** | Port Priority | [optional] 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 

## Methods

### NewSimplifiedMetadataEntity

`func NewSimplifiedMetadataEntity() *SimplifiedMetadataEntity`

NewSimplifiedMetadataEntity instantiates a new SimplifiedMetadataEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedMetadataEntityWithDefaults

`func NewSimplifiedMetadataEntityWithDefaults() *SimplifiedMetadataEntity`

NewSimplifiedMetadataEntityWithDefaults instantiates a new SimplifiedMetadataEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *SimplifiedMetadataEntity) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SimplifiedMetadataEntity) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SimplifiedMetadataEntity) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SimplifiedMetadataEntity) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *SimplifiedMetadataEntity) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SimplifiedMetadataEntity) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SimplifiedMetadataEntity) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SimplifiedMetadataEntity) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *SimplifiedMetadataEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedMetadataEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedMetadataEntity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SimplifiedMetadataEntity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCvpId

`func (o *SimplifiedMetadataEntity) GetCvpId() int32`

GetCvpId returns the CvpId field if non-nil, zero value otherwise.

### GetCvpIdOk

`func (o *SimplifiedMetadataEntity) GetCvpIdOk() (*int32, bool)`

GetCvpIdOk returns a tuple with the CvpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvpId

`func (o *SimplifiedMetadataEntity) SetCvpId(v int32)`

SetCvpId sets CvpId field to given value.

### HasCvpId

`func (o *SimplifiedMetadataEntity) HasCvpId() bool`

HasCvpId returns a boolean if a field has been set.

### GetBandwidth

`func (o *SimplifiedMetadataEntity) GetBandwidth() float32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *SimplifiedMetadataEntity) GetBandwidthOk() (*float32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *SimplifiedMetadataEntity) SetBandwidth(v float32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *SimplifiedMetadataEntity) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetPortName

`func (o *SimplifiedMetadataEntity) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *SimplifiedMetadataEntity) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *SimplifiedMetadataEntity) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *SimplifiedMetadataEntity) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetEncapsulationProtocolType

`func (o *SimplifiedMetadataEntity) GetEncapsulationProtocolType() string`

GetEncapsulationProtocolType returns the EncapsulationProtocolType field if non-nil, zero value otherwise.

### GetEncapsulationProtocolTypeOk

`func (o *SimplifiedMetadataEntity) GetEncapsulationProtocolTypeOk() (*string, bool)`

GetEncapsulationProtocolTypeOk returns a tuple with the EncapsulationProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulationProtocolType

`func (o *SimplifiedMetadataEntity) SetEncapsulationProtocolType(v string)`

SetEncapsulationProtocolType sets EncapsulationProtocolType field to given value.

### HasEncapsulationProtocolType

`func (o *SimplifiedMetadataEntity) HasEncapsulationProtocolType() bool`

HasEncapsulationProtocolType returns a boolean if a field has been set.

### GetAccountName

`func (o *SimplifiedMetadataEntity) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *SimplifiedMetadataEntity) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *SimplifiedMetadataEntity) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *SimplifiedMetadataEntity) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetPriority

`func (o *SimplifiedMetadataEntity) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SimplifiedMetadataEntity) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SimplifiedMetadataEntity) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SimplifiedMetadataEntity) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetLocation

`func (o *SimplifiedMetadataEntity) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SimplifiedMetadataEntity) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SimplifiedMetadataEntity) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SimplifiedMetadataEntity) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


