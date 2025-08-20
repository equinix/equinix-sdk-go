# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | Account Number | 
**Name** | **string** | Name of Account | 
**IsCreditHold** | Pointer to **bool** | Is credit Hold account | [optional] 
**IsPOBearing** | Pointer to **bool** | Is PO Bearing account | [optional] 
**Cabinets** | Pointer to [**[]Cabinets**](Cabinets.md) | List of Cabinets | [optional] 

## Methods

### NewAccount

`func NewAccount(number string, name string, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *Account) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Account) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Account) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.


### GetIsCreditHold

`func (o *Account) GetIsCreditHold() bool`

GetIsCreditHold returns the IsCreditHold field if non-nil, zero value otherwise.

### GetIsCreditHoldOk

`func (o *Account) GetIsCreditHoldOk() (*bool, bool)`

GetIsCreditHoldOk returns a tuple with the IsCreditHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCreditHold

`func (o *Account) SetIsCreditHold(v bool)`

SetIsCreditHold sets IsCreditHold field to given value.

### HasIsCreditHold

`func (o *Account) HasIsCreditHold() bool`

HasIsCreditHold returns a boolean if a field has been set.

### GetIsPOBearing

`func (o *Account) GetIsPOBearing() bool`

GetIsPOBearing returns the IsPOBearing field if non-nil, zero value otherwise.

### GetIsPOBearingOk

`func (o *Account) GetIsPOBearingOk() (*bool, bool)`

GetIsPOBearingOk returns a tuple with the IsPOBearing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPOBearing

`func (o *Account) SetIsPOBearing(v bool)`

SetIsPOBearing sets IsPOBearing field to given value.

### HasIsPOBearing

`func (o *Account) HasIsPOBearing() bool`

HasIsPOBearing returns a boolean if a field has been set.

### GetCabinets

`func (o *Account) GetCabinets() []Cabinets`

GetCabinets returns the Cabinets field if non-nil, zero value otherwise.

### GetCabinetsOk

`func (o *Account) GetCabinetsOk() (*[]Cabinets, bool)`

GetCabinetsOk returns a tuple with the Cabinets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinets

`func (o *Account) SetCabinets(v []Cabinets)`

SetCabinets sets Cabinets field to given value.

### HasCabinets

`func (o *Account) HasCabinets() bool`

HasCabinets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


