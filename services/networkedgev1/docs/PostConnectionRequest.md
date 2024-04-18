# PostConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryName** | Pointer to **string** |  | [optional] 
**VirtualDeviceUUID** | Pointer to **string** |  | [optional] 
**ProfileUUID** | Pointer to **string** |  | [optional] 
**AuthorizationKey** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **int32** |  | [optional] 
**SpeedUnit** | Pointer to **string** |  | [optional] 
**Notifications** | Pointer to **[]string** |  | [optional] 
**PurchaseOrderNumber** | Pointer to **string** |  | [optional] 
**SellerMetroCode** | Pointer to **string** |  | [optional] 
**InterfaceId** | Pointer to **int32** |  | [optional] 
**SecondaryName** | Pointer to **string** |  | [optional] 
**NamedTag** | Pointer to **string** |  | [optional] 
**SecondaryVirtualDeviceUUID** | Pointer to **string** |  | [optional] 
**SecondaryProfileUUID** | Pointer to **string** |  | [optional] 
**SecondaryAuthorizationKey** | Pointer to **string** |  | [optional] 
**SecondarySellerMetroCode** | Pointer to **string** |  | [optional] 
**SecondarySpeed** | Pointer to **int32** |  | [optional] 
**SecondarySpeedUnit** | Pointer to **string** |  | [optional] 
**SecondaryNotifications** | Pointer to **[]string** |  | [optional] 
**SecondaryInterfaceId** | Pointer to **int32** |  | [optional] 
**PrimaryZSideVlanCTag** | Pointer to **int32** |  | [optional] 
**SecondaryZSideVlanCTag** | Pointer to **int32** |  | [optional] 
**PrimaryZSidePortUUID** | Pointer to **string** |  | [optional] 
**PrimaryZSideVlanSTag** | Pointer to **int32** |  | [optional] 
**SecondaryZSidePortUUID** | Pointer to **string** |  | [optional] 
**SecondaryZSideVlanSTag** | Pointer to **int32** |  | [optional] 

## Methods

### NewPostConnectionRequest

`func NewPostConnectionRequest() *PostConnectionRequest`

NewPostConnectionRequest instantiates a new PostConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostConnectionRequestWithDefaults

`func NewPostConnectionRequestWithDefaults() *PostConnectionRequest`

NewPostConnectionRequestWithDefaults instantiates a new PostConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryName

`func (o *PostConnectionRequest) GetPrimaryName() string`

GetPrimaryName returns the PrimaryName field if non-nil, zero value otherwise.

### GetPrimaryNameOk

`func (o *PostConnectionRequest) GetPrimaryNameOk() (*string, bool)`

GetPrimaryNameOk returns a tuple with the PrimaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryName

`func (o *PostConnectionRequest) SetPrimaryName(v string)`

SetPrimaryName sets PrimaryName field to given value.

### HasPrimaryName

`func (o *PostConnectionRequest) HasPrimaryName() bool`

HasPrimaryName returns a boolean if a field has been set.

### GetVirtualDeviceUUID

`func (o *PostConnectionRequest) GetVirtualDeviceUUID() string`

GetVirtualDeviceUUID returns the VirtualDeviceUUID field if non-nil, zero value otherwise.

### GetVirtualDeviceUUIDOk

`func (o *PostConnectionRequest) GetVirtualDeviceUUIDOk() (*string, bool)`

GetVirtualDeviceUUIDOk returns a tuple with the VirtualDeviceUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceUUID

`func (o *PostConnectionRequest) SetVirtualDeviceUUID(v string)`

SetVirtualDeviceUUID sets VirtualDeviceUUID field to given value.

### HasVirtualDeviceUUID

`func (o *PostConnectionRequest) HasVirtualDeviceUUID() bool`

HasVirtualDeviceUUID returns a boolean if a field has been set.

### GetProfileUUID

`func (o *PostConnectionRequest) GetProfileUUID() string`

GetProfileUUID returns the ProfileUUID field if non-nil, zero value otherwise.

### GetProfileUUIDOk

`func (o *PostConnectionRequest) GetProfileUUIDOk() (*string, bool)`

GetProfileUUIDOk returns a tuple with the ProfileUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUUID

`func (o *PostConnectionRequest) SetProfileUUID(v string)`

SetProfileUUID sets ProfileUUID field to given value.

### HasProfileUUID

`func (o *PostConnectionRequest) HasProfileUUID() bool`

HasProfileUUID returns a boolean if a field has been set.

### GetAuthorizationKey

`func (o *PostConnectionRequest) GetAuthorizationKey() string`

GetAuthorizationKey returns the AuthorizationKey field if non-nil, zero value otherwise.

### GetAuthorizationKeyOk

`func (o *PostConnectionRequest) GetAuthorizationKeyOk() (*string, bool)`

GetAuthorizationKeyOk returns a tuple with the AuthorizationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationKey

`func (o *PostConnectionRequest) SetAuthorizationKey(v string)`

SetAuthorizationKey sets AuthorizationKey field to given value.

### HasAuthorizationKey

`func (o *PostConnectionRequest) HasAuthorizationKey() bool`

HasAuthorizationKey returns a boolean if a field has been set.

### GetSpeed

`func (o *PostConnectionRequest) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *PostConnectionRequest) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *PostConnectionRequest) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *PostConnectionRequest) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetSpeedUnit

`func (o *PostConnectionRequest) GetSpeedUnit() string`

GetSpeedUnit returns the SpeedUnit field if non-nil, zero value otherwise.

### GetSpeedUnitOk

`func (o *PostConnectionRequest) GetSpeedUnitOk() (*string, bool)`

GetSpeedUnitOk returns a tuple with the SpeedUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedUnit

`func (o *PostConnectionRequest) SetSpeedUnit(v string)`

SetSpeedUnit sets SpeedUnit field to given value.

### HasSpeedUnit

`func (o *PostConnectionRequest) HasSpeedUnit() bool`

HasSpeedUnit returns a boolean if a field has been set.

### GetNotifications

`func (o *PostConnectionRequest) GetNotifications() []string`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *PostConnectionRequest) GetNotificationsOk() (*[]string, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *PostConnectionRequest) SetNotifications(v []string)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *PostConnectionRequest) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *PostConnectionRequest) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *PostConnectionRequest) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *PostConnectionRequest) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *PostConnectionRequest) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetSellerMetroCode

`func (o *PostConnectionRequest) GetSellerMetroCode() string`

GetSellerMetroCode returns the SellerMetroCode field if non-nil, zero value otherwise.

### GetSellerMetroCodeOk

`func (o *PostConnectionRequest) GetSellerMetroCodeOk() (*string, bool)`

GetSellerMetroCodeOk returns a tuple with the SellerMetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerMetroCode

`func (o *PostConnectionRequest) SetSellerMetroCode(v string)`

SetSellerMetroCode sets SellerMetroCode field to given value.

### HasSellerMetroCode

`func (o *PostConnectionRequest) HasSellerMetroCode() bool`

HasSellerMetroCode returns a boolean if a field has been set.

### GetInterfaceId

`func (o *PostConnectionRequest) GetInterfaceId() int32`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *PostConnectionRequest) GetInterfaceIdOk() (*int32, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *PostConnectionRequest) SetInterfaceId(v int32)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *PostConnectionRequest) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetSecondaryName

`func (o *PostConnectionRequest) GetSecondaryName() string`

GetSecondaryName returns the SecondaryName field if non-nil, zero value otherwise.

### GetSecondaryNameOk

`func (o *PostConnectionRequest) GetSecondaryNameOk() (*string, bool)`

GetSecondaryNameOk returns a tuple with the SecondaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryName

`func (o *PostConnectionRequest) SetSecondaryName(v string)`

SetSecondaryName sets SecondaryName field to given value.

### HasSecondaryName

`func (o *PostConnectionRequest) HasSecondaryName() bool`

HasSecondaryName returns a boolean if a field has been set.

### GetNamedTag

`func (o *PostConnectionRequest) GetNamedTag() string`

GetNamedTag returns the NamedTag field if non-nil, zero value otherwise.

### GetNamedTagOk

`func (o *PostConnectionRequest) GetNamedTagOk() (*string, bool)`

GetNamedTagOk returns a tuple with the NamedTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedTag

`func (o *PostConnectionRequest) SetNamedTag(v string)`

SetNamedTag sets NamedTag field to given value.

### HasNamedTag

`func (o *PostConnectionRequest) HasNamedTag() bool`

HasNamedTag returns a boolean if a field has been set.

### GetSecondaryVirtualDeviceUUID

`func (o *PostConnectionRequest) GetSecondaryVirtualDeviceUUID() string`

GetSecondaryVirtualDeviceUUID returns the SecondaryVirtualDeviceUUID field if non-nil, zero value otherwise.

### GetSecondaryVirtualDeviceUUIDOk

`func (o *PostConnectionRequest) GetSecondaryVirtualDeviceUUIDOk() (*string, bool)`

GetSecondaryVirtualDeviceUUIDOk returns a tuple with the SecondaryVirtualDeviceUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryVirtualDeviceUUID

`func (o *PostConnectionRequest) SetSecondaryVirtualDeviceUUID(v string)`

SetSecondaryVirtualDeviceUUID sets SecondaryVirtualDeviceUUID field to given value.

### HasSecondaryVirtualDeviceUUID

`func (o *PostConnectionRequest) HasSecondaryVirtualDeviceUUID() bool`

HasSecondaryVirtualDeviceUUID returns a boolean if a field has been set.

### GetSecondaryProfileUUID

`func (o *PostConnectionRequest) GetSecondaryProfileUUID() string`

GetSecondaryProfileUUID returns the SecondaryProfileUUID field if non-nil, zero value otherwise.

### GetSecondaryProfileUUIDOk

`func (o *PostConnectionRequest) GetSecondaryProfileUUIDOk() (*string, bool)`

GetSecondaryProfileUUIDOk returns a tuple with the SecondaryProfileUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryProfileUUID

`func (o *PostConnectionRequest) SetSecondaryProfileUUID(v string)`

SetSecondaryProfileUUID sets SecondaryProfileUUID field to given value.

### HasSecondaryProfileUUID

`func (o *PostConnectionRequest) HasSecondaryProfileUUID() bool`

HasSecondaryProfileUUID returns a boolean if a field has been set.

### GetSecondaryAuthorizationKey

`func (o *PostConnectionRequest) GetSecondaryAuthorizationKey() string`

GetSecondaryAuthorizationKey returns the SecondaryAuthorizationKey field if non-nil, zero value otherwise.

### GetSecondaryAuthorizationKeyOk

`func (o *PostConnectionRequest) GetSecondaryAuthorizationKeyOk() (*string, bool)`

GetSecondaryAuthorizationKeyOk returns a tuple with the SecondaryAuthorizationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuthorizationKey

`func (o *PostConnectionRequest) SetSecondaryAuthorizationKey(v string)`

SetSecondaryAuthorizationKey sets SecondaryAuthorizationKey field to given value.

### HasSecondaryAuthorizationKey

`func (o *PostConnectionRequest) HasSecondaryAuthorizationKey() bool`

HasSecondaryAuthorizationKey returns a boolean if a field has been set.

### GetSecondarySellerMetroCode

`func (o *PostConnectionRequest) GetSecondarySellerMetroCode() string`

GetSecondarySellerMetroCode returns the SecondarySellerMetroCode field if non-nil, zero value otherwise.

### GetSecondarySellerMetroCodeOk

`func (o *PostConnectionRequest) GetSecondarySellerMetroCodeOk() (*string, bool)`

GetSecondarySellerMetroCodeOk returns a tuple with the SecondarySellerMetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySellerMetroCode

`func (o *PostConnectionRequest) SetSecondarySellerMetroCode(v string)`

SetSecondarySellerMetroCode sets SecondarySellerMetroCode field to given value.

### HasSecondarySellerMetroCode

`func (o *PostConnectionRequest) HasSecondarySellerMetroCode() bool`

HasSecondarySellerMetroCode returns a boolean if a field has been set.

### GetSecondarySpeed

`func (o *PostConnectionRequest) GetSecondarySpeed() int32`

GetSecondarySpeed returns the SecondarySpeed field if non-nil, zero value otherwise.

### GetSecondarySpeedOk

`func (o *PostConnectionRequest) GetSecondarySpeedOk() (*int32, bool)`

GetSecondarySpeedOk returns a tuple with the SecondarySpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySpeed

`func (o *PostConnectionRequest) SetSecondarySpeed(v int32)`

SetSecondarySpeed sets SecondarySpeed field to given value.

### HasSecondarySpeed

`func (o *PostConnectionRequest) HasSecondarySpeed() bool`

HasSecondarySpeed returns a boolean if a field has been set.

### GetSecondarySpeedUnit

`func (o *PostConnectionRequest) GetSecondarySpeedUnit() string`

GetSecondarySpeedUnit returns the SecondarySpeedUnit field if non-nil, zero value otherwise.

### GetSecondarySpeedUnitOk

`func (o *PostConnectionRequest) GetSecondarySpeedUnitOk() (*string, bool)`

GetSecondarySpeedUnitOk returns a tuple with the SecondarySpeedUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySpeedUnit

`func (o *PostConnectionRequest) SetSecondarySpeedUnit(v string)`

SetSecondarySpeedUnit sets SecondarySpeedUnit field to given value.

### HasSecondarySpeedUnit

`func (o *PostConnectionRequest) HasSecondarySpeedUnit() bool`

HasSecondarySpeedUnit returns a boolean if a field has been set.

### GetSecondaryNotifications

`func (o *PostConnectionRequest) GetSecondaryNotifications() []string`

GetSecondaryNotifications returns the SecondaryNotifications field if non-nil, zero value otherwise.

### GetSecondaryNotificationsOk

`func (o *PostConnectionRequest) GetSecondaryNotificationsOk() (*[]string, bool)`

GetSecondaryNotificationsOk returns a tuple with the SecondaryNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryNotifications

`func (o *PostConnectionRequest) SetSecondaryNotifications(v []string)`

SetSecondaryNotifications sets SecondaryNotifications field to given value.

### HasSecondaryNotifications

`func (o *PostConnectionRequest) HasSecondaryNotifications() bool`

HasSecondaryNotifications returns a boolean if a field has been set.

### GetSecondaryInterfaceId

`func (o *PostConnectionRequest) GetSecondaryInterfaceId() int32`

GetSecondaryInterfaceId returns the SecondaryInterfaceId field if non-nil, zero value otherwise.

### GetSecondaryInterfaceIdOk

`func (o *PostConnectionRequest) GetSecondaryInterfaceIdOk() (*int32, bool)`

GetSecondaryInterfaceIdOk returns a tuple with the SecondaryInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryInterfaceId

`func (o *PostConnectionRequest) SetSecondaryInterfaceId(v int32)`

SetSecondaryInterfaceId sets SecondaryInterfaceId field to given value.

### HasSecondaryInterfaceId

`func (o *PostConnectionRequest) HasSecondaryInterfaceId() bool`

HasSecondaryInterfaceId returns a boolean if a field has been set.

### GetPrimaryZSideVlanCTag

`func (o *PostConnectionRequest) GetPrimaryZSideVlanCTag() int32`

GetPrimaryZSideVlanCTag returns the PrimaryZSideVlanCTag field if non-nil, zero value otherwise.

### GetPrimaryZSideVlanCTagOk

`func (o *PostConnectionRequest) GetPrimaryZSideVlanCTagOk() (*int32, bool)`

GetPrimaryZSideVlanCTagOk returns a tuple with the PrimaryZSideVlanCTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryZSideVlanCTag

`func (o *PostConnectionRequest) SetPrimaryZSideVlanCTag(v int32)`

SetPrimaryZSideVlanCTag sets PrimaryZSideVlanCTag field to given value.

### HasPrimaryZSideVlanCTag

`func (o *PostConnectionRequest) HasPrimaryZSideVlanCTag() bool`

HasPrimaryZSideVlanCTag returns a boolean if a field has been set.

### GetSecondaryZSideVlanCTag

`func (o *PostConnectionRequest) GetSecondaryZSideVlanCTag() int32`

GetSecondaryZSideVlanCTag returns the SecondaryZSideVlanCTag field if non-nil, zero value otherwise.

### GetSecondaryZSideVlanCTagOk

`func (o *PostConnectionRequest) GetSecondaryZSideVlanCTagOk() (*int32, bool)`

GetSecondaryZSideVlanCTagOk returns a tuple with the SecondaryZSideVlanCTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryZSideVlanCTag

`func (o *PostConnectionRequest) SetSecondaryZSideVlanCTag(v int32)`

SetSecondaryZSideVlanCTag sets SecondaryZSideVlanCTag field to given value.

### HasSecondaryZSideVlanCTag

`func (o *PostConnectionRequest) HasSecondaryZSideVlanCTag() bool`

HasSecondaryZSideVlanCTag returns a boolean if a field has been set.

### GetPrimaryZSidePortUUID

`func (o *PostConnectionRequest) GetPrimaryZSidePortUUID() string`

GetPrimaryZSidePortUUID returns the PrimaryZSidePortUUID field if non-nil, zero value otherwise.

### GetPrimaryZSidePortUUIDOk

`func (o *PostConnectionRequest) GetPrimaryZSidePortUUIDOk() (*string, bool)`

GetPrimaryZSidePortUUIDOk returns a tuple with the PrimaryZSidePortUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryZSidePortUUID

`func (o *PostConnectionRequest) SetPrimaryZSidePortUUID(v string)`

SetPrimaryZSidePortUUID sets PrimaryZSidePortUUID field to given value.

### HasPrimaryZSidePortUUID

`func (o *PostConnectionRequest) HasPrimaryZSidePortUUID() bool`

HasPrimaryZSidePortUUID returns a boolean if a field has been set.

### GetPrimaryZSideVlanSTag

`func (o *PostConnectionRequest) GetPrimaryZSideVlanSTag() int32`

GetPrimaryZSideVlanSTag returns the PrimaryZSideVlanSTag field if non-nil, zero value otherwise.

### GetPrimaryZSideVlanSTagOk

`func (o *PostConnectionRequest) GetPrimaryZSideVlanSTagOk() (*int32, bool)`

GetPrimaryZSideVlanSTagOk returns a tuple with the PrimaryZSideVlanSTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryZSideVlanSTag

`func (o *PostConnectionRequest) SetPrimaryZSideVlanSTag(v int32)`

SetPrimaryZSideVlanSTag sets PrimaryZSideVlanSTag field to given value.

### HasPrimaryZSideVlanSTag

`func (o *PostConnectionRequest) HasPrimaryZSideVlanSTag() bool`

HasPrimaryZSideVlanSTag returns a boolean if a field has been set.

### GetSecondaryZSidePortUUID

`func (o *PostConnectionRequest) GetSecondaryZSidePortUUID() string`

GetSecondaryZSidePortUUID returns the SecondaryZSidePortUUID field if non-nil, zero value otherwise.

### GetSecondaryZSidePortUUIDOk

`func (o *PostConnectionRequest) GetSecondaryZSidePortUUIDOk() (*string, bool)`

GetSecondaryZSidePortUUIDOk returns a tuple with the SecondaryZSidePortUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryZSidePortUUID

`func (o *PostConnectionRequest) SetSecondaryZSidePortUUID(v string)`

SetSecondaryZSidePortUUID sets SecondaryZSidePortUUID field to given value.

### HasSecondaryZSidePortUUID

`func (o *PostConnectionRequest) HasSecondaryZSidePortUUID() bool`

HasSecondaryZSidePortUUID returns a boolean if a field has been set.

### GetSecondaryZSideVlanSTag

`func (o *PostConnectionRequest) GetSecondaryZSideVlanSTag() int32`

GetSecondaryZSideVlanSTag returns the SecondaryZSideVlanSTag field if non-nil, zero value otherwise.

### GetSecondaryZSideVlanSTagOk

`func (o *PostConnectionRequest) GetSecondaryZSideVlanSTagOk() (*int32, bool)`

GetSecondaryZSideVlanSTagOk returns a tuple with the SecondaryZSideVlanSTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryZSideVlanSTag

`func (o *PostConnectionRequest) SetSecondaryZSideVlanSTag(v int32)`

SetSecondaryZSideVlanSTag sets SecondaryZSideVlanSTag field to given value.

### HasSecondaryZSideVlanSTag

`func (o *PostConnectionRequest) HasSecondaryZSideVlanSTag() bool`

HasSecondaryZSideVlanSTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


