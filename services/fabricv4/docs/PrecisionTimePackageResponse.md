# PrecisionTimePackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Type** | [**PrecisionTimePackageResponseType**](PrecisionTimePackageResponseType.md) |  | 
**Code** | [**GetTimeServicesPackageByCodePackageCodeParameter**](GetTimeServicesPackageByCodePackageCodeParameter.md) |  | 
**Bandwidth** | **int32** | Connection bandwidth in Mbps. | 
**ClientsPerSecondMax** | Pointer to **int32** | Max. number of clients that can be synchronized per second at a packet rate of 1 per second. | [optional] 
**RedundancySupported** | Pointer to **bool** | Is Redundant virtual connection supported for the package code. | [optional] 
**MultiSubnetSupported** | Pointer to **bool** | Is Multiple subnet supported for the package code. | [optional] 
**AccuracySlaUnit** | Pointer to **string** | Accuracy SLA unit. | [optional] 
**AccuracySla** | Pointer to **int32** | Accuracy SLA for the package code, -1 value denotes the accuracySla is not published. | [optional] 
**AccuracySlaMin** | Pointer to **int32** | Typical minimum Accuracy for the package code. | [optional] 
**AccuracySlaMax** | Pointer to **int32** | Typical maximum Accuracy for the package code. | [optional] 
**Changelog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewPrecisionTimePackageResponse

`func NewPrecisionTimePackageResponse(type_ PrecisionTimePackageResponseType, code GetTimeServicesPackageByCodePackageCodeParameter, bandwidth int32, ) *PrecisionTimePackageResponse`

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


### GetCode

`func (o *PrecisionTimePackageResponse) GetCode() GetTimeServicesPackageByCodePackageCodeParameter`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PrecisionTimePackageResponse) GetCodeOk() (*GetTimeServicesPackageByCodePackageCodeParameter, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PrecisionTimePackageResponse) SetCode(v GetTimeServicesPackageByCodePackageCodeParameter)`

SetCode sets Code field to given value.


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

### GetAccuracySlaUnit

`func (o *PrecisionTimePackageResponse) GetAccuracySlaUnit() string`

GetAccuracySlaUnit returns the AccuracySlaUnit field if non-nil, zero value otherwise.

### GetAccuracySlaUnitOk

`func (o *PrecisionTimePackageResponse) GetAccuracySlaUnitOk() (*string, bool)`

GetAccuracySlaUnitOk returns a tuple with the AccuracySlaUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracySlaUnit

`func (o *PrecisionTimePackageResponse) SetAccuracySlaUnit(v string)`

SetAccuracySlaUnit sets AccuracySlaUnit field to given value.

### HasAccuracySlaUnit

`func (o *PrecisionTimePackageResponse) HasAccuracySlaUnit() bool`

HasAccuracySlaUnit returns a boolean if a field has been set.

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

### GetAccuracySlaMin

`func (o *PrecisionTimePackageResponse) GetAccuracySlaMin() int32`

GetAccuracySlaMin returns the AccuracySlaMin field if non-nil, zero value otherwise.

### GetAccuracySlaMinOk

`func (o *PrecisionTimePackageResponse) GetAccuracySlaMinOk() (*int32, bool)`

GetAccuracySlaMinOk returns a tuple with the AccuracySlaMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracySlaMin

`func (o *PrecisionTimePackageResponse) SetAccuracySlaMin(v int32)`

SetAccuracySlaMin sets AccuracySlaMin field to given value.

### HasAccuracySlaMin

`func (o *PrecisionTimePackageResponse) HasAccuracySlaMin() bool`

HasAccuracySlaMin returns a boolean if a field has been set.

### GetAccuracySlaMax

`func (o *PrecisionTimePackageResponse) GetAccuracySlaMax() int32`

GetAccuracySlaMax returns the AccuracySlaMax field if non-nil, zero value otherwise.

### GetAccuracySlaMaxOk

`func (o *PrecisionTimePackageResponse) GetAccuracySlaMaxOk() (*int32, bool)`

GetAccuracySlaMaxOk returns a tuple with the AccuracySlaMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracySlaMax

`func (o *PrecisionTimePackageResponse) SetAccuracySlaMax(v int32)`

SetAccuracySlaMax sets AccuracySlaMax field to given value.

### HasAccuracySlaMax

`func (o *PrecisionTimePackageResponse) HasAccuracySlaMax() bool`

HasAccuracySlaMax returns a boolean if a field has been set.

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


