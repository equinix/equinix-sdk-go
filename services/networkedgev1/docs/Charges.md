# Charges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the charge, whether it is for the virtual device, the device license, or the additional bandwidth. | [optional] 
**MonthlyRecurringCharges** | Pointer to **string** | The monthly charges. | [optional] 

## Methods

### NewCharges

`func NewCharges() *Charges`

NewCharges instantiates a new Charges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargesWithDefaults

`func NewChargesWithDefaults() *Charges`

NewChargesWithDefaults instantiates a new Charges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Charges) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Charges) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Charges) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Charges) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMonthlyRecurringCharges

`func (o *Charges) GetMonthlyRecurringCharges() string`

GetMonthlyRecurringCharges returns the MonthlyRecurringCharges field if non-nil, zero value otherwise.

### GetMonthlyRecurringChargesOk

`func (o *Charges) GetMonthlyRecurringChargesOk() (*string, bool)`

GetMonthlyRecurringChargesOk returns a tuple with the MonthlyRecurringCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyRecurringCharges

`func (o *Charges) SetMonthlyRecurringCharges(v string)`

SetMonthlyRecurringCharges sets MonthlyRecurringCharges field to given value.

### HasMonthlyRecurringCharges

`func (o *Charges) HasMonthlyRecurringCharges() bool`

HasMonthlyRecurringCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


