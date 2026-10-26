# ExchangeServiceOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurchaseOrderNumber** | Pointer to **string** | Purchase order number reference. | [optional] 
**OrderId** | Pointer to **string** | Equinix-assigned order identifier. | [optional] [readonly] 
**CustomerReferenceId** | Pointer to **string** | Your own reference for this Exchange Service. | [optional] 
**OrderNumber** | Pointer to **string** | Equinix order reference number. | [optional] 
**OrderLineItemNumber** | Pointer to **string** | Order Reference Number for Sales Assisted Order | [optional] [readonly] 
**DomainOrderNumber** | Pointer to **string** | product domain Order Number for Sales Assisted Order | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned order identifier, this is a derived response attribute | [optional] [readonly] 

## Methods

### NewExchangeServiceOrder

`func NewExchangeServiceOrder() *ExchangeServiceOrder`

NewExchangeServiceOrder instantiates a new ExchangeServiceOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeServiceOrderWithDefaults

`func NewExchangeServiceOrderWithDefaults() *ExchangeServiceOrder`

NewExchangeServiceOrderWithDefaults instantiates a new ExchangeServiceOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurchaseOrderNumber

`func (o *ExchangeServiceOrder) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *ExchangeServiceOrder) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *ExchangeServiceOrder) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *ExchangeServiceOrder) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetOrderId

`func (o *ExchangeServiceOrder) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ExchangeServiceOrder) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ExchangeServiceOrder) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ExchangeServiceOrder) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetCustomerReferenceId

`func (o *ExchangeServiceOrder) GetCustomerReferenceId() string`

GetCustomerReferenceId returns the CustomerReferenceId field if non-nil, zero value otherwise.

### GetCustomerReferenceIdOk

`func (o *ExchangeServiceOrder) GetCustomerReferenceIdOk() (*string, bool)`

GetCustomerReferenceIdOk returns a tuple with the CustomerReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReferenceId

`func (o *ExchangeServiceOrder) SetCustomerReferenceId(v string)`

SetCustomerReferenceId sets CustomerReferenceId field to given value.

### HasCustomerReferenceId

`func (o *ExchangeServiceOrder) HasCustomerReferenceId() bool`

HasCustomerReferenceId returns a boolean if a field has been set.

### GetOrderNumber

`func (o *ExchangeServiceOrder) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *ExchangeServiceOrder) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *ExchangeServiceOrder) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *ExchangeServiceOrder) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetOrderLineItemNumber

`func (o *ExchangeServiceOrder) GetOrderLineItemNumber() string`

GetOrderLineItemNumber returns the OrderLineItemNumber field if non-nil, zero value otherwise.

### GetOrderLineItemNumberOk

`func (o *ExchangeServiceOrder) GetOrderLineItemNumberOk() (*string, bool)`

GetOrderLineItemNumberOk returns a tuple with the OrderLineItemNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLineItemNumber

`func (o *ExchangeServiceOrder) SetOrderLineItemNumber(v string)`

SetOrderLineItemNumber sets OrderLineItemNumber field to given value.

### HasOrderLineItemNumber

`func (o *ExchangeServiceOrder) HasOrderLineItemNumber() bool`

HasOrderLineItemNumber returns a boolean if a field has been set.

### GetDomainOrderNumber

`func (o *ExchangeServiceOrder) GetDomainOrderNumber() string`

GetDomainOrderNumber returns the DomainOrderNumber field if non-nil, zero value otherwise.

### GetDomainOrderNumberOk

`func (o *ExchangeServiceOrder) GetDomainOrderNumberOk() (*string, bool)`

GetDomainOrderNumberOk returns a tuple with the DomainOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainOrderNumber

`func (o *ExchangeServiceOrder) SetDomainOrderNumber(v string)`

SetDomainOrderNumber sets DomainOrderNumber field to given value.

### HasDomainOrderNumber

`func (o *ExchangeServiceOrder) HasDomainOrderNumber() bool`

HasDomainOrderNumber returns a boolean if a field has been set.

### GetUuid

`func (o *ExchangeServiceOrder) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ExchangeServiceOrder) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ExchangeServiceOrder) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ExchangeServiceOrder) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


