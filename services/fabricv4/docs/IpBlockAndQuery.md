# IpBlockAndQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | **string** |  | 
**Operator** | [**IpBlockAndQueryOperator**](IpBlockAndQueryOperator.md) |  | 
**Values** | **[]string** |  | 

## Methods

### NewIpBlockAndQuery

`func NewIpBlockAndQuery(property string, operator IpBlockAndQueryOperator, values []string, ) *IpBlockAndQuery`

NewIpBlockAndQuery instantiates a new IpBlockAndQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlockAndQueryWithDefaults

`func NewIpBlockAndQueryWithDefaults() *IpBlockAndQuery`

NewIpBlockAndQueryWithDefaults instantiates a new IpBlockAndQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *IpBlockAndQuery) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *IpBlockAndQuery) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *IpBlockAndQuery) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetOperator

`func (o *IpBlockAndQuery) GetOperator() IpBlockAndQueryOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *IpBlockAndQuery) GetOperatorOk() (*IpBlockAndQueryOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *IpBlockAndQuery) SetOperator(v IpBlockAndQueryOperator)`

SetOperator sets Operator field to given value.


### GetValues

`func (o *IpBlockAndQuery) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *IpBlockAndQuery) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *IpBlockAndQuery) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


