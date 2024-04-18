# UserPublicKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Username. | [optional] 
**KeyName** | Pointer to **string** | Key name. This field may be required for some vendors. The keyName must be an existing keyName associated with an existing keyValue. To set up a new keyName and keyValue pair, call Create Public Key. | [optional] 

## Methods

### NewUserPublicKeyRequest

`func NewUserPublicKeyRequest() *UserPublicKeyRequest`

NewUserPublicKeyRequest instantiates a new UserPublicKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPublicKeyRequestWithDefaults

`func NewUserPublicKeyRequestWithDefaults() *UserPublicKeyRequest`

NewUserPublicKeyRequestWithDefaults instantiates a new UserPublicKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserPublicKeyRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserPublicKeyRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserPublicKeyRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserPublicKeyRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetKeyName

`func (o *UserPublicKeyRequest) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *UserPublicKeyRequest) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *UserPublicKeyRequest) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *UserPublicKeyRequest) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


