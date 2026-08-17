# AppLinkAttachServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeoScope** | **string** | Geo scope for the App Service | 
**DestinationIp** | **string** | Target IP for forwarding API requests | 

## Methods

### NewAppLinkAttachServiceRequest

`func NewAppLinkAttachServiceRequest(geoScope string, destinationIp string, ) *AppLinkAttachServiceRequest`

NewAppLinkAttachServiceRequest instantiates a new AppLinkAttachServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkAttachServiceRequestWithDefaults

`func NewAppLinkAttachServiceRequestWithDefaults() *AppLinkAttachServiceRequest`

NewAppLinkAttachServiceRequestWithDefaults instantiates a new AppLinkAttachServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeoScope

`func (o *AppLinkAttachServiceRequest) GetGeoScope() string`

GetGeoScope returns the GeoScope field if non-nil, zero value otherwise.

### GetGeoScopeOk

`func (o *AppLinkAttachServiceRequest) GetGeoScopeOk() (*string, bool)`

GetGeoScopeOk returns a tuple with the GeoScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoScope

`func (o *AppLinkAttachServiceRequest) SetGeoScope(v string)`

SetGeoScope sets GeoScope field to given value.


### GetDestinationIp

`func (o *AppLinkAttachServiceRequest) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *AppLinkAttachServiceRequest) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *AppLinkAttachServiceRequest) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


