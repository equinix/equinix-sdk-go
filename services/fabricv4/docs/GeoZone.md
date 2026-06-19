# GeoZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code assigned to a geographic zone. | [optional] 
**Name** | Pointer to **string** | Name of a geographic zone. | [optional] 
**Description** | Pointer to **string** | Description of a geographic zone. | [optional] 

## Methods

### NewGeoZone

`func NewGeoZone() *GeoZone`

NewGeoZone instantiates a new GeoZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoZoneWithDefaults

`func NewGeoZoneWithDefaults() *GeoZone`

NewGeoZoneWithDefaults instantiates a new GeoZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GeoZone) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GeoZone) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GeoZone) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GeoZone) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GeoZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeoZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeoZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GeoZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GeoZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GeoZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GeoZone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GeoZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


