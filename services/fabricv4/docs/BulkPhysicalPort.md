# BulkPhysicalPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PhysicalPort**](PhysicalPort.md) | add physical ports to virtual port | [optional] 

## Methods

### NewBulkPhysicalPort

`func NewBulkPhysicalPort() *BulkPhysicalPort`

NewBulkPhysicalPort instantiates a new BulkPhysicalPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkPhysicalPortWithDefaults

`func NewBulkPhysicalPortWithDefaults() *BulkPhysicalPort`

NewBulkPhysicalPortWithDefaults instantiates a new BulkPhysicalPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BulkPhysicalPort) GetData() []PhysicalPort`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BulkPhysicalPort) GetDataOk() (*[]PhysicalPort, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BulkPhysicalPort) SetData(v []PhysicalPort)`

SetData sets Data field to given value.

### HasData

`func (o *BulkPhysicalPort) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


