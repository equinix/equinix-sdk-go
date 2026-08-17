# AppSubscriptionTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppService** | Pointer to [**AppSubscriptionTargetAppService**](AppSubscriptionTargetAppService.md) |  | [optional] 
**GeoScope** | **string** | Geo scope | 
**Prioritization** | [**AppSubscriptionPrioritization**](AppSubscriptionPrioritization.md) |  | 

## Methods

### NewAppSubscriptionTarget

`func NewAppSubscriptionTarget(geoScope string, prioritization AppSubscriptionPrioritization, ) *AppSubscriptionTarget`

NewAppSubscriptionTarget instantiates a new AppSubscriptionTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSubscriptionTargetWithDefaults

`func NewAppSubscriptionTargetWithDefaults() *AppSubscriptionTarget`

NewAppSubscriptionTargetWithDefaults instantiates a new AppSubscriptionTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppService

`func (o *AppSubscriptionTarget) GetAppService() AppSubscriptionTargetAppService`

GetAppService returns the AppService field if non-nil, zero value otherwise.

### GetAppServiceOk

`func (o *AppSubscriptionTarget) GetAppServiceOk() (*AppSubscriptionTargetAppService, bool)`

GetAppServiceOk returns a tuple with the AppService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppService

`func (o *AppSubscriptionTarget) SetAppService(v AppSubscriptionTargetAppService)`

SetAppService sets AppService field to given value.

### HasAppService

`func (o *AppSubscriptionTarget) HasAppService() bool`

HasAppService returns a boolean if a field has been set.

### GetGeoScope

`func (o *AppSubscriptionTarget) GetGeoScope() string`

GetGeoScope returns the GeoScope field if non-nil, zero value otherwise.

### GetGeoScopeOk

`func (o *AppSubscriptionTarget) GetGeoScopeOk() (*string, bool)`

GetGeoScopeOk returns a tuple with the GeoScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoScope

`func (o *AppSubscriptionTarget) SetGeoScope(v string)`

SetGeoScope sets GeoScope field to given value.


### GetPrioritization

`func (o *AppSubscriptionTarget) GetPrioritization() AppSubscriptionPrioritization`

GetPrioritization returns the Prioritization field if non-nil, zero value otherwise.

### GetPrioritizationOk

`func (o *AppSubscriptionTarget) GetPrioritizationOk() (*AppSubscriptionPrioritization, bool)`

GetPrioritizationOk returns a tuple with the Prioritization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritization

`func (o *AppSubscriptionTarget) SetPrioritization(v AppSubscriptionPrioritization)`

SetPrioritization sets Prioritization field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


