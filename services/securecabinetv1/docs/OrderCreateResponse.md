# OrderCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderNumber** | **string** | Order Number | 

## Methods

### NewOrderCreateResponse

`func NewOrderCreateResponse(orderNumber string, ) *OrderCreateResponse`

NewOrderCreateResponse instantiates a new OrderCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateResponseWithDefaults

`func NewOrderCreateResponseWithDefaults() *OrderCreateResponse`

NewOrderCreateResponseWithDefaults instantiates a new OrderCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderNumber

`func (o *OrderCreateResponse) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *OrderCreateResponse) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *OrderCreateResponse) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


