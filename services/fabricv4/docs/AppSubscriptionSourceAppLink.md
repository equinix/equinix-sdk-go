# AppSubscriptionSourceAppLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] [readonly] 
**Type** | Pointer to **string** | App Link Type | [optional] 
**Uuid** | **string** | Equinix-assigned access point identifier | 

## Methods

### NewAppSubscriptionSourceAppLink

`func NewAppSubscriptionSourceAppLink(uuid string, ) *AppSubscriptionSourceAppLink`

NewAppSubscriptionSourceAppLink instantiates a new AppSubscriptionSourceAppLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSubscriptionSourceAppLinkWithDefaults

`func NewAppSubscriptionSourceAppLinkWithDefaults() *AppSubscriptionSourceAppLink`

NewAppSubscriptionSourceAppLinkWithDefaults instantiates a new AppSubscriptionSourceAppLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AppSubscriptionSourceAppLink) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppSubscriptionSourceAppLink) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppSubscriptionSourceAppLink) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppSubscriptionSourceAppLink) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AppSubscriptionSourceAppLink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppSubscriptionSourceAppLink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppSubscriptionSourceAppLink) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppSubscriptionSourceAppLink) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *AppSubscriptionSourceAppLink) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppSubscriptionSourceAppLink) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppSubscriptionSourceAppLink) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


