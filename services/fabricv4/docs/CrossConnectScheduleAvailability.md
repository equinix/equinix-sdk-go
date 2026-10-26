# CrossConnectScheduleAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Schedule URI. | [optional] [readonly] 
**Type** | Pointer to [**CrossConnectScheduleAvailabilityType**](CrossConnectScheduleAvailabilityType.md) |  | [optional] 
**Code** | Pointer to [**ScheduleType**](ScheduleType.md) |  | [optional] [readonly] 
**ASide** | Pointer to [**CrossConnectScheduleAvailabilityASide**](CrossConnectScheduleAvailabilityASide.md) |  | [optional] 
**ZSide** | Pointer to [**CrossConnectScheduleAvailabilityZSide**](CrossConnectScheduleAvailabilityZSide.md) |  | [optional] 
**StartDateTime** | Pointer to **time.Time** | Start of the schedule window. | [optional] [readonly] 
**EndDateTime** | Pointer to **time.Time** | End of the schedule window. | [optional] [readonly] 
**UnavailableDates** | Pointer to **[]string** | Dates within the schedule window when the IBX is unavailable, such as the holidays observed at the IBX.  | [optional] [readonly] 

## Methods

### NewCrossConnectScheduleAvailability

`func NewCrossConnectScheduleAvailability() *CrossConnectScheduleAvailability`

NewCrossConnectScheduleAvailability instantiates a new CrossConnectScheduleAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectScheduleAvailabilityWithDefaults

`func NewCrossConnectScheduleAvailabilityWithDefaults() *CrossConnectScheduleAvailability`

NewCrossConnectScheduleAvailabilityWithDefaults instantiates a new CrossConnectScheduleAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CrossConnectScheduleAvailability) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CrossConnectScheduleAvailability) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CrossConnectScheduleAvailability) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CrossConnectScheduleAvailability) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *CrossConnectScheduleAvailability) GetType() CrossConnectScheduleAvailabilityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CrossConnectScheduleAvailability) GetTypeOk() (*CrossConnectScheduleAvailabilityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CrossConnectScheduleAvailability) SetType(v CrossConnectScheduleAvailabilityType)`

SetType sets Type field to given value.

### HasType

`func (o *CrossConnectScheduleAvailability) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *CrossConnectScheduleAvailability) GetCode() ScheduleType`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CrossConnectScheduleAvailability) GetCodeOk() (*ScheduleType, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CrossConnectScheduleAvailability) SetCode(v ScheduleType)`

SetCode sets Code field to given value.

### HasCode

`func (o *CrossConnectScheduleAvailability) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetASide

`func (o *CrossConnectScheduleAvailability) GetASide() CrossConnectScheduleAvailabilityASide`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *CrossConnectScheduleAvailability) GetASideOk() (*CrossConnectScheduleAvailabilityASide, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *CrossConnectScheduleAvailability) SetASide(v CrossConnectScheduleAvailabilityASide)`

SetASide sets ASide field to given value.

### HasASide

`func (o *CrossConnectScheduleAvailability) HasASide() bool`

HasASide returns a boolean if a field has been set.

### GetZSide

`func (o *CrossConnectScheduleAvailability) GetZSide() CrossConnectScheduleAvailabilityZSide`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *CrossConnectScheduleAvailability) GetZSideOk() (*CrossConnectScheduleAvailabilityZSide, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *CrossConnectScheduleAvailability) SetZSide(v CrossConnectScheduleAvailabilityZSide)`

SetZSide sets ZSide field to given value.

### HasZSide

`func (o *CrossConnectScheduleAvailability) HasZSide() bool`

HasZSide returns a boolean if a field has been set.

### GetStartDateTime

`func (o *CrossConnectScheduleAvailability) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *CrossConnectScheduleAvailability) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *CrossConnectScheduleAvailability) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *CrossConnectScheduleAvailability) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetEndDateTime

`func (o *CrossConnectScheduleAvailability) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *CrossConnectScheduleAvailability) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *CrossConnectScheduleAvailability) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *CrossConnectScheduleAvailability) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### GetUnavailableDates

`func (o *CrossConnectScheduleAvailability) GetUnavailableDates() []string`

GetUnavailableDates returns the UnavailableDates field if non-nil, zero value otherwise.

### GetUnavailableDatesOk

`func (o *CrossConnectScheduleAvailability) GetUnavailableDatesOk() (*[]string, bool)`

GetUnavailableDatesOk returns a tuple with the UnavailableDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableDates

`func (o *CrossConnectScheduleAvailability) SetUnavailableDates(v []string)`

SetUnavailableDates sets UnavailableDates field to given value.

### HasUnavailableDates

`func (o *CrossConnectScheduleAvailability) HasUnavailableDates() bool`

HasUnavailableDates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


