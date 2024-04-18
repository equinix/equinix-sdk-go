# GetValidateAuthKeyRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Primary** | Pointer to [**GetValidateAuthkeyresPrimary**](GetValidateAuthkeyresPrimary.md) |  | [optional] 
**Secondary** | Pointer to [**GetValidateAuthkeyresSecondary**](GetValidateAuthkeyresSecondary.md) |  | [optional] 

## Methods

### NewGetValidateAuthKeyRes

`func NewGetValidateAuthKeyRes() *GetValidateAuthKeyRes`

NewGetValidateAuthKeyRes instantiates a new GetValidateAuthKeyRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetValidateAuthKeyResWithDefaults

`func NewGetValidateAuthKeyResWithDefaults() *GetValidateAuthKeyRes`

NewGetValidateAuthKeyResWithDefaults instantiates a new GetValidateAuthKeyRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetValidateAuthKeyRes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetValidateAuthKeyRes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetValidateAuthKeyRes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetValidateAuthKeyRes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *GetValidateAuthKeyRes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetValidateAuthKeyRes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetValidateAuthKeyRes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetValidateAuthKeyRes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrimary

`func (o *GetValidateAuthKeyRes) GetPrimary() GetValidateAuthkeyresPrimary`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *GetValidateAuthKeyRes) GetPrimaryOk() (*GetValidateAuthkeyresPrimary, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *GetValidateAuthKeyRes) SetPrimary(v GetValidateAuthkeyresPrimary)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *GetValidateAuthKeyRes) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetSecondary

`func (o *GetValidateAuthKeyRes) GetSecondary() GetValidateAuthkeyresSecondary`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *GetValidateAuthKeyRes) GetSecondaryOk() (*GetValidateAuthkeyresSecondary, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *GetValidateAuthKeyRes) SetSecondary(v GetValidateAuthkeyresSecondary)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *GetValidateAuthKeyRes) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


