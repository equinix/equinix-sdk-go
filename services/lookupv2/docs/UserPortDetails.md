# UserPortDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortNumber** | Pointer to **float32** | Port number that is currently in use | [optional] 
**SerialNumber** | Pointer to **string** | Serial number or cable id | [optional] 
**ConnectionServicesName** | Pointer to [**ConnectionServices**](ConnectionServices.md) |  | [optional] 
**ZSideProviderName** | Pointer to **string** | Name of Z-Side service provider | [optional] 
**CircuitId** | Pointer to **string** |  | [optional] 

## Methods

### NewUserPortDetails

`func NewUserPortDetails() *UserPortDetails`

NewUserPortDetails instantiates a new UserPortDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPortDetailsWithDefaults

`func NewUserPortDetailsWithDefaults() *UserPortDetails`

NewUserPortDetailsWithDefaults instantiates a new UserPortDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortNumber

`func (o *UserPortDetails) GetPortNumber() float32`

GetPortNumber returns the PortNumber field if non-nil, zero value otherwise.

### GetPortNumberOk

`func (o *UserPortDetails) GetPortNumberOk() (*float32, bool)`

GetPortNumberOk returns a tuple with the PortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNumber

`func (o *UserPortDetails) SetPortNumber(v float32)`

SetPortNumber sets PortNumber field to given value.

### HasPortNumber

`func (o *UserPortDetails) HasPortNumber() bool`

HasPortNumber returns a boolean if a field has been set.

### GetSerialNumber

`func (o *UserPortDetails) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *UserPortDetails) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *UserPortDetails) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *UserPortDetails) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetConnectionServicesName

`func (o *UserPortDetails) GetConnectionServicesName() ConnectionServices`

GetConnectionServicesName returns the ConnectionServicesName field if non-nil, zero value otherwise.

### GetConnectionServicesNameOk

`func (o *UserPortDetails) GetConnectionServicesNameOk() (*ConnectionServices, bool)`

GetConnectionServicesNameOk returns a tuple with the ConnectionServicesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionServicesName

`func (o *UserPortDetails) SetConnectionServicesName(v ConnectionServices)`

SetConnectionServicesName sets ConnectionServicesName field to given value.

### HasConnectionServicesName

`func (o *UserPortDetails) HasConnectionServicesName() bool`

HasConnectionServicesName returns a boolean if a field has been set.

### GetZSideProviderName

`func (o *UserPortDetails) GetZSideProviderName() string`

GetZSideProviderName returns the ZSideProviderName field if non-nil, zero value otherwise.

### GetZSideProviderNameOk

`func (o *UserPortDetails) GetZSideProviderNameOk() (*string, bool)`

GetZSideProviderNameOk returns a tuple with the ZSideProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSideProviderName

`func (o *UserPortDetails) SetZSideProviderName(v string)`

SetZSideProviderName sets ZSideProviderName field to given value.

### HasZSideProviderName

`func (o *UserPortDetails) HasZSideProviderName() bool`

HasZSideProviderName returns a boolean if a field has been set.

### GetCircuitId

`func (o *UserPortDetails) GetCircuitId() string`

GetCircuitId returns the CircuitId field if non-nil, zero value otherwise.

### GetCircuitIdOk

`func (o *UserPortDetails) GetCircuitIdOk() (*string, bool)`

GetCircuitIdOk returns a tuple with the CircuitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitId

`func (o *UserPortDetails) SetCircuitId(v string)`

SetCircuitId sets CircuitId field to given value.

### HasCircuitId

`func (o *UserPortDetails) HasCircuitId() bool`

HasCircuitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


