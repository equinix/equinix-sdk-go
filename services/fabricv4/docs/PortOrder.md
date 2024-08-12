# PortOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurchaseOrder** | Pointer to [**PortOrderPurchaseOrder**](PortOrderPurchaseOrder.md) |  | [optional] 
**OrderId** | Pointer to **string** | Order Identification | [optional] 
**CustomerReferenceId** | Pointer to **string** | Customer order reference Id | [optional] 
**OrderNumber** | Pointer to **string** | Order Reference Number | [optional] 
**Uuid** | Pointer to **string** | Equinix-assigned order identifier | [optional] 
**Signature** | Pointer to [**PortOrderSignature**](PortOrderSignature.md) |  | [optional] 

## Methods

### NewPortOrder

`func NewPortOrder() *PortOrder`

NewPortOrder instantiates a new PortOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortOrderWithDefaults

`func NewPortOrderWithDefaults() *PortOrder`

NewPortOrderWithDefaults instantiates a new PortOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurchaseOrder

`func (o *PortOrder) GetPurchaseOrder() PortOrderPurchaseOrder`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *PortOrder) GetPurchaseOrderOk() (*PortOrderPurchaseOrder, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *PortOrder) SetPurchaseOrder(v PortOrderPurchaseOrder)`

SetPurchaseOrder sets PurchaseOrder field to given value.

### HasPurchaseOrder

`func (o *PortOrder) HasPurchaseOrder() bool`

HasPurchaseOrder returns a boolean if a field has been set.

### GetOrderId

`func (o *PortOrder) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PortOrder) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PortOrder) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PortOrder) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetCustomerReferenceId

`func (o *PortOrder) GetCustomerReferenceId() string`

GetCustomerReferenceId returns the CustomerReferenceId field if non-nil, zero value otherwise.

### GetCustomerReferenceIdOk

`func (o *PortOrder) GetCustomerReferenceIdOk() (*string, bool)`

GetCustomerReferenceIdOk returns a tuple with the CustomerReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReferenceId

`func (o *PortOrder) SetCustomerReferenceId(v string)`

SetCustomerReferenceId sets CustomerReferenceId field to given value.

### HasCustomerReferenceId

`func (o *PortOrder) HasCustomerReferenceId() bool`

HasCustomerReferenceId returns a boolean if a field has been set.

### GetOrderNumber

`func (o *PortOrder) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *PortOrder) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *PortOrder) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *PortOrder) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetUuid

`func (o *PortOrder) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PortOrder) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PortOrder) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PortOrder) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetSignature

`func (o *PortOrder) GetSignature() PortOrderSignature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PortOrder) GetSignatureOk() (*PortOrderSignature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PortOrder) SetSignature(v PortOrderSignature)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *PortOrder) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


