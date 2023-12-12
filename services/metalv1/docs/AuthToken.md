# AuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** | Available only for API keys | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**AuthTokenProject**](AuthTokenProject.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to [**AuthTokenUser**](AuthTokenUser.md) |  | [optional] 

## Methods

### NewAuthToken

`func NewAuthToken() *AuthToken`

NewAuthToken instantiates a new AuthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenWithDefaults

`func NewAuthTokenWithDefaults() *AuthToken`

NewAuthTokenWithDefaults instantiates a new AuthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AuthToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *AuthToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AuthToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *AuthToken) GetProject() AuthTokenProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AuthToken) GetProjectOk() (*AuthTokenProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AuthToken) SetProject(v AuthTokenProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *AuthToken) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetReadOnly

`func (o *AuthToken) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *AuthToken) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *AuthToken) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *AuthToken) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetToken

`func (o *AuthToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthToken) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthToken) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthToken) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *AuthToken) GetUser() AuthTokenUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthToken) GetUserOk() (*AuthTokenUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthToken) SetUser(v AuthTokenUser)`

SetUser sets User field to given value.

### HasUser

`func (o *AuthToken) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


