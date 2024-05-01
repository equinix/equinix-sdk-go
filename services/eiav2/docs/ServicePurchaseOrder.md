# ServicePurchaseOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**PurchaseOrderType**](PurchaseOrderType.md) |  | [optional] 
**Number** | **string** | Purchase order number  | 
**Amount** | Pointer to **float32** | Amount | [optional] 
**StartDate** | Pointer to **string** | Start date in YYYY-MM-DD format | [optional] 
**EndDate** | Pointer to **string** | End date in YYYY-MM-DD format | [optional] 
**Description** | Pointer to **string** | Purchase order description | [optional] 
**Attachment** | Pointer to [**ServicePurchaseOrderAttachment**](ServicePurchaseOrderAttachment.md) |  | [optional] 

## Methods

### NewServicePurchaseOrder

`func NewServicePurchaseOrder(number string, ) *ServicePurchaseOrder`

NewServicePurchaseOrder instantiates a new ServicePurchaseOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePurchaseOrderWithDefaults

`func NewServicePurchaseOrderWithDefaults() *ServicePurchaseOrder`

NewServicePurchaseOrderWithDefaults instantiates a new ServicePurchaseOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServicePurchaseOrder) GetType() PurchaseOrderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServicePurchaseOrder) GetTypeOk() (*PurchaseOrderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServicePurchaseOrder) SetType(v PurchaseOrderType)`

SetType sets Type field to given value.

### HasType

`func (o *ServicePurchaseOrder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNumber

`func (o *ServicePurchaseOrder) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ServicePurchaseOrder) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ServicePurchaseOrder) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetAmount

`func (o *ServicePurchaseOrder) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ServicePurchaseOrder) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ServicePurchaseOrder) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ServicePurchaseOrder) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetStartDate

`func (o *ServicePurchaseOrder) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ServicePurchaseOrder) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ServicePurchaseOrder) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ServicePurchaseOrder) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ServicePurchaseOrder) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ServicePurchaseOrder) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ServicePurchaseOrder) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ServicePurchaseOrder) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDescription

`func (o *ServicePurchaseOrder) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicePurchaseOrder) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicePurchaseOrder) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicePurchaseOrder) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttachment

`func (o *ServicePurchaseOrder) GetAttachment() ServicePurchaseOrderAttachment`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *ServicePurchaseOrder) GetAttachmentOk() (*ServicePurchaseOrderAttachment, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *ServicePurchaseOrder) SetAttachment(v ServicePurchaseOrderAttachment)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *ServicePurchaseOrder) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


