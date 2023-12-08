# PortVlanAssignmentBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ErrorMessages** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Port** | Pointer to [**Port**](Port.md) |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**State** | Pointer to [**PortVlanAssignmentBatchState**](PortVlanAssignmentBatchState.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**VlanAssignments** | Pointer to [**[]PortVlanAssignmentBatchVlanAssignmentsInner**](PortVlanAssignmentBatchVlanAssignmentsInner.md) |  | [optional] 
**Project** | Pointer to [**Href**](Href.md) |  | [optional] 

## Methods

### NewPortVlanAssignmentBatch

`func NewPortVlanAssignmentBatch() *PortVlanAssignmentBatch`

NewPortVlanAssignmentBatch instantiates a new PortVlanAssignmentBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortVlanAssignmentBatchWithDefaults

`func NewPortVlanAssignmentBatchWithDefaults() *PortVlanAssignmentBatch`

NewPortVlanAssignmentBatchWithDefaults instantiates a new PortVlanAssignmentBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PortVlanAssignmentBatch) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortVlanAssignmentBatch) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortVlanAssignmentBatch) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortVlanAssignmentBatch) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetErrorMessages

`func (o *PortVlanAssignmentBatch) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *PortVlanAssignmentBatch) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *PortVlanAssignmentBatch) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *PortVlanAssignmentBatch) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetId

`func (o *PortVlanAssignmentBatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortVlanAssignmentBatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortVlanAssignmentBatch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortVlanAssignmentBatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPort

`func (o *PortVlanAssignmentBatch) GetPort() Port`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PortVlanAssignmentBatch) GetPortOk() (*Port, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PortVlanAssignmentBatch) SetPort(v Port)`

SetPort sets Port field to given value.

### HasPort

`func (o *PortVlanAssignmentBatch) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetQuantity

`func (o *PortVlanAssignmentBatch) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PortVlanAssignmentBatch) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PortVlanAssignmentBatch) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PortVlanAssignmentBatch) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetState

`func (o *PortVlanAssignmentBatch) GetState() PortVlanAssignmentBatchState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PortVlanAssignmentBatch) GetStateOk() (*PortVlanAssignmentBatchState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PortVlanAssignmentBatch) SetState(v PortVlanAssignmentBatchState)`

SetState sets State field to given value.

### HasState

`func (o *PortVlanAssignmentBatch) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PortVlanAssignmentBatch) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortVlanAssignmentBatch) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortVlanAssignmentBatch) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PortVlanAssignmentBatch) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVlanAssignments

`func (o *PortVlanAssignmentBatch) GetVlanAssignments() []PortVlanAssignmentBatchVlanAssignmentsInner`

GetVlanAssignments returns the VlanAssignments field if non-nil, zero value otherwise.

### GetVlanAssignmentsOk

`func (o *PortVlanAssignmentBatch) GetVlanAssignmentsOk() (*[]PortVlanAssignmentBatchVlanAssignmentsInner, bool)`

GetVlanAssignmentsOk returns a tuple with the VlanAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanAssignments

`func (o *PortVlanAssignmentBatch) SetVlanAssignments(v []PortVlanAssignmentBatchVlanAssignmentsInner)`

SetVlanAssignments sets VlanAssignments field to given value.

### HasVlanAssignments

`func (o *PortVlanAssignmentBatch) HasVlanAssignments() bool`

HasVlanAssignments returns a boolean if a field has been set.

### GetProject

`func (o *PortVlanAssignmentBatch) GetProject() Href`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PortVlanAssignmentBatch) GetProjectOk() (*Href, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PortVlanAssignmentBatch) SetProject(v Href)`

SetProject sets Project field to given value.

### HasProject

`func (o *PortVlanAssignmentBatch) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


