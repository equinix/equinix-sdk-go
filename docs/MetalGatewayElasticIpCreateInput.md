# MetalGatewayElasticIpCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | An IP address (or IP Address range) contained within one of the project&#39;s IP Reservations | 
**NextHop** | **string** | An IP address contained within the Metal Gateways&#39; IP Reservation range. | 
**Customdata** | Pointer to **map[string]interface{}** | Optional User-defined JSON object value. | [optional] 
**Tags** | Pointer to **[]string** | Optional list of User-defined tags. Can be used by users to provide additional details or context regarding the purpose or usage of this resource. | [optional] 

## Methods

### NewMetalGatewayElasticIpCreateInput

`func NewMetalGatewayElasticIpCreateInput(address string, nextHop string, ) *MetalGatewayElasticIpCreateInput`

NewMetalGatewayElasticIpCreateInput instantiates a new MetalGatewayElasticIpCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetalGatewayElasticIpCreateInputWithDefaults

`func NewMetalGatewayElasticIpCreateInputWithDefaults() *MetalGatewayElasticIpCreateInput`

NewMetalGatewayElasticIpCreateInputWithDefaults instantiates a new MetalGatewayElasticIpCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MetalGatewayElasticIpCreateInput) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MetalGatewayElasticIpCreateInput) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MetalGatewayElasticIpCreateInput) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNextHop

`func (o *MetalGatewayElasticIpCreateInput) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *MetalGatewayElasticIpCreateInput) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *MetalGatewayElasticIpCreateInput) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.


### GetCustomdata

`func (o *MetalGatewayElasticIpCreateInput) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *MetalGatewayElasticIpCreateInput) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *MetalGatewayElasticIpCreateInput) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *MetalGatewayElasticIpCreateInput) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetTags

`func (o *MetalGatewayElasticIpCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetalGatewayElasticIpCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetalGatewayElasticIpCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetalGatewayElasticIpCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


