# Metro

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | The Canonical URL at which the resource resides. | [optional] 
**Type** | Pointer to **string** | Indicator of a Fabric Metro | [optional] 
**Code** | Pointer to **string** | Code Assigned to an Equinix IBX data center in a specified metropolitan area. | [optional] 
**Region** | Pointer to **string** | Board geographic area in which the data center is located | [optional] 
**Name** | Pointer to **string** | Name of the region in which the data center is located. | [optional] 
**Country** | Pointer to **string** | Country code in which the data center is located. | [optional] 
**EquinixAsn** | Pointer to **int64** | Autonomous system number (ASN) for a specified Fabric metro. The ASN is a unique identifier that carries the network routing protocol and exchanges that data with other internal systems via border gateway protocol. | [optional] 
**LocalVCBandwidthMax** | Pointer to **int64** | This field holds Max Connection speed with in the metro | [optional] 
**GeoCoordinates** | Pointer to [**GeoCoordinates**](GeoCoordinates.md) |  | [optional] 
**ConnectedMetros** | Pointer to [**[]ConnectedMetro**](ConnectedMetro.md) |  | [optional] 
**Services** | Pointer to [**[]Services**](Services.md) |  | [optional] 
**GeoScopes** | Pointer to [**[]GeoScopeType**](GeoScopeType.md) | List of supported geographic boundaries of a Fabric Metro. | [optional] 

## Methods

### NewMetro

`func NewMetro() *Metro`

NewMetro instantiates a new Metro object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetroWithDefaults

`func NewMetroWithDefaults() *Metro`

NewMetroWithDefaults instantiates a new Metro object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Metro) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Metro) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Metro) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Metro) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *Metro) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Metro) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Metro) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Metro) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *Metro) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Metro) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Metro) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Metro) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRegion

`func (o *Metro) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Metro) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Metro) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Metro) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetName

`func (o *Metro) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Metro) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Metro) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Metro) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCountry

`func (o *Metro) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Metro) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Metro) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Metro) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEquinixAsn

`func (o *Metro) GetEquinixAsn() int64`

GetEquinixAsn returns the EquinixAsn field if non-nil, zero value otherwise.

### GetEquinixAsnOk

`func (o *Metro) GetEquinixAsnOk() (*int64, bool)`

GetEquinixAsnOk returns a tuple with the EquinixAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixAsn

`func (o *Metro) SetEquinixAsn(v int64)`

SetEquinixAsn sets EquinixAsn field to given value.

### HasEquinixAsn

`func (o *Metro) HasEquinixAsn() bool`

HasEquinixAsn returns a boolean if a field has been set.

### GetLocalVCBandwidthMax

`func (o *Metro) GetLocalVCBandwidthMax() int64`

GetLocalVCBandwidthMax returns the LocalVCBandwidthMax field if non-nil, zero value otherwise.

### GetLocalVCBandwidthMaxOk

`func (o *Metro) GetLocalVCBandwidthMaxOk() (*int64, bool)`

GetLocalVCBandwidthMaxOk returns a tuple with the LocalVCBandwidthMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVCBandwidthMax

`func (o *Metro) SetLocalVCBandwidthMax(v int64)`

SetLocalVCBandwidthMax sets LocalVCBandwidthMax field to given value.

### HasLocalVCBandwidthMax

`func (o *Metro) HasLocalVCBandwidthMax() bool`

HasLocalVCBandwidthMax returns a boolean if a field has been set.

### GetGeoCoordinates

`func (o *Metro) GetGeoCoordinates() GeoCoordinates`

GetGeoCoordinates returns the GeoCoordinates field if non-nil, zero value otherwise.

### GetGeoCoordinatesOk

`func (o *Metro) GetGeoCoordinatesOk() (*GeoCoordinates, bool)`

GetGeoCoordinatesOk returns a tuple with the GeoCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoCoordinates

`func (o *Metro) SetGeoCoordinates(v GeoCoordinates)`

SetGeoCoordinates sets GeoCoordinates field to given value.

### HasGeoCoordinates

`func (o *Metro) HasGeoCoordinates() bool`

HasGeoCoordinates returns a boolean if a field has been set.

### GetConnectedMetros

`func (o *Metro) GetConnectedMetros() []ConnectedMetro`

GetConnectedMetros returns the ConnectedMetros field if non-nil, zero value otherwise.

### GetConnectedMetrosOk

`func (o *Metro) GetConnectedMetrosOk() (*[]ConnectedMetro, bool)`

GetConnectedMetrosOk returns a tuple with the ConnectedMetros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedMetros

`func (o *Metro) SetConnectedMetros(v []ConnectedMetro)`

SetConnectedMetros sets ConnectedMetros field to given value.

### HasConnectedMetros

`func (o *Metro) HasConnectedMetros() bool`

HasConnectedMetros returns a boolean if a field has been set.

### GetServices

`func (o *Metro) GetServices() []Services`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Metro) GetServicesOk() (*[]Services, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Metro) SetServices(v []Services)`

SetServices sets Services field to given value.

### HasServices

`func (o *Metro) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetGeoScopes

`func (o *Metro) GetGeoScopes() []GeoScopeType`

GetGeoScopes returns the GeoScopes field if non-nil, zero value otherwise.

### GetGeoScopesOk

`func (o *Metro) GetGeoScopesOk() (*[]GeoScopeType, bool)`

GetGeoScopesOk returns a tuple with the GeoScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoScopes

`func (o *Metro) SetGeoScopes(v []GeoScopeType)`

SetGeoScopes sets GeoScopes field to given value.

### HasGeoScopes

`func (o *Metro) HasGeoScopes() bool`

HasGeoScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


