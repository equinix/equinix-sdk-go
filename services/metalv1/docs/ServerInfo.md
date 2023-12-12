# ServerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facility** | Pointer to **string** |  | [optional] 
**Metro** | Pointer to **string** | The metro ID or code to check the capacity in. | [optional] 
**Plan** | Pointer to **string** | The plan ID or slug to check the capacity of. | [optional] 
**Quantity** | Pointer to **string** | The number of servers to check the capacity of. | [optional] 

## Methods

### NewServerInfo

`func NewServerInfo() *ServerInfo`

NewServerInfo instantiates a new ServerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInfoWithDefaults

`func NewServerInfoWithDefaults() *ServerInfo`

NewServerInfoWithDefaults instantiates a new ServerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacility

`func (o *ServerInfo) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *ServerInfo) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *ServerInfo) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *ServerInfo) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetMetro

`func (o *ServerInfo) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *ServerInfo) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *ServerInfo) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *ServerInfo) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetPlan

`func (o *ServerInfo) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ServerInfo) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ServerInfo) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ServerInfo) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetQuantity

`func (o *ServerInfo) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ServerInfo) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ServerInfo) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ServerInfo) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


