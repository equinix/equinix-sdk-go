# PurchaseOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** | Purchase Order number. | [optional] 
**Amount** | Pointer to **string** | Purchase Order amount. | [optional] 
**StartDate** | Pointer to **string** | Date the Purchase Order takes effect. | [optional] 
**EndDate** | Pointer to **string** | Date the Purchase Order expires. | [optional] 
**AttachmentId** | Pointer to **string** | Identifier of the uploaded purchase order document. | [optional] 
**Type** | Pointer to [**PurchaseOrderType**](PurchaseOrderType.md) |  | [optional] 

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

### GetAmount

`func (o *PurchaseOrder) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PurchaseOrder) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PurchaseOrder) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PurchaseOrder) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetStartDate

`func (o *PurchaseOrder) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PurchaseOrder) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PurchaseOrder) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PurchaseOrder) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *PurchaseOrder) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PurchaseOrder) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PurchaseOrder) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PurchaseOrder) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetAttachmentId

`func (o *PurchaseOrder) GetAttachmentId() string`

GetAttachmentId returns the AttachmentId field if non-nil, zero value otherwise.

### GetAttachmentIdOk

`func (o *PurchaseOrder) GetAttachmentIdOk() (*string, bool)`

GetAttachmentIdOk returns a tuple with the AttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentId

`func (o *PurchaseOrder) SetAttachmentId(v string)`

SetAttachmentId sets AttachmentId field to given value.

### HasAttachmentId

`func (o *PurchaseOrder) HasAttachmentId() bool`

HasAttachmentId returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


