# OrdersBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **string** | Unique identifier of the order. | [optional] 
**AccountName** | Pointer to **string** | Customer account name | [optional] 
**AccountNumber** | Pointer to **string** | Customer account number | [optional] 
**QuoteRequestType** | Pointer to [**OrdersBaseQuoteRequestType**](OrdersBaseQuoteRequestType.md) |  | [optional] 
**Contacts** | Pointer to [**[]Contacts**](Contacts.md) | The related party associated with the ticket. | [optional] 
**Status** | Pointer to [**OrdersBaseStatus**](OrdersBaseStatus.md) |  | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | The order created date and time in UTC timezone. | [optional] 
**UpdatedDateTime** | Pointer to **time.Time** | The order updated date and time in UTC timezone. | [optional] 
**ClosedDateTime** | Pointer to **time.Time** | The order closed date and time in UTC timezone. | [optional] 
**EstimatedCompletionDateTime** | Pointer to **time.Time** | The estimated completion date and time in UTC timezone | [optional] 
**CurrencyCode** | Pointer to [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**Channel** | Pointer to [**OrdersBaseChannel**](OrdersBaseChannel.md) |  | [optional] 
**SubChannel** | Pointer to [**OrdersBaseSubChannel**](OrdersBaseSubChannel.md) |  | [optional] 
**Notes** | Pointer to [**[]Note**](Note.md) | The notes that are associated to the orders. | [optional] 
**AdditionalInfo** | Pointer to [**[]ProductAdditionalInfoInner**](ProductAdditionalInfoInner.md) | This section is reserved to display product specific information | [optional] 
**CustomerReferenceId** | Pointer to **string** | Customer Reference Number / External Reference Number | [optional] 
**Cancellable** | Pointer to **bool** | When &#x60;true&#x60;, order can be cancelled. | [optional] [default to false]
**Modifiable** | Pointer to **bool** | When &#x60;true&#x60;, order can be modified. | [optional] [default to false]

## Methods

### NewOrdersBase

`func NewOrdersBase() *OrdersBase`

NewOrdersBase instantiates a new OrdersBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersBaseWithDefaults

`func NewOrdersBaseWithDefaults() *OrdersBase`

NewOrdersBaseWithDefaults instantiates a new OrdersBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *OrdersBase) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrdersBase) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrdersBase) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrdersBase) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetAccountName

`func (o *OrdersBase) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *OrdersBase) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *OrdersBase) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *OrdersBase) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountNumber

`func (o *OrdersBase) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *OrdersBase) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *OrdersBase) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *OrdersBase) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetQuoteRequestType

`func (o *OrdersBase) GetQuoteRequestType() OrdersBaseQuoteRequestType`

GetQuoteRequestType returns the QuoteRequestType field if non-nil, zero value otherwise.

### GetQuoteRequestTypeOk

`func (o *OrdersBase) GetQuoteRequestTypeOk() (*OrdersBaseQuoteRequestType, bool)`

GetQuoteRequestTypeOk returns a tuple with the QuoteRequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteRequestType

`func (o *OrdersBase) SetQuoteRequestType(v OrdersBaseQuoteRequestType)`

SetQuoteRequestType sets QuoteRequestType field to given value.

### HasQuoteRequestType

`func (o *OrdersBase) HasQuoteRequestType() bool`

HasQuoteRequestType returns a boolean if a field has been set.

### GetContacts

`func (o *OrdersBase) GetContacts() []Contacts`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *OrdersBase) GetContactsOk() (*[]Contacts, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *OrdersBase) SetContacts(v []Contacts)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *OrdersBase) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetStatus

`func (o *OrdersBase) GetStatus() OrdersBaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrdersBase) GetStatusOk() (*OrdersBaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrdersBase) SetStatus(v OrdersBaseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrdersBase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *OrdersBase) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *OrdersBase) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *OrdersBase) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *OrdersBase) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *OrdersBase) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *OrdersBase) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *OrdersBase) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.

### HasUpdatedDateTime

`func (o *OrdersBase) HasUpdatedDateTime() bool`

HasUpdatedDateTime returns a boolean if a field has been set.

### GetClosedDateTime

`func (o *OrdersBase) GetClosedDateTime() time.Time`

GetClosedDateTime returns the ClosedDateTime field if non-nil, zero value otherwise.

### GetClosedDateTimeOk

`func (o *OrdersBase) GetClosedDateTimeOk() (*time.Time, bool)`

GetClosedDateTimeOk returns a tuple with the ClosedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedDateTime

`func (o *OrdersBase) SetClosedDateTime(v time.Time)`

SetClosedDateTime sets ClosedDateTime field to given value.

### HasClosedDateTime

`func (o *OrdersBase) HasClosedDateTime() bool`

HasClosedDateTime returns a boolean if a field has been set.

### GetEstimatedCompletionDateTime

`func (o *OrdersBase) GetEstimatedCompletionDateTime() time.Time`

GetEstimatedCompletionDateTime returns the EstimatedCompletionDateTime field if non-nil, zero value otherwise.

### GetEstimatedCompletionDateTimeOk

`func (o *OrdersBase) GetEstimatedCompletionDateTimeOk() (*time.Time, bool)`

GetEstimatedCompletionDateTimeOk returns a tuple with the EstimatedCompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCompletionDateTime

`func (o *OrdersBase) SetEstimatedCompletionDateTime(v time.Time)`

SetEstimatedCompletionDateTime sets EstimatedCompletionDateTime field to given value.

### HasEstimatedCompletionDateTime

`func (o *OrdersBase) HasEstimatedCompletionDateTime() bool`

HasEstimatedCompletionDateTime returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *OrdersBase) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *OrdersBase) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *OrdersBase) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *OrdersBase) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetChannel

`func (o *OrdersBase) GetChannel() OrdersBaseChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *OrdersBase) GetChannelOk() (*OrdersBaseChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *OrdersBase) SetChannel(v OrdersBaseChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *OrdersBase) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetSubChannel

`func (o *OrdersBase) GetSubChannel() OrdersBaseSubChannel`

GetSubChannel returns the SubChannel field if non-nil, zero value otherwise.

### GetSubChannelOk

`func (o *OrdersBase) GetSubChannelOk() (*OrdersBaseSubChannel, bool)`

GetSubChannelOk returns a tuple with the SubChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubChannel

`func (o *OrdersBase) SetSubChannel(v OrdersBaseSubChannel)`

SetSubChannel sets SubChannel field to given value.

### HasSubChannel

`func (o *OrdersBase) HasSubChannel() bool`

HasSubChannel returns a boolean if a field has been set.

### GetNotes

`func (o *OrdersBase) GetNotes() []Note`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrdersBase) GetNotesOk() (*[]Note, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrdersBase) SetNotes(v []Note)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrdersBase) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *OrdersBase) GetAdditionalInfo() []ProductAdditionalInfoInner`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *OrdersBase) GetAdditionalInfoOk() (*[]ProductAdditionalInfoInner, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *OrdersBase) SetAdditionalInfo(v []ProductAdditionalInfoInner)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *OrdersBase) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetCustomerReferenceId

`func (o *OrdersBase) GetCustomerReferenceId() string`

GetCustomerReferenceId returns the CustomerReferenceId field if non-nil, zero value otherwise.

### GetCustomerReferenceIdOk

`func (o *OrdersBase) GetCustomerReferenceIdOk() (*string, bool)`

GetCustomerReferenceIdOk returns a tuple with the CustomerReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReferenceId

`func (o *OrdersBase) SetCustomerReferenceId(v string)`

SetCustomerReferenceId sets CustomerReferenceId field to given value.

### HasCustomerReferenceId

`func (o *OrdersBase) HasCustomerReferenceId() bool`

HasCustomerReferenceId returns a boolean if a field has been set.

### GetCancellable

`func (o *OrdersBase) GetCancellable() bool`

GetCancellable returns the Cancellable field if non-nil, zero value otherwise.

### GetCancellableOk

`func (o *OrdersBase) GetCancellableOk() (*bool, bool)`

GetCancellableOk returns a tuple with the Cancellable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellable

`func (o *OrdersBase) SetCancellable(v bool)`

SetCancellable sets Cancellable field to given value.

### HasCancellable

`func (o *OrdersBase) HasCancellable() bool`

HasCancellable returns a boolean if a field has been set.

### GetModifiable

`func (o *OrdersBase) GetModifiable() bool`

GetModifiable returns the Modifiable field if non-nil, zero value otherwise.

### GetModifiableOk

`func (o *OrdersBase) GetModifiableOk() (*bool, bool)`

GetModifiableOk returns a tuple with the Modifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiable

`func (o *OrdersBase) SetModifiable(v bool)`

SetModifiable sets Modifiable field to given value.

### HasModifiable

`func (o *OrdersBase) HasModifiable() bool`

HasModifiable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


