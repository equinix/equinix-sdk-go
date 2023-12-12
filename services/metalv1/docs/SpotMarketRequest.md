# SpotMarketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DevicesMax** | Pointer to **int32** |  | [optional] 
**DevicesMin** | Pointer to **int32** |  | [optional] 
**EndAt** | Pointer to **time.Time** |  | [optional] 
**Facilities** | Pointer to [**Href**](Href.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Instances** | Pointer to [**Href**](Href.md) |  | [optional] 
**MaxBidPrice** | Pointer to **float32** |  | [optional] 
**Metro** | Pointer to [**SpotMarketRequestMetro**](SpotMarketRequestMetro.md) |  | [optional] 
**Project** | Pointer to [**Href**](Href.md) |  | [optional] 

## Methods

### NewSpotMarketRequest

`func NewSpotMarketRequest() *SpotMarketRequest`

NewSpotMarketRequest instantiates a new SpotMarketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotMarketRequestWithDefaults

`func NewSpotMarketRequestWithDefaults() *SpotMarketRequest`

NewSpotMarketRequestWithDefaults instantiates a new SpotMarketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SpotMarketRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpotMarketRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpotMarketRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SpotMarketRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDevicesMax

`func (o *SpotMarketRequest) GetDevicesMax() int32`

GetDevicesMax returns the DevicesMax field if non-nil, zero value otherwise.

### GetDevicesMaxOk

`func (o *SpotMarketRequest) GetDevicesMaxOk() (*int32, bool)`

GetDevicesMaxOk returns a tuple with the DevicesMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesMax

`func (o *SpotMarketRequest) SetDevicesMax(v int32)`

SetDevicesMax sets DevicesMax field to given value.

### HasDevicesMax

`func (o *SpotMarketRequest) HasDevicesMax() bool`

HasDevicesMax returns a boolean if a field has been set.

### GetDevicesMin

`func (o *SpotMarketRequest) GetDevicesMin() int32`

GetDevicesMin returns the DevicesMin field if non-nil, zero value otherwise.

### GetDevicesMinOk

`func (o *SpotMarketRequest) GetDevicesMinOk() (*int32, bool)`

GetDevicesMinOk returns a tuple with the DevicesMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesMin

`func (o *SpotMarketRequest) SetDevicesMin(v int32)`

SetDevicesMin sets DevicesMin field to given value.

### HasDevicesMin

`func (o *SpotMarketRequest) HasDevicesMin() bool`

HasDevicesMin returns a boolean if a field has been set.

### GetEndAt

`func (o *SpotMarketRequest) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *SpotMarketRequest) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *SpotMarketRequest) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *SpotMarketRequest) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### GetFacilities

`func (o *SpotMarketRequest) GetFacilities() Href`

GetFacilities returns the Facilities field if non-nil, zero value otherwise.

### GetFacilitiesOk

`func (o *SpotMarketRequest) GetFacilitiesOk() (*Href, bool)`

GetFacilitiesOk returns a tuple with the Facilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilities

`func (o *SpotMarketRequest) SetFacilities(v Href)`

SetFacilities sets Facilities field to given value.

### HasFacilities

`func (o *SpotMarketRequest) HasFacilities() bool`

HasFacilities returns a boolean if a field has been set.

### GetHref

`func (o *SpotMarketRequest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SpotMarketRequest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SpotMarketRequest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SpotMarketRequest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *SpotMarketRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpotMarketRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpotMarketRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpotMarketRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstances

`func (o *SpotMarketRequest) GetInstances() Href`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *SpotMarketRequest) GetInstancesOk() (*Href, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *SpotMarketRequest) SetInstances(v Href)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *SpotMarketRequest) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMaxBidPrice

`func (o *SpotMarketRequest) GetMaxBidPrice() float32`

GetMaxBidPrice returns the MaxBidPrice field if non-nil, zero value otherwise.

### GetMaxBidPriceOk

`func (o *SpotMarketRequest) GetMaxBidPriceOk() (*float32, bool)`

GetMaxBidPriceOk returns a tuple with the MaxBidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBidPrice

`func (o *SpotMarketRequest) SetMaxBidPrice(v float32)`

SetMaxBidPrice sets MaxBidPrice field to given value.

### HasMaxBidPrice

`func (o *SpotMarketRequest) HasMaxBidPrice() bool`

HasMaxBidPrice returns a boolean if a field has been set.

### GetMetro

`func (o *SpotMarketRequest) GetMetro() SpotMarketRequestMetro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *SpotMarketRequest) GetMetroOk() (*SpotMarketRequestMetro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *SpotMarketRequest) SetMetro(v SpotMarketRequestMetro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *SpotMarketRequest) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetProject

`func (o *SpotMarketRequest) GetProject() Href`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SpotMarketRequest) GetProjectOk() (*Href, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SpotMarketRequest) SetProject(v Href)`

SetProject sets Project field to given value.

### HasProject

`func (o *SpotMarketRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


