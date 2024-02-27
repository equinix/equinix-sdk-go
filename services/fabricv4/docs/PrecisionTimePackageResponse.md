# PrecisionTimePackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Code** | [**PrecisionTimePackageResponseCode**](PrecisionTimePackageResponseCode.md) |  | 
**Type** | Pointer to [**PrecisionTimePackageResponseType**](PrecisionTimePackageResponseType.md) |  | [optional] 
**Bandwidth** | Pointer to **int32** |  | [optional] 
**ClientsPerSecondMax** | Pointer to **int32** |  | [optional] 
**RedundancySupported** | Pointer to **bool** |  | [optional] 
**MultiSubnetSupported** | Pointer to **bool** |  | [optional] 
**AccuracyUnit** | Pointer to **string** |  | [optional] 
**AccuracySla** | Pointer to **int32** |  | [optional] 
**AccuracyAvgMin** | Pointer to **int32** |  | [optional] 
**AccuracyAvgMax** | Pointer to **int32** |  | [optional] 
**Changelog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewPrecisionTimePackageResponse

`func NewPrecisionTimePackageResponse(code PrecisionTimePackageResponseCode, ) *PrecisionTimePackageResponse`

NewPrecisionTimePackageResponse instantiates a new PrecisionTimePackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrecisionTimePackageResponseWithDefaults

`func NewPrecisionTimePackageResponseWithDefaults() *PrecisionTimePackageResponse`

NewPrecisionTimePackageResponseWithDefaults instantiates a new PrecisionTimePackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PrecisionTimePackageResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PrecisionTimePackageResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PrecisionTimePackageResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PrecisionTimePackageResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetCode

`func (o *PrecisionTimePackageResponse) GetCode() PrecisionTimePackageResponseCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PrecisionTimePackageResponse) GetCodeOk() (*PrecisionTimePackageResponseCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PrecisionTimePackageResponse) SetCode(v PrecisionTimePackageResponseCode)`

SetCode sets Code field to given value.


### GetType

`func (o *PrecisionTimePackageResponse) GetType() PrecisionTimePackageResponseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrecisionTimePackageResponse) GetTypeOk() (*PrecisionTimePackageResponseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrecisionTimePackageResponse) SetType(v PrecisionTimePackageResponseType)`

SetType sets Type field to given value.

### HasType

`func (o *PrecisionTimePackageResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBandwidth

`func (o *PrecisionTimePackageResponse) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *PrecisionTimePackageResponse) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *PrecisionTimePackageResponse) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *PrecisionTimePackageResponse) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetClientsPerSecondMax

`func (o *PrecisionTimePackageResponse) GetClientsPerSecondMax() int32`

GetClientsPerSecondMax returns the ClientsPerSecondMax field if non-nil, zero value otherwise.

### GetClientsPerSecondMaxOk

`func (o *PrecisionTimePackageResponse) GetClientsPerSecondMaxOk() (*int32, bool)`

GetClientsPerSecondMaxOk returns a tuple with the ClientsPerSecondMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsPerSecondMax

`func (o *PrecisionTimePackageResponse) SetClientsPerSecondMax(v int32)`

SetClientsPerSecondMax sets ClientsPerSecondMax field to given value.

### HasClientsPerSecondMax

`func (o *PrecisionTimePackageResponse) HasClientsPerSecondMax() bool`

HasClientsPerSecondMax returns a boolean if a field has been set.

### GetRedundancySupported

`func (o *PrecisionTimePackageResponse) GetRedundancySupported() bool`

GetRedundancySupported returns the RedundancySupported field if non-nil, zero value otherwise.

### GetRedundancySupportedOk

`func (o *PrecisionTimePackageResponse) GetRedundancySupportedOk() (*bool, bool)`

GetRedundancySupportedOk returns a tuple with the RedundancySupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancySupported

`func (o *PrecisionTimePackageResponse) SetRedundancySupported(v bool)`

SetRedundancySupported sets RedundancySupported field to given value.

### HasRedundancySupported

`func (o *PrecisionTimePackageResponse) HasRedundancySupported() bool`

HasRedundancySupported returns a boolean if a field has been set.

### GetMultiSubnetSupported

`func (o *PrecisionTimePackageResponse) GetMultiSubnetSupported() bool`

GetMultiSubnetSupported returns the MultiSubnetSupported field if non-nil, zero value otherwise.

### GetMultiSubnetSupportedOk

`func (o *PrecisionTimePackageResponse) GetMultiSubnetSupportedOk() (*bool, bool)`

GetMultiSubnetSupportedOk returns a tuple with the MultiSubnetSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSubnetSupported

`func (o *PrecisionTimePackageResponse) SetMultiSubnetSupported(v bool)`

SetMultiSubnetSupported sets MultiSubnetSupported field to given value.

### HasMultiSubnetSupported

`func (o *PrecisionTimePackageResponse) HasMultiSubnetSupported() bool`

HasMultiSubnetSupported returns a boolean if a field has been set.

### GetAccuracyUnit

`func (o *PrecisionTimePackageResponse) GetAccuracyUnit() string`

GetAccuracyUnit returns the AccuracyUnit field if non-nil, zero value otherwise.

### GetAccuracyUnitOk

`func (o *PrecisionTimePackageResponse) GetAccuracyUnitOk() (*string, bool)`

GetAccuracyUnitOk returns a tuple with the AccuracyUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracyUnit

`func (o *PrecisionTimePackageResponse) SetAccuracyUnit(v string)`

SetAccuracyUnit sets AccuracyUnit field to given value.

### HasAccuracyUnit

`func (o *PrecisionTimePackageResponse) HasAccuracyUnit() bool`

HasAccuracyUnit returns a boolean if a field has been set.

### GetAccuracySla

`func (o *PrecisionTimePackageResponse) GetAccuracySla() int32`

GetAccuracySla returns the AccuracySla field if non-nil, zero value otherwise.

### GetAccuracySlaOk

`func (o *PrecisionTimePackageResponse) GetAccuracySlaOk() (*int32, bool)`

GetAccuracySlaOk returns a tuple with the AccuracySla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracySla

`func (o *PrecisionTimePackageResponse) SetAccuracySla(v int32)`

SetAccuracySla sets AccuracySla field to given value.

### HasAccuracySla

`func (o *PrecisionTimePackageResponse) HasAccuracySla() bool`

HasAccuracySla returns a boolean if a field has been set.

### GetAccuracyAvgMin

`func (o *PrecisionTimePackageResponse) GetAccuracyAvgMin() int32`

GetAccuracyAvgMin returns the AccuracyAvgMin field if non-nil, zero value otherwise.

### GetAccuracyAvgMinOk

`func (o *PrecisionTimePackageResponse) GetAccuracyAvgMinOk() (*int32, bool)`

GetAccuracyAvgMinOk returns a tuple with the AccuracyAvgMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracyAvgMin

`func (o *PrecisionTimePackageResponse) SetAccuracyAvgMin(v int32)`

SetAccuracyAvgMin sets AccuracyAvgMin field to given value.

### HasAccuracyAvgMin

`func (o *PrecisionTimePackageResponse) HasAccuracyAvgMin() bool`

HasAccuracyAvgMin returns a boolean if a field has been set.

### GetAccuracyAvgMax

`func (o *PrecisionTimePackageResponse) GetAccuracyAvgMax() int32`

GetAccuracyAvgMax returns the AccuracyAvgMax field if non-nil, zero value otherwise.

### GetAccuracyAvgMaxOk

`func (o *PrecisionTimePackageResponse) GetAccuracyAvgMaxOk() (*int32, bool)`

GetAccuracyAvgMaxOk returns a tuple with the AccuracyAvgMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracyAvgMax

`func (o *PrecisionTimePackageResponse) SetAccuracyAvgMax(v int32)`

SetAccuracyAvgMax sets AccuracyAvgMax field to given value.

### HasAccuracyAvgMax

`func (o *PrecisionTimePackageResponse) HasAccuracyAvgMax() bool`

HasAccuracyAvgMax returns a boolean if a field has been set.

### GetChangelog

`func (o *PrecisionTimePackageResponse) GetChangelog() Changelog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *PrecisionTimePackageResponse) GetChangelogOk() (*Changelog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *PrecisionTimePackageResponse) SetChangelog(v Changelog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *PrecisionTimePackageResponse) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


