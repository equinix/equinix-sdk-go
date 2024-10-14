# BulkPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Port**](Port.md) | Ports that are part of BulkPort | [optional] 

## Methods

### NewBulkPort

`func NewBulkPort() *BulkPort`

NewBulkPort instantiates a new BulkPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkPortWithDefaults

`func NewBulkPortWithDefaults() *BulkPort`

NewBulkPortWithDefaults instantiates a new BulkPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BulkPort) GetData() []Port`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BulkPort) GetDataOk() (*[]Port, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BulkPort) SetData(v []Port)`

SetData sets Data field to given value.

### HasData

`func (o *BulkPort) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


