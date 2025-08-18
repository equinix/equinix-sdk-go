# MoveJumperCableRequestServiceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | [**MoveJumperCableRequestServiceDetailsQuantity**](MoveJumperCableRequestServiceDetailsQuantity.md) |  | 
**CableId** | Pointer to **string** | Enter Cable ID or ‘None’ if not applicable. This field is mandatory if &#39;Quantity&#39; is 1. | [optional] 
**CurrentDeviceDetails** | Pointer to [**Device**](Device.md) |  | [optional] 
**NewDeviceDetails** | Pointer to [**Device**](Device.md) |  | [optional] 
**ScopeOfWork** | **string** | Enter any additional details that will help our technicians execute your request. You may also attach your scope of work as a document if you exceed the character limit in this field. | 

## Methods

### NewMoveJumperCableRequestServiceDetails

`func NewMoveJumperCableRequestServiceDetails(quantity MoveJumperCableRequestServiceDetailsQuantity, scopeOfWork string, ) *MoveJumperCableRequestServiceDetails`

NewMoveJumperCableRequestServiceDetails instantiates a new MoveJumperCableRequestServiceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveJumperCableRequestServiceDetailsWithDefaults

`func NewMoveJumperCableRequestServiceDetailsWithDefaults() *MoveJumperCableRequestServiceDetails`

NewMoveJumperCableRequestServiceDetailsWithDefaults instantiates a new MoveJumperCableRequestServiceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *MoveJumperCableRequestServiceDetails) GetQuantity() MoveJumperCableRequestServiceDetailsQuantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MoveJumperCableRequestServiceDetails) GetQuantityOk() (*MoveJumperCableRequestServiceDetailsQuantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MoveJumperCableRequestServiceDetails) SetQuantity(v MoveJumperCableRequestServiceDetailsQuantity)`

SetQuantity sets Quantity field to given value.


### GetCableId

`func (o *MoveJumperCableRequestServiceDetails) GetCableId() string`

GetCableId returns the CableId field if non-nil, zero value otherwise.

### GetCableIdOk

`func (o *MoveJumperCableRequestServiceDetails) GetCableIdOk() (*string, bool)`

GetCableIdOk returns a tuple with the CableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableId

`func (o *MoveJumperCableRequestServiceDetails) SetCableId(v string)`

SetCableId sets CableId field to given value.

### HasCableId

`func (o *MoveJumperCableRequestServiceDetails) HasCableId() bool`

HasCableId returns a boolean if a field has been set.

### GetCurrentDeviceDetails

`func (o *MoveJumperCableRequestServiceDetails) GetCurrentDeviceDetails() Device`

GetCurrentDeviceDetails returns the CurrentDeviceDetails field if non-nil, zero value otherwise.

### GetCurrentDeviceDetailsOk

`func (o *MoveJumperCableRequestServiceDetails) GetCurrentDeviceDetailsOk() (*Device, bool)`

GetCurrentDeviceDetailsOk returns a tuple with the CurrentDeviceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDeviceDetails

`func (o *MoveJumperCableRequestServiceDetails) SetCurrentDeviceDetails(v Device)`

SetCurrentDeviceDetails sets CurrentDeviceDetails field to given value.

### HasCurrentDeviceDetails

`func (o *MoveJumperCableRequestServiceDetails) HasCurrentDeviceDetails() bool`

HasCurrentDeviceDetails returns a boolean if a field has been set.

### GetNewDeviceDetails

`func (o *MoveJumperCableRequestServiceDetails) GetNewDeviceDetails() Device`

GetNewDeviceDetails returns the NewDeviceDetails field if non-nil, zero value otherwise.

### GetNewDeviceDetailsOk

`func (o *MoveJumperCableRequestServiceDetails) GetNewDeviceDetailsOk() (*Device, bool)`

GetNewDeviceDetailsOk returns a tuple with the NewDeviceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDeviceDetails

`func (o *MoveJumperCableRequestServiceDetails) SetNewDeviceDetails(v Device)`

SetNewDeviceDetails sets NewDeviceDetails field to given value.

### HasNewDeviceDetails

`func (o *MoveJumperCableRequestServiceDetails) HasNewDeviceDetails() bool`

HasNewDeviceDetails returns a boolean if a field has been set.

### GetScopeOfWork

`func (o *MoveJumperCableRequestServiceDetails) GetScopeOfWork() string`

GetScopeOfWork returns the ScopeOfWork field if non-nil, zero value otherwise.

### GetScopeOfWorkOk

`func (o *MoveJumperCableRequestServiceDetails) GetScopeOfWorkOk() (*string, bool)`

GetScopeOfWorkOk returns a tuple with the ScopeOfWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOfWork

`func (o *MoveJumperCableRequestServiceDetails) SetScopeOfWork(v string)`

SetScopeOfWork sets ScopeOfWork field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


