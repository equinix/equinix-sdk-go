# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Address2** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Coordinates** | Pointer to [**Coordinates**](Coordinates.md) |  | [optional] 
**Country** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 
**ZipCode** | **string** |  | 

## Methods

### NewAddress

`func NewAddress(address string, country string, zipCode string, ) *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Address) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Address) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Address) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddress2

`func (o *Address) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *Address) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *Address) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *Address) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *Address) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Address) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCoordinates

`func (o *Address) GetCoordinates() Coordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *Address) GetCoordinatesOk() (*Coordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *Address) SetCoordinates(v Coordinates)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *Address) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetCountry

`func (o *Address) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetState

`func (o *Address) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Address) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Address) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Address) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZipCode

`func (o *Address) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *Address) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *Address) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


