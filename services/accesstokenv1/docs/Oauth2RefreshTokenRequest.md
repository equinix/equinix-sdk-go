# Oauth2RefreshTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | API Consumer Key available under \&quot;My Apps\&quot; in developer portal | 
**ClientSecret** | **string** | API Consumer secret available under \&quot;My Apps\&quot; in developer portal | 
**RefreshToken** | **string** | The OAuth2 refresh_token retrieved from the previous successful Access Token API call | 

## Methods

### NewOauth2RefreshTokenRequest

`func NewOauth2RefreshTokenRequest(clientId string, clientSecret string, refreshToken string, ) *Oauth2RefreshTokenRequest`

NewOauth2RefreshTokenRequest instantiates a new Oauth2RefreshTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2RefreshTokenRequestWithDefaults

`func NewOauth2RefreshTokenRequestWithDefaults() *Oauth2RefreshTokenRequest`

NewOauth2RefreshTokenRequestWithDefaults instantiates a new Oauth2RefreshTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *Oauth2RefreshTokenRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Oauth2RefreshTokenRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Oauth2RefreshTokenRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *Oauth2RefreshTokenRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Oauth2RefreshTokenRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Oauth2RefreshTokenRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetRefreshToken

`func (o *Oauth2RefreshTokenRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *Oauth2RefreshTokenRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *Oauth2RefreshTokenRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


