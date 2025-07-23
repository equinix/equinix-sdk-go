# OrderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrawCapacity** | **float64** | Maximum, combined power draw of all cabinets in your order, measured in kVA. Applicable values, in 0.5 increments, depend on the IBX data center. | 
**FabricPort** | **bool** | Indicates if a single, primary Fabric port should be included and delivered to one of the ordered cabinets. | 
**NumberOfCabinets** | **int32** | The number of ordered cabinets. | 
**CabinetDimensions** | [**Dimensions**](Dimensions.md) |  | 
**Pdus** | **bool** | Indicates if an Equinix-recommended set of two PDUs, for single-phase circuit, per cabinet should be installed. | 

## Methods

### NewOrderItem

`func NewOrderItem(drawCapacity float64, fabricPort bool, numberOfCabinets int32, cabinetDimensions Dimensions, pdus bool, ) *OrderItem`

NewOrderItem instantiates a new OrderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemWithDefaults

`func NewOrderItemWithDefaults() *OrderItem`

NewOrderItemWithDefaults instantiates a new OrderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrawCapacity

`func (o *OrderItem) GetDrawCapacity() float64`

GetDrawCapacity returns the DrawCapacity field if non-nil, zero value otherwise.

### GetDrawCapacityOk

`func (o *OrderItem) GetDrawCapacityOk() (*float64, bool)`

GetDrawCapacityOk returns a tuple with the DrawCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawCapacity

`func (o *OrderItem) SetDrawCapacity(v float64)`

SetDrawCapacity sets DrawCapacity field to given value.


### GetFabricPort

`func (o *OrderItem) GetFabricPort() bool`

GetFabricPort returns the FabricPort field if non-nil, zero value otherwise.

### GetFabricPortOk

`func (o *OrderItem) GetFabricPortOk() (*bool, bool)`

GetFabricPortOk returns a tuple with the FabricPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricPort

`func (o *OrderItem) SetFabricPort(v bool)`

SetFabricPort sets FabricPort field to given value.


### GetNumberOfCabinets

`func (o *OrderItem) GetNumberOfCabinets() int32`

GetNumberOfCabinets returns the NumberOfCabinets field if non-nil, zero value otherwise.

### GetNumberOfCabinetsOk

`func (o *OrderItem) GetNumberOfCabinetsOk() (*int32, bool)`

GetNumberOfCabinetsOk returns a tuple with the NumberOfCabinets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCabinets

`func (o *OrderItem) SetNumberOfCabinets(v int32)`

SetNumberOfCabinets sets NumberOfCabinets field to given value.


### GetCabinetDimensions

`func (o *OrderItem) GetCabinetDimensions() Dimensions`

GetCabinetDimensions returns the CabinetDimensions field if non-nil, zero value otherwise.

### GetCabinetDimensionsOk

`func (o *OrderItem) GetCabinetDimensionsOk() (*Dimensions, bool)`

GetCabinetDimensionsOk returns a tuple with the CabinetDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinetDimensions

`func (o *OrderItem) SetCabinetDimensions(v Dimensions)`

SetCabinetDimensions sets CabinetDimensions field to given value.


### GetPdus

`func (o *OrderItem) GetPdus() bool`

GetPdus returns the Pdus field if non-nil, zero value otherwise.

### GetPdusOk

`func (o *OrderItem) GetPdusOk() (*bool, bool)`

GetPdusOk returns a tuple with the Pdus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdus

`func (o *OrderItem) SetPdus(v bool)`

SetPdus sets Pdus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


