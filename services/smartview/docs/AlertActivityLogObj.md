# AlertActivityLogObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**[]AlertDto2**](AlertDto2.md) |  | [optional] 
**TotalCount** | Pointer to **float32** | totalCount | [optional] 

## Methods

### NewAlertActivityLogObj

`func NewAlertActivityLogObj() *AlertActivityLogObj`

NewAlertActivityLogObj instantiates a new AlertActivityLogObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertActivityLogObjWithDefaults

`func NewAlertActivityLogObjWithDefaults() *AlertActivityLogObj`

NewAlertActivityLogObjWithDefaults instantiates a new AlertActivityLogObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *AlertActivityLogObj) GetAlerts() []AlertDto2`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *AlertActivityLogObj) GetAlertsOk() (*[]AlertDto2, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *AlertActivityLogObj) SetAlerts(v []AlertDto2)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *AlertActivityLogObj) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetTotalCount

`func (o *AlertActivityLogObj) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AlertActivityLogObj) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AlertActivityLogObj) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *AlertActivityLogObj) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


