# SelfServiceReservationItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetroId** | Pointer to **string** | Metro code or ID of where the Self Service Reservation Item should be provisioned in. | [optional] 
**PlanId** | Pointer to **string** | The plan slug of the Self Service Reservation Item. | [optional] 
**Quantity** | Pointer to **int32** | The number of Self Service Reservation Items. | [optional] 
**Term** | Pointer to **string** | The duration of Self Service Reservation Item. | [optional] 

## Methods

### NewSelfServiceReservationItemRequest

`func NewSelfServiceReservationItemRequest() *SelfServiceReservationItemRequest`

NewSelfServiceReservationItemRequest instantiates a new SelfServiceReservationItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfServiceReservationItemRequestWithDefaults

`func NewSelfServiceReservationItemRequestWithDefaults() *SelfServiceReservationItemRequest`

NewSelfServiceReservationItemRequestWithDefaults instantiates a new SelfServiceReservationItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetroId

`func (o *SelfServiceReservationItemRequest) GetMetroId() string`

GetMetroId returns the MetroId field if non-nil, zero value otherwise.

### GetMetroIdOk

`func (o *SelfServiceReservationItemRequest) GetMetroIdOk() (*string, bool)`

GetMetroIdOk returns a tuple with the MetroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroId

`func (o *SelfServiceReservationItemRequest) SetMetroId(v string)`

SetMetroId sets MetroId field to given value.

### HasMetroId

`func (o *SelfServiceReservationItemRequest) HasMetroId() bool`

HasMetroId returns a boolean if a field has been set.

### GetPlanId

`func (o *SelfServiceReservationItemRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SelfServiceReservationItemRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SelfServiceReservationItemRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *SelfServiceReservationItemRequest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetQuantity

`func (o *SelfServiceReservationItemRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SelfServiceReservationItemRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SelfServiceReservationItemRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SelfServiceReservationItemRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTerm

`func (o *SelfServiceReservationItemRequest) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *SelfServiceReservationItemRequest) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *SelfServiceReservationItemRequest) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *SelfServiceReservationItemRequest) HasTerm() bool`

HasTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


