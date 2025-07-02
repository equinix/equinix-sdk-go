# TokenExchangeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  | 
**IssuedTokenType** | [**TokenExchangeResponseIssuedTokenType**](TokenExchangeResponseIssuedTokenType.md) |  | 
**TokenType** | [**TokenExchangeResponseTokenType**](TokenExchangeResponseTokenType.md) |  | 
**ExpiresIn** | **float32** |  | 

## Methods

### NewTokenExchangeResponse

`func NewTokenExchangeResponse(accessToken string, issuedTokenType TokenExchangeResponseIssuedTokenType, tokenType TokenExchangeResponseTokenType, expiresIn float32, ) *TokenExchangeResponse`

NewTokenExchangeResponse instantiates a new TokenExchangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenExchangeResponseWithDefaults

`func NewTokenExchangeResponseWithDefaults() *TokenExchangeResponse`

NewTokenExchangeResponseWithDefaults instantiates a new TokenExchangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *TokenExchangeResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenExchangeResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenExchangeResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetIssuedTokenType

`func (o *TokenExchangeResponse) GetIssuedTokenType() TokenExchangeResponseIssuedTokenType`

GetIssuedTokenType returns the IssuedTokenType field if non-nil, zero value otherwise.

### GetIssuedTokenTypeOk

`func (o *TokenExchangeResponse) GetIssuedTokenTypeOk() (*TokenExchangeResponseIssuedTokenType, bool)`

GetIssuedTokenTypeOk returns a tuple with the IssuedTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTokenType

`func (o *TokenExchangeResponse) SetIssuedTokenType(v TokenExchangeResponseIssuedTokenType)`

SetIssuedTokenType sets IssuedTokenType field to given value.


### GetTokenType

`func (o *TokenExchangeResponse) GetTokenType() TokenExchangeResponseTokenType`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TokenExchangeResponse) GetTokenTypeOk() (*TokenExchangeResponseTokenType, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TokenExchangeResponse) SetTokenType(v TokenExchangeResponseTokenType)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *TokenExchangeResponse) GetExpiresIn() float32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *TokenExchangeResponse) GetExpiresInOk() (*float32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *TokenExchangeResponse) SetExpiresIn(v float32)`

SetExpiresIn sets ExpiresIn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


