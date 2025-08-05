# Orders

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
**Details** | Pointer to [**[]OrdersLine**](OrdersLine.md) |  | [optional] 

## Methods

### NewOrders

`func NewOrders() *Orders`

NewOrders instantiates a new Orders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrdersWithDefaults

`func NewOrdersWithDefaults() *Orders`

NewOrdersWithDefaults instantiates a new Orders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *Orders) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Orders) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Orders) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *Orders) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetAccountName

`func (o *Orders) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Orders) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Orders) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *Orders) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountNumber

`func (o *Orders) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Orders) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Orders) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *Orders) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetQuoteRequestType

`func (o *Orders) GetQuoteRequestType() OrdersBaseQuoteRequestType`

GetQuoteRequestType returns the QuoteRequestType field if non-nil, zero value otherwise.

### GetQuoteRequestTypeOk

`func (o *Orders) GetQuoteRequestTypeOk() (*OrdersBaseQuoteRequestType, bool)`

GetQuoteRequestTypeOk returns a tuple with the QuoteRequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteRequestType

`func (o *Orders) SetQuoteRequestType(v OrdersBaseQuoteRequestType)`

SetQuoteRequestType sets QuoteRequestType field to given value.

### HasQuoteRequestType

`func (o *Orders) HasQuoteRequestType() bool`

HasQuoteRequestType returns a boolean if a field has been set.

### GetContacts

`func (o *Orders) GetContacts() []Contacts`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *Orders) GetContactsOk() (*[]Contacts, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *Orders) SetContacts(v []Contacts)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *Orders) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetStatus

`func (o *Orders) GetStatus() OrdersBaseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Orders) GetStatusOk() (*OrdersBaseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Orders) SetStatus(v OrdersBaseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Orders) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *Orders) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Orders) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Orders) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Orders) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *Orders) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *Orders) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *Orders) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.

### HasUpdatedDateTime

`func (o *Orders) HasUpdatedDateTime() bool`

HasUpdatedDateTime returns a boolean if a field has been set.

### GetClosedDateTime

`func (o *Orders) GetClosedDateTime() time.Time`

GetClosedDateTime returns the ClosedDateTime field if non-nil, zero value otherwise.

### GetClosedDateTimeOk

`func (o *Orders) GetClosedDateTimeOk() (*time.Time, bool)`

GetClosedDateTimeOk returns a tuple with the ClosedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedDateTime

`func (o *Orders) SetClosedDateTime(v time.Time)`

SetClosedDateTime sets ClosedDateTime field to given value.

### HasClosedDateTime

`func (o *Orders) HasClosedDateTime() bool`

HasClosedDateTime returns a boolean if a field has been set.

### GetEstimatedCompletionDateTime

`func (o *Orders) GetEstimatedCompletionDateTime() time.Time`

GetEstimatedCompletionDateTime returns the EstimatedCompletionDateTime field if non-nil, zero value otherwise.

### GetEstimatedCompletionDateTimeOk

`func (o *Orders) GetEstimatedCompletionDateTimeOk() (*time.Time, bool)`

GetEstimatedCompletionDateTimeOk returns a tuple with the EstimatedCompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCompletionDateTime

`func (o *Orders) SetEstimatedCompletionDateTime(v time.Time)`

SetEstimatedCompletionDateTime sets EstimatedCompletionDateTime field to given value.

### HasEstimatedCompletionDateTime

`func (o *Orders) HasEstimatedCompletionDateTime() bool`

HasEstimatedCompletionDateTime returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *Orders) GetCurrencyCode() CurrencyCode`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Orders) GetCurrencyCodeOk() (*CurrencyCode, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Orders) SetCurrencyCode(v CurrencyCode)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *Orders) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetChannel

`func (o *Orders) GetChannel() OrdersBaseChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Orders) GetChannelOk() (*OrdersBaseChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Orders) SetChannel(v OrdersBaseChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Orders) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetSubChannel

`func (o *Orders) GetSubChannel() OrdersBaseSubChannel`

GetSubChannel returns the SubChannel field if non-nil, zero value otherwise.

### GetSubChannelOk

`func (o *Orders) GetSubChannelOk() (*OrdersBaseSubChannel, bool)`

GetSubChannelOk returns a tuple with the SubChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubChannel

`func (o *Orders) SetSubChannel(v OrdersBaseSubChannel)`

SetSubChannel sets SubChannel field to given value.

### HasSubChannel

`func (o *Orders) HasSubChannel() bool`

HasSubChannel returns a boolean if a field has been set.

### GetNotes

`func (o *Orders) GetNotes() []Note`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Orders) GetNotesOk() (*[]Note, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Orders) SetNotes(v []Note)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Orders) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *Orders) GetAdditionalInfo() []ProductAdditionalInfoInner`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *Orders) GetAdditionalInfoOk() (*[]ProductAdditionalInfoInner, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *Orders) SetAdditionalInfo(v []ProductAdditionalInfoInner)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *Orders) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetCustomerReferenceId

`func (o *Orders) GetCustomerReferenceId() string`

GetCustomerReferenceId returns the CustomerReferenceId field if non-nil, zero value otherwise.

### GetCustomerReferenceIdOk

`func (o *Orders) GetCustomerReferenceIdOk() (*string, bool)`

GetCustomerReferenceIdOk returns a tuple with the CustomerReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReferenceId

`func (o *Orders) SetCustomerReferenceId(v string)`

SetCustomerReferenceId sets CustomerReferenceId field to given value.

### HasCustomerReferenceId

`func (o *Orders) HasCustomerReferenceId() bool`

HasCustomerReferenceId returns a boolean if a field has been set.

### GetCancellable

`func (o *Orders) GetCancellable() bool`

GetCancellable returns the Cancellable field if non-nil, zero value otherwise.

### GetCancellableOk

`func (o *Orders) GetCancellableOk() (*bool, bool)`

GetCancellableOk returns a tuple with the Cancellable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellable

`func (o *Orders) SetCancellable(v bool)`

SetCancellable sets Cancellable field to given value.

### HasCancellable

`func (o *Orders) HasCancellable() bool`

HasCancellable returns a boolean if a field has been set.

### GetModifiable

`func (o *Orders) GetModifiable() bool`

GetModifiable returns the Modifiable field if non-nil, zero value otherwise.

### GetModifiableOk

`func (o *Orders) GetModifiableOk() (*bool, bool)`

GetModifiableOk returns a tuple with the Modifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiable

`func (o *Orders) SetModifiable(v bool)`

SetModifiable sets Modifiable field to given value.

### HasModifiable

`func (o *Orders) HasModifiable() bool`

HasModifiable returns a boolean if a field has been set.

### GetDetails

`func (o *Orders) GetDetails() []OrdersLine`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Orders) GetDetailsOk() (*[]OrdersLine, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Orders) SetDetails(v []OrdersLine)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Orders) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


