# AsnSearchSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**AsnSearchSortCriteriaDirection**](AsnSearchSortCriteriaDirection.md) |  | [optional] [default to ASNSEARCHSORTCRITERIADIRECTION_DESC]
**Property** | Pointer to **string** | Possible field names to use on sorting | [optional] [default to "/changeLog/updatedDateTime"]

## Methods

### NewAsnSearchSortCriteria

`func NewAsnSearchSortCriteria() *AsnSearchSortCriteria`

NewAsnSearchSortCriteria instantiates a new AsnSearchSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsnSearchSortCriteriaWithDefaults

`func NewAsnSearchSortCriteriaWithDefaults() *AsnSearchSortCriteria`

NewAsnSearchSortCriteriaWithDefaults instantiates a new AsnSearchSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *AsnSearchSortCriteria) GetDirection() AsnSearchSortCriteriaDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AsnSearchSortCriteria) GetDirectionOk() (*AsnSearchSortCriteriaDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AsnSearchSortCriteria) SetDirection(v AsnSearchSortCriteriaDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *AsnSearchSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *AsnSearchSortCriteria) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AsnSearchSortCriteria) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AsnSearchSortCriteria) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *AsnSearchSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


