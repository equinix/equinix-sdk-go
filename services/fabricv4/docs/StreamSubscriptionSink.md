# StreamSubscriptionSink

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

## Methods

### NewStreamSubscriptionSink

`func NewStreamSubscriptionSink() *StreamSubscriptionSink`

NewStreamSubscriptionSink instantiates a new StreamSubscriptionSink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSubscriptionSinkWithDefaults

`func NewStreamSubscriptionSinkWithDefaults() *StreamSubscriptionSink`

NewStreamSubscriptionSinkWithDefaults instantiates a new StreamSubscriptionSink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *StreamSubscriptionSink) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *StreamSubscriptionSink) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *StreamSubscriptionSink) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *StreamSubscriptionSink) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetType

`func (o *StreamSubscriptionSink) GetType() StreamSubscriptionSinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamSubscriptionSink) GetTypeOk() (*StreamSubscriptionSinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamSubscriptionSink) SetType(v StreamSubscriptionSinkType)`

SetType sets Type field to given value.

### HasType

`func (o *StreamSubscriptionSink) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBatchEnabled

`func (o *StreamSubscriptionSink) GetBatchEnabled() bool`

GetBatchEnabled returns the BatchEnabled field if non-nil, zero value otherwise.

### GetBatchEnabledOk

`func (o *StreamSubscriptionSink) GetBatchEnabledOk() (*bool, bool)`

GetBatchEnabledOk returns a tuple with the BatchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchEnabled

`func (o *StreamSubscriptionSink) SetBatchEnabled(v bool)`

SetBatchEnabled sets BatchEnabled field to given value.

### HasBatchEnabled

`func (o *StreamSubscriptionSink) HasBatchEnabled() bool`

HasBatchEnabled returns a boolean if a field has been set.

### GetBatchSizeMax

`func (o *StreamSubscriptionSink) GetBatchSizeMax() int32`

GetBatchSizeMax returns the BatchSizeMax field if non-nil, zero value otherwise.

### GetBatchSizeMaxOk

`func (o *StreamSubscriptionSink) GetBatchSizeMaxOk() (*int32, bool)`

GetBatchSizeMaxOk returns a tuple with the BatchSizeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSizeMax

`func (o *StreamSubscriptionSink) SetBatchSizeMax(v int32)`

SetBatchSizeMax sets BatchSizeMax field to given value.

### HasBatchSizeMax

`func (o *StreamSubscriptionSink) HasBatchSizeMax() bool`

HasBatchSizeMax returns a boolean if a field has been set.

### GetBatchWaitTimeMax

`func (o *StreamSubscriptionSink) GetBatchWaitTimeMax() int32`

GetBatchWaitTimeMax returns the BatchWaitTimeMax field if non-nil, zero value otherwise.

### GetBatchWaitTimeMaxOk

`func (o *StreamSubscriptionSink) GetBatchWaitTimeMaxOk() (*int32, bool)`

GetBatchWaitTimeMaxOk returns a tuple with the BatchWaitTimeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchWaitTimeMax

`func (o *StreamSubscriptionSink) SetBatchWaitTimeMax(v int32)`

SetBatchWaitTimeMax sets BatchWaitTimeMax field to given value.

### HasBatchWaitTimeMax

`func (o *StreamSubscriptionSink) HasBatchWaitTimeMax() bool`

HasBatchWaitTimeMax returns a boolean if a field has been set.

### GetCredential

`func (o *StreamSubscriptionSink) GetCredential() StreamSubscriptionSinkCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *StreamSubscriptionSink) GetCredentialOk() (*StreamSubscriptionSinkCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *StreamSubscriptionSink) SetCredential(v StreamSubscriptionSinkCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *StreamSubscriptionSink) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetSettings

`func (o *StreamSubscriptionSink) GetSettings() StreamSubscriptionSinkSetting`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *StreamSubscriptionSink) GetSettingsOk() (*StreamSubscriptionSinkSetting, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *StreamSubscriptionSink) SetSettings(v StreamSubscriptionSinkSetting)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *StreamSubscriptionSink) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


