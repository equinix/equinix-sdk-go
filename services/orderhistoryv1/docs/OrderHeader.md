# OrderHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderNumber** | Pointer to **string** | Order number | [optional] 
**Type** | Pointer to **[]string** | List Of Proudcts Ordered in the Order. | [optional] 
**OrderStatus** | Pointer to [**OrderHeaderStatusEnum**](OrderHeaderStatusEnum.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | Order Created DateTime in ISO Date Format | [optional] 
**Account** | Pointer to [**Account**](Account.md) |  | [optional] 
**OrderingContact** | Pointer to [**OrderingContact**](OrderingContact.md) |  | [optional] 
**NotificationContact** | Pointer to [**NotificationContact**](NotificationContact.md) |  | [optional] 
**Ibx** | Pointer to **[]string** | List of ibxs for the order lines. | [optional] 
**CustomerReferenceNumbers** | Pointer to **[]string** | List of customer reference numbers for the order lines. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) |  | [optional] 

## Methods

### NewOrderHeader

`func NewOrderHeader() *OrderHeader`

NewOrderHeader instantiates a new OrderHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderHeaderWithDefaults

`func NewOrderHeaderWithDefaults() *OrderHeader`

NewOrderHeaderWithDefaults instantiates a new OrderHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderNumber

`func (o *OrderHeader) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *OrderHeader) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *OrderHeader) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *OrderHeader) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetType

`func (o *OrderHeader) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderHeader) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderHeader) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderHeader) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrderStatus

`func (o *OrderHeader) GetOrderStatus() OrderHeaderStatusEnum`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *OrderHeader) GetOrderStatusOk() (*OrderHeaderStatusEnum, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *OrderHeader) SetOrderStatus(v OrderHeaderStatusEnum)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *OrderHeader) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrderHeader) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderHeader) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderHeader) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrderHeader) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAccount

`func (o *OrderHeader) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OrderHeader) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OrderHeader) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *OrderHeader) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOrderingContact

`func (o *OrderHeader) GetOrderingContact() OrderingContact`

GetOrderingContact returns the OrderingContact field if non-nil, zero value otherwise.

### GetOrderingContactOk

`func (o *OrderHeader) GetOrderingContactOk() (*OrderingContact, bool)`

GetOrderingContactOk returns a tuple with the OrderingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderingContact

`func (o *OrderHeader) SetOrderingContact(v OrderingContact)`

SetOrderingContact sets OrderingContact field to given value.

### HasOrderingContact

`func (o *OrderHeader) HasOrderingContact() bool`

HasOrderingContact returns a boolean if a field has been set.

### GetNotificationContact

`func (o *OrderHeader) GetNotificationContact() NotificationContact`

GetNotificationContact returns the NotificationContact field if non-nil, zero value otherwise.

### GetNotificationContactOk

`func (o *OrderHeader) GetNotificationContactOk() (*NotificationContact, bool)`

GetNotificationContactOk returns a tuple with the NotificationContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationContact

`func (o *OrderHeader) SetNotificationContact(v NotificationContact)`

SetNotificationContact sets NotificationContact field to given value.

### HasNotificationContact

`func (o *OrderHeader) HasNotificationContact() bool`

HasNotificationContact returns a boolean if a field has been set.

### GetIbx

`func (o *OrderHeader) GetIbx() []string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *OrderHeader) GetIbxOk() (*[]string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *OrderHeader) SetIbx(v []string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *OrderHeader) HasIbx() bool`

HasIbx returns a boolean if a field has been set.

### GetCustomerReferenceNumbers

`func (o *OrderHeader) GetCustomerReferenceNumbers() []string`

GetCustomerReferenceNumbers returns the CustomerReferenceNumbers field if non-nil, zero value otherwise.

### GetCustomerReferenceNumbersOk

`func (o *OrderHeader) GetCustomerReferenceNumbersOk() (*[]string, bool)`

GetCustomerReferenceNumbersOk returns a tuple with the CustomerReferenceNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReferenceNumbers

`func (o *OrderHeader) SetCustomerReferenceNumbers(v []string)`

SetCustomerReferenceNumbers sets CustomerReferenceNumbers field to given value.

### HasCustomerReferenceNumbers

`func (o *OrderHeader) HasCustomerReferenceNumbers() bool`

HasCustomerReferenceNumbers returns a boolean if a field has been set.

### GetLinks

`func (o *OrderHeader) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrderHeader) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrderHeader) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrderHeader) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


