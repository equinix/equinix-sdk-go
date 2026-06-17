# InternetAccessBilling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**InternetAccessBillingType**](InternetAccessBillingType.md) |  | 
**Enabled** | Pointer to **bool** | Indicates whether the billing is enabled | [optional] 
**StartDate** | Pointer to **time.Time** | The start date for the billing period | [optional] 

## Methods

### NewInternetAccessBilling

`func NewInternetAccessBilling(type_ InternetAccessBillingType, ) *InternetAccessBilling`

NewInternetAccessBilling instantiates a new InternetAccessBilling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessBillingWithDefaults

`func NewInternetAccessBillingWithDefaults() *InternetAccessBilling`

NewInternetAccessBillingWithDefaults instantiates a new InternetAccessBilling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InternetAccessBilling) GetType() InternetAccessBillingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternetAccessBilling) GetTypeOk() (*InternetAccessBillingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternetAccessBilling) SetType(v InternetAccessBillingType)`

SetType sets Type field to given value.


### GetEnabled

`func (o *InternetAccessBilling) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InternetAccessBilling) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InternetAccessBilling) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InternetAccessBilling) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStartDate

`func (o *InternetAccessBilling) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InternetAccessBilling) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InternetAccessBilling) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InternetAccessBilling) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


