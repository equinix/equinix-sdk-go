# ProviderSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**ProviderExpression**](ProviderExpression.md) |  | [optional] 

## Methods

### NewProviderSearchRequest

`func NewProviderSearchRequest() *ProviderSearchRequest`

NewProviderSearchRequest instantiates a new ProviderSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderSearchRequestWithDefaults

`func NewProviderSearchRequestWithDefaults() *ProviderSearchRequest`

NewProviderSearchRequestWithDefaults instantiates a new ProviderSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ProviderSearchRequest) GetFilter() ProviderExpression`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ProviderSearchRequest) GetFilterOk() (*ProviderExpression, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ProviderSearchRequest) SetFilter(v ProviderExpression)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ProviderSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


