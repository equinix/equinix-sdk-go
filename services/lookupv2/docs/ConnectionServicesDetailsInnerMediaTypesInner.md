# ConnectionServicesDetailsInnerMediaTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**MediaTypes**](MediaTypes.md) |  | [optional] 
**ProtocolTypes** | Pointer to [**[]ConnectionServicesDetailsInnerMediaTypesInnerProtocolTypesInner**](ConnectionServicesDetailsInnerMediaTypesInnerProtocolTypesInner.md) |  | [optional] 
**CircuitCounts** | Pointer to **[]int32** | Intra-Faciltiy Cable (IFC) circuit count options available for the respective connection service. E.g. If the &#39;circuitCount&#39; is empty, it means there are no available IFC circuits. If &#39;3,6&#39; appears in the circuit count, it means that the IFC circuits options available are 3 circuits and 6 circuits. | [optional] 

## Methods

### NewConnectionServicesDetailsInnerMediaTypesInner

`func NewConnectionServicesDetailsInnerMediaTypesInner() *ConnectionServicesDetailsInnerMediaTypesInner`

NewConnectionServicesDetailsInnerMediaTypesInner instantiates a new ConnectionServicesDetailsInnerMediaTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionServicesDetailsInnerMediaTypesInnerWithDefaults

`func NewConnectionServicesDetailsInnerMediaTypesInnerWithDefaults() *ConnectionServicesDetailsInnerMediaTypesInner`

NewConnectionServicesDetailsInnerMediaTypesInnerWithDefaults instantiates a new ConnectionServicesDetailsInnerMediaTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectionServicesDetailsInnerMediaTypesInner) GetName() MediaTypes`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionServicesDetailsInnerMediaTypesInner) GetNameOk() (*MediaTypes, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionServicesDetailsInnerMediaTypesInner) SetName(v MediaTypes)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionServicesDetailsInnerMediaTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtocolTypes

`func (o *ConnectionServicesDetailsInnerMediaTypesInner) GetProtocolTypes() []ConnectionServicesDetailsInnerMediaTypesInnerProtocolTypesInner`

GetProtocolTypes returns the ProtocolTypes field if non-nil, zero value otherwise.

### GetProtocolTypesOk

`func (o *ConnectionServicesDetailsInnerMediaTypesInner) GetProtocolTypesOk() (*[]ConnectionServicesDetailsInnerMediaTypesInnerProtocolTypesInner, bool)`

GetProtocolTypesOk returns a tuple with the ProtocolTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolTypes

`func (o *ConnectionServicesDetailsInnerMediaTypesInner) SetProtocolTypes(v []ConnectionServicesDetailsInnerMediaTypesInnerProtocolTypesInner)`

SetProtocolTypes sets ProtocolTypes field to given value.

### HasProtocolTypes

`func (o *ConnectionServicesDetailsInnerMediaTypesInner) HasProtocolTypes() bool`

HasProtocolTypes returns a boolean if a field has been set.

### GetCircuitCounts

`func (o *ConnectionServicesDetailsInnerMediaTypesInner) GetCircuitCounts() []int32`

GetCircuitCounts returns the CircuitCounts field if non-nil, zero value otherwise.

### GetCircuitCountsOk

`func (o *ConnectionServicesDetailsInnerMediaTypesInner) GetCircuitCountsOk() (*[]int32, bool)`

GetCircuitCountsOk returns a tuple with the CircuitCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitCounts

`func (o *ConnectionServicesDetailsInnerMediaTypesInner) SetCircuitCounts(v []int32)`

SetCircuitCounts sets CircuitCounts field to given value.

### HasCircuitCounts

`func (o *ConnectionServicesDetailsInnerMediaTypesInner) HasCircuitCounts() bool`

HasCircuitCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


