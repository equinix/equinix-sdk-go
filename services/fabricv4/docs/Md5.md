# Md5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**Md5Type**](Md5Type.md) |  | [optional] 
**KeyNumber** | Pointer to **int32** | The authentication Key ID. | [optional] 
**Key** | Pointer to **string** | The plaintext authentication key. Must be Base64 encoded. For ASCII type, the key must contain printable ASCII characters, range 10-20 characters. For HEX type, range should be 10-40 characters. | [optional] 

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

### GetKeyNumber

`func (o *Md5) GetKeyNumber() int32`

GetKeyNumber returns the KeyNumber field if non-nil, zero value otherwise.

### GetKeyNumberOk

`func (o *Md5) GetKeyNumberOk() (*int32, bool)`

GetKeyNumberOk returns a tuple with the KeyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyNumber

`func (o *Md5) SetKeyNumber(v int32)`

SetKeyNumber sets KeyNumber field to given value.

### HasKeyNumber

`func (o *Md5) HasKeyNumber() bool`

HasKeyNumber returns a boolean if a field has been set.

### GetKey

`func (o *Md5) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Md5) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Md5) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Md5) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


