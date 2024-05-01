# RoutingProtocolPeeringConnectionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | Ip block URI | 
**Uuid** | **string** |  | 
**Type** | [**ConnectionType**](ConnectionType.md) |  | 

## Methods

### NewRoutingProtocolPeeringConnectionItem

`func NewRoutingProtocolPeeringConnectionItem(href string, uuid string, type_ ConnectionType, ) *RoutingProtocolPeeringConnectionItem`

NewRoutingProtocolPeeringConnectionItem instantiates a new RoutingProtocolPeeringConnectionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolPeeringConnectionItemWithDefaults

`func NewRoutingProtocolPeeringConnectionItemWithDefaults() *RoutingProtocolPeeringConnectionItem`

NewRoutingProtocolPeeringConnectionItemWithDefaults instantiates a new RoutingProtocolPeeringConnectionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *RoutingProtocolPeeringConnectionItem) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RoutingProtocolPeeringConnectionItem) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RoutingProtocolPeeringConnectionItem) SetHref(v string)`

SetHref sets Href field to given value.


### GetUuid

`func (o *RoutingProtocolPeeringConnectionItem) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoutingProtocolPeeringConnectionItem) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoutingProtocolPeeringConnectionItem) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *RoutingProtocolPeeringConnectionItem) GetType() ConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolPeeringConnectionItem) GetTypeOk() (*ConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolPeeringConnectionItem) SetType(v ConnectionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


