# InternetAccessOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurchaseOrderNumber** | Pointer to **string** | Purchase order number | [optional] 
**CustomerReferenceNumber** | Pointer to **string** | Customer reference number | [optional] 
**BillingTier** | Pointer to **string** | Billing tier for connection bandwidth | [optional] 
**OrderId** | Pointer to **string** | Order Identification | [optional] 
**OrderNumber** | Pointer to **string** | Order Reference Number | [optional] 
**TermLength** | Pointer to **int32** | Term length in months, valid values are 1, 12, 24, 36 where 1 is the default value (for on-demand case). | [optional] [default to 1]
**ContractedBandwidth** | Pointer to **int32** | Contracted bandwidth | [optional] 
**Href** | Pointer to **string** | Order URI | [optional] 

## Methods

### NewInternetAccessOrder

`func NewInternetAccessOrder() *InternetAccessOrder`

NewInternetAccessOrder instantiates a new InternetAccessOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessOrderWithDefaults

`func NewInternetAccessOrderWithDefaults() *InternetAccessOrder`

NewInternetAccessOrderWithDefaults instantiates a new InternetAccessOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurchaseOrderNumber

`func (o *InternetAccessOrder) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *InternetAccessOrder) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *InternetAccessOrder) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *InternetAccessOrder) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetCustomerReferenceNumber

`func (o *InternetAccessOrder) GetCustomerReferenceNumber() string`

GetCustomerReferenceNumber returns the CustomerReferenceNumber field if non-nil, zero value otherwise.

### GetCustomerReferenceNumberOk

`func (o *InternetAccessOrder) GetCustomerReferenceNumberOk() (*string, bool)`

GetCustomerReferenceNumberOk returns a tuple with the CustomerReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReferenceNumber

`func (o *InternetAccessOrder) SetCustomerReferenceNumber(v string)`

SetCustomerReferenceNumber sets CustomerReferenceNumber field to given value.

### HasCustomerReferenceNumber

`func (o *InternetAccessOrder) HasCustomerReferenceNumber() bool`

HasCustomerReferenceNumber returns a boolean if a field has been set.

### GetBillingTier

`func (o *InternetAccessOrder) GetBillingTier() string`

GetBillingTier returns the BillingTier field if non-nil, zero value otherwise.

### GetBillingTierOk

`func (o *InternetAccessOrder) GetBillingTierOk() (*string, bool)`

GetBillingTierOk returns a tuple with the BillingTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTier

`func (o *InternetAccessOrder) SetBillingTier(v string)`

SetBillingTier sets BillingTier field to given value.

### HasBillingTier

`func (o *InternetAccessOrder) HasBillingTier() bool`

HasBillingTier returns a boolean if a field has been set.

### GetOrderId

`func (o *InternetAccessOrder) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *InternetAccessOrder) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *InternetAccessOrder) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *InternetAccessOrder) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderNumber

`func (o *InternetAccessOrder) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *InternetAccessOrder) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *InternetAccessOrder) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *InternetAccessOrder) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetTermLength

`func (o *InternetAccessOrder) GetTermLength() int32`

GetTermLength returns the TermLength field if non-nil, zero value otherwise.

### GetTermLengthOk

`func (o *InternetAccessOrder) GetTermLengthOk() (*int32, bool)`

GetTermLengthOk returns a tuple with the TermLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermLength

`func (o *InternetAccessOrder) SetTermLength(v int32)`

SetTermLength sets TermLength field to given value.

### HasTermLength

`func (o *InternetAccessOrder) HasTermLength() bool`

HasTermLength returns a boolean if a field has been set.

### GetContractedBandwidth

`func (o *InternetAccessOrder) GetContractedBandwidth() int32`

GetContractedBandwidth returns the ContractedBandwidth field if non-nil, zero value otherwise.

### GetContractedBandwidthOk

`func (o *InternetAccessOrder) GetContractedBandwidthOk() (*int32, bool)`

GetContractedBandwidthOk returns a tuple with the ContractedBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractedBandwidth

`func (o *InternetAccessOrder) SetContractedBandwidth(v int32)`

SetContractedBandwidth sets ContractedBandwidth field to given value.

### HasContractedBandwidth

`func (o *InternetAccessOrder) HasContractedBandwidth() bool`

HasContractedBandwidth returns a boolean if a field has been set.

### GetHref

`func (o *InternetAccessOrder) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *InternetAccessOrder) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *InternetAccessOrder) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *InternetAccessOrder) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


