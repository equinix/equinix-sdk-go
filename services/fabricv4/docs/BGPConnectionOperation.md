# BGPConnectionOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationalStatus** | Pointer to [**BGPConnectionOperationOperationalStatus**](BGPConnectionOperationOperationalStatus.md) |  | [optional] 
**OpStatusChangedAt** | Pointer to **time.Time** | Last BGP State Update by Date and Time | [optional] 

## Methods

### NewBGPConnectionOperation

`func NewBGPConnectionOperation() *BGPConnectionOperation`

NewBGPConnectionOperation instantiates a new BGPConnectionOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBGPConnectionOperationWithDefaults

`func NewBGPConnectionOperationWithDefaults() *BGPConnectionOperation`

NewBGPConnectionOperationWithDefaults instantiates a new BGPConnectionOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationalStatus

`func (o *BGPConnectionOperation) GetOperationalStatus() BGPConnectionOperationOperationalStatus`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *BGPConnectionOperation) GetOperationalStatusOk() (*BGPConnectionOperationOperationalStatus, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *BGPConnectionOperation) SetOperationalStatus(v BGPConnectionOperationOperationalStatus)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *BGPConnectionOperation) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetOpStatusChangedAt

`func (o *BGPConnectionOperation) GetOpStatusChangedAt() time.Time`

GetOpStatusChangedAt returns the OpStatusChangedAt field if non-nil, zero value otherwise.

### GetOpStatusChangedAtOk

`func (o *BGPConnectionOperation) GetOpStatusChangedAtOk() (*time.Time, bool)`

GetOpStatusChangedAtOk returns a tuple with the OpStatusChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpStatusChangedAt

`func (o *BGPConnectionOperation) SetOpStatusChangedAt(v time.Time)`

SetOpStatusChangedAt sets OpStatusChangedAt field to given value.

### HasOpStatusChangedAt

`func (o *BGPConnectionOperation) HasOpStatusChangedAt() bool`

HasOpStatusChangedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


