# Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsIotCoreChannelConfiguration** | Pointer to [**AwsIotCoreChannelConfiguration**](AwsIotCoreChannelConfiguration.md) |  | [optional] 
**AzureChannelConfiguration** | Pointer to [**AzureChannelConfiguration**](AzureChannelConfiguration.md) |  | [optional] 
**ChannelType** | Pointer to [**ChannelChannelType**](ChannelChannelType.md) |  | [optional] 
**WebhookChannelConfiguration** | Pointer to [**WebhookChannelConfiguration**](WebhookChannelConfiguration.md) |  | [optional] 

## Methods

### NewChannel

`func NewChannel() *Channel`

NewChannel instantiates a new Channel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelWithDefaults

`func NewChannelWithDefaults() *Channel`

NewChannelWithDefaults instantiates a new Channel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsIotCoreChannelConfiguration

`func (o *Channel) GetAwsIotCoreChannelConfiguration() AwsIotCoreChannelConfiguration`

GetAwsIotCoreChannelConfiguration returns the AwsIotCoreChannelConfiguration field if non-nil, zero value otherwise.

### GetAwsIotCoreChannelConfigurationOk

`func (o *Channel) GetAwsIotCoreChannelConfigurationOk() (*AwsIotCoreChannelConfiguration, bool)`

GetAwsIotCoreChannelConfigurationOk returns a tuple with the AwsIotCoreChannelConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsIotCoreChannelConfiguration

`func (o *Channel) SetAwsIotCoreChannelConfiguration(v AwsIotCoreChannelConfiguration)`

SetAwsIotCoreChannelConfiguration sets AwsIotCoreChannelConfiguration field to given value.

### HasAwsIotCoreChannelConfiguration

`func (o *Channel) HasAwsIotCoreChannelConfiguration() bool`

HasAwsIotCoreChannelConfiguration returns a boolean if a field has been set.

### GetAzureChannelConfiguration

`func (o *Channel) GetAzureChannelConfiguration() AzureChannelConfiguration`

GetAzureChannelConfiguration returns the AzureChannelConfiguration field if non-nil, zero value otherwise.

### GetAzureChannelConfigurationOk

`func (o *Channel) GetAzureChannelConfigurationOk() (*AzureChannelConfiguration, bool)`

GetAzureChannelConfigurationOk returns a tuple with the AzureChannelConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureChannelConfiguration

`func (o *Channel) SetAzureChannelConfiguration(v AzureChannelConfiguration)`

SetAzureChannelConfiguration sets AzureChannelConfiguration field to given value.

### HasAzureChannelConfiguration

`func (o *Channel) HasAzureChannelConfiguration() bool`

HasAzureChannelConfiguration returns a boolean if a field has been set.

### GetChannelType

`func (o *Channel) GetChannelType() ChannelChannelType`

GetChannelType returns the ChannelType field if non-nil, zero value otherwise.

### GetChannelTypeOk

`func (o *Channel) GetChannelTypeOk() (*ChannelChannelType, bool)`

GetChannelTypeOk returns a tuple with the ChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelType

`func (o *Channel) SetChannelType(v ChannelChannelType)`

SetChannelType sets ChannelType field to given value.

### HasChannelType

`func (o *Channel) HasChannelType() bool`

HasChannelType returns a boolean if a field has been set.

### GetWebhookChannelConfiguration

`func (o *Channel) GetWebhookChannelConfiguration() WebhookChannelConfiguration`

GetWebhookChannelConfiguration returns the WebhookChannelConfiguration field if non-nil, zero value otherwise.

### GetWebhookChannelConfigurationOk

`func (o *Channel) GetWebhookChannelConfigurationOk() (*WebhookChannelConfiguration, bool)`

GetWebhookChannelConfigurationOk returns a tuple with the WebhookChannelConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookChannelConfiguration

`func (o *Channel) SetWebhookChannelConfiguration(v WebhookChannelConfiguration)`

SetWebhookChannelConfiguration sets WebhookChannelConfiguration field to given value.

### HasWebhookChannelConfiguration

`func (o *Channel) HasWebhookChannelConfiguration() bool`

HasWebhookChannelConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


