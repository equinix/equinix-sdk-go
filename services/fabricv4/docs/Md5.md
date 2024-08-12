# Md5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**Md5Type**](Md5Type.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewMd5

`func NewMd5() *Md5`

NewMd5 instantiates a new Md5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMd5WithDefaults

`func NewMd5WithDefaults() *Md5`

NewMd5WithDefaults instantiates a new Md5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Md5) GetType() Md5Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Md5) GetTypeOk() (*Md5Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Md5) SetType(v Md5Type)`

SetType sets Type field to given value.

### HasType

`func (o *Md5) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Md5) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Md5) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Md5) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Md5) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPassword

`func (o *Md5) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Md5) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Md5) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Md5) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


