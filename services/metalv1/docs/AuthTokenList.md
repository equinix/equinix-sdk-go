# AuthTokenList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeys** | Pointer to [**[]AuthToken**](AuthToken.md) |  | [optional] 

## Methods

### NewAuthTokenList

`func NewAuthTokenList() *AuthTokenList`

NewAuthTokenList instantiates a new AuthTokenList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenListWithDefaults

`func NewAuthTokenListWithDefaults() *AuthTokenList`

NewAuthTokenListWithDefaults instantiates a new AuthTokenList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeys

`func (o *AuthTokenList) GetApiKeys() []AuthToken`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *AuthTokenList) GetApiKeysOk() (*[]AuthToken, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *AuthTokenList) SetApiKeys(v []AuthToken)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *AuthTokenList) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


