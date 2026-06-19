# ServiceProfileLastMileAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Require** | Pointer to **bool** | Whether the address is required for provisioning and ordering. | [optional] 
**DataType** | Pointer to **string** | Expected format of the address. | [optional] 
**Description** | Pointer to **string** | Additional information about the address requirement. | [optional] 

## Methods

### NewServiceProfileLastMileAddress

`func NewServiceProfileLastMileAddress() *ServiceProfileLastMileAddress`

NewServiceProfileLastMileAddress instantiates a new ServiceProfileLastMileAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileLastMileAddressWithDefaults

`func NewServiceProfileLastMileAddressWithDefaults() *ServiceProfileLastMileAddress`

NewServiceProfileLastMileAddressWithDefaults instantiates a new ServiceProfileLastMileAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequire

`func (o *ServiceProfileLastMileAddress) GetRequire() bool`

GetRequire returns the Require field if non-nil, zero value otherwise.

### GetRequireOk

`func (o *ServiceProfileLastMileAddress) GetRequireOk() (*bool, bool)`

GetRequireOk returns a tuple with the Require field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequire

`func (o *ServiceProfileLastMileAddress) SetRequire(v bool)`

SetRequire sets Require field to given value.

### HasRequire

`func (o *ServiceProfileLastMileAddress) HasRequire() bool`

HasRequire returns a boolean if a field has been set.

### GetDataType

`func (o *ServiceProfileLastMileAddress) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ServiceProfileLastMileAddress) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ServiceProfileLastMileAddress) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *ServiceProfileLastMileAddress) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceProfileLastMileAddress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceProfileLastMileAddress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceProfileLastMileAddress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceProfileLastMileAddress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


