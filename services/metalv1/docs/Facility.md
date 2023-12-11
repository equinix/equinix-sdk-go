# Facility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**[]FacilityFeaturesInner**](FacilityFeaturesInner.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IpRanges** | Pointer to **[]string** | IP ranges registered in facility. Can be used for GeoIP location | [optional] 
**Metro** | Pointer to [**DeviceMetro**](DeviceMetro.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewFacility

`func NewFacility() *Facility`

NewFacility instantiates a new Facility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacilityWithDefaults

`func NewFacilityWithDefaults() *Facility`

NewFacilityWithDefaults instantiates a new Facility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Facility) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Facility) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Facility) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Facility) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCode

`func (o *Facility) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Facility) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Facility) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Facility) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFeatures

`func (o *Facility) GetFeatures() []FacilityFeaturesInner`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Facility) GetFeaturesOk() (*[]FacilityFeaturesInner, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Facility) SetFeatures(v []FacilityFeaturesInner)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Facility) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetId

`func (o *Facility) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Facility) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Facility) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Facility) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpRanges

`func (o *Facility) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *Facility) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *Facility) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *Facility) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### GetMetro

`func (o *Facility) GetMetro() DeviceMetro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *Facility) GetMetroOk() (*DeviceMetro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *Facility) SetMetro(v DeviceMetro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *Facility) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetName

`func (o *Facility) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Facility) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Facility) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Facility) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


