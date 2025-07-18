# StreamSubscriptionSinkSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIndex** | Pointer to **string** | event index | [optional] 
**MetricIndex** | Pointer to **string** | metric index | [optional] 
**Source** | Pointer to **string** | source | [optional] 
**ApplicationKey** | Pointer to **string** | Application key | [optional] 
**EventUri** | Pointer to **string** | event uri | [optional] 
**MetricUri** | Pointer to **string** | metric uri | [optional] 
**Format** | Pointer to [**StreamSubscriptionSinkSettingFormat**](StreamSubscriptionSinkSettingFormat.md) |  | [optional] [default to STREAMSUBSCRIPTIONSINKSETTINGFORMAT_CLOUDEVENT]

## Methods

### NewStreamSubscriptionSinkSetting

`func NewStreamSubscriptionSinkSetting() *StreamSubscriptionSinkSetting`

NewStreamSubscriptionSinkSetting instantiates a new StreamSubscriptionSinkSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSubscriptionSinkSettingWithDefaults

`func NewStreamSubscriptionSinkSettingWithDefaults() *StreamSubscriptionSinkSetting`

NewStreamSubscriptionSinkSettingWithDefaults instantiates a new StreamSubscriptionSinkSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIndex

`func (o *StreamSubscriptionSinkSetting) GetEventIndex() string`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *StreamSubscriptionSinkSetting) GetEventIndexOk() (*string, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *StreamSubscriptionSinkSetting) SetEventIndex(v string)`

SetEventIndex sets EventIndex field to given value.

### HasEventIndex

`func (o *StreamSubscriptionSinkSetting) HasEventIndex() bool`

HasEventIndex returns a boolean if a field has been set.

### GetMetricIndex

`func (o *StreamSubscriptionSinkSetting) GetMetricIndex() string`

GetMetricIndex returns the MetricIndex field if non-nil, zero value otherwise.

### GetMetricIndexOk

`func (o *StreamSubscriptionSinkSetting) GetMetricIndexOk() (*string, bool)`

GetMetricIndexOk returns a tuple with the MetricIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricIndex

`func (o *StreamSubscriptionSinkSetting) SetMetricIndex(v string)`

SetMetricIndex sets MetricIndex field to given value.

### HasMetricIndex

`func (o *StreamSubscriptionSinkSetting) HasMetricIndex() bool`

HasMetricIndex returns a boolean if a field has been set.

### GetSource

`func (o *StreamSubscriptionSinkSetting) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StreamSubscriptionSinkSetting) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StreamSubscriptionSinkSetting) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StreamSubscriptionSinkSetting) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetApplicationKey

`func (o *StreamSubscriptionSinkSetting) GetApplicationKey() string`

GetApplicationKey returns the ApplicationKey field if non-nil, zero value otherwise.

### GetApplicationKeyOk

`func (o *StreamSubscriptionSinkSetting) GetApplicationKeyOk() (*string, bool)`

GetApplicationKeyOk returns a tuple with the ApplicationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKey

`func (o *StreamSubscriptionSinkSetting) SetApplicationKey(v string)`

SetApplicationKey sets ApplicationKey field to given value.

### HasApplicationKey

`func (o *StreamSubscriptionSinkSetting) HasApplicationKey() bool`

HasApplicationKey returns a boolean if a field has been set.

### GetEventUri

`func (o *StreamSubscriptionSinkSetting) GetEventUri() string`

GetEventUri returns the EventUri field if non-nil, zero value otherwise.

### GetEventUriOk

`func (o *StreamSubscriptionSinkSetting) GetEventUriOk() (*string, bool)`

GetEventUriOk returns a tuple with the EventUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUri

`func (o *StreamSubscriptionSinkSetting) SetEventUri(v string)`

SetEventUri sets EventUri field to given value.

### HasEventUri

`func (o *StreamSubscriptionSinkSetting) HasEventUri() bool`

HasEventUri returns a boolean if a field has been set.

### GetMetricUri

`func (o *StreamSubscriptionSinkSetting) GetMetricUri() string`

GetMetricUri returns the MetricUri field if non-nil, zero value otherwise.

### GetMetricUriOk

`func (o *StreamSubscriptionSinkSetting) GetMetricUriOk() (*string, bool)`

GetMetricUriOk returns a tuple with the MetricUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricUri

`func (o *StreamSubscriptionSinkSetting) SetMetricUri(v string)`

SetMetricUri sets MetricUri field to given value.

### HasMetricUri

`func (o *StreamSubscriptionSinkSetting) HasMetricUri() bool`

HasMetricUri returns a boolean if a field has been set.

### GetFormat

`func (o *StreamSubscriptionSinkSetting) GetFormat() StreamSubscriptionSinkSettingFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *StreamSubscriptionSinkSetting) GetFormatOk() (*StreamSubscriptionSinkSettingFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *StreamSubscriptionSinkSetting) SetFormat(v StreamSubscriptionSinkSettingFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *StreamSubscriptionSinkSetting) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


