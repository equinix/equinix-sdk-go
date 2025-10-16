# ServiceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** | The unique ID of the service. | [optional] 
**ServiceName** | Pointer to **string** | The name of the service. | [optional] 
**OperationNeededToPerform** | Pointer to **string** | The operation you must perform to restore the backup successfully. UNSUPPORTED- You cannot restore this backup. DELETE- You need to delete this service to restore the backup. NONE- You do not need to change anything to restore the backup. | [optional] 

## Methods

### NewServiceInfo

`func NewServiceInfo() *ServiceInfo`

NewServiceInfo instantiates a new ServiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInfoWithDefaults

`func NewServiceInfoWithDefaults() *ServiceInfo`

NewServiceInfoWithDefaults instantiates a new ServiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *ServiceInfo) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceInfo) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceInfo) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServiceInfo) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *ServiceInfo) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceInfo) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceInfo) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceInfo) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetOperationNeededToPerform

`func (o *ServiceInfo) GetOperationNeededToPerform() string`

GetOperationNeededToPerform returns the OperationNeededToPerform field if non-nil, zero value otherwise.

### GetOperationNeededToPerformOk

`func (o *ServiceInfo) GetOperationNeededToPerformOk() (*string, bool)`

GetOperationNeededToPerformOk returns a tuple with the OperationNeededToPerform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationNeededToPerform

`func (o *ServiceInfo) SetOperationNeededToPerform(v string)`

SetOperationNeededToPerform sets OperationNeededToPerform field to given value.

### HasOperationNeededToPerform

`func (o *ServiceInfo) HasOperationNeededToPerform() bool`

HasOperationNeededToPerform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


