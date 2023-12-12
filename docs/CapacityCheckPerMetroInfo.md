# CapacityCheckPerMetroInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Returns true if there is enough capacity in the metro to fulfill the quantity set. Returns false if there is not enough. | [optional] 
**Metro** | Pointer to **string** | The metro ID or code sent to check capacity. | [optional] 
**Plan** | Pointer to **string** | The plan ID or slug sent to check capacity. | [optional] 
**Quantity** | Pointer to **string** | The number of servers sent to check capacity. | [optional] 

## Methods

### NewCapacityCheckPerMetroInfo

`func NewCapacityCheckPerMetroInfo() *CapacityCheckPerMetroInfo`

NewCapacityCheckPerMetroInfo instantiates a new CapacityCheckPerMetroInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityCheckPerMetroInfoWithDefaults

`func NewCapacityCheckPerMetroInfoWithDefaults() *CapacityCheckPerMetroInfo`

NewCapacityCheckPerMetroInfoWithDefaults instantiates a new CapacityCheckPerMetroInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *CapacityCheckPerMetroInfo) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CapacityCheckPerMetroInfo) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CapacityCheckPerMetroInfo) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CapacityCheckPerMetroInfo) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetMetro

`func (o *CapacityCheckPerMetroInfo) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *CapacityCheckPerMetroInfo) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *CapacityCheckPerMetroInfo) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *CapacityCheckPerMetroInfo) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetPlan

`func (o *CapacityCheckPerMetroInfo) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CapacityCheckPerMetroInfo) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CapacityCheckPerMetroInfo) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CapacityCheckPerMetroInfo) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetQuantity

`func (o *CapacityCheckPerMetroInfo) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CapacityCheckPerMetroInfo) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CapacityCheckPerMetroInfo) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CapacityCheckPerMetroInfo) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


