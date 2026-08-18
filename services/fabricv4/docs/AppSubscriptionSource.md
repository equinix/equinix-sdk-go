# AppSubscriptionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppLink** | Pointer to [**AppSubscriptionSourceAppLink**](AppSubscriptionSourceAppLink.md) |  | [optional] 
**IpSubnets** | **[]string** | List of IP subnets in CIDR notation | 

## Methods

### NewAppSubscriptionSource

`func NewAppSubscriptionSource(ipSubnets []string, ) *AppSubscriptionSource`

NewAppSubscriptionSource instantiates a new AppSubscriptionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSubscriptionSourceWithDefaults

`func NewAppSubscriptionSourceWithDefaults() *AppSubscriptionSource`

NewAppSubscriptionSourceWithDefaults instantiates a new AppSubscriptionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppLink

`func (o *AppSubscriptionSource) GetAppLink() AppSubscriptionSourceAppLink`

GetAppLink returns the AppLink field if non-nil, zero value otherwise.

### GetAppLinkOk

`func (o *AppSubscriptionSource) GetAppLinkOk() (*AppSubscriptionSourceAppLink, bool)`

GetAppLinkOk returns a tuple with the AppLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLink

`func (o *AppSubscriptionSource) SetAppLink(v AppSubscriptionSourceAppLink)`

SetAppLink sets AppLink field to given value.

### HasAppLink

`func (o *AppSubscriptionSource) HasAppLink() bool`

HasAppLink returns a boolean if a field has been set.

### GetIpSubnets

`func (o *AppSubscriptionSource) GetIpSubnets() []string`

GetIpSubnets returns the IpSubnets field if non-nil, zero value otherwise.

### GetIpSubnetsOk

`func (o *AppSubscriptionSource) GetIpSubnetsOk() (*[]string, bool)`

GetIpSubnetsOk returns a tuple with the IpSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSubnets

`func (o *AppSubscriptionSource) SetIpSubnets(v []string)`

SetIpSubnets sets IpSubnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


