# VirtualConnectionTimeServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | Connection URI | [readonly] 
**Type** | **string** | Connection Type. | 
**Uuid** | **string** | Connection UUID. | 
**ASide** | Pointer to [**VirtualConnectionSide**](VirtualConnectionSide.md) |  | [optional] 
**ZSide** | Pointer to [**VirtualConnectionSide**](VirtualConnectionSide.md) |  | [optional] 

## Methods

### NewVirtualConnectionTimeServiceResponse

`func NewVirtualConnectionTimeServiceResponse(href string, type_ string, uuid string, ) *VirtualConnectionTimeServiceResponse`

NewVirtualConnectionTimeServiceResponse instantiates a new VirtualConnectionTimeServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualConnectionTimeServiceResponseWithDefaults

`func NewVirtualConnectionTimeServiceResponseWithDefaults() *VirtualConnectionTimeServiceResponse`

NewVirtualConnectionTimeServiceResponseWithDefaults instantiates a new VirtualConnectionTimeServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *VirtualConnectionTimeServiceResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VirtualConnectionTimeServiceResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VirtualConnectionTimeServiceResponse) SetHref(v string)`

SetHref sets Href field to given value.


### GetType

`func (o *VirtualConnectionTimeServiceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualConnectionTimeServiceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualConnectionTimeServiceResponse) SetType(v string)`

SetType sets Type field to given value.


### GetUuid

`func (o *VirtualConnectionTimeServiceResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualConnectionTimeServiceResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualConnectionTimeServiceResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetASide

`func (o *VirtualConnectionTimeServiceResponse) GetASide() VirtualConnectionSide`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *VirtualConnectionTimeServiceResponse) GetASideOk() (*VirtualConnectionSide, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *VirtualConnectionTimeServiceResponse) SetASide(v VirtualConnectionSide)`

SetASide sets ASide field to given value.

### HasASide

`func (o *VirtualConnectionTimeServiceResponse) HasASide() bool`

HasASide returns a boolean if a field has been set.

### GetZSide

`func (o *VirtualConnectionTimeServiceResponse) GetZSide() VirtualConnectionSide`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *VirtualConnectionTimeServiceResponse) GetZSideOk() (*VirtualConnectionSide, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *VirtualConnectionTimeServiceResponse) SetZSide(v VirtualConnectionSide)`

SetZSide sets ZSide field to given value.

### HasZSide

`func (o *VirtualConnectionTimeServiceResponse) HasZSide() bool`

HasZSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


