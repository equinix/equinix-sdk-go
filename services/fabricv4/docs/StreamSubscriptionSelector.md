# StreamSubscriptionSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Include** | Pointer to **[]string** |  | [optional] 
**Except** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStreamSubscriptionSelector

`func NewStreamSubscriptionSelector() *StreamSubscriptionSelector`

NewStreamSubscriptionSelector instantiates a new StreamSubscriptionSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSubscriptionSelectorWithDefaults

`func NewStreamSubscriptionSelectorWithDefaults() *StreamSubscriptionSelector`

NewStreamSubscriptionSelectorWithDefaults instantiates a new StreamSubscriptionSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInclude

`func (o *StreamSubscriptionSelector) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *StreamSubscriptionSelector) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *StreamSubscriptionSelector) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *StreamSubscriptionSelector) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### GetExcept

`func (o *StreamSubscriptionSelector) GetExcept() []string`

GetExcept returns the Except field if non-nil, zero value otherwise.

### GetExceptOk

`func (o *StreamSubscriptionSelector) GetExceptOk() (*[]string, bool)`

GetExceptOk returns a tuple with the Except field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcept

`func (o *StreamSubscriptionSelector) SetExcept(v []string)`

SetExcept sets Except field to given value.

### HasExcept

`func (o *StreamSubscriptionSelector) HasExcept() bool`

HasExcept returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


