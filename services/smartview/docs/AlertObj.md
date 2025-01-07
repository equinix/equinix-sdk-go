# AlertObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**[]AlertDto1**](AlertDto1.md) |  | [optional] 
**TotalCount** | Pointer to **float32** | totalCount | [optional] 

## Methods

### NewAlertObj

`func NewAlertObj() *AlertObj`

NewAlertObj instantiates a new AlertObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertObjWithDefaults

`func NewAlertObjWithDefaults() *AlertObj`

NewAlertObjWithDefaults instantiates a new AlertObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *AlertObj) GetAlerts() []AlertDto1`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *AlertObj) GetAlertsOk() (*[]AlertDto1, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *AlertObj) SetAlerts(v []AlertDto1)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *AlertObj) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetTotalCount

`func (o *AlertObj) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AlertObj) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AlertObj) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *AlertObj) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


