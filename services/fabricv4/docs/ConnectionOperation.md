# ConnectionOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderStatus** | Pointer to [**ProviderStatus**](ProviderStatus.md) |  | [optional] 
**EquinixStatus** | Pointer to [**EquinixStatus**](EquinixStatus.md) |  | [optional] 
**OperationalStatus** | Pointer to [**ConnectionOperationOperationalStatus**](ConnectionOperationOperationalStatus.md) |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) |  | [optional] 
**OpStatusChangedAt** | Pointer to **time.Time** | When connection transitioned into current operational status | [optional] 

## Methods

### NewConnectionOperation

`func NewConnectionOperation() *ConnectionOperation`

NewConnectionOperation instantiates a new ConnectionOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionOperationWithDefaults

`func NewConnectionOperationWithDefaults() *ConnectionOperation`

NewConnectionOperationWithDefaults instantiates a new ConnectionOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderStatus

`func (o *ConnectionOperation) GetProviderStatus() ProviderStatus`

GetProviderStatus returns the ProviderStatus field if non-nil, zero value otherwise.

### GetProviderStatusOk

`func (o *ConnectionOperation) GetProviderStatusOk() (*ProviderStatus, bool)`

GetProviderStatusOk returns a tuple with the ProviderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderStatus

`func (o *ConnectionOperation) SetProviderStatus(v ProviderStatus)`

SetProviderStatus sets ProviderStatus field to given value.

### HasProviderStatus

`func (o *ConnectionOperation) HasProviderStatus() bool`

HasProviderStatus returns a boolean if a field has been set.

### GetEquinixStatus

`func (o *ConnectionOperation) GetEquinixStatus() EquinixStatus`

GetEquinixStatus returns the EquinixStatus field if non-nil, zero value otherwise.

### GetEquinixStatusOk

`func (o *ConnectionOperation) GetEquinixStatusOk() (*EquinixStatus, bool)`

GetEquinixStatusOk returns a tuple with the EquinixStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixStatus

`func (o *ConnectionOperation) SetEquinixStatus(v EquinixStatus)`

SetEquinixStatus sets EquinixStatus field to given value.

### HasEquinixStatus

`func (o *ConnectionOperation) HasEquinixStatus() bool`

HasEquinixStatus returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *ConnectionOperation) GetOperationalStatus() ConnectionOperationOperationalStatus`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *ConnectionOperation) GetOperationalStatusOk() (*ConnectionOperationOperationalStatus, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *ConnectionOperation) SetOperationalStatus(v ConnectionOperationOperationalStatus)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *ConnectionOperation) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetErrors

`func (o *ConnectionOperation) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ConnectionOperation) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ConnectionOperation) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ConnectionOperation) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetOpStatusChangedAt

`func (o *ConnectionOperation) GetOpStatusChangedAt() time.Time`

GetOpStatusChangedAt returns the OpStatusChangedAt field if non-nil, zero value otherwise.

### GetOpStatusChangedAtOk

`func (o *ConnectionOperation) GetOpStatusChangedAtOk() (*time.Time, bool)`

GetOpStatusChangedAtOk returns a tuple with the OpStatusChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpStatusChangedAt

`func (o *ConnectionOperation) SetOpStatusChangedAt(v time.Time)`

SetOpStatusChangedAt sets OpStatusChangedAt field to given value.

### HasOpStatusChangedAt

`func (o *ConnectionOperation) HasOpStatusChangedAt() bool`

HasOpStatusChangedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


