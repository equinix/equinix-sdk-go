# PurchaseOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** | Purchase Order Number | [optional] 
**PurchaseOrderType** | [**PurchaseOrderPurchaseOrderType**](PurchaseOrderPurchaseOrderType.md) |  | 

## Methods

### NewPurchaseOrder

`func NewPurchaseOrder(purchaseOrderType PurchaseOrderPurchaseOrderType, ) *PurchaseOrder`

NewPurchaseOrder instantiates a new PurchaseOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderWithDefaults

`func NewPurchaseOrderWithDefaults() *PurchaseOrder`

NewPurchaseOrderWithDefaults instantiates a new PurchaseOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *PurchaseOrder) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PurchaseOrder) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PurchaseOrder) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PurchaseOrder) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPurchaseOrderType

`func (o *PurchaseOrder) GetPurchaseOrderType() PurchaseOrderPurchaseOrderType`

GetPurchaseOrderType returns the PurchaseOrderType field if non-nil, zero value otherwise.

### GetPurchaseOrderTypeOk

`func (o *PurchaseOrder) GetPurchaseOrderTypeOk() (*PurchaseOrderPurchaseOrderType, bool)`

GetPurchaseOrderTypeOk returns a tuple with the PurchaseOrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderType

`func (o *PurchaseOrder) SetPurchaseOrderType(v PurchaseOrderPurchaseOrderType)`

SetPurchaseOrderType sets PurchaseOrderType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


