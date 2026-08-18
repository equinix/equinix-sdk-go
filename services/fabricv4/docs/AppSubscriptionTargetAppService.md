# AppSubscriptionTargetAppService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] [readonly] 
**Type** | Pointer to **string** | App Service Type | [optional] 
**Uuid** | **string** | Equinix-assigned access point identifier | 

## Methods

### NewAppSubscriptionTargetAppService

`func NewAppSubscriptionTargetAppService(uuid string, ) *AppSubscriptionTargetAppService`

NewAppSubscriptionTargetAppService instantiates a new AppSubscriptionTargetAppService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSubscriptionTargetAppServiceWithDefaults

`func NewAppSubscriptionTargetAppServiceWithDefaults() *AppSubscriptionTargetAppService`

NewAppSubscriptionTargetAppServiceWithDefaults instantiates a new AppSubscriptionTargetAppService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AppSubscriptionTargetAppService) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppSubscriptionTargetAppService) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppSubscriptionTargetAppService) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppSubscriptionTargetAppService) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AppSubscriptionTargetAppService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppSubscriptionTargetAppService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppSubscriptionTargetAppService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AppSubscriptionTargetAppService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *AppSubscriptionTargetAppService) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppSubscriptionTargetAppService) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppSubscriptionTargetAppService) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


