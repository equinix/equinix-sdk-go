# SrvRegisterClientPost200ResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | An OAuth 2.0 client id for an EaaS service. | 
**ClientSecret** | **string** |  | 
**ClientSecretExpiresAt** | **int32** |  | 

## Methods

### NewSrvRegisterClientPost200ResponseResult

`func NewSrvRegisterClientPost200ResponseResult(clientId string, clientSecret string, clientSecretExpiresAt int32, ) *SrvRegisterClientPost200ResponseResult`

NewSrvRegisterClientPost200ResponseResult instantiates a new SrvRegisterClientPost200ResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrvRegisterClientPost200ResponseResultWithDefaults

`func NewSrvRegisterClientPost200ResponseResultWithDefaults() *SrvRegisterClientPost200ResponseResult`

NewSrvRegisterClientPost200ResponseResultWithDefaults instantiates a new SrvRegisterClientPost200ResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SrvRegisterClientPost200ResponseResult) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SrvRegisterClientPost200ResponseResult) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SrvRegisterClientPost200ResponseResult) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *SrvRegisterClientPost200ResponseResult) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *SrvRegisterClientPost200ResponseResult) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *SrvRegisterClientPost200ResponseResult) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetClientSecretExpiresAt

`func (o *SrvRegisterClientPost200ResponseResult) GetClientSecretExpiresAt() int32`

GetClientSecretExpiresAt returns the ClientSecretExpiresAt field if non-nil, zero value otherwise.

### GetClientSecretExpiresAtOk

`func (o *SrvRegisterClientPost200ResponseResult) GetClientSecretExpiresAtOk() (*int32, bool)`

GetClientSecretExpiresAtOk returns a tuple with the ClientSecretExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretExpiresAt

`func (o *SrvRegisterClientPost200ResponseResult) SetClientSecretExpiresAt(v int32)`

SetClientSecretExpiresAt sets ClientSecretExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


