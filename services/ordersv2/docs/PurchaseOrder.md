# PurchaseOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**PurchaseOrderType**](PurchaseOrderType.md) |  | [optional] 
**Number** | Pointer to **string** | Purchase order number | [optional] 
**ClosingDate** | Pointer to **string** | Closing date of purchase order. | [optional] 

## Methods

### NewPurchaseOrder

`func NewPurchaseOrder() *PurchaseOrder`

NewPurchaseOrder instantiates a new PurchaseOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderWithDefaults

`func NewPurchaseOrderWithDefaults() *PurchaseOrder`

NewPurchaseOrderWithDefaults instantiates a new PurchaseOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PurchaseOrder) GetType() PurchaseOrderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PurchaseOrder) GetTypeOk() (*PurchaseOrderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PurchaseOrder) SetType(v PurchaseOrderType)`

SetType sets Type field to given value.

### HasType

`func (o *PurchaseOrder) HasType() bool`

HasType returns a boolean if a field has been set.

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

### GetClosingDate

`func (o *PurchaseOrder) GetClosingDate() string`

GetClosingDate returns the ClosingDate field if non-nil, zero value otherwise.

### GetClosingDateOk

`func (o *PurchaseOrder) GetClosingDateOk() (*string, bool)`

GetClosingDateOk returns a tuple with the ClosingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingDate

`func (o *PurchaseOrder) SetClosingDate(v string)`

SetClosingDate sets ClosingDate field to given value.

### HasClosingDate

`func (o *PurchaseOrder) HasClosingDate() bool`

HasClosingDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


