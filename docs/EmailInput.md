# EmailInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Default** | Pointer to **bool** |  | [optional] 

## Methods

### NewEmailInput

`func NewEmailInput(address string, ) *EmailInput`

NewEmailInput instantiates a new EmailInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailInputWithDefaults

`func NewEmailInputWithDefaults() *EmailInput`

NewEmailInputWithDefaults instantiates a new EmailInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *EmailInput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EmailInput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EmailInput) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetDefault

`func (o *EmailInput) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *EmailInput) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *EmailInput) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *EmailInput) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


