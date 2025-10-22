# MetricFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Possible field names to use on filters:  * &#x60;/subject&#x60; - Metric subject description (required and limited to 1 value)  * &#x60;/name&#x60; - Metric names (required)  * &#x60;/dataPoints/endDateTime&#x60; - Time of Metrics  | [optional] 
**Operator** | Pointer to **string** | Possible operators to use on filters:  * &#x60;&#x3D;&#x60; - equal  * &#x60;&gt;&#x60; - greater than  * &#x60;&gt;&#x3D;&#x60; - greater than or equal to  * &#x60;&lt;&#x60; - less than  * &#x60;&lt;&#x3D;&#x60; - less than or equal to  * &#x60;BETWEEN&#x60; - between  * &#x60;IN&#x60; - in  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMetricFilter

`func NewMetricFilter() *MetricFilter`

NewMetricFilter instantiates a new MetricFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricFilterWithDefaults

`func NewMetricFilterWithDefaults() *MetricFilter`

NewMetricFilterWithDefaults instantiates a new MetricFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *MetricFilter) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *MetricFilter) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *MetricFilter) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *MetricFilter) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetOperator

`func (o *MetricFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MetricFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MetricFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *MetricFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *MetricFilter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MetricFilter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MetricFilter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *MetricFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


