# CrossConnectEndCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDisclosed** | Pointer to **bool** | Indicates if the end customer is disclosed. | [optional] 
**Name** | Pointer to **string** | End customer name. | [optional] 
**MdmId** | Pointer to **string** | Master Data Management identifier of the end customer, as held by Equinix.  | [optional] 

## Methods

### NewCrossConnectEndCustomer

`func NewCrossConnectEndCustomer() *CrossConnectEndCustomer`

NewCrossConnectEndCustomer instantiates a new CrossConnectEndCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectEndCustomerWithDefaults

`func NewCrossConnectEndCustomerWithDefaults() *CrossConnectEndCustomer`

NewCrossConnectEndCustomerWithDefaults instantiates a new CrossConnectEndCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDisclosed

`func (o *CrossConnectEndCustomer) GetIsDisclosed() bool`

GetIsDisclosed returns the IsDisclosed field if non-nil, zero value otherwise.

### GetIsDisclosedOk

`func (o *CrossConnectEndCustomer) GetIsDisclosedOk() (*bool, bool)`

GetIsDisclosedOk returns a tuple with the IsDisclosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisclosed

`func (o *CrossConnectEndCustomer) SetIsDisclosed(v bool)`

SetIsDisclosed sets IsDisclosed field to given value.

### HasIsDisclosed

`func (o *CrossConnectEndCustomer) HasIsDisclosed() bool`

HasIsDisclosed returns a boolean if a field has been set.

### GetName

`func (o *CrossConnectEndCustomer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrossConnectEndCustomer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrossConnectEndCustomer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CrossConnectEndCustomer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMdmId

`func (o *CrossConnectEndCustomer) GetMdmId() string`

GetMdmId returns the MdmId field if non-nil, zero value otherwise.

### GetMdmIdOk

`func (o *CrossConnectEndCustomer) GetMdmIdOk() (*string, bool)`

GetMdmIdOk returns a tuple with the MdmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmId

`func (o *CrossConnectEndCustomer) SetMdmId(v string)`

SetMdmId sets MdmId field to given value.

### HasMdmId

`func (o *CrossConnectEndCustomer) HasMdmId() bool`

HasMdmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


