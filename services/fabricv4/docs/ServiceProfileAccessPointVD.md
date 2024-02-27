# ServiceProfileAccessPointVD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ServiceProfileAccessPointVDType**](ServiceProfileAccessPointVDType.md) |  | 
**Uuid** | **string** |  | 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 
**InterfaceUuid** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceProfileAccessPointVD

`func NewServiceProfileAccessPointVD(type_ ServiceProfileAccessPointVDType, uuid string, ) *ServiceProfileAccessPointVD`

NewServiceProfileAccessPointVD instantiates a new ServiceProfileAccessPointVD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileAccessPointVDWithDefaults

`func NewServiceProfileAccessPointVDWithDefaults() *ServiceProfileAccessPointVD`

NewServiceProfileAccessPointVDWithDefaults instantiates a new ServiceProfileAccessPointVD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceProfileAccessPointVD) GetType() ServiceProfileAccessPointVDType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceProfileAccessPointVD) GetTypeOk() (*ServiceProfileAccessPointVDType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceProfileAccessPointVD) SetType(v ServiceProfileAccessPointVDType)`

SetType sets Type field to given value.


### GetUuid

`func (o *ServiceProfileAccessPointVD) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceProfileAccessPointVD) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceProfileAccessPointVD) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetLocation

`func (o *ServiceProfileAccessPointVD) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ServiceProfileAccessPointVD) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ServiceProfileAccessPointVD) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ServiceProfileAccessPointVD) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetInterfaceUuid

`func (o *ServiceProfileAccessPointVD) GetInterfaceUuid() string`

GetInterfaceUuid returns the InterfaceUuid field if non-nil, zero value otherwise.

### GetInterfaceUuidOk

`func (o *ServiceProfileAccessPointVD) GetInterfaceUuidOk() (*string, bool)`

GetInterfaceUuidOk returns a tuple with the InterfaceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceUuid

`func (o *ServiceProfileAccessPointVD) SetInterfaceUuid(v string)`

SetInterfaceUuid sets InterfaceUuid field to given value.

### HasInterfaceUuid

`func (o *ServiceProfileAccessPointVD) HasInterfaceUuid() bool`

HasInterfaceUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


