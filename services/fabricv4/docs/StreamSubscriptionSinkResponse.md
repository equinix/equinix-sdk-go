# StreamSubscriptionSinkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | Pointer to **string** | any publicly reachable http endpoint | [optional] 
**Type** | Pointer to [**StreamSubscriptionSinkType**](StreamSubscriptionSinkType.md) |  | [optional] 
**BatchEnabled** | Pointer to **bool** | batch mode on/off | [optional] 
**BatchSizeMax** | Pointer to **int32** | maximum batch size | [optional] 
**BatchWaitTimeMax** | Pointer to **int32** | maximum batch waiting time | [optional] 
**Credential** | Pointer to [**StreamSubscriptionSinkCredential**](StreamSubscriptionSinkCredential.md) |  | [optional] 
**Settings** | Pointer to [**StreamSubscriptionSinkSetting**](StreamSubscriptionSinkSetting.md) |  | [optional] 
**Host** | Pointer to **string** | sink host | [optional] 

## Methods

### NewStreamSubscriptionSinkResponse

`func NewStreamSubscriptionSinkResponse() *StreamSubscriptionSinkResponse`

NewStreamSubscriptionSinkResponse instantiates a new StreamSubscriptionSinkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSubscriptionSinkResponseWithDefaults

`func NewStreamSubscriptionSinkResponseWithDefaults() *StreamSubscriptionSinkResponse`

NewStreamSubscriptionSinkResponseWithDefaults instantiates a new StreamSubscriptionSinkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *StreamSubscriptionSinkResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *StreamSubscriptionSinkResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *StreamSubscriptionSinkResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *StreamSubscriptionSinkResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetType

`func (o *StreamSubscriptionSinkResponse) GetType() StreamSubscriptionSinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamSubscriptionSinkResponse) GetTypeOk() (*StreamSubscriptionSinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamSubscriptionSinkResponse) SetType(v StreamSubscriptionSinkType)`

SetType sets Type field to given value.

### HasType

`func (o *StreamSubscriptionSinkResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBatchEnabled

`func (o *StreamSubscriptionSinkResponse) GetBatchEnabled() bool`

GetBatchEnabled returns the BatchEnabled field if non-nil, zero value otherwise.

### GetBatchEnabledOk

`func (o *StreamSubscriptionSinkResponse) GetBatchEnabledOk() (*bool, bool)`

GetBatchEnabledOk returns a tuple with the BatchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchEnabled

`func (o *StreamSubscriptionSinkResponse) SetBatchEnabled(v bool)`

SetBatchEnabled sets BatchEnabled field to given value.

### HasBatchEnabled

`func (o *StreamSubscriptionSinkResponse) HasBatchEnabled() bool`

HasBatchEnabled returns a boolean if a field has been set.

### GetBatchSizeMax

`func (o *StreamSubscriptionSinkResponse) GetBatchSizeMax() int32`

GetBatchSizeMax returns the BatchSizeMax field if non-nil, zero value otherwise.

### GetBatchSizeMaxOk

`func (o *StreamSubscriptionSinkResponse) GetBatchSizeMaxOk() (*int32, bool)`

GetBatchSizeMaxOk returns a tuple with the BatchSizeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSizeMax

`func (o *StreamSubscriptionSinkResponse) SetBatchSizeMax(v int32)`

SetBatchSizeMax sets BatchSizeMax field to given value.

### HasBatchSizeMax

`func (o *StreamSubscriptionSinkResponse) HasBatchSizeMax() bool`

HasBatchSizeMax returns a boolean if a field has been set.

### GetBatchWaitTimeMax

`func (o *StreamSubscriptionSinkResponse) GetBatchWaitTimeMax() int32`

GetBatchWaitTimeMax returns the BatchWaitTimeMax field if non-nil, zero value otherwise.

### GetBatchWaitTimeMaxOk

`func (o *StreamSubscriptionSinkResponse) GetBatchWaitTimeMaxOk() (*int32, bool)`

GetBatchWaitTimeMaxOk returns a tuple with the BatchWaitTimeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchWaitTimeMax

`func (o *StreamSubscriptionSinkResponse) SetBatchWaitTimeMax(v int32)`

SetBatchWaitTimeMax sets BatchWaitTimeMax field to given value.

### HasBatchWaitTimeMax

`func (o *StreamSubscriptionSinkResponse) HasBatchWaitTimeMax() bool`

HasBatchWaitTimeMax returns a boolean if a field has been set.

### GetCredential

`func (o *StreamSubscriptionSinkResponse) GetCredential() StreamSubscriptionSinkCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *StreamSubscriptionSinkResponse) GetCredentialOk() (*StreamSubscriptionSinkCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *StreamSubscriptionSinkResponse) SetCredential(v StreamSubscriptionSinkCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *StreamSubscriptionSinkResponse) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetSettings

`func (o *StreamSubscriptionSinkResponse) GetSettings() StreamSubscriptionSinkSetting`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *StreamSubscriptionSinkResponse) GetSettingsOk() (*StreamSubscriptionSinkSetting, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *StreamSubscriptionSinkResponse) SetSettings(v StreamSubscriptionSinkSetting)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *StreamSubscriptionSinkResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetHost

`func (o *StreamSubscriptionSinkResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StreamSubscriptionSinkResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StreamSubscriptionSinkResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *StreamSubscriptionSinkResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


