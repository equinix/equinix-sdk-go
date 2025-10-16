# PublicKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyName** | **string** | Key name | 
**KeyValue** | **string** | Key value | 
**KeyType** | Pointer to **string** | Key type, whether RSA or DSA. Default is RSA. | [optional] 

## Methods

### NewPublicKeyRequest

`func NewPublicKeyRequest(keyName string, keyValue string, ) *PublicKeyRequest`

NewPublicKeyRequest instantiates a new PublicKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicKeyRequestWithDefaults

`func NewPublicKeyRequestWithDefaults() *PublicKeyRequest`

NewPublicKeyRequestWithDefaults instantiates a new PublicKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyName

`func (o *PublicKeyRequest) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *PublicKeyRequest) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *PublicKeyRequest) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetKeyValue

`func (o *PublicKeyRequest) GetKeyValue() string`

GetKeyValue returns the KeyValue field if non-nil, zero value otherwise.

### GetKeyValueOk

`func (o *PublicKeyRequest) GetKeyValueOk() (*string, bool)`

GetKeyValueOk returns a tuple with the KeyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyValue

`func (o *PublicKeyRequest) SetKeyValue(v string)`

SetKeyValue sets KeyValue field to given value.


### GetKeyType

`func (o *PublicKeyRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *PublicKeyRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *PublicKeyRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *PublicKeyRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


