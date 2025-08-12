# StreamSubscriptionOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventsDeliveredCount** | Pointer to **int32** | count of delivered events | [optional] 
**MetricsDeliveredCount** | Pointer to **int32** | count of delivered metrics | [optional] 
**AlertsDeliveredCount** | Pointer to **int32** | count of delivered alerts | [optional] 
**LastSuccessfulDeliveryDateTime** | Pointer to **time.Time** | last successful date time of delivered event, metric, or alert | [optional] 
**SuspendedDateTime** | Pointer to **time.Time** | suspended date time of stream subscription delivery for event, metric, or alert | [optional] 
**Errors** | Pointer to [**[]StreamSubscriptionOperationErrors**](StreamSubscriptionOperationErrors.md) | List of error information for stream subscription delivery | [optional] 

## Methods

### NewStreamSubscriptionOperation

`func NewStreamSubscriptionOperation() *StreamSubscriptionOperation`

NewStreamSubscriptionOperation instantiates a new StreamSubscriptionOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSubscriptionOperationWithDefaults

`func NewStreamSubscriptionOperationWithDefaults() *StreamSubscriptionOperation`

NewStreamSubscriptionOperationWithDefaults instantiates a new StreamSubscriptionOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventsDeliveredCount

`func (o *StreamSubscriptionOperation) GetEventsDeliveredCount() int32`

GetEventsDeliveredCount returns the EventsDeliveredCount field if non-nil, zero value otherwise.

### GetEventsDeliveredCountOk

`func (o *StreamSubscriptionOperation) GetEventsDeliveredCountOk() (*int32, bool)`

GetEventsDeliveredCountOk returns a tuple with the EventsDeliveredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsDeliveredCount

`func (o *StreamSubscriptionOperation) SetEventsDeliveredCount(v int32)`

SetEventsDeliveredCount sets EventsDeliveredCount field to given value.

### HasEventsDeliveredCount

`func (o *StreamSubscriptionOperation) HasEventsDeliveredCount() bool`

HasEventsDeliveredCount returns a boolean if a field has been set.

### GetMetricsDeliveredCount

`func (o *StreamSubscriptionOperation) GetMetricsDeliveredCount() int32`

GetMetricsDeliveredCount returns the MetricsDeliveredCount field if non-nil, zero value otherwise.

### GetMetricsDeliveredCountOk

`func (o *StreamSubscriptionOperation) GetMetricsDeliveredCountOk() (*int32, bool)`

GetMetricsDeliveredCountOk returns a tuple with the MetricsDeliveredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsDeliveredCount

`func (o *StreamSubscriptionOperation) SetMetricsDeliveredCount(v int32)`

SetMetricsDeliveredCount sets MetricsDeliveredCount field to given value.

### HasMetricsDeliveredCount

`func (o *StreamSubscriptionOperation) HasMetricsDeliveredCount() bool`

HasMetricsDeliveredCount returns a boolean if a field has been set.

### GetAlertsDeliveredCount

`func (o *StreamSubscriptionOperation) GetAlertsDeliveredCount() int32`

GetAlertsDeliveredCount returns the AlertsDeliveredCount field if non-nil, zero value otherwise.

### GetAlertsDeliveredCountOk

`func (o *StreamSubscriptionOperation) GetAlertsDeliveredCountOk() (*int32, bool)`

GetAlertsDeliveredCountOk returns a tuple with the AlertsDeliveredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsDeliveredCount

`func (o *StreamSubscriptionOperation) SetAlertsDeliveredCount(v int32)`

SetAlertsDeliveredCount sets AlertsDeliveredCount field to given value.

### HasAlertsDeliveredCount

`func (o *StreamSubscriptionOperation) HasAlertsDeliveredCount() bool`

HasAlertsDeliveredCount returns a boolean if a field has been set.

### GetLastSuccessfulDeliveryDateTime

`func (o *StreamSubscriptionOperation) GetLastSuccessfulDeliveryDateTime() time.Time`

GetLastSuccessfulDeliveryDateTime returns the LastSuccessfulDeliveryDateTime field if non-nil, zero value otherwise.

### GetLastSuccessfulDeliveryDateTimeOk

`func (o *StreamSubscriptionOperation) GetLastSuccessfulDeliveryDateTimeOk() (*time.Time, bool)`

GetLastSuccessfulDeliveryDateTimeOk returns a tuple with the LastSuccessfulDeliveryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulDeliveryDateTime

`func (o *StreamSubscriptionOperation) SetLastSuccessfulDeliveryDateTime(v time.Time)`

SetLastSuccessfulDeliveryDateTime sets LastSuccessfulDeliveryDateTime field to given value.

### HasLastSuccessfulDeliveryDateTime

`func (o *StreamSubscriptionOperation) HasLastSuccessfulDeliveryDateTime() bool`

HasLastSuccessfulDeliveryDateTime returns a boolean if a field has been set.

### GetSuspendedDateTime

`func (o *StreamSubscriptionOperation) GetSuspendedDateTime() time.Time`

GetSuspendedDateTime returns the SuspendedDateTime field if non-nil, zero value otherwise.

### GetSuspendedDateTimeOk

`func (o *StreamSubscriptionOperation) GetSuspendedDateTimeOk() (*time.Time, bool)`

GetSuspendedDateTimeOk returns a tuple with the SuspendedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedDateTime

`func (o *StreamSubscriptionOperation) SetSuspendedDateTime(v time.Time)`

SetSuspendedDateTime sets SuspendedDateTime field to given value.

### HasSuspendedDateTime

`func (o *StreamSubscriptionOperation) HasSuspendedDateTime() bool`

HasSuspendedDateTime returns a boolean if a field has been set.

### GetErrors

`func (o *StreamSubscriptionOperation) GetErrors() []StreamSubscriptionOperationErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *StreamSubscriptionOperation) GetErrorsOk() (*[]StreamSubscriptionOperationErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *StreamSubscriptionOperation) SetErrors(v []StreamSubscriptionOperationErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *StreamSubscriptionOperation) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


