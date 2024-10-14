# PortOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationalStatus** | Pointer to [**PortOperationOperationalStatus**](PortOperationOperationalStatus.md) |  | [optional] 
**ConnectionCount** | Pointer to **int64** | Total number of connections. | [optional] 
**EvplVCCount** | Pointer to **int64** | Total number of connections. | [optional] 
**FgVCCount** | Pointer to **int64** | Total number of connections. | [optional] 
**AccessVCCount** | Pointer to **int64** | Total number of connections. | [optional] 
**OpStatusChangedAt** | Pointer to **time.Time** | Date and time at which port availability changed. | [optional] 

## Methods

### NewPortOperation

`func NewPortOperation() *PortOperation`

NewPortOperation instantiates a new PortOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortOperationWithDefaults

`func NewPortOperationWithDefaults() *PortOperation`

NewPortOperationWithDefaults instantiates a new PortOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationalStatus

`func (o *PortOperation) GetOperationalStatus() PortOperationOperationalStatus`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *PortOperation) GetOperationalStatusOk() (*PortOperationOperationalStatus, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *PortOperation) SetOperationalStatus(v PortOperationOperationalStatus)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *PortOperation) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetConnectionCount

`func (o *PortOperation) GetConnectionCount() int64`

GetConnectionCount returns the ConnectionCount field if non-nil, zero value otherwise.

### GetConnectionCountOk

`func (o *PortOperation) GetConnectionCountOk() (*int64, bool)`

GetConnectionCountOk returns a tuple with the ConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCount

`func (o *PortOperation) SetConnectionCount(v int64)`

SetConnectionCount sets ConnectionCount field to given value.

### HasConnectionCount

`func (o *PortOperation) HasConnectionCount() bool`

HasConnectionCount returns a boolean if a field has been set.

### GetEvplVCCount

`func (o *PortOperation) GetEvplVCCount() int64`

GetEvplVCCount returns the EvplVCCount field if non-nil, zero value otherwise.

### GetEvplVCCountOk

`func (o *PortOperation) GetEvplVCCountOk() (*int64, bool)`

GetEvplVCCountOk returns a tuple with the EvplVCCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvplVCCount

`func (o *PortOperation) SetEvplVCCount(v int64)`

SetEvplVCCount sets EvplVCCount field to given value.

### HasEvplVCCount

`func (o *PortOperation) HasEvplVCCount() bool`

HasEvplVCCount returns a boolean if a field has been set.

### GetFgVCCount

`func (o *PortOperation) GetFgVCCount() int64`

GetFgVCCount returns the FgVCCount field if non-nil, zero value otherwise.

### GetFgVCCountOk

`func (o *PortOperation) GetFgVCCountOk() (*int64, bool)`

GetFgVCCountOk returns a tuple with the FgVCCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFgVCCount

`func (o *PortOperation) SetFgVCCount(v int64)`

SetFgVCCount sets FgVCCount field to given value.

### HasFgVCCount

`func (o *PortOperation) HasFgVCCount() bool`

HasFgVCCount returns a boolean if a field has been set.

### GetAccessVCCount

`func (o *PortOperation) GetAccessVCCount() int64`

GetAccessVCCount returns the AccessVCCount field if non-nil, zero value otherwise.

### GetAccessVCCountOk

`func (o *PortOperation) GetAccessVCCountOk() (*int64, bool)`

GetAccessVCCountOk returns a tuple with the AccessVCCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessVCCount

`func (o *PortOperation) SetAccessVCCount(v int64)`

SetAccessVCCount sets AccessVCCount field to given value.

### HasAccessVCCount

`func (o *PortOperation) HasAccessVCCount() bool`

HasAccessVCCount returns a boolean if a field has been set.

### GetOpStatusChangedAt

`func (o *PortOperation) GetOpStatusChangedAt() time.Time`

GetOpStatusChangedAt returns the OpStatusChangedAt field if non-nil, zero value otherwise.

### GetOpStatusChangedAtOk

`func (o *PortOperation) GetOpStatusChangedAtOk() (*time.Time, bool)`

GetOpStatusChangedAtOk returns a tuple with the OpStatusChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpStatusChangedAt

`func (o *PortOperation) SetOpStatusChangedAt(v time.Time)`

SetOpStatusChangedAt sets OpStatusChangedAt field to given value.

### HasOpStatusChangedAt

`func (o *PortOperation) HasOpStatusChangedAt() bool`

HasOpStatusChangedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


