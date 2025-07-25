# OrderCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | Identifier of a billing account, with permission to order colocation products. | 
**CustomerReference** | Pointer to **string** | Supplementary identifier for your discretionary use. For example, your internal identifier. | [optional] 
**EndCustomerName** | Pointer to **string** | End customer name. Applicable and required if a reseller billing account is used. | [optional] 
**IbxCode** | **string** | IBX data center identifier. | 
**ContractTerm** | [**ContractTerm**](ContractTerm.md) |  | 
**OrderItem** | [**OrderItem**](OrderItem.md) |  | 
**PurchaseOrderNumber** | Pointer to **string** | Purchase order number. Optional, unless the billing account requires it. | [optional] 
**TechnicalContact** | Pointer to [**ContactDetails**](ContactDetails.md) |  | [optional] 

## Methods

### NewOrderCreateRequest

`func NewOrderCreateRequest(accountNumber string, ibxCode string, contractTerm ContractTerm, orderItem OrderItem, ) *OrderCreateRequest`

NewOrderCreateRequest instantiates a new OrderCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateRequestWithDefaults

`func NewOrderCreateRequestWithDefaults() *OrderCreateRequest`

NewOrderCreateRequestWithDefaults instantiates a new OrderCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *OrderCreateRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *OrderCreateRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *OrderCreateRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetCustomerReference

`func (o *OrderCreateRequest) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *OrderCreateRequest) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *OrderCreateRequest) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *OrderCreateRequest) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetEndCustomerName

`func (o *OrderCreateRequest) GetEndCustomerName() string`

GetEndCustomerName returns the EndCustomerName field if non-nil, zero value otherwise.

### GetEndCustomerNameOk

`func (o *OrderCreateRequest) GetEndCustomerNameOk() (*string, bool)`

GetEndCustomerNameOk returns a tuple with the EndCustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomerName

`func (o *OrderCreateRequest) SetEndCustomerName(v string)`

SetEndCustomerName sets EndCustomerName field to given value.

### HasEndCustomerName

`func (o *OrderCreateRequest) HasEndCustomerName() bool`

HasEndCustomerName returns a boolean if a field has been set.

### GetIbxCode

`func (o *OrderCreateRequest) GetIbxCode() string`

GetIbxCode returns the IbxCode field if non-nil, zero value otherwise.

### GetIbxCodeOk

`func (o *OrderCreateRequest) GetIbxCodeOk() (*string, bool)`

GetIbxCodeOk returns a tuple with the IbxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbxCode

`func (o *OrderCreateRequest) SetIbxCode(v string)`

SetIbxCode sets IbxCode field to given value.


### GetContractTerm

`func (o *OrderCreateRequest) GetContractTerm() ContractTerm`

GetContractTerm returns the ContractTerm field if non-nil, zero value otherwise.

### GetContractTermOk

`func (o *OrderCreateRequest) GetContractTermOk() (*ContractTerm, bool)`

GetContractTermOk returns a tuple with the ContractTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractTerm

`func (o *OrderCreateRequest) SetContractTerm(v ContractTerm)`

SetContractTerm sets ContractTerm field to given value.


### GetOrderItem

`func (o *OrderCreateRequest) GetOrderItem() OrderItem`

GetOrderItem returns the OrderItem field if non-nil, zero value otherwise.

### GetOrderItemOk

`func (o *OrderCreateRequest) GetOrderItemOk() (*OrderItem, bool)`

GetOrderItemOk returns a tuple with the OrderItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItem

`func (o *OrderCreateRequest) SetOrderItem(v OrderItem)`

SetOrderItem sets OrderItem field to given value.


### GetPurchaseOrderNumber

`func (o *OrderCreateRequest) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *OrderCreateRequest) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *OrderCreateRequest) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *OrderCreateRequest) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetTechnicalContact

`func (o *OrderCreateRequest) GetTechnicalContact() ContactDetails`

GetTechnicalContact returns the TechnicalContact field if non-nil, zero value otherwise.

### GetTechnicalContactOk

`func (o *OrderCreateRequest) GetTechnicalContactOk() (*ContactDetails, bool)`

GetTechnicalContactOk returns a tuple with the TechnicalContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContact

`func (o *OrderCreateRequest) SetTechnicalContact(v ContactDetails)`

SetTechnicalContact sets TechnicalContact field to given value.

### HasTechnicalContact

`func (o *OrderCreateRequest) HasTechnicalContact() bool`

HasTechnicalContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


