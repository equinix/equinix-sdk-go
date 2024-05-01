# ServiceOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Contacts** | Pointer to [**[]ServiceOrderContact**](ServiceOrderContact.md) |  | [optional] 
**Draft** | Pointer to **bool** |  | [optional] [default to false]
**PurchaseOrder** | Pointer to [**ServicePurchaseOrder**](ServicePurchaseOrder.md) |  | [optional] 
**ReferenceNumber** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to [**OrderSignatureRequest**](OrderSignatureRequest.md) |  | [optional] 

## Methods

### NewServiceOrderRequest

`func NewServiceOrderRequest() *ServiceOrderRequest`

NewServiceOrderRequest instantiates a new ServiceOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOrderRequestWithDefaults

`func NewServiceOrderRequestWithDefaults() *ServiceOrderRequest`

NewServiceOrderRequestWithDefaults instantiates a new ServiceOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *ServiceOrderRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceOrderRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceOrderRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceOrderRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetContacts

`func (o *ServiceOrderRequest) GetContacts() []ServiceOrderContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ServiceOrderRequest) GetContactsOk() (*[]ServiceOrderContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ServiceOrderRequest) SetContacts(v []ServiceOrderContact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *ServiceOrderRequest) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetDraft

`func (o *ServiceOrderRequest) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *ServiceOrderRequest) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *ServiceOrderRequest) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *ServiceOrderRequest) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetPurchaseOrder

`func (o *ServiceOrderRequest) GetPurchaseOrder() ServicePurchaseOrder`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *ServiceOrderRequest) GetPurchaseOrderOk() (*ServicePurchaseOrder, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *ServiceOrderRequest) SetPurchaseOrder(v ServicePurchaseOrder)`

SetPurchaseOrder sets PurchaseOrder field to given value.

### HasPurchaseOrder

`func (o *ServiceOrderRequest) HasPurchaseOrder() bool`

HasPurchaseOrder returns a boolean if a field has been set.

### GetReferenceNumber

`func (o *ServiceOrderRequest) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *ServiceOrderRequest) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *ServiceOrderRequest) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *ServiceOrderRequest) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### GetSignature

`func (o *ServiceOrderRequest) GetSignature() OrderSignatureRequest`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ServiceOrderRequest) GetSignatureOk() (*OrderSignatureRequest, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ServiceOrderRequest) SetSignature(v OrderSignatureRequest)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ServiceOrderRequest) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


