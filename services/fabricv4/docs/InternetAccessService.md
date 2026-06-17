# InternetAccessService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | Service URL path | 
**Type** | [**InternetAccessServiceType**](InternetAccessServiceType.md) |  | 
**Uuid** | **string** | Unique identifier for the EIA Service | 
**Name** | **string** | The name of the EIA Service | 
**Bandwidth** | Pointer to **int32** | Bandwidth of the service | [optional] 
**BandwidthCommit** | Pointer to **int32** | Minimum bandwidth commit for burst billing variant of the service | [optional] 
**State** | [**InternetAccessServiceState**](InternetAccessServiceState.md) |  | 
**Change** | [**InternetAccessChange**](InternetAccessChange.md) |  | 
**Locations** | Pointer to [**[]InternetAccessLocation**](InternetAccessLocation.md) | List of locations associated with the service | [optional] 
**RoutingProtocol** | [**InternetAccessRoutingProtocol**](InternetAccessRoutingProtocol.md) |  | 
**Billing** | [**InternetAccessBilling**](InternetAccessBilling.md) |  | 
**Account** | [**InternetAccessAccount**](InternetAccessAccount.md) |  | 
**Project** | [**Project**](Project.md) |  | 
**Order** | [**InternetAccessOrder**](InternetAccessOrder.md) |  | 
**ChangeLog** | [**Changelog**](Changelog.md) |  | 
**UseCase** | [**InternetAccessUseCase**](InternetAccessUseCase.md) |  | 

## Methods

### NewInternetAccessService

`func NewInternetAccessService(href string, type_ InternetAccessServiceType, uuid string, name string, state InternetAccessServiceState, change InternetAccessChange, routingProtocol InternetAccessRoutingProtocol, billing InternetAccessBilling, account InternetAccessAccount, project Project, order InternetAccessOrder, changeLog Changelog, useCase InternetAccessUseCase, ) *InternetAccessService`

NewInternetAccessService instantiates a new InternetAccessService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessServiceWithDefaults

`func NewInternetAccessServiceWithDefaults() *InternetAccessService`

NewInternetAccessServiceWithDefaults instantiates a new InternetAccessService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *InternetAccessService) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *InternetAccessService) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *InternetAccessService) SetHref(v string)`

SetHref sets Href field to given value.


### GetType

`func (o *InternetAccessService) GetType() InternetAccessServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternetAccessService) GetTypeOk() (*InternetAccessServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternetAccessService) SetType(v InternetAccessServiceType)`

SetType sets Type field to given value.


### GetUuid

`func (o *InternetAccessService) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InternetAccessService) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InternetAccessService) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *InternetAccessService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternetAccessService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternetAccessService) SetName(v string)`

SetName sets Name field to given value.


### GetBandwidth

`func (o *InternetAccessService) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *InternetAccessService) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *InternetAccessService) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *InternetAccessService) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetBandwidthCommit

`func (o *InternetAccessService) GetBandwidthCommit() int32`

GetBandwidthCommit returns the BandwidthCommit field if non-nil, zero value otherwise.

### GetBandwidthCommitOk

`func (o *InternetAccessService) GetBandwidthCommitOk() (*int32, bool)`

GetBandwidthCommitOk returns a tuple with the BandwidthCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthCommit

`func (o *InternetAccessService) SetBandwidthCommit(v int32)`

SetBandwidthCommit sets BandwidthCommit field to given value.

### HasBandwidthCommit

`func (o *InternetAccessService) HasBandwidthCommit() bool`

HasBandwidthCommit returns a boolean if a field has been set.

### GetState

`func (o *InternetAccessService) GetState() InternetAccessServiceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InternetAccessService) GetStateOk() (*InternetAccessServiceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InternetAccessService) SetState(v InternetAccessServiceState)`

SetState sets State field to given value.


### GetChange

`func (o *InternetAccessService) GetChange() InternetAccessChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *InternetAccessService) GetChangeOk() (*InternetAccessChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *InternetAccessService) SetChange(v InternetAccessChange)`

SetChange sets Change field to given value.


### GetLocations

`func (o *InternetAccessService) GetLocations() []InternetAccessLocation`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *InternetAccessService) GetLocationsOk() (*[]InternetAccessLocation, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *InternetAccessService) SetLocations(v []InternetAccessLocation)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *InternetAccessService) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetRoutingProtocol

`func (o *InternetAccessService) GetRoutingProtocol() InternetAccessRoutingProtocol`

GetRoutingProtocol returns the RoutingProtocol field if non-nil, zero value otherwise.

### GetRoutingProtocolOk

`func (o *InternetAccessService) GetRoutingProtocolOk() (*InternetAccessRoutingProtocol, bool)`

GetRoutingProtocolOk returns a tuple with the RoutingProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingProtocol

`func (o *InternetAccessService) SetRoutingProtocol(v InternetAccessRoutingProtocol)`

SetRoutingProtocol sets RoutingProtocol field to given value.


### GetBilling

`func (o *InternetAccessService) GetBilling() InternetAccessBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *InternetAccessService) GetBillingOk() (*InternetAccessBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *InternetAccessService) SetBilling(v InternetAccessBilling)`

SetBilling sets Billing field to given value.


### GetAccount

`func (o *InternetAccessService) GetAccount() InternetAccessAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InternetAccessService) GetAccountOk() (*InternetAccessAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InternetAccessService) SetAccount(v InternetAccessAccount)`

SetAccount sets Account field to given value.


### GetProject

`func (o *InternetAccessService) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *InternetAccessService) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *InternetAccessService) SetProject(v Project)`

SetProject sets Project field to given value.


### GetOrder

`func (o *InternetAccessService) GetOrder() InternetAccessOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InternetAccessService) GetOrderOk() (*InternetAccessOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InternetAccessService) SetOrder(v InternetAccessOrder)`

SetOrder sets Order field to given value.


### GetChangeLog

`func (o *InternetAccessService) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *InternetAccessService) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *InternetAccessService) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.


### GetUseCase

`func (o *InternetAccessService) GetUseCase() InternetAccessUseCase`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *InternetAccessService) GetUseCaseOk() (*InternetAccessUseCase, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *InternetAccessService) SetUseCase(v InternetAccessUseCase)`

SetUseCase sets UseCase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


