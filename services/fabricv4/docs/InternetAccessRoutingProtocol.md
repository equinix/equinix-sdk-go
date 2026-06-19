# InternetAccessRoutingProtocol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**InternetAccessRoutingProtocolType**](InternetAccessRoutingProtocolType.md) |  | 
**CustomerRoutes** | [**[]InternetAccessCustomerRoute**](InternetAccessCustomerRoute.md) |  | 
**Connections** | [**[]InternetAccessConnection**](InternetAccessConnection.md) |  | 

## Methods

### NewInternetAccessRoutingProtocol

`func NewInternetAccessRoutingProtocol(type_ InternetAccessRoutingProtocolType, customerRoutes []InternetAccessCustomerRoute, connections []InternetAccessConnection, ) *InternetAccessRoutingProtocol`

NewInternetAccessRoutingProtocol instantiates a new InternetAccessRoutingProtocol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessRoutingProtocolWithDefaults

`func NewInternetAccessRoutingProtocolWithDefaults() *InternetAccessRoutingProtocol`

NewInternetAccessRoutingProtocolWithDefaults instantiates a new InternetAccessRoutingProtocol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InternetAccessRoutingProtocol) GetType() InternetAccessRoutingProtocolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternetAccessRoutingProtocol) GetTypeOk() (*InternetAccessRoutingProtocolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternetAccessRoutingProtocol) SetType(v InternetAccessRoutingProtocolType)`

SetType sets Type field to given value.


### GetCustomerRoutes

`func (o *InternetAccessRoutingProtocol) GetCustomerRoutes() []InternetAccessCustomerRoute`

GetCustomerRoutes returns the CustomerRoutes field if non-nil, zero value otherwise.

### GetCustomerRoutesOk

`func (o *InternetAccessRoutingProtocol) GetCustomerRoutesOk() (*[]InternetAccessCustomerRoute, bool)`

GetCustomerRoutesOk returns a tuple with the CustomerRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoutes

`func (o *InternetAccessRoutingProtocol) SetCustomerRoutes(v []InternetAccessCustomerRoute)`

SetCustomerRoutes sets CustomerRoutes field to given value.


### GetConnections

`func (o *InternetAccessRoutingProtocol) GetConnections() []InternetAccessConnection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *InternetAccessRoutingProtocol) GetConnectionsOk() (*[]InternetAccessConnection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *InternetAccessRoutingProtocol) SetConnections(v []InternetAccessConnection)`

SetConnections sets Connections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


