# CreateAsnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | [**Project**](Project.md) |  | 
**Number** | Pointer to **int64** | ASN number used for public ASN creation. | [optional] 
**RirRegion** | Pointer to [**AsnRirRegion**](AsnRirRegion.md) |  | [optional] 
**OwnershipChallengeDetails** | Pointer to [**AsnOwnershipChallengeDetails**](AsnOwnershipChallengeDetails.md) |  | [optional] 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 
**Range** | Pointer to [**AsnRange**](AsnRange.md) |  | [optional] 

## Methods

### NewCreateAsnRequest

`func NewCreateAsnRequest(project Project, ) *CreateAsnRequest`

NewCreateAsnRequest instantiates a new CreateAsnRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAsnRequestWithDefaults

`func NewCreateAsnRequestWithDefaults() *CreateAsnRequest`

NewCreateAsnRequestWithDefaults instantiates a new CreateAsnRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *CreateAsnRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CreateAsnRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CreateAsnRequest) SetProject(v Project)`

SetProject sets Project field to given value.


### GetNumber

`func (o *CreateAsnRequest) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CreateAsnRequest) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CreateAsnRequest) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CreateAsnRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetRirRegion

`func (o *CreateAsnRequest) GetRirRegion() AsnRirRegion`

GetRirRegion returns the RirRegion field if non-nil, zero value otherwise.

### GetRirRegionOk

`func (o *CreateAsnRequest) GetRirRegionOk() (*AsnRirRegion, bool)`

GetRirRegionOk returns a tuple with the RirRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRirRegion

`func (o *CreateAsnRequest) SetRirRegion(v AsnRirRegion)`

SetRirRegion sets RirRegion field to given value.

### HasRirRegion

`func (o *CreateAsnRequest) HasRirRegion() bool`

HasRirRegion returns a boolean if a field has been set.

### GetOwnershipChallengeDetails

`func (o *CreateAsnRequest) GetOwnershipChallengeDetails() AsnOwnershipChallengeDetails`

GetOwnershipChallengeDetails returns the OwnershipChallengeDetails field if non-nil, zero value otherwise.

### GetOwnershipChallengeDetailsOk

`func (o *CreateAsnRequest) GetOwnershipChallengeDetailsOk() (*AsnOwnershipChallengeDetails, bool)`

GetOwnershipChallengeDetailsOk returns a tuple with the OwnershipChallengeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipChallengeDetails

`func (o *CreateAsnRequest) SetOwnershipChallengeDetails(v AsnOwnershipChallengeDetails)`

SetOwnershipChallengeDetails sets OwnershipChallengeDetails field to given value.

### HasOwnershipChallengeDetails

`func (o *CreateAsnRequest) HasOwnershipChallengeDetails() bool`

HasOwnershipChallengeDetails returns a boolean if a field has been set.

### GetLocation

`func (o *CreateAsnRequest) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateAsnRequest) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateAsnRequest) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CreateAsnRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRange

`func (o *CreateAsnRequest) GetRange() AsnRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CreateAsnRequest) GetRangeOk() (*AsnRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CreateAsnRequest) SetRange(v AsnRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *CreateAsnRequest) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


