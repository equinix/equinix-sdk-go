# ApiServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Route** | Pointer to **string** | service routes | [optional] 
**Status** | Pointer to **string** | service status | [optional] 
**ChangedDateTime** | Pointer to **string** | service status change date | [optional] 

## Methods

### NewApiServices

`func NewApiServices() *ApiServices`

NewApiServices instantiates a new ApiServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServicesWithDefaults

`func NewApiServicesWithDefaults() *ApiServices`

NewApiServicesWithDefaults instantiates a new ApiServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoute

`func (o *ApiServices) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *ApiServices) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *ApiServices) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *ApiServices) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetStatus

`func (o *ApiServices) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiServices) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiServices) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiServices) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetChangedDateTime

`func (o *ApiServices) GetChangedDateTime() string`

GetChangedDateTime returns the ChangedDateTime field if non-nil, zero value otherwise.

### GetChangedDateTimeOk

`func (o *ApiServices) GetChangedDateTimeOk() (*string, bool)`

GetChangedDateTimeOk returns a tuple with the ChangedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedDateTime

`func (o *ApiServices) SetChangedDateTime(v string)`

SetChangedDateTime sets ChangedDateTime field to given value.

### HasChangedDateTime

`func (o *ApiServices) HasChangedDateTime() bool`

HasChangedDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


