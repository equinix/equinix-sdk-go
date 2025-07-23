# ProductsAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ibx** | **string** | IBX data center identifier. | 
**MaximumNumberOfCabinetsToOrder** | **int32** | Maximum number of cabinets that can be ordered in one order in this ibx. | 
**MinimumDrawCapacityPerCabinet** | **float64** | The minimum power draw capacity per ordered cabinet. | 
**MaximumDrawCapacityPerCabinet** | **float64** | The maximum power draw capacity per ordered cabinet. | 
**CabinetDimensions** | [**Dimensions**](Dimensions.md) |  | 
**AcCircuitConfiguration** | [**AcCircuitConfig**](AcCircuitConfig.md) |  | 
**PduConfiguration** | Pointer to [**PduConfig**](PduConfig.md) |  | [optional] 
**FabricPortSpeed** | Pointer to [**FabricPortSpeed**](FabricPortSpeed.md) |  | [optional] 

## Methods

### NewProductsAvailability

`func NewProductsAvailability(ibx string, maximumNumberOfCabinetsToOrder int32, minimumDrawCapacityPerCabinet float64, maximumDrawCapacityPerCabinet float64, cabinetDimensions Dimensions, acCircuitConfiguration AcCircuitConfig, ) *ProductsAvailability`

NewProductsAvailability instantiates a new ProductsAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsAvailabilityWithDefaults

`func NewProductsAvailabilityWithDefaults() *ProductsAvailability`

NewProductsAvailabilityWithDefaults instantiates a new ProductsAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIbx

`func (o *ProductsAvailability) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *ProductsAvailability) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *ProductsAvailability) SetIbx(v string)`

SetIbx sets Ibx field to given value.


### GetMaximumNumberOfCabinetsToOrder

`func (o *ProductsAvailability) GetMaximumNumberOfCabinetsToOrder() int32`

GetMaximumNumberOfCabinetsToOrder returns the MaximumNumberOfCabinetsToOrder field if non-nil, zero value otherwise.

### GetMaximumNumberOfCabinetsToOrderOk

`func (o *ProductsAvailability) GetMaximumNumberOfCabinetsToOrderOk() (*int32, bool)`

GetMaximumNumberOfCabinetsToOrderOk returns a tuple with the MaximumNumberOfCabinetsToOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNumberOfCabinetsToOrder

`func (o *ProductsAvailability) SetMaximumNumberOfCabinetsToOrder(v int32)`

SetMaximumNumberOfCabinetsToOrder sets MaximumNumberOfCabinetsToOrder field to given value.


### GetMinimumDrawCapacityPerCabinet

`func (o *ProductsAvailability) GetMinimumDrawCapacityPerCabinet() float64`

GetMinimumDrawCapacityPerCabinet returns the MinimumDrawCapacityPerCabinet field if non-nil, zero value otherwise.

### GetMinimumDrawCapacityPerCabinetOk

`func (o *ProductsAvailability) GetMinimumDrawCapacityPerCabinetOk() (*float64, bool)`

GetMinimumDrawCapacityPerCabinetOk returns a tuple with the MinimumDrawCapacityPerCabinet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumDrawCapacityPerCabinet

`func (o *ProductsAvailability) SetMinimumDrawCapacityPerCabinet(v float64)`

SetMinimumDrawCapacityPerCabinet sets MinimumDrawCapacityPerCabinet field to given value.


### GetMaximumDrawCapacityPerCabinet

`func (o *ProductsAvailability) GetMaximumDrawCapacityPerCabinet() float64`

GetMaximumDrawCapacityPerCabinet returns the MaximumDrawCapacityPerCabinet field if non-nil, zero value otherwise.

### GetMaximumDrawCapacityPerCabinetOk

`func (o *ProductsAvailability) GetMaximumDrawCapacityPerCabinetOk() (*float64, bool)`

GetMaximumDrawCapacityPerCabinetOk returns a tuple with the MaximumDrawCapacityPerCabinet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDrawCapacityPerCabinet

`func (o *ProductsAvailability) SetMaximumDrawCapacityPerCabinet(v float64)`

SetMaximumDrawCapacityPerCabinet sets MaximumDrawCapacityPerCabinet field to given value.


### GetCabinetDimensions

`func (o *ProductsAvailability) GetCabinetDimensions() Dimensions`

GetCabinetDimensions returns the CabinetDimensions field if non-nil, zero value otherwise.

### GetCabinetDimensionsOk

`func (o *ProductsAvailability) GetCabinetDimensionsOk() (*Dimensions, bool)`

GetCabinetDimensionsOk returns a tuple with the CabinetDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinetDimensions

`func (o *ProductsAvailability) SetCabinetDimensions(v Dimensions)`

SetCabinetDimensions sets CabinetDimensions field to given value.


### GetAcCircuitConfiguration

`func (o *ProductsAvailability) GetAcCircuitConfiguration() AcCircuitConfig`

GetAcCircuitConfiguration returns the AcCircuitConfiguration field if non-nil, zero value otherwise.

### GetAcCircuitConfigurationOk

`func (o *ProductsAvailability) GetAcCircuitConfigurationOk() (*AcCircuitConfig, bool)`

GetAcCircuitConfigurationOk returns a tuple with the AcCircuitConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcCircuitConfiguration

`func (o *ProductsAvailability) SetAcCircuitConfiguration(v AcCircuitConfig)`

SetAcCircuitConfiguration sets AcCircuitConfiguration field to given value.


### GetPduConfiguration

`func (o *ProductsAvailability) GetPduConfiguration() PduConfig`

GetPduConfiguration returns the PduConfiguration field if non-nil, zero value otherwise.

### GetPduConfigurationOk

`func (o *ProductsAvailability) GetPduConfigurationOk() (*PduConfig, bool)`

GetPduConfigurationOk returns a tuple with the PduConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduConfiguration

`func (o *ProductsAvailability) SetPduConfiguration(v PduConfig)`

SetPduConfiguration sets PduConfiguration field to given value.

### HasPduConfiguration

`func (o *ProductsAvailability) HasPduConfiguration() bool`

HasPduConfiguration returns a boolean if a field has been set.

### GetFabricPortSpeed

`func (o *ProductsAvailability) GetFabricPortSpeed() FabricPortSpeed`

GetFabricPortSpeed returns the FabricPortSpeed field if non-nil, zero value otherwise.

### GetFabricPortSpeedOk

`func (o *ProductsAvailability) GetFabricPortSpeedOk() (*FabricPortSpeed, bool)`

GetFabricPortSpeedOk returns a tuple with the FabricPortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricPortSpeed

`func (o *ProductsAvailability) SetFabricPortSpeed(v FabricPortSpeed)`

SetFabricPortSpeed sets FabricPortSpeed field to given value.

### HasFabricPortSpeed

`func (o *ProductsAvailability) HasFabricPortSpeed() bool`

HasFabricPortSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


