# AuthContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authtype** | Pointer to [**AuthContextAuthtype**](AuthContextAuthtype.md) |  | [optional] 
**Authid** | Pointer to [**AuthContextAuthid**](AuthContextAuthid.md) |  | [optional] 
**Name** | Pointer to **string** | Cloud Event username | [optional] 
**Email** | Pointer to **string** | Cloud Event email | [optional] 

## Methods

### NewAuthContext

`func NewAuthContext() *AuthContext`

NewAuthContext instantiates a new AuthContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthContextWithDefaults

`func NewAuthContextWithDefaults() *AuthContext`

NewAuthContextWithDefaults instantiates a new AuthContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthtype

`func (o *AuthContext) GetAuthtype() AuthContextAuthtype`

GetAuthtype returns the Authtype field if non-nil, zero value otherwise.

### GetAuthtypeOk

`func (o *AuthContext) GetAuthtypeOk() (*AuthContextAuthtype, bool)`

GetAuthtypeOk returns a tuple with the Authtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthtype

`func (o *AuthContext) SetAuthtype(v AuthContextAuthtype)`

SetAuthtype sets Authtype field to given value.

### HasAuthtype

`func (o *AuthContext) HasAuthtype() bool`

HasAuthtype returns a boolean if a field has been set.

### GetAuthid

`func (o *AuthContext) GetAuthid() AuthContextAuthid`

GetAuthid returns the Authid field if non-nil, zero value otherwise.

### GetAuthidOk

`func (o *AuthContext) GetAuthidOk() (*AuthContextAuthid, bool)`

GetAuthidOk returns a tuple with the Authid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthid

`func (o *AuthContext) SetAuthid(v AuthContextAuthid)`

SetAuthid sets Authid field to given value.

### HasAuthid

`func (o *AuthContext) HasAuthid() bool`

HasAuthid returns a boolean if a field has been set.

### GetName

`func (o *AuthContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthContext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthContext) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *AuthContext) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthContext) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthContext) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthContext) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


