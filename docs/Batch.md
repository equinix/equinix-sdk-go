# Batch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Devices** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**ErrorMessages** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**Href**](Href.md) |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBatch

`func NewBatch() *Batch`

NewBatch instantiates a new Batch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchWithDefaults

`func NewBatchWithDefaults() *Batch`

NewBatchWithDefaults instantiates a new Batch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Batch) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Batch) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Batch) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Batch) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDevices

`func (o *Batch) GetDevices() []Href`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *Batch) GetDevicesOk() (*[]Href, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *Batch) SetDevices(v []Href)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *Batch) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetErrorMessages

`func (o *Batch) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *Batch) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *Batch) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *Batch) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetId

`func (o *Batch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Batch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Batch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Batch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *Batch) GetProject() Href`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Batch) GetProjectOk() (*Href, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Batch) SetProject(v Href)`

SetProject sets Project field to given value.

### HasProject

`func (o *Batch) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetQuantity

`func (o *Batch) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Batch) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Batch) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Batch) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetState

`func (o *Batch) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Batch) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Batch) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Batch) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Batch) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Batch) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Batch) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Batch) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


