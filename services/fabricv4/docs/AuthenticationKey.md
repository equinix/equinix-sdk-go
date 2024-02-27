# AuthenticationKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Required** | Pointer to **bool** |  | [optional] [default to false]
**Label** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthenticationKey

`func NewAuthenticationKey() *AuthenticationKey`

NewAuthenticationKey instantiates a new AuthenticationKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationKeyWithDefaults

`func NewAuthenticationKeyWithDefaults() *AuthenticationKey`

NewAuthenticationKeyWithDefaults instantiates a new AuthenticationKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequired

`func (o *AuthenticationKey) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AuthenticationKey) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AuthenticationKey) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *AuthenticationKey) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetLabel

`func (o *AuthenticationKey) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AuthenticationKey) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AuthenticationKey) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AuthenticationKey) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *AuthenticationKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthenticationKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthenticationKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthenticationKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


