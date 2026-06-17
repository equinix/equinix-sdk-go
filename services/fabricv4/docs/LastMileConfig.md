# LastMileConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceType** | Pointer to **string** | Last mile service type | [optional] 
**Bandwidth** | Pointer to **int32** | Last mile bandwidth in Mbps | [optional] 
**Address** | Pointer to **string** | Last mile address | [optional] 
**Notifications** | Pointer to [**[]LastMileNotificationInfo**](LastMileNotificationInfo.md) | Last mile notification contacts | [optional] 

## Methods

### NewLastMileConfig

`func NewLastMileConfig() *LastMileConfig`

NewLastMileConfig instantiates a new LastMileConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastMileConfigWithDefaults

`func NewLastMileConfigWithDefaults() *LastMileConfig`

NewLastMileConfigWithDefaults instantiates a new LastMileConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceType

`func (o *LastMileConfig) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *LastMileConfig) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *LastMileConfig) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *LastMileConfig) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetBandwidth

`func (o *LastMileConfig) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *LastMileConfig) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *LastMileConfig) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *LastMileConfig) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetAddress

`func (o *LastMileConfig) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LastMileConfig) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LastMileConfig) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LastMileConfig) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNotifications

`func (o *LastMileConfig) GetNotifications() []LastMileNotificationInfo`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *LastMileConfig) GetNotificationsOk() (*[]LastMileNotificationInfo, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *LastMileConfig) SetNotifications(v []LastMileNotificationInfo)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *LastMileConfig) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


