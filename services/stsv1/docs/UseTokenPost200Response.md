# UseTokenPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  | 
**IssuedTokenType** | [**UseTokenPost200ResponseIssuedTokenType**](UseTokenPost200ResponseIssuedTokenType.md) |  | 
**TokenType** | [**UseTokenPost200ResponseTokenType**](UseTokenPost200ResponseTokenType.md) |  | 
**ExpiresIn** | **float32** |  | 

## Methods

### NewUseTokenPost200Response

`func NewUseTokenPost200Response(accessToken string, issuedTokenType UseTokenPost200ResponseIssuedTokenType, tokenType UseTokenPost200ResponseTokenType, expiresIn float32, ) *UseTokenPost200Response`

NewUseTokenPost200Response instantiates a new UseTokenPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUseTokenPost200ResponseWithDefaults

`func NewUseTokenPost200ResponseWithDefaults() *UseTokenPost200Response`

NewUseTokenPost200ResponseWithDefaults instantiates a new UseTokenPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *UseTokenPost200Response) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *UseTokenPost200Response) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *UseTokenPost200Response) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetIssuedTokenType

`func (o *UseTokenPost200Response) GetIssuedTokenType() UseTokenPost200ResponseIssuedTokenType`

GetIssuedTokenType returns the IssuedTokenType field if non-nil, zero value otherwise.

### GetIssuedTokenTypeOk

`func (o *UseTokenPost200Response) GetIssuedTokenTypeOk() (*UseTokenPost200ResponseIssuedTokenType, bool)`

GetIssuedTokenTypeOk returns a tuple with the IssuedTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTokenType

`func (o *UseTokenPost200Response) SetIssuedTokenType(v UseTokenPost200ResponseIssuedTokenType)`

SetIssuedTokenType sets IssuedTokenType field to given value.


### GetTokenType

`func (o *UseTokenPost200Response) GetTokenType() UseTokenPost200ResponseTokenType`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *UseTokenPost200Response) GetTokenTypeOk() (*UseTokenPost200ResponseTokenType, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *UseTokenPost200Response) SetTokenType(v UseTokenPost200ResponseTokenType)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *UseTokenPost200Response) GetExpiresIn() float32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *UseTokenPost200Response) GetExpiresInOk() (*float32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *UseTokenPost200Response) SetExpiresIn(v float32)`

SetExpiresIn sets ExpiresIn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


