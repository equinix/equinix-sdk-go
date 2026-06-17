# InternetAccessRoutingProtocolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**InternetAccessRoutingProtocolType**](InternetAccessRoutingProtocolType.md) |  | 
**CustomerRoutes** | [**[]InternetAccessCustomerRouteRequest**](InternetAccessCustomerRouteRequest.md) |  | 

## Methods

### NewInternetAccessRoutingProtocolRequest

`func NewInternetAccessRoutingProtocolRequest(type_ InternetAccessRoutingProtocolType, customerRoutes []InternetAccessCustomerRouteRequest, ) *InternetAccessRoutingProtocolRequest`

NewInternetAccessRoutingProtocolRequest instantiates a new InternetAccessRoutingProtocolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessRoutingProtocolRequestWithDefaults

`func NewInternetAccessRoutingProtocolRequestWithDefaults() *InternetAccessRoutingProtocolRequest`

NewInternetAccessRoutingProtocolRequestWithDefaults instantiates a new InternetAccessRoutingProtocolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InternetAccessRoutingProtocolRequest) GetType() InternetAccessRoutingProtocolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternetAccessRoutingProtocolRequest) GetTypeOk() (*InternetAccessRoutingProtocolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternetAccessRoutingProtocolRequest) SetType(v InternetAccessRoutingProtocolType)`

SetType sets Type field to given value.


### GetCustomerRoutes

`func (o *InternetAccessRoutingProtocolRequest) GetCustomerRoutes() []InternetAccessCustomerRouteRequest`

GetCustomerRoutes returns the CustomerRoutes field if non-nil, zero value otherwise.

### GetCustomerRoutesOk

`func (o *InternetAccessRoutingProtocolRequest) GetCustomerRoutesOk() (*[]InternetAccessCustomerRouteRequest, bool)`

GetCustomerRoutesOk returns a tuple with the CustomerRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRoutes

`func (o *InternetAccessRoutingProtocolRequest) SetCustomerRoutes(v []InternetAccessCustomerRouteRequest)`

SetCustomerRoutes sets CustomerRoutes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


