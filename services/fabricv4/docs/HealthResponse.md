# HealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | The Canonical URL at which the resource resides. | [optional] 
**Version** | Pointer to **string** | Indicator of a version | [optional] 
**Release** | Pointer to **string** | release details. | [optional] 
**State** | Pointer to **string** | status of a service | [optional] 
**ApiServices** | Pointer to [**ApiServices**](ApiServices.md) |  | [optional] 

## Methods

### NewHealthResponse

`func NewHealthResponse() *HealthResponse`

NewHealthResponse instantiates a new HealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthResponseWithDefaults

`func NewHealthResponseWithDefaults() *HealthResponse`

NewHealthResponseWithDefaults instantiates a new HealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *HealthResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *HealthResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *HealthResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *HealthResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetVersion

`func (o *HealthResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HealthResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HealthResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HealthResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRelease

`func (o *HealthResponse) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *HealthResponse) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *HealthResponse) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *HealthResponse) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetState

`func (o *HealthResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HealthResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HealthResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HealthResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetApiServices

`func (o *HealthResponse) GetApiServices() ApiServices`

GetApiServices returns the ApiServices field if non-nil, zero value otherwise.

### GetApiServicesOk

`func (o *HealthResponse) GetApiServicesOk() (*ApiServices, bool)`

GetApiServicesOk returns a tuple with the ApiServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServices

`func (o *HealthResponse) SetApiServices(v ApiServices)`

SetApiServices sets ApiServices field to given value.

### HasApiServices

`func (o *HealthResponse) HasApiServices() bool`

HasApiServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


