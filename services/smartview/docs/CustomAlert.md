# CustomAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** |  | 
**Ibx** | **[]string** |  | 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomAlert

`func NewCustomAlert(accountNumber string, ibx []string, ) *CustomAlert`

NewCustomAlert instantiates a new CustomAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAlertWithDefaults

`func NewCustomAlertWithDefaults() *CustomAlert`

NewCustomAlertWithDefaults instantiates a new CustomAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *CustomAlert) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CustomAlert) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CustomAlert) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetIbx

`func (o *CustomAlert) GetIbx() []string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *CustomAlert) GetIbxOk() (*[]string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *CustomAlert) SetIbx(v []string)`

SetIbx sets Ibx field to given value.


### GetUser

`func (o *CustomAlert) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CustomAlert) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CustomAlert) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CustomAlert) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


