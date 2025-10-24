# AllowedInterfaceProfiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **float32** | Allowed interface count | [optional] 
**Interfaces** | Pointer to [**[]InterfaceDetails**](InterfaceDetails.md) |  | [optional] 
**Default** | Pointer to **bool** | Whether this will be the default interface count if you do not provide a number. | [optional] 

## Methods

### NewAllowedInterfaceProfiles

`func NewAllowedInterfaceProfiles() *AllowedInterfaceProfiles`

NewAllowedInterfaceProfiles instantiates a new AllowedInterfaceProfiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedInterfaceProfilesWithDefaults

`func NewAllowedInterfaceProfilesWithDefaults() *AllowedInterfaceProfiles`

NewAllowedInterfaceProfilesWithDefaults instantiates a new AllowedInterfaceProfiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AllowedInterfaceProfiles) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AllowedInterfaceProfiles) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AllowedInterfaceProfiles) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AllowedInterfaceProfiles) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetInterfaces

`func (o *AllowedInterfaceProfiles) GetInterfaces() []InterfaceDetails`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *AllowedInterfaceProfiles) GetInterfacesOk() (*[]InterfaceDetails, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *AllowedInterfaceProfiles) SetInterfaces(v []InterfaceDetails)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *AllowedInterfaceProfiles) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetDefault

`func (o *AllowedInterfaceProfiles) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *AllowedInterfaceProfiles) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *AllowedInterfaceProfiles) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *AllowedInterfaceProfiles) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


