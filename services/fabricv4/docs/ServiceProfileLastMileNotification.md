# ServiceProfileLastMileNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | Pointer to **string** | Expected format of the contact information, such as email or phone number. | [optional] 
**Label** | Pointer to **string** | Type of contact information, such as ordering or technical support. | [optional] 
**Required** | Pointer to **bool** | Whether this contact information is required for provisioning and ordering. | [optional] 

## Methods

### NewServiceProfileLastMileNotification

`func NewServiceProfileLastMileNotification() *ServiceProfileLastMileNotification`

NewServiceProfileLastMileNotification instantiates a new ServiceProfileLastMileNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileLastMileNotificationWithDefaults

`func NewServiceProfileLastMileNotificationWithDefaults() *ServiceProfileLastMileNotification`

NewServiceProfileLastMileNotificationWithDefaults instantiates a new ServiceProfileLastMileNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *ServiceProfileLastMileNotification) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ServiceProfileLastMileNotification) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ServiceProfileLastMileNotification) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *ServiceProfileLastMileNotification) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetLabel

`func (o *ServiceProfileLastMileNotification) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServiceProfileLastMileNotification) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServiceProfileLastMileNotification) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServiceProfileLastMileNotification) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRequired

`func (o *ServiceProfileLastMileNotification) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ServiceProfileLastMileNotification) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ServiceProfileLastMileNotification) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ServiceProfileLastMileNotification) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


