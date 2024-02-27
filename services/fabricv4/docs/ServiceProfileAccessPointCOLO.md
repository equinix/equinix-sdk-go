# ServiceProfileAccessPointCOLO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ServiceProfileAccessPointCOLOType**](ServiceProfileAccessPointCOLOType.md) |  | 
**Uuid** | **string** |  | 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 
**SellerRegion** | Pointer to **string** |  | [optional] 
**SellerRegionDescription** | Pointer to **string** |  | [optional] 
**CrossConnectId** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceProfileAccessPointCOLO

`func NewServiceProfileAccessPointCOLO(type_ ServiceProfileAccessPointCOLOType, uuid string, ) *ServiceProfileAccessPointCOLO`

NewServiceProfileAccessPointCOLO instantiates a new ServiceProfileAccessPointCOLO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileAccessPointCOLOWithDefaults

`func NewServiceProfileAccessPointCOLOWithDefaults() *ServiceProfileAccessPointCOLO`

NewServiceProfileAccessPointCOLOWithDefaults instantiates a new ServiceProfileAccessPointCOLO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceProfileAccessPointCOLO) GetType() ServiceProfileAccessPointCOLOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceProfileAccessPointCOLO) GetTypeOk() (*ServiceProfileAccessPointCOLOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceProfileAccessPointCOLO) SetType(v ServiceProfileAccessPointCOLOType)`

SetType sets Type field to given value.


### GetUuid

`func (o *ServiceProfileAccessPointCOLO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceProfileAccessPointCOLO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceProfileAccessPointCOLO) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetLocation

`func (o *ServiceProfileAccessPointCOLO) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ServiceProfileAccessPointCOLO) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ServiceProfileAccessPointCOLO) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ServiceProfileAccessPointCOLO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSellerRegion

`func (o *ServiceProfileAccessPointCOLO) GetSellerRegion() string`

GetSellerRegion returns the SellerRegion field if non-nil, zero value otherwise.

### GetSellerRegionOk

`func (o *ServiceProfileAccessPointCOLO) GetSellerRegionOk() (*string, bool)`

GetSellerRegionOk returns a tuple with the SellerRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerRegion

`func (o *ServiceProfileAccessPointCOLO) SetSellerRegion(v string)`

SetSellerRegion sets SellerRegion field to given value.

### HasSellerRegion

`func (o *ServiceProfileAccessPointCOLO) HasSellerRegion() bool`

HasSellerRegion returns a boolean if a field has been set.

### GetSellerRegionDescription

`func (o *ServiceProfileAccessPointCOLO) GetSellerRegionDescription() string`

GetSellerRegionDescription returns the SellerRegionDescription field if non-nil, zero value otherwise.

### GetSellerRegionDescriptionOk

`func (o *ServiceProfileAccessPointCOLO) GetSellerRegionDescriptionOk() (*string, bool)`

GetSellerRegionDescriptionOk returns a tuple with the SellerRegionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerRegionDescription

`func (o *ServiceProfileAccessPointCOLO) SetSellerRegionDescription(v string)`

SetSellerRegionDescription sets SellerRegionDescription field to given value.

### HasSellerRegionDescription

`func (o *ServiceProfileAccessPointCOLO) HasSellerRegionDescription() bool`

HasSellerRegionDescription returns a boolean if a field has been set.

### GetCrossConnectId

`func (o *ServiceProfileAccessPointCOLO) GetCrossConnectId() string`

GetCrossConnectId returns the CrossConnectId field if non-nil, zero value otherwise.

### GetCrossConnectIdOk

`func (o *ServiceProfileAccessPointCOLO) GetCrossConnectIdOk() (*string, bool)`

GetCrossConnectIdOk returns a tuple with the CrossConnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossConnectId

`func (o *ServiceProfileAccessPointCOLO) SetCrossConnectId(v string)`

SetCrossConnectId sets CrossConnectId field to given value.

### HasCrossConnectId

`func (o *ServiceProfileAccessPointCOLO) HasCrossConnectId() bool`

HasCrossConnectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


