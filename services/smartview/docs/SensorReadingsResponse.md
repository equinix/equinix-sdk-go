# SensorReadingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SensorReading**](SensorReading.md) | List of data objects | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewSensorReadingsResponse

`func NewSensorReadingsResponse(data []SensorReading, pagination Pagination, ) *SensorReadingsResponse`

NewSensorReadingsResponse instantiates a new SensorReadingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorReadingsResponseWithDefaults

`func NewSensorReadingsResponseWithDefaults() *SensorReadingsResponse`

NewSensorReadingsResponseWithDefaults instantiates a new SensorReadingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SensorReadingsResponse) GetData() []SensorReading`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SensorReadingsResponse) GetDataOk() (*[]SensorReading, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SensorReadingsResponse) SetData(v []SensorReading)`

SetData sets Data field to given value.


### GetPagination

`func (o *SensorReadingsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SensorReadingsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SensorReadingsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


