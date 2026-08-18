# AppServiceAttachedAppSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] [readonly] 
**Type** | [**AppSubscriptionType**](AppSubscriptionType.md) |  | [default to APPSUBSCRIPTIONTYPE_APP_SUBSCRIPTION]
**Uuid** | **string** | Equinix-assigned access point identifier | 
**State** | [**AppSubscriptionState**](AppSubscriptionState.md) |  | 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 

## Methods

### NewAppServiceAttachedAppSubscription

`func NewAppServiceAttachedAppSubscription(type_ AppSubscriptionType, uuid string, state AppSubscriptionState, ) *AppServiceAttachedAppSubscription`

NewAppServiceAttachedAppSubscription instantiates a new AppServiceAttachedAppSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceAttachedAppSubscriptionWithDefaults

`func NewAppServiceAttachedAppSubscriptionWithDefaults() *AppServiceAttachedAppSubscription`

NewAppServiceAttachedAppSubscriptionWithDefaults instantiates a new AppServiceAttachedAppSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AppServiceAttachedAppSubscription) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppServiceAttachedAppSubscription) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppServiceAttachedAppSubscription) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppServiceAttachedAppSubscription) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AppServiceAttachedAppSubscription) GetType() AppSubscriptionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppServiceAttachedAppSubscription) GetTypeOk() (*AppSubscriptionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppServiceAttachedAppSubscription) SetType(v AppSubscriptionType)`

SetType sets Type field to given value.


### GetUuid

`func (o *AppServiceAttachedAppSubscription) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppServiceAttachedAppSubscription) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppServiceAttachedAppSubscription) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetState

`func (o *AppServiceAttachedAppSubscription) GetState() AppSubscriptionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppServiceAttachedAppSubscription) GetStateOk() (*AppSubscriptionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppServiceAttachedAppSubscription) SetState(v AppSubscriptionState)`

SetState sets State field to given value.


### GetAccount

`func (o *AppServiceAttachedAppSubscription) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AppServiceAttachedAppSubscription) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AppServiceAttachedAppSubscription) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AppServiceAttachedAppSubscription) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


