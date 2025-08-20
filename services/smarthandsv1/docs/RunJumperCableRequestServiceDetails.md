# RunJumperCableRequestServiceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | [**MoveJumperCableRequestServiceDetailsQuantity**](MoveJumperCableRequestServiceDetailsQuantity.md) |  | 
**JumperType** | Pointer to [**RunJumperCableRequestServiceDetailsJumperType**](RunJumperCableRequestServiceDetailsJumperType.md) |  | [optional] 
**MediaType** | Pointer to [**RunJumperCableRequestServiceDetailsMediaType**](RunJumperCableRequestServiceDetailsMediaType.md) |  | [optional] 
**Connector** | Pointer to [**RunJumperCableRequestServiceDetailsConnector**](RunJumperCableRequestServiceDetailsConnector.md) |  | [optional] 
**CableId** | Pointer to **string** | Cable ID | [optional] 
**ProvideTxRxLightLevels** | Pointer to **bool** | Provide Tx/Rx Light Levels | [optional] 
**DeviceDetails** | Pointer to [**[]Device**](Device.md) |  | [optional] 
**ScopeOfWork** | **string** | Enter any additional details that will help our technicians execute your request. You may also attach your scope of work as a document if you exceed the character limit in this field. | 

## Methods

### NewRunJumperCableRequestServiceDetails

`func NewRunJumperCableRequestServiceDetails(quantity MoveJumperCableRequestServiceDetailsQuantity, scopeOfWork string, ) *RunJumperCableRequestServiceDetails`

NewRunJumperCableRequestServiceDetails instantiates a new RunJumperCableRequestServiceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunJumperCableRequestServiceDetailsWithDefaults

`func NewRunJumperCableRequestServiceDetailsWithDefaults() *RunJumperCableRequestServiceDetails`

NewRunJumperCableRequestServiceDetailsWithDefaults instantiates a new RunJumperCableRequestServiceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *RunJumperCableRequestServiceDetails) GetQuantity() MoveJumperCableRequestServiceDetailsQuantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *RunJumperCableRequestServiceDetails) GetQuantityOk() (*MoveJumperCableRequestServiceDetailsQuantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *RunJumperCableRequestServiceDetails) SetQuantity(v MoveJumperCableRequestServiceDetailsQuantity)`

SetQuantity sets Quantity field to given value.


### GetJumperType

`func (o *RunJumperCableRequestServiceDetails) GetJumperType() RunJumperCableRequestServiceDetailsJumperType`

GetJumperType returns the JumperType field if non-nil, zero value otherwise.

### GetJumperTypeOk

`func (o *RunJumperCableRequestServiceDetails) GetJumperTypeOk() (*RunJumperCableRequestServiceDetailsJumperType, bool)`

GetJumperTypeOk returns a tuple with the JumperType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumperType

`func (o *RunJumperCableRequestServiceDetails) SetJumperType(v RunJumperCableRequestServiceDetailsJumperType)`

SetJumperType sets JumperType field to given value.

### HasJumperType

`func (o *RunJumperCableRequestServiceDetails) HasJumperType() bool`

HasJumperType returns a boolean if a field has been set.

### GetMediaType

`func (o *RunJumperCableRequestServiceDetails) GetMediaType() RunJumperCableRequestServiceDetailsMediaType`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *RunJumperCableRequestServiceDetails) GetMediaTypeOk() (*RunJumperCableRequestServiceDetailsMediaType, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *RunJumperCableRequestServiceDetails) SetMediaType(v RunJumperCableRequestServiceDetailsMediaType)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *RunJumperCableRequestServiceDetails) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetConnector

`func (o *RunJumperCableRequestServiceDetails) GetConnector() RunJumperCableRequestServiceDetailsConnector`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *RunJumperCableRequestServiceDetails) GetConnectorOk() (*RunJumperCableRequestServiceDetailsConnector, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *RunJumperCableRequestServiceDetails) SetConnector(v RunJumperCableRequestServiceDetailsConnector)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *RunJumperCableRequestServiceDetails) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetCableId

`func (o *RunJumperCableRequestServiceDetails) GetCableId() string`

GetCableId returns the CableId field if non-nil, zero value otherwise.

### GetCableIdOk

`func (o *RunJumperCableRequestServiceDetails) GetCableIdOk() (*string, bool)`

GetCableIdOk returns a tuple with the CableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableId

`func (o *RunJumperCableRequestServiceDetails) SetCableId(v string)`

SetCableId sets CableId field to given value.

### HasCableId

`func (o *RunJumperCableRequestServiceDetails) HasCableId() bool`

HasCableId returns a boolean if a field has been set.

### GetProvideTxRxLightLevels

`func (o *RunJumperCableRequestServiceDetails) GetProvideTxRxLightLevels() bool`

GetProvideTxRxLightLevels returns the ProvideTxRxLightLevels field if non-nil, zero value otherwise.

### GetProvideTxRxLightLevelsOk

`func (o *RunJumperCableRequestServiceDetails) GetProvideTxRxLightLevelsOk() (*bool, bool)`

GetProvideTxRxLightLevelsOk returns a tuple with the ProvideTxRxLightLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvideTxRxLightLevels

`func (o *RunJumperCableRequestServiceDetails) SetProvideTxRxLightLevels(v bool)`

SetProvideTxRxLightLevels sets ProvideTxRxLightLevels field to given value.

### HasProvideTxRxLightLevels

`func (o *RunJumperCableRequestServiceDetails) HasProvideTxRxLightLevels() bool`

HasProvideTxRxLightLevels returns a boolean if a field has been set.

### GetDeviceDetails

`func (o *RunJumperCableRequestServiceDetails) GetDeviceDetails() []Device`

GetDeviceDetails returns the DeviceDetails field if non-nil, zero value otherwise.

### GetDeviceDetailsOk

`func (o *RunJumperCableRequestServiceDetails) GetDeviceDetailsOk() (*[]Device, bool)`

GetDeviceDetailsOk returns a tuple with the DeviceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDetails

`func (o *RunJumperCableRequestServiceDetails) SetDeviceDetails(v []Device)`

SetDeviceDetails sets DeviceDetails field to given value.

### HasDeviceDetails

`func (o *RunJumperCableRequestServiceDetails) HasDeviceDetails() bool`

HasDeviceDetails returns a boolean if a field has been set.

### GetScopeOfWork

`func (o *RunJumperCableRequestServiceDetails) GetScopeOfWork() string`

GetScopeOfWork returns the ScopeOfWork field if non-nil, zero value otherwise.

### GetScopeOfWorkOk

`func (o *RunJumperCableRequestServiceDetails) GetScopeOfWorkOk() (*string, bool)`

GetScopeOfWorkOk returns a tuple with the ScopeOfWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOfWork

`func (o *RunJumperCableRequestServiceDetails) SetScopeOfWork(v string)`

SetScopeOfWork sets ScopeOfWork field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


