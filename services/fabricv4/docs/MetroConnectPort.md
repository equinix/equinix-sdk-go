# MetroConnectPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Port Type | [optional] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 

## Methods

### NewMetroConnectPort

`func NewMetroConnectPort() *MetroConnectPort`

NewMetroConnectPort instantiates a new MetroConnectPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetroConnectPortWithDefaults

`func NewMetroConnectPortWithDefaults() *MetroConnectPort`

NewMetroConnectPortWithDefaults instantiates a new MetroConnectPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *MetroConnectPort) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MetroConnectPort) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MetroConnectPort) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *MetroConnectPort) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *MetroConnectPort) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetroConnectPort) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetroConnectPort) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetroConnectPort) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *MetroConnectPort) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MetroConnectPort) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MetroConnectPort) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MetroConnectPort) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetLocation

`func (o *MetroConnectPort) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MetroConnectPort) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MetroConnectPort) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MetroConnectPort) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


