# AppSubscriptionSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppLink** | Pointer to [**AppSubscriptionSourceRequestAppLink**](AppSubscriptionSourceRequestAppLink.md) |  | [optional] 
**IpSubnets** | **[]string** | List of IP subnets in CIDR notation | 

## Methods

### NewAppSubscriptionSourceRequest

`func NewAppSubscriptionSourceRequest(ipSubnets []string, ) *AppSubscriptionSourceRequest`

NewAppSubscriptionSourceRequest instantiates a new AppSubscriptionSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSubscriptionSourceRequestWithDefaults

`func NewAppSubscriptionSourceRequestWithDefaults() *AppSubscriptionSourceRequest`

NewAppSubscriptionSourceRequestWithDefaults instantiates a new AppSubscriptionSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppLink

`func (o *AppSubscriptionSourceRequest) GetAppLink() AppSubscriptionSourceRequestAppLink`

GetAppLink returns the AppLink field if non-nil, zero value otherwise.

### GetAppLinkOk

`func (o *AppSubscriptionSourceRequest) GetAppLinkOk() (*AppSubscriptionSourceRequestAppLink, bool)`

GetAppLinkOk returns a tuple with the AppLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLink

`func (o *AppSubscriptionSourceRequest) SetAppLink(v AppSubscriptionSourceRequestAppLink)`

SetAppLink sets AppLink field to given value.

### HasAppLink

`func (o *AppSubscriptionSourceRequest) HasAppLink() bool`

HasAppLink returns a boolean if a field has been set.

### GetIpSubnets

`func (o *AppSubscriptionSourceRequest) GetIpSubnets() []string`

GetIpSubnets returns the IpSubnets field if non-nil, zero value otherwise.

### GetIpSubnetsOk

`func (o *AppSubscriptionSourceRequest) GetIpSubnetsOk() (*[]string, bool)`

GetIpSubnetsOk returns a tuple with the IpSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSubnets

`func (o *AppSubscriptionSourceRequest) SetIpSubnets(v []string)`

SetIpSubnets sets IpSubnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


