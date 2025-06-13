# VPC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**VPCType**](VPCType.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**VPCState**](VPCState.md) |  | [optional] 

## Methods

### NewVPC

`func NewVPC() *VPC`

NewVPC instantiates a new VPC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPCWithDefaults

`func NewVPCWithDefaults() *VPC`

NewVPCWithDefaults instantiates a new VPC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VPC) GetType() VPCType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VPC) GetTypeOk() (*VPCType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VPC) SetType(v VPCType)`

SetType sets Type field to given value.

### HasType

`func (o *VPC) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *VPC) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VPC) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VPC) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VPC) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *VPC) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPC) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPC) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPC) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *VPC) GetState() VPCState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VPC) GetStateOk() (*VPCState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VPC) SetState(v VPCState)`

SetState sets State field to given value.

### HasState

`func (o *VPC) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


