# Oauth2TokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **string** | Deprecated - The Equinix username used to access portals | [optional] 
**UserPassword** | Pointer to **string** | Deprecated - The Equinix user password used to access portals | [optional] 
**ClientId** | **string** | API Consumer Key available under \&quot;My Apps\&quot; in developer portal | 
**ClientSecret** | **string** | API Consumer secret available under \&quot;My Apps\&quot; in developer portal | 
**GrantType** | Pointer to **string** | The OAuth2 grant type used for authorization. Supported values are \&quot;password\&quot; &amp; \&quot;client_credentials\&quot;. user_name and password is not considered in case this value is \&quot;client_credentials\&quot;. If the grant_type is not passed, by default it would consider \&quot;password\&quot; type in which user_name and password is required. Note that the password grant type is deprecated. Recommended to use grant_type of &#39;client_credentials&#39; instead. | [optional] 
**PasswordEncoding** | Pointer to **string** | For enhanced security, you may encrypt the password value while requesting for an access_token. Currently only \&quot;md5-b64\&quot; hashing is supported. Any other value would treat password value as raw string | [optional] [default to "none"]

## Methods

### NewOauth2TokenRequest

`func NewOauth2TokenRequest(clientId string, clientSecret string, ) *Oauth2TokenRequest`

NewOauth2TokenRequest instantiates a new Oauth2TokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2TokenRequestWithDefaults

`func NewOauth2TokenRequestWithDefaults() *Oauth2TokenRequest`

NewOauth2TokenRequestWithDefaults instantiates a new Oauth2TokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *Oauth2TokenRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *Oauth2TokenRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *Oauth2TokenRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *Oauth2TokenRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserPassword

`func (o *Oauth2TokenRequest) GetUserPassword() string`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *Oauth2TokenRequest) GetUserPasswordOk() (*string, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *Oauth2TokenRequest) SetUserPassword(v string)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *Oauth2TokenRequest) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.

### GetClientId

`func (o *Oauth2TokenRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Oauth2TokenRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Oauth2TokenRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *Oauth2TokenRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Oauth2TokenRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Oauth2TokenRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetGrantType

`func (o *Oauth2TokenRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *Oauth2TokenRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *Oauth2TokenRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *Oauth2TokenRequest) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetPasswordEncoding

`func (o *Oauth2TokenRequest) GetPasswordEncoding() string`

GetPasswordEncoding returns the PasswordEncoding field if non-nil, zero value otherwise.

### GetPasswordEncodingOk

`func (o *Oauth2TokenRequest) GetPasswordEncodingOk() (*string, bool)`

GetPasswordEncodingOk returns a tuple with the PasswordEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordEncoding

`func (o *Oauth2TokenRequest) SetPasswordEncoding(v string)`

SetPasswordEncoding sets PasswordEncoding field to given value.

### HasPasswordEncoding

`func (o *Oauth2TokenRequest) HasPasswordEncoding() bool`

HasPasswordEncoding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


