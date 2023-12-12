# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress** | Pointer to [**PaymentMethodBillingAddress**](PaymentMethodBillingAddress.md) |  | [optional] 
**CardType** | Pointer to **string** |  | [optional] 
**CardholderName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedByUser** | Pointer to [**Href**](Href.md) |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**ExpirationMonth** | Pointer to **string** |  | [optional] 
**ExpirationYear** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**Href**](Href.md) |  | [optional] 
**Projects** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPaymentMethod

`func NewPaymentMethod() *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAddress

`func (o *PaymentMethod) GetBillingAddress() PaymentMethodBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *PaymentMethod) GetBillingAddressOk() (*PaymentMethodBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *PaymentMethod) SetBillingAddress(v PaymentMethodBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *PaymentMethod) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCardType

`func (o *PaymentMethod) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *PaymentMethod) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *PaymentMethod) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *PaymentMethod) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetCardholderName

`func (o *PaymentMethod) GetCardholderName() string`

GetCardholderName returns the CardholderName field if non-nil, zero value otherwise.

### GetCardholderNameOk

`func (o *PaymentMethod) GetCardholderNameOk() (*string, bool)`

GetCardholderNameOk returns a tuple with the CardholderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderName

`func (o *PaymentMethod) SetCardholderName(v string)`

SetCardholderName sets CardholderName field to given value.

### HasCardholderName

`func (o *PaymentMethod) HasCardholderName() bool`

HasCardholderName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaymentMethod) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentMethod) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentMethod) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaymentMethod) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *PaymentMethod) GetCreatedByUser() Href`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *PaymentMethod) GetCreatedByUserOk() (*Href, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *PaymentMethod) SetCreatedByUser(v Href)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *PaymentMethod) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetDefault

`func (o *PaymentMethod) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PaymentMethod) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PaymentMethod) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PaymentMethod) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEmail

`func (o *PaymentMethod) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PaymentMethod) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PaymentMethod) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PaymentMethod) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExpirationMonth

`func (o *PaymentMethod) GetExpirationMonth() string`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *PaymentMethod) GetExpirationMonthOk() (*string, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *PaymentMethod) SetExpirationMonth(v string)`

SetExpirationMonth sets ExpirationMonth field to given value.

### HasExpirationMonth

`func (o *PaymentMethod) HasExpirationMonth() bool`

HasExpirationMonth returns a boolean if a field has been set.

### GetExpirationYear

`func (o *PaymentMethod) GetExpirationYear() string`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *PaymentMethod) GetExpirationYearOk() (*string, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *PaymentMethod) SetExpirationYear(v string)`

SetExpirationYear sets ExpirationYear field to given value.

### HasExpirationYear

`func (o *PaymentMethod) HasExpirationYear() bool`

HasExpirationYear returns a boolean if a field has been set.

### GetId

`func (o *PaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethod) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PaymentMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentMethod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentMethod) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *PaymentMethod) GetOrganization() Href`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PaymentMethod) GetOrganizationOk() (*Href, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PaymentMethod) SetOrganization(v Href)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PaymentMethod) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProjects

`func (o *PaymentMethod) GetProjects() []Href`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *PaymentMethod) GetProjectsOk() (*[]Href, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *PaymentMethod) SetProjects(v []Href)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *PaymentMethod) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetType

`func (o *PaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PaymentMethod) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentMethod) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentMethod) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PaymentMethod) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


