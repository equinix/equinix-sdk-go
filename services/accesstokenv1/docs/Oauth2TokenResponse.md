# Oauth2TokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | the OAuth2 Bearer token required to access Equinix API&#39;s | 
**TokenTimeout** | **string** | The max validity in seconds for this access_token | [default to "3599 (60 minutes)"]
**UserName** | **string** | The Equinix user account to which this token is associated with | 
**TokenType** | **string** | The type of access_token returned in this response | [default to "Bearer"]
**RefreshToken** | Pointer to **string** | The OAuth2 refresh token which could be used to retrieve a new access_token without user credentials. The refresh_token has a max validity of 60 days after which a new Access Token request must be made. Refresh token is Not returned for \&quot;client_credentials\&quot; grant type | [optional] 
**RefreshTokenTimeout** | Pointer to **string** | The max validity in seconds for the refresh_token. This field is not retruned for the \&quot;client_credentials\&quot; grant type | [optional] [default to "5182560 (60 days)"]

## Methods

### NewOauth2TokenResponse

`func NewOauth2TokenResponse(accessToken string, tokenTimeout string, userName string, tokenType string, ) *Oauth2TokenResponse`

NewOauth2TokenResponse instantiates a new Oauth2TokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2TokenResponseWithDefaults

`func NewOauth2TokenResponseWithDefaults() *Oauth2TokenResponse`

NewOauth2TokenResponseWithDefaults instantiates a new Oauth2TokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *Oauth2TokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Oauth2TokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Oauth2TokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTokenTimeout

`func (o *Oauth2TokenResponse) GetTokenTimeout() string`

GetTokenTimeout returns the TokenTimeout field if non-nil, zero value otherwise.

### GetTokenTimeoutOk

`func (o *Oauth2TokenResponse) GetTokenTimeoutOk() (*string, bool)`

GetTokenTimeoutOk returns a tuple with the TokenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTimeout

`func (o *Oauth2TokenResponse) SetTokenTimeout(v string)`

SetTokenTimeout sets TokenTimeout field to given value.


### GetUserName

`func (o *Oauth2TokenResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *Oauth2TokenResponse) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *Oauth2TokenResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetTokenType

`func (o *Oauth2TokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *Oauth2TokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *Oauth2TokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetRefreshToken

`func (o *Oauth2TokenResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *Oauth2TokenResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *Oauth2TokenResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *Oauth2TokenResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetRefreshTokenTimeout

`func (o *Oauth2TokenResponse) GetRefreshTokenTimeout() string`

GetRefreshTokenTimeout returns the RefreshTokenTimeout field if non-nil, zero value otherwise.

### GetRefreshTokenTimeoutOk

`func (o *Oauth2TokenResponse) GetRefreshTokenTimeoutOk() (*string, bool)`

GetRefreshTokenTimeoutOk returns a tuple with the RefreshTokenTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenTimeout

`func (o *Oauth2TokenResponse) SetRefreshTokenTimeout(v string)`

SetRefreshTokenTimeout sets RefreshTokenTimeout field to given value.

### HasRefreshTokenTimeout

`func (o *Oauth2TokenResponse) HasRefreshTokenTimeout() bool`

HasRefreshTokenTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


