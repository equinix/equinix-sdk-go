# OrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderNumber** | Pointer to **string** | Order Number. | [optional] 
**OrderLineItemNumber** | Pointer to **string** | Order Line Item Number. | [optional] 
**PurchaseOrder** | Pointer to [**PurchaseOrder**](PurchaseOrder.md) |  | [optional] 
**CustomerReferenceId** | Pointer to **string** | Customer Reference Identifier. | [optional] 

## Methods

### NewOrderResponse

`func NewOrderResponse() *OrderResponse`

NewOrderResponse instantiates a new OrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResponseWithDefaults

`func NewOrderResponseWithDefaults() *OrderResponse`

NewOrderResponseWithDefaults instantiates a new OrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderNumber

`func (o *OrderResponse) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *OrderResponse) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *OrderResponse) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *OrderResponse) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetOrderLineItemNumber

`func (o *OrderResponse) GetOrderLineItemNumber() string`

GetOrderLineItemNumber returns the OrderLineItemNumber field if non-nil, zero value otherwise.

### GetOrderLineItemNumberOk

`func (o *OrderResponse) GetOrderLineItemNumberOk() (*string, bool)`

GetOrderLineItemNumberOk returns a tuple with the OrderLineItemNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLineItemNumber

`func (o *OrderResponse) SetOrderLineItemNumber(v string)`

SetOrderLineItemNumber sets OrderLineItemNumber field to given value.

### HasOrderLineItemNumber

`func (o *OrderResponse) HasOrderLineItemNumber() bool`

HasOrderLineItemNumber returns a boolean if a field has been set.

### GetPurchaseOrder

`func (o *OrderResponse) GetPurchaseOrder() PurchaseOrder`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *OrderResponse) GetPurchaseOrderOk() (*PurchaseOrder, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *OrderResponse) SetPurchaseOrder(v PurchaseOrder)`

SetPurchaseOrder sets PurchaseOrder field to given value.

### HasPurchaseOrder

`func (o *OrderResponse) HasPurchaseOrder() bool`

HasPurchaseOrder returns a boolean if a field has been set.

### GetCustomerReferenceId

`func (o *OrderResponse) GetCustomerReferenceId() string`

GetCustomerReferenceId returns the CustomerReferenceId field if non-nil, zero value otherwise.

### GetCustomerReferenceIdOk

`func (o *OrderResponse) GetCustomerReferenceIdOk() (*string, bool)`

GetCustomerReferenceIdOk returns a tuple with the CustomerReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReferenceId

`func (o *OrderResponse) SetCustomerReferenceId(v string)`

SetCustomerReferenceId sets CustomerReferenceId field to given value.

### HasCustomerReferenceId

`func (o *OrderResponse) HasCustomerReferenceId() bool`

HasCustomerReferenceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


