# ServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | [**ServiceTypeV2**](ServiceTypeV2.md) |  | 
**Connections** | **[]string** | Collection of service connections uuids | 
**RoutingProtocol** | [**RoutingProtocolRequest**](RoutingProtocolRequest.md) |  | 
**Order** | Pointer to [**ServiceOrderRequest**](ServiceOrderRequest.md) |  | [optional] 

## Methods

### NewServiceRequest

`func NewServiceRequest(type_ ServiceTypeV2, connections []string, routingProtocol RoutingProtocolRequest, ) *ServiceRequest`

NewServiceRequest instantiates a new ServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRequestWithDefaults

`func NewServiceRequestWithDefaults() *ServiceRequest`

NewServiceRequestWithDefaults instantiates a new ServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *ServiceRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *ServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ServiceRequest) GetType() ServiceTypeV2`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceRequest) GetTypeOk() (*ServiceTypeV2, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceRequest) SetType(v ServiceTypeV2)`

SetType sets Type field to given value.


### GetConnections

`func (o *ServiceRequest) GetConnections() []string`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ServiceRequest) GetConnectionsOk() (*[]string, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ServiceRequest) SetConnections(v []string)`

SetConnections sets Connections field to given value.


### GetRoutingProtocol

`func (o *ServiceRequest) GetRoutingProtocol() RoutingProtocolRequest`

GetRoutingProtocol returns the RoutingProtocol field if non-nil, zero value otherwise.

### GetRoutingProtocolOk

`func (o *ServiceRequest) GetRoutingProtocolOk() (*RoutingProtocolRequest, bool)`

GetRoutingProtocolOk returns a tuple with the RoutingProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingProtocol

`func (o *ServiceRequest) SetRoutingProtocol(v RoutingProtocolRequest)`

SetRoutingProtocol sets RoutingProtocol field to given value.


### GetOrder

`func (o *ServiceRequest) GetOrder() ServiceOrderRequest`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ServiceRequest) GetOrderOk() (*ServiceOrderRequest, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ServiceRequest) SetOrder(v ServiceOrderRequest)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ServiceRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


