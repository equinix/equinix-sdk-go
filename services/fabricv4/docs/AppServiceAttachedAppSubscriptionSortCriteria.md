# AppServiceAttachedAppSubscriptionSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**AppServiceAttachedAppSubscriptionSortDirection**](AppServiceAttachedAppSubscriptionSortDirection.md) |  | [optional] [default to APPSERVICEATTACHEDAPPSUBSCRIPTIONSORTDIRECTION_DESC]
**Property** | Pointer to **string** | Possible field names to use on &#39;400_InvalidSorting&#39;:   * &#x60;/uuid&#x60; - App Subscription UUID   * &#x60;/state&#x60; - App Subscription lifecycle state  | [optional] [default to "/uuid"]

## Methods

### NewAppServiceAttachedAppSubscriptionSortCriteria

`func NewAppServiceAttachedAppSubscriptionSortCriteria() *AppServiceAttachedAppSubscriptionSortCriteria`

NewAppServiceAttachedAppSubscriptionSortCriteria instantiates a new AppServiceAttachedAppSubscriptionSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceAttachedAppSubscriptionSortCriteriaWithDefaults

`func NewAppServiceAttachedAppSubscriptionSortCriteriaWithDefaults() *AppServiceAttachedAppSubscriptionSortCriteria`

NewAppServiceAttachedAppSubscriptionSortCriteriaWithDefaults instantiates a new AppServiceAttachedAppSubscriptionSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *AppServiceAttachedAppSubscriptionSortCriteria) GetDirection() AppServiceAttachedAppSubscriptionSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AppServiceAttachedAppSubscriptionSortCriteria) GetDirectionOk() (*AppServiceAttachedAppSubscriptionSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AppServiceAttachedAppSubscriptionSortCriteria) SetDirection(v AppServiceAttachedAppSubscriptionSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *AppServiceAttachedAppSubscriptionSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *AppServiceAttachedAppSubscriptionSortCriteria) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AppServiceAttachedAppSubscriptionSortCriteria) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AppServiceAttachedAppSubscriptionSortCriteria) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *AppServiceAttachedAppSubscriptionSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


