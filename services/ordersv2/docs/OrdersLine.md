# OrdersLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineId** | Pointer to **string** | Unique identifier associated with line item.  | [optional] 
**ParentlineId** | Pointer to **string** | Unique identifier associated with &#x60;lineId&#x60;to denote hierarchy of the product. When this field is null, item considered to be at root level. | [optional] 
**RootlineId** | Pointer to **string** | Unique identifier to specify root item. Useful when product order hierarchy exceeds 2. | [optional] 
**Status** | Pointer to **string** | The current status of the order or service | [optional] 
**Description** | Pointer to **string** | Description of the orders. | [optional] 
**BillingStartDateTime** | Pointer to **time.Time** | Billing commencement date and time in UTC timezone | [optional] 
**EstimatedCompletionDateTime** | Pointer to **time.Time** | The estimated completion date and time in UTC timezone | [optional] 
**UnitPricing** | Pointer to [**[]Price**](Price.md) | Unit Pricing details | [optional] 
**TotalPricing** | Pointer to [**[]Price**](Price.md) | Total Pricing details | [optional] 
**ProductType** | Pointer to [**OrdersLineProductType**](OrdersLineProductType.md) |  | [optional] 
**ProductCode** | Pointer to **string** | Product code associated with a particular product | [optional] 
**ProductName** | Pointer to **string** | Product name | [optional] 
**Ibx** | Pointer to **string** | Location code of Equinix datacenter | [optional] 
**Cage** | Pointer to **string** | Cage ID | [optional] 
**Cabinets** | Pointer to **[]string** | One or more cabinets ID | [optional] 
**RequestType** | Pointer to [**OrdersLineRequestType**](OrdersLineRequestType.md) |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | File(s) attached to the Orders. To learn about including attachments in your request, see POST Attachments API. | [optional] 
**AdditionalInfo** | Pointer to [**[]ProductAdditionalInfoInner**](ProductAdditionalInfoInner.md) | This section is reserved to display product specific information | [optional] 
**PurchaseOrder** | Pointer to [**PurchaseOrder**](PurchaseOrder.md) |  | [optional] 
**Cancellable** | Pointer to **bool** | When &#x60;true&#x60;, order can be cancelled | [optional] [default to false]
**Modifiable** | Pointer to **bool** | When &#x60;true&#x60;, order can be modified | [optional] 

## Methods

### NewOrdersLine

`func NewOrdersLine() *OrdersLine`

NewOrdersLine instantiates a new OrdersLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersLineWithDefaults

`func NewOrdersLineWithDefaults() *OrdersLine`

NewOrdersLineWithDefaults instantiates a new OrdersLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineId

`func (o *OrdersLine) GetLineId() string`

GetLineId returns the LineId field if non-nil, zero value otherwise.

### GetLineIdOk

`func (o *OrdersLine) GetLineIdOk() (*string, bool)`

GetLineIdOk returns a tuple with the LineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineId

`func (o *OrdersLine) SetLineId(v string)`

SetLineId sets LineId field to given value.

### HasLineId

`func (o *OrdersLine) HasLineId() bool`

HasLineId returns a boolean if a field has been set.

### GetParentlineId

`func (o *OrdersLine) GetParentlineId() string`

GetParentlineId returns the ParentlineId field if non-nil, zero value otherwise.

### GetParentlineIdOk

`func (o *OrdersLine) GetParentlineIdOk() (*string, bool)`

GetParentlineIdOk returns a tuple with the ParentlineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentlineId

`func (o *OrdersLine) SetParentlineId(v string)`

SetParentlineId sets ParentlineId field to given value.

### HasParentlineId

`func (o *OrdersLine) HasParentlineId() bool`

HasParentlineId returns a boolean if a field has been set.

### GetRootlineId

`func (o *OrdersLine) GetRootlineId() string`

GetRootlineId returns the RootlineId field if non-nil, zero value otherwise.

### GetRootlineIdOk

`func (o *OrdersLine) GetRootlineIdOk() (*string, bool)`

GetRootlineIdOk returns a tuple with the RootlineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootlineId

`func (o *OrdersLine) SetRootlineId(v string)`

SetRootlineId sets RootlineId field to given value.

### HasRootlineId

`func (o *OrdersLine) HasRootlineId() bool`

HasRootlineId returns a boolean if a field has been set.

### GetStatus

`func (o *OrdersLine) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrdersLine) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrdersLine) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrdersLine) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *OrdersLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrdersLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrdersLine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrdersLine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBillingStartDateTime

`func (o *OrdersLine) GetBillingStartDateTime() time.Time`

GetBillingStartDateTime returns the BillingStartDateTime field if non-nil, zero value otherwise.

### GetBillingStartDateTimeOk

`func (o *OrdersLine) GetBillingStartDateTimeOk() (*time.Time, bool)`

GetBillingStartDateTimeOk returns a tuple with the BillingStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartDateTime

`func (o *OrdersLine) SetBillingStartDateTime(v time.Time)`

SetBillingStartDateTime sets BillingStartDateTime field to given value.

### HasBillingStartDateTime

`func (o *OrdersLine) HasBillingStartDateTime() bool`

HasBillingStartDateTime returns a boolean if a field has been set.

### GetEstimatedCompletionDateTime

`func (o *OrdersLine) GetEstimatedCompletionDateTime() time.Time`

GetEstimatedCompletionDateTime returns the EstimatedCompletionDateTime field if non-nil, zero value otherwise.

### GetEstimatedCompletionDateTimeOk

`func (o *OrdersLine) GetEstimatedCompletionDateTimeOk() (*time.Time, bool)`

GetEstimatedCompletionDateTimeOk returns a tuple with the EstimatedCompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCompletionDateTime

`func (o *OrdersLine) SetEstimatedCompletionDateTime(v time.Time)`

SetEstimatedCompletionDateTime sets EstimatedCompletionDateTime field to given value.

### HasEstimatedCompletionDateTime

`func (o *OrdersLine) HasEstimatedCompletionDateTime() bool`

HasEstimatedCompletionDateTime returns a boolean if a field has been set.

### GetUnitPricing

`func (o *OrdersLine) GetUnitPricing() []Price`

GetUnitPricing returns the UnitPricing field if non-nil, zero value otherwise.

### GetUnitPricingOk

`func (o *OrdersLine) GetUnitPricingOk() (*[]Price, bool)`

GetUnitPricingOk returns a tuple with the UnitPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPricing

`func (o *OrdersLine) SetUnitPricing(v []Price)`

SetUnitPricing sets UnitPricing field to given value.

### HasUnitPricing

`func (o *OrdersLine) HasUnitPricing() bool`

HasUnitPricing returns a boolean if a field has been set.

### GetTotalPricing

`func (o *OrdersLine) GetTotalPricing() []Price`

GetTotalPricing returns the TotalPricing field if non-nil, zero value otherwise.

### GetTotalPricingOk

`func (o *OrdersLine) GetTotalPricingOk() (*[]Price, bool)`

GetTotalPricingOk returns a tuple with the TotalPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPricing

`func (o *OrdersLine) SetTotalPricing(v []Price)`

SetTotalPricing sets TotalPricing field to given value.

### HasTotalPricing

`func (o *OrdersLine) HasTotalPricing() bool`

HasTotalPricing returns a boolean if a field has been set.

### GetProductType

`func (o *OrdersLine) GetProductType() OrdersLineProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *OrdersLine) GetProductTypeOk() (*OrdersLineProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *OrdersLine) SetProductType(v OrdersLineProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *OrdersLine) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetProductCode

`func (o *OrdersLine) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *OrdersLine) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *OrdersLine) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *OrdersLine) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetProductName

`func (o *OrdersLine) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *OrdersLine) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *OrdersLine) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *OrdersLine) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetIbx

`func (o *OrdersLine) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *OrdersLine) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *OrdersLine) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *OrdersLine) HasIbx() bool`

HasIbx returns a boolean if a field has been set.

### GetCage

`func (o *OrdersLine) GetCage() string`

GetCage returns the Cage field if non-nil, zero value otherwise.

### GetCageOk

`func (o *OrdersLine) GetCageOk() (*string, bool)`

GetCageOk returns a tuple with the Cage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCage

`func (o *OrdersLine) SetCage(v string)`

SetCage sets Cage field to given value.

### HasCage

`func (o *OrdersLine) HasCage() bool`

HasCage returns a boolean if a field has been set.

### GetCabinets

`func (o *OrdersLine) GetCabinets() []string`

GetCabinets returns the Cabinets field if non-nil, zero value otherwise.

### GetCabinetsOk

`func (o *OrdersLine) GetCabinetsOk() (*[]string, bool)`

GetCabinetsOk returns a tuple with the Cabinets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinets

`func (o *OrdersLine) SetCabinets(v []string)`

SetCabinets sets Cabinets field to given value.

### HasCabinets

`func (o *OrdersLine) HasCabinets() bool`

HasCabinets returns a boolean if a field has been set.

### GetRequestType

`func (o *OrdersLine) GetRequestType() OrdersLineRequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *OrdersLine) GetRequestTypeOk() (*OrdersLineRequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *OrdersLine) SetRequestType(v OrdersLineRequestType)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *OrdersLine) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetAttachments

`func (o *OrdersLine) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *OrdersLine) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *OrdersLine) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *OrdersLine) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *OrdersLine) GetAdditionalInfo() []ProductAdditionalInfoInner`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *OrdersLine) GetAdditionalInfoOk() (*[]ProductAdditionalInfoInner, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *OrdersLine) SetAdditionalInfo(v []ProductAdditionalInfoInner)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *OrdersLine) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetPurchaseOrder

`func (o *OrdersLine) GetPurchaseOrder() PurchaseOrder`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *OrdersLine) GetPurchaseOrderOk() (*PurchaseOrder, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *OrdersLine) SetPurchaseOrder(v PurchaseOrder)`

SetPurchaseOrder sets PurchaseOrder field to given value.

### HasPurchaseOrder

`func (o *OrdersLine) HasPurchaseOrder() bool`

HasPurchaseOrder returns a boolean if a field has been set.

### GetCancellable

`func (o *OrdersLine) GetCancellable() bool`

GetCancellable returns the Cancellable field if non-nil, zero value otherwise.

### GetCancellableOk

`func (o *OrdersLine) GetCancellableOk() (*bool, bool)`

GetCancellableOk returns a tuple with the Cancellable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellable

`func (o *OrdersLine) SetCancellable(v bool)`

SetCancellable sets Cancellable field to given value.

### HasCancellable

`func (o *OrdersLine) HasCancellable() bool`

HasCancellable returns a boolean if a field has been set.

### GetModifiable

`func (o *OrdersLine) GetModifiable() bool`

GetModifiable returns the Modifiable field if non-nil, zero value otherwise.

### GetModifiableOk

`func (o *OrdersLine) GetModifiableOk() (*bool, bool)`

GetModifiableOk returns a tuple with the Modifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiable

`func (o *OrdersLine) SetModifiable(v bool)`

SetModifiable sets Modifiable field to given value.

### HasModifiable

`func (o *OrdersLine) HasModifiable() bool`

HasModifiable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


