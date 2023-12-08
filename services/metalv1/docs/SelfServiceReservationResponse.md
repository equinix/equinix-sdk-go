# SelfServiceReservationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Item** | Pointer to [**[]SelfServiceReservationItemResponse**](SelfServiceReservationItemResponse.md) |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Period** | Pointer to [**CreateSelfServiceReservationRequestPeriod**](CreateSelfServiceReservationRequestPeriod.md) |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalCost** | Pointer to **int32** |  | [optional] 

## Methods

### NewSelfServiceReservationResponse

`func NewSelfServiceReservationResponse() *SelfServiceReservationResponse`

NewSelfServiceReservationResponse instantiates a new SelfServiceReservationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfServiceReservationResponseWithDefaults

`func NewSelfServiceReservationResponseWithDefaults() *SelfServiceReservationResponse`

NewSelfServiceReservationResponseWithDefaults instantiates a new SelfServiceReservationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SelfServiceReservationResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SelfServiceReservationResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SelfServiceReservationResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SelfServiceReservationResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetItem

`func (o *SelfServiceReservationResponse) GetItem() []SelfServiceReservationItemResponse`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *SelfServiceReservationResponse) GetItemOk() (*[]SelfServiceReservationItemResponse, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *SelfServiceReservationResponse) SetItem(v []SelfServiceReservationItemResponse)`

SetItem sets Item field to given value.

### HasItem

`func (o *SelfServiceReservationResponse) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetNotes

`func (o *SelfServiceReservationResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SelfServiceReservationResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SelfServiceReservationResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SelfServiceReservationResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOrganization

`func (o *SelfServiceReservationResponse) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SelfServiceReservationResponse) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SelfServiceReservationResponse) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SelfServiceReservationResponse) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationId

`func (o *SelfServiceReservationResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SelfServiceReservationResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SelfServiceReservationResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SelfServiceReservationResponse) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPeriod

`func (o *SelfServiceReservationResponse) GetPeriod() CreateSelfServiceReservationRequestPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *SelfServiceReservationResponse) GetPeriodOk() (*CreateSelfServiceReservationRequestPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *SelfServiceReservationResponse) SetPeriod(v CreateSelfServiceReservationRequestPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *SelfServiceReservationResponse) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetProject

`func (o *SelfServiceReservationResponse) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SelfServiceReservationResponse) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SelfServiceReservationResponse) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SelfServiceReservationResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *SelfServiceReservationResponse) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SelfServiceReservationResponse) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SelfServiceReservationResponse) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *SelfServiceReservationResponse) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStartDate

`func (o *SelfServiceReservationResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SelfServiceReservationResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SelfServiceReservationResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SelfServiceReservationResponse) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *SelfServiceReservationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SelfServiceReservationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SelfServiceReservationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SelfServiceReservationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalCost

`func (o *SelfServiceReservationResponse) GetTotalCost() int32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *SelfServiceReservationResponse) GetTotalCostOk() (*int32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *SelfServiceReservationResponse) SetTotalCost(v int32)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *SelfServiceReservationResponse) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


