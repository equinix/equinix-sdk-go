# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**Uuid** | **string** |  | 
**Type** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**Contacts** | Pointer to [**[]Contact**](Contact.md) |  | [optional] 
**Draft** | Pointer to **bool** |  | [optional] [default to false]
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 
**PurchaseOrder** | Pointer to [**OrderPurchaseOrder**](OrderPurchaseOrder.md) |  | [optional] 
**ReferenceNumber** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to [**OrderSignature**](OrderSignature.md) |  | [optional] 
**State** | Pointer to [**OrderState**](OrderState.md) |  | [optional] 
**ChangeLog** | Pointer to [**OrderChangeLog**](OrderChangeLog.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOrder

`func NewOrder(href string, uuid string, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Order) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Order) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Order) SetHref(v string)`

SetHref sets Href field to given value.


### GetUuid

`func (o *Order) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Order) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Order) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *Order) GetType() OrderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Order) GetTypeOk() (*OrderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Order) SetType(v OrderType)`

SetType sets Type field to given value.

### HasType

`func (o *Order) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContacts

`func (o *Order) GetContacts() []Contact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *Order) GetContactsOk() (*[]Contact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *Order) SetContacts(v []Contact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *Order) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetDraft

`func (o *Order) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *Order) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *Order) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *Order) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetLinks

`func (o *Order) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Order) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Order) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Order) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPurchaseOrder

`func (o *Order) GetPurchaseOrder() OrderPurchaseOrder`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *Order) GetPurchaseOrderOk() (*OrderPurchaseOrder, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *Order) SetPurchaseOrder(v OrderPurchaseOrder)`

SetPurchaseOrder sets PurchaseOrder field to given value.

### HasPurchaseOrder

`func (o *Order) HasPurchaseOrder() bool`

HasPurchaseOrder returns a boolean if a field has been set.

### GetReferenceNumber

`func (o *Order) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *Order) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *Order) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *Order) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### GetSignature

`func (o *Order) GetSignature() OrderSignature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Order) GetSignatureOk() (*OrderSignature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Order) SetSignature(v OrderSignature)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *Order) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetState

`func (o *Order) GetState() OrderState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Order) GetStateOk() (*OrderState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Order) SetState(v OrderState)`

SetState sets State field to given value.

### HasState

`func (o *Order) HasState() bool`

HasState returns a boolean if a field has been set.

### GetChangeLog

`func (o *Order) GetChangeLog() OrderChangeLog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *Order) GetChangeLogOk() (*OrderChangeLog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *Order) SetChangeLog(v OrderChangeLog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *Order) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetTags

`func (o *Order) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Order) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Order) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Order) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


