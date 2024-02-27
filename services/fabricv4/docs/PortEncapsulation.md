# PortEncapsulation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**PortEncapsulationType**](PortEncapsulationType.md) |  | [optional] 
**TagProtocolId** | Pointer to **string** | Port encapsulation tag protocol identifier | [optional] 

## Methods

### NewPortEncapsulation

`func NewPortEncapsulation() *PortEncapsulation`

NewPortEncapsulation instantiates a new PortEncapsulation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortEncapsulationWithDefaults

`func NewPortEncapsulationWithDefaults() *PortEncapsulation`

NewPortEncapsulationWithDefaults instantiates a new PortEncapsulation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PortEncapsulation) GetType() PortEncapsulationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PortEncapsulation) GetTypeOk() (*PortEncapsulationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PortEncapsulation) SetType(v PortEncapsulationType)`

SetType sets Type field to given value.

### HasType

`func (o *PortEncapsulation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTagProtocolId

`func (o *PortEncapsulation) GetTagProtocolId() string`

GetTagProtocolId returns the TagProtocolId field if non-nil, zero value otherwise.

### GetTagProtocolIdOk

`func (o *PortEncapsulation) GetTagProtocolIdOk() (*string, bool)`

GetTagProtocolIdOk returns a tuple with the TagProtocolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagProtocolId

`func (o *PortEncapsulation) SetTagProtocolId(v string)`

SetTagProtocolId sets TagProtocolId field to given value.

### HasTagProtocolId

`func (o *PortEncapsulation) HasTagProtocolId() bool`

HasTagProtocolId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


