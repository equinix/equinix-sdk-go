# SystemAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** |  | 
**Ibx** | **[]string** |  | 

## Methods

### NewSystemAlert

`func NewSystemAlert(accountNumber string, ibx []string, ) *SystemAlert`

NewSystemAlert instantiates a new SystemAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemAlertWithDefaults

`func NewSystemAlertWithDefaults() *SystemAlert`

NewSystemAlertWithDefaults instantiates a new SystemAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *SystemAlert) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *SystemAlert) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *SystemAlert) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetIbx

`func (o *SystemAlert) GetIbx() []string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *SystemAlert) GetIbxOk() (*[]string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *SystemAlert) SetIbx(v []string)`

SetIbx sets Ibx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


