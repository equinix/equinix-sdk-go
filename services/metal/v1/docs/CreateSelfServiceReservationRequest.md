# CreateSelfServiceReservationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**[]SelfServiceReservationItemRequest**](SelfServiceReservationItemRequest.md) |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Period** | Pointer to [**CreateSelfServiceReservationRequestPeriod**](CreateSelfServiceReservationRequestPeriod.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateSelfServiceReservationRequest

`func NewCreateSelfServiceReservationRequest() *CreateSelfServiceReservationRequest`

NewCreateSelfServiceReservationRequest instantiates a new CreateSelfServiceReservationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSelfServiceReservationRequestWithDefaults

`func NewCreateSelfServiceReservationRequestWithDefaults() *CreateSelfServiceReservationRequest`

NewCreateSelfServiceReservationRequestWithDefaults instantiates a new CreateSelfServiceReservationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *CreateSelfServiceReservationRequest) GetItem() []SelfServiceReservationItemRequest`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *CreateSelfServiceReservationRequest) GetItemOk() (*[]SelfServiceReservationItemRequest, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *CreateSelfServiceReservationRequest) SetItem(v []SelfServiceReservationItemRequest)`

SetItem sets Item field to given value.

### HasItem

`func (o *CreateSelfServiceReservationRequest) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetNotes

`func (o *CreateSelfServiceReservationRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateSelfServiceReservationRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateSelfServiceReservationRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateSelfServiceReservationRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPeriod

`func (o *CreateSelfServiceReservationRequest) GetPeriod() CreateSelfServiceReservationRequestPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CreateSelfServiceReservationRequest) GetPeriodOk() (*CreateSelfServiceReservationRequestPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CreateSelfServiceReservationRequest) SetPeriod(v CreateSelfServiceReservationRequestPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CreateSelfServiceReservationRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateSelfServiceReservationRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateSelfServiceReservationRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateSelfServiceReservationRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateSelfServiceReservationRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


