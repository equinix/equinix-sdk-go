# WebhookChannelConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchSize** | Pointer to **int32** |  | [optional] 
**NumberOfConcurrentCalls** | Pointer to **int32** |  | [optional] 
**NumberOfRetries** | Pointer to **int32** |  | [optional] 
**SslCertificate** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewWebhookChannelConfiguration

`func NewWebhookChannelConfiguration(url string, ) *WebhookChannelConfiguration`

NewWebhookChannelConfiguration instantiates a new WebhookChannelConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookChannelConfigurationWithDefaults

`func NewWebhookChannelConfigurationWithDefaults() *WebhookChannelConfiguration`

NewWebhookChannelConfigurationWithDefaults instantiates a new WebhookChannelConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchSize

`func (o *WebhookChannelConfiguration) GetBatchSize() int32`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *WebhookChannelConfiguration) GetBatchSizeOk() (*int32, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *WebhookChannelConfiguration) SetBatchSize(v int32)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *WebhookChannelConfiguration) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetNumberOfConcurrentCalls

`func (o *WebhookChannelConfiguration) GetNumberOfConcurrentCalls() int32`

GetNumberOfConcurrentCalls returns the NumberOfConcurrentCalls field if non-nil, zero value otherwise.

### GetNumberOfConcurrentCallsOk

`func (o *WebhookChannelConfiguration) GetNumberOfConcurrentCallsOk() (*int32, bool)`

GetNumberOfConcurrentCallsOk returns a tuple with the NumberOfConcurrentCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfConcurrentCalls

`func (o *WebhookChannelConfiguration) SetNumberOfConcurrentCalls(v int32)`

SetNumberOfConcurrentCalls sets NumberOfConcurrentCalls field to given value.

### HasNumberOfConcurrentCalls

`func (o *WebhookChannelConfiguration) HasNumberOfConcurrentCalls() bool`

HasNumberOfConcurrentCalls returns a boolean if a field has been set.

### GetNumberOfRetries

`func (o *WebhookChannelConfiguration) GetNumberOfRetries() int32`

GetNumberOfRetries returns the NumberOfRetries field if non-nil, zero value otherwise.

### GetNumberOfRetriesOk

`func (o *WebhookChannelConfiguration) GetNumberOfRetriesOk() (*int32, bool)`

GetNumberOfRetriesOk returns a tuple with the NumberOfRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfRetries

`func (o *WebhookChannelConfiguration) SetNumberOfRetries(v int32)`

SetNumberOfRetries sets NumberOfRetries field to given value.

### HasNumberOfRetries

`func (o *WebhookChannelConfiguration) HasNumberOfRetries() bool`

HasNumberOfRetries returns a boolean if a field has been set.

### GetSslCertificate

`func (o *WebhookChannelConfiguration) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *WebhookChannelConfiguration) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *WebhookChannelConfiguration) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *WebhookChannelConfiguration) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookChannelConfiguration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookChannelConfiguration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookChannelConfiguration) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


