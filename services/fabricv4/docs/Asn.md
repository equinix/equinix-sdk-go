# Asn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** | Unique identifier for the ASN record | [optional] 
**State** | Pointer to [**AsnState**](AsnState.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Number** | Pointer to **int64** | Autonomous System Number | [optional] 
**RirRegion** | Pointer to [**AsnRirRegion**](AsnRirRegion.md) |  | [optional] 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 
**Range** | Pointer to [**AsnRange**](AsnRange.md) |  | [optional] 
**OwnershipChallengeDetails** | Pointer to [**AsnOwnershipChallengeDetails**](AsnOwnershipChallengeDetails.md) |  | [optional] 
**Change** | Pointer to [**AsnChange**](AsnChange.md) |  | [optional] 
**Assets** | Pointer to [**[]AsnAsset**](AsnAsset.md) |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) | List of errors | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewAsn

`func NewAsn() *Asn`

NewAsn instantiates a new Asn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsnWithDefaults

`func NewAsnWithDefaults() *Asn`

NewAsnWithDefaults instantiates a new Asn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Asn) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Asn) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Asn) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Asn) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *Asn) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Asn) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Asn) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Asn) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *Asn) GetState() AsnState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Asn) GetStateOk() (*AsnState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Asn) SetState(v AsnState)`

SetState sets State field to given value.

### HasState

`func (o *Asn) HasState() bool`

HasState returns a boolean if a field has been set.

### GetProject

`func (o *Asn) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Asn) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Asn) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *Asn) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetNumber

`func (o *Asn) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Asn) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Asn) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Asn) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetRirRegion

`func (o *Asn) GetRirRegion() AsnRirRegion`

GetRirRegion returns the RirRegion field if non-nil, zero value otherwise.

### GetRirRegionOk

`func (o *Asn) GetRirRegionOk() (*AsnRirRegion, bool)`

GetRirRegionOk returns a tuple with the RirRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegion

`func (o *Asn) SetRirRegion(v AsnRirRegion)`

SetRirRegion sets RirRegion field to given value.

### HasRirRegion

`func (o *Asn) HasRirRegion() bool`

HasRirRegion returns a boolean if a field has been set.

### GetLocation

`func (o *Asn) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Asn) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Asn) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Asn) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRange

`func (o *Asn) GetRange() AsnRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Asn) GetRangeOk() (*AsnRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Asn) SetRange(v AsnRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *Asn) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetOwnershipChallengeDetails

`func (o *Asn) GetOwnershipChallengeDetails() AsnOwnershipChallengeDetails`

GetOwnershipChallengeDetails returns the OwnershipChallengeDetails field if non-nil, zero value otherwise.

### GetOwnershipChallengeDetailsOk

`func (o *Asn) GetOwnershipChallengeDetailsOk() (*AsnOwnershipChallengeDetails, bool)`

GetOwnershipChallengeDetailsOk returns a tuple with the OwnershipChallengeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipChallengeDetails

`func (o *Asn) SetOwnershipChallengeDetails(v AsnOwnershipChallengeDetails)`

SetOwnershipChallengeDetails sets OwnershipChallengeDetails field to given value.

### HasOwnershipChallengeDetails

`func (o *Asn) HasOwnershipChallengeDetails() bool`

HasOwnershipChallengeDetails returns a boolean if a field has been set.

### GetChange

`func (o *Asn) GetChange() AsnChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *Asn) GetChangeOk() (*AsnChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *Asn) SetChange(v AsnChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *Asn) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetAssets

`func (o *Asn) GetAssets() []AsnAsset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *Asn) GetAssetsOk() (*[]AsnAsset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *Asn) SetAssets(v []AsnAsset)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *Asn) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetErrors

`func (o *Asn) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Asn) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Asn) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Asn) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetChangeLog

`func (o *Asn) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *Asn) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *Asn) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *Asn) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


