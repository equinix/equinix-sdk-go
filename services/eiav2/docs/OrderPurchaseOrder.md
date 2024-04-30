# OrderPurchaseOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**Number** | **string** |  | 
**Type** | Pointer to [**OrderPurchaseOrderAllOfType**](OrderPurchaseOrderAllOfType.md) |  | [optional] 

## Methods

### NewOrderPurchaseOrder

`func NewOrderPurchaseOrder(href string, number string, ) *OrderPurchaseOrder`

NewOrderPurchaseOrder instantiates a new OrderPurchaseOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderPurchaseOrderWithDefaults

`func NewOrderPurchaseOrderWithDefaults() *OrderPurchaseOrder`

NewOrderPurchaseOrderWithDefaults instantiates a new OrderPurchaseOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *OrderPurchaseOrder) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *OrderPurchaseOrder) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *OrderPurchaseOrder) SetHref(v string)`

SetHref sets Href field to given value.


### GetNumber

`func (o *OrderPurchaseOrder) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OrderPurchaseOrder) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OrderPurchaseOrder) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetType

`func (o *OrderPurchaseOrder) GetType() OrderPurchaseOrderAllOfType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderPurchaseOrder) GetTypeOk() (*OrderPurchaseOrderAllOfType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderPurchaseOrder) SetType(v OrderPurchaseOrderAllOfType)`

SetType sets Type field to given value.

### HasType

`func (o *OrderPurchaseOrder) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


