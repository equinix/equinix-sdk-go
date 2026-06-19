# InternetAccessPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**InternetAccessServiceType**](InternetAccessServiceType.md) |  | 
**Name** | **string** | The name of the EIA Service | 
**Bandwidth** | Pointer to **int32** | Bandwidth of the service | [optional] 
**BandwidthCommit** | Pointer to **int32** | Minimum bandwidth commit for burst billing variant of the service | [optional] 
**RoutingProtocol** | [**InternetAccessRoutingProtocolRequest**](InternetAccessRoutingProtocolRequest.md) |  | 
**Order** | Pointer to [**InternetAccessOrderRequest**](InternetAccessOrderRequest.md) |  | [optional] 
**Billing** | [**InternetAccessPostRequestBilling**](InternetAccessPostRequestBilling.md) |  | 
**Project** | [**Project**](Project.md) |  | 
**Account** | [**InternetAccessAccount**](InternetAccessAccount.md) |  | 

## Methods

### NewInternetAccessPostRequest

`func NewInternetAccessPostRequest(type_ InternetAccessServiceType, name string, routingProtocol InternetAccessRoutingProtocolRequest, billing InternetAccessPostRequestBilling, project Project, account InternetAccessAccount, ) *InternetAccessPostRequest`

NewInternetAccessPostRequest instantiates a new InternetAccessPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessPostRequestWithDefaults

`func NewInternetAccessPostRequestWithDefaults() *InternetAccessPostRequest`

NewInternetAccessPostRequestWithDefaults instantiates a new InternetAccessPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InternetAccessPostRequest) GetType() InternetAccessServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternetAccessPostRequest) GetTypeOk() (*InternetAccessServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternetAccessPostRequest) SetType(v InternetAccessServiceType)`

SetType sets Type field to given value.


### GetName

`func (o *InternetAccessPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternetAccessPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternetAccessPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetBandwidth

`func (o *InternetAccessPostRequest) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *InternetAccessPostRequest) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *InternetAccessPostRequest) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *InternetAccessPostRequest) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetBandwidthCommit

`func (o *InternetAccessPostRequest) GetBandwidthCommit() int32`

GetBandwidthCommit returns the BandwidthCommit field if non-nil, zero value otherwise.

### GetBandwidthCommitOk

`func (o *InternetAccessPostRequest) GetBandwidthCommitOk() (*int32, bool)`

GetBandwidthCommitOk returns a tuple with the BandwidthCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthCommit

`func (o *InternetAccessPostRequest) SetBandwidthCommit(v int32)`

SetBandwidthCommit sets BandwidthCommit field to given value.

### HasBandwidthCommit

`func (o *InternetAccessPostRequest) HasBandwidthCommit() bool`

HasBandwidthCommit returns a boolean if a field has been set.

### GetRoutingProtocol

`func (o *InternetAccessPostRequest) GetRoutingProtocol() InternetAccessRoutingProtocolRequest`

GetRoutingProtocol returns the RoutingProtocol field if non-nil, zero value otherwise.

### GetRoutingProtocolOk

`func (o *InternetAccessPostRequest) GetRoutingProtocolOk() (*InternetAccessRoutingProtocolRequest, bool)`

GetRoutingProtocolOk returns a tuple with the RoutingProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingProtocol

`func (o *InternetAccessPostRequest) SetRoutingProtocol(v InternetAccessRoutingProtocolRequest)`

SetRoutingProtocol sets RoutingProtocol field to given value.


### GetOrder

`func (o *InternetAccessPostRequest) GetOrder() InternetAccessOrderRequest`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InternetAccessPostRequest) GetOrderOk() (*InternetAccessOrderRequest, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InternetAccessPostRequest) SetOrder(v InternetAccessOrderRequest)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *InternetAccessPostRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetBilling

`func (o *InternetAccessPostRequest) GetBilling() InternetAccessPostRequestBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *InternetAccessPostRequest) GetBillingOk() (*InternetAccessPostRequestBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *InternetAccessPostRequest) SetBilling(v InternetAccessPostRequestBilling)`

SetBilling sets Billing field to given value.


### GetProject

`func (o *InternetAccessPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *InternetAccessPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *InternetAccessPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.


### GetAccount

`func (o *InternetAccessPostRequest) GetAccount() InternetAccessAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InternetAccessPostRequest) GetAccountOk() (*InternetAccessAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InternetAccessPostRequest) SetAccount(v InternetAccessAccount)`

SetAccount sets Account field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


