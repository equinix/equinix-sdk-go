# GETConnectionByUuidResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerOrganizationName** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**VlanSTag** | Pointer to **int32** |  | [optional] 
**PortUUID** | Pointer to **string** |  | [optional] 
**PortName** | Pointer to **string** |  | [optional] 
**AsideEncapsulation** | Pointer to **string** |  | [optional] 
**MetroCode** | Pointer to **string** |  | [optional] 
**MetroDescription** | Pointer to **string** |  | [optional] 
**ProviderStatus** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**BillingTier** | Pointer to **string** |  | [optional] 
**AuthorizationKey** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **int32** |  | [optional] 
**SpeedUnit** | Pointer to **string** |  | [optional] 
**RedundancyType** | Pointer to **string** |  | [optional] 
**RedundancyGroup** | Pointer to **string** |  | [optional] 
**SellerMetroCode** | Pointer to **string** |  | [optional] 
**SellerMetroDescription** | Pointer to **string** |  | [optional] 
**SellerServiceName** | Pointer to **string** |  | [optional] 
**SellerServiceUUID** | Pointer to **string** |  | [optional] 
**SellerOrganizationName** | Pointer to **string** |  | [optional] 
**Notifications** | Pointer to **[]string** |  | [optional] 
**PurchaseOrderNumber** | Pointer to **string** |  | [optional] 
**NamedTag** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedByFullName** | Pointer to **string** |  | [optional] 
**CreatedByEmail** | Pointer to **string** |  | [optional] 
**LastUpdatedBy** | Pointer to **string** |  | [optional] 
**LastUpdatedDate** | Pointer to **string** |  | [optional] 
**LastUpdatedByFullName** | Pointer to **string** |  | [optional] 
**LastUpdatedByEmail** | Pointer to **string** |  | [optional] 
**ZSidePortName** | Pointer to **string** |  | [optional] 
**ZSidePortUUID** | Pointer to **string** |  | [optional] 
**ZSideVlanCTag** | Pointer to **int32** |  | [optional] 
**ZSideVlanSTag** | Pointer to **int32** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**Self** | Pointer to **bool** |  | [optional] 
**RedundantUuid** | Pointer to **string** |  | [optional] 

## Methods

### NewGETConnectionByUuidResponse

`func NewGETConnectionByUuidResponse() *GETConnectionByUuidResponse`

NewGETConnectionByUuidResponse instantiates a new GETConnectionByUuidResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETConnectionByUuidResponseWithDefaults

`func NewGETConnectionByUuidResponseWithDefaults() *GETConnectionByUuidResponse`

NewGETConnectionByUuidResponseWithDefaults instantiates a new GETConnectionByUuidResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerOrganizationName

`func (o *GETConnectionByUuidResponse) GetBuyerOrganizationName() string`

GetBuyerOrganizationName returns the BuyerOrganizationName field if non-nil, zero value otherwise.

### GetBuyerOrganizationNameOk

`func (o *GETConnectionByUuidResponse) GetBuyerOrganizationNameOk() (*string, bool)`

GetBuyerOrganizationNameOk returns a tuple with the BuyerOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerOrganizationName

`func (o *GETConnectionByUuidResponse) SetBuyerOrganizationName(v string)`

SetBuyerOrganizationName sets BuyerOrganizationName field to given value.

### HasBuyerOrganizationName

`func (o *GETConnectionByUuidResponse) HasBuyerOrganizationName() bool`

HasBuyerOrganizationName returns a boolean if a field has been set.

### GetUuid

`func (o *GETConnectionByUuidResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GETConnectionByUuidResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GETConnectionByUuidResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GETConnectionByUuidResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *GETConnectionByUuidResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETConnectionByUuidResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETConnectionByUuidResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GETConnectionByUuidResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVlanSTag

`func (o *GETConnectionByUuidResponse) GetVlanSTag() int32`

GetVlanSTag returns the VlanSTag field if non-nil, zero value otherwise.

### GetVlanSTagOk

`func (o *GETConnectionByUuidResponse) GetVlanSTagOk() (*int32, bool)`

GetVlanSTagOk returns a tuple with the VlanSTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSTag

`func (o *GETConnectionByUuidResponse) SetVlanSTag(v int32)`

SetVlanSTag sets VlanSTag field to given value.

### HasVlanSTag

`func (o *GETConnectionByUuidResponse) HasVlanSTag() bool`

HasVlanSTag returns a boolean if a field has been set.

### GetPortUUID

`func (o *GETConnectionByUuidResponse) GetPortUUID() string`

GetPortUUID returns the PortUUID field if non-nil, zero value otherwise.

### GetPortUUIDOk

`func (o *GETConnectionByUuidResponse) GetPortUUIDOk() (*string, bool)`

GetPortUUIDOk returns a tuple with the PortUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUUID

`func (o *GETConnectionByUuidResponse) SetPortUUID(v string)`

SetPortUUID sets PortUUID field to given value.

### HasPortUUID

`func (o *GETConnectionByUuidResponse) HasPortUUID() bool`

HasPortUUID returns a boolean if a field has been set.

### GetPortName

`func (o *GETConnectionByUuidResponse) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *GETConnectionByUuidResponse) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *GETConnectionByUuidResponse) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *GETConnectionByUuidResponse) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetAsideEncapsulation

`func (o *GETConnectionByUuidResponse) GetAsideEncapsulation() string`

GetAsideEncapsulation returns the AsideEncapsulation field if non-nil, zero value otherwise.

### GetAsideEncapsulationOk

`func (o *GETConnectionByUuidResponse) GetAsideEncapsulationOk() (*string, bool)`

GetAsideEncapsulationOk returns a tuple with the AsideEncapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsideEncapsulation

`func (o *GETConnectionByUuidResponse) SetAsideEncapsulation(v string)`

SetAsideEncapsulation sets AsideEncapsulation field to given value.

### HasAsideEncapsulation

`func (o *GETConnectionByUuidResponse) HasAsideEncapsulation() bool`

HasAsideEncapsulation returns a boolean if a field has been set.

### GetMetroCode

`func (o *GETConnectionByUuidResponse) GetMetroCode() string`

GetMetroCode returns the MetroCode field if non-nil, zero value otherwise.

### GetMetroCodeOk

`func (o *GETConnectionByUuidResponse) GetMetroCodeOk() (*string, bool)`

GetMetroCodeOk returns a tuple with the MetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCode

`func (o *GETConnectionByUuidResponse) SetMetroCode(v string)`

SetMetroCode sets MetroCode field to given value.

### HasMetroCode

`func (o *GETConnectionByUuidResponse) HasMetroCode() bool`

HasMetroCode returns a boolean if a field has been set.

### GetMetroDescription

`func (o *GETConnectionByUuidResponse) GetMetroDescription() string`

GetMetroDescription returns the MetroDescription field if non-nil, zero value otherwise.

### GetMetroDescriptionOk

`func (o *GETConnectionByUuidResponse) GetMetroDescriptionOk() (*string, bool)`

GetMetroDescriptionOk returns a tuple with the MetroDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroDescription

`func (o *GETConnectionByUuidResponse) SetMetroDescription(v string)`

SetMetroDescription sets MetroDescription field to given value.

### HasMetroDescription

`func (o *GETConnectionByUuidResponse) HasMetroDescription() bool`

HasMetroDescription returns a boolean if a field has been set.

### GetProviderStatus

`func (o *GETConnectionByUuidResponse) GetProviderStatus() string`

GetProviderStatus returns the ProviderStatus field if non-nil, zero value otherwise.

### GetProviderStatusOk

`func (o *GETConnectionByUuidResponse) GetProviderStatusOk() (*string, bool)`

GetProviderStatusOk returns a tuple with the ProviderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderStatus

`func (o *GETConnectionByUuidResponse) SetProviderStatus(v string)`

SetProviderStatus sets ProviderStatus field to given value.

### HasProviderStatus

`func (o *GETConnectionByUuidResponse) HasProviderStatus() bool`

HasProviderStatus returns a boolean if a field has been set.

### GetStatus

`func (o *GETConnectionByUuidResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETConnectionByUuidResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETConnectionByUuidResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETConnectionByUuidResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBillingTier

`func (o *GETConnectionByUuidResponse) GetBillingTier() string`

GetBillingTier returns the BillingTier field if non-nil, zero value otherwise.

### GetBillingTierOk

`func (o *GETConnectionByUuidResponse) GetBillingTierOk() (*string, bool)`

GetBillingTierOk returns a tuple with the BillingTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTier

`func (o *GETConnectionByUuidResponse) SetBillingTier(v string)`

SetBillingTier sets BillingTier field to given value.

### HasBillingTier

`func (o *GETConnectionByUuidResponse) HasBillingTier() bool`

HasBillingTier returns a boolean if a field has been set.

### GetAuthorizationKey

`func (o *GETConnectionByUuidResponse) GetAuthorizationKey() string`

GetAuthorizationKey returns the AuthorizationKey field if non-nil, zero value otherwise.

### GetAuthorizationKeyOk

`func (o *GETConnectionByUuidResponse) GetAuthorizationKeyOk() (*string, bool)`

GetAuthorizationKeyOk returns a tuple with the AuthorizationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationKey

`func (o *GETConnectionByUuidResponse) SetAuthorizationKey(v string)`

SetAuthorizationKey sets AuthorizationKey field to given value.

### HasAuthorizationKey

`func (o *GETConnectionByUuidResponse) HasAuthorizationKey() bool`

HasAuthorizationKey returns a boolean if a field has been set.

### GetSpeed

`func (o *GETConnectionByUuidResponse) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GETConnectionByUuidResponse) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GETConnectionByUuidResponse) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *GETConnectionByUuidResponse) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetSpeedUnit

`func (o *GETConnectionByUuidResponse) GetSpeedUnit() string`

GetSpeedUnit returns the SpeedUnit field if non-nil, zero value otherwise.

### GetSpeedUnitOk

`func (o *GETConnectionByUuidResponse) GetSpeedUnitOk() (*string, bool)`

GetSpeedUnitOk returns a tuple with the SpeedUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedUnit

`func (o *GETConnectionByUuidResponse) SetSpeedUnit(v string)`

SetSpeedUnit sets SpeedUnit field to given value.

### HasSpeedUnit

`func (o *GETConnectionByUuidResponse) HasSpeedUnit() bool`

HasSpeedUnit returns a boolean if a field has been set.

### GetRedundancyType

`func (o *GETConnectionByUuidResponse) GetRedundancyType() string`

GetRedundancyType returns the RedundancyType field if non-nil, zero value otherwise.

### GetRedundancyTypeOk

`func (o *GETConnectionByUuidResponse) GetRedundancyTypeOk() (*string, bool)`

GetRedundancyTypeOk returns a tuple with the RedundancyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyType

`func (o *GETConnectionByUuidResponse) SetRedundancyType(v string)`

SetRedundancyType sets RedundancyType field to given value.

### HasRedundancyType

`func (o *GETConnectionByUuidResponse) HasRedundancyType() bool`

HasRedundancyType returns a boolean if a field has been set.

### GetRedundancyGroup

`func (o *GETConnectionByUuidResponse) GetRedundancyGroup() string`

GetRedundancyGroup returns the RedundancyGroup field if non-nil, zero value otherwise.

### GetRedundancyGroupOk

`func (o *GETConnectionByUuidResponse) GetRedundancyGroupOk() (*string, bool)`

GetRedundancyGroupOk returns a tuple with the RedundancyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyGroup

`func (o *GETConnectionByUuidResponse) SetRedundancyGroup(v string)`

SetRedundancyGroup sets RedundancyGroup field to given value.

### HasRedundancyGroup

`func (o *GETConnectionByUuidResponse) HasRedundancyGroup() bool`

HasRedundancyGroup returns a boolean if a field has been set.

### GetSellerMetroCode

`func (o *GETConnectionByUuidResponse) GetSellerMetroCode() string`

GetSellerMetroCode returns the SellerMetroCode field if non-nil, zero value otherwise.

### GetSellerMetroCodeOk

`func (o *GETConnectionByUuidResponse) GetSellerMetroCodeOk() (*string, bool)`

GetSellerMetroCodeOk returns a tuple with the SellerMetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerMetroCode

`func (o *GETConnectionByUuidResponse) SetSellerMetroCode(v string)`

SetSellerMetroCode sets SellerMetroCode field to given value.

### HasSellerMetroCode

`func (o *GETConnectionByUuidResponse) HasSellerMetroCode() bool`

HasSellerMetroCode returns a boolean if a field has been set.

### GetSellerMetroDescription

`func (o *GETConnectionByUuidResponse) GetSellerMetroDescription() string`

GetSellerMetroDescription returns the SellerMetroDescription field if non-nil, zero value otherwise.

### GetSellerMetroDescriptionOk

`func (o *GETConnectionByUuidResponse) GetSellerMetroDescriptionOk() (*string, bool)`

GetSellerMetroDescriptionOk returns a tuple with the SellerMetroDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerMetroDescription

`func (o *GETConnectionByUuidResponse) SetSellerMetroDescription(v string)`

SetSellerMetroDescription sets SellerMetroDescription field to given value.

### HasSellerMetroDescription

`func (o *GETConnectionByUuidResponse) HasSellerMetroDescription() bool`

HasSellerMetroDescription returns a boolean if a field has been set.

### GetSellerServiceName

`func (o *GETConnectionByUuidResponse) GetSellerServiceName() string`

GetSellerServiceName returns the SellerServiceName field if non-nil, zero value otherwise.

### GetSellerServiceNameOk

`func (o *GETConnectionByUuidResponse) GetSellerServiceNameOk() (*string, bool)`

GetSellerServiceNameOk returns a tuple with the SellerServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerServiceName

`func (o *GETConnectionByUuidResponse) SetSellerServiceName(v string)`

SetSellerServiceName sets SellerServiceName field to given value.

### HasSellerServiceName

`func (o *GETConnectionByUuidResponse) HasSellerServiceName() bool`

HasSellerServiceName returns a boolean if a field has been set.

### GetSellerServiceUUID

`func (o *GETConnectionByUuidResponse) GetSellerServiceUUID() string`

GetSellerServiceUUID returns the SellerServiceUUID field if non-nil, zero value otherwise.

### GetSellerServiceUUIDOk

`func (o *GETConnectionByUuidResponse) GetSellerServiceUUIDOk() (*string, bool)`

GetSellerServiceUUIDOk returns a tuple with the SellerServiceUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerServiceUUID

`func (o *GETConnectionByUuidResponse) SetSellerServiceUUID(v string)`

SetSellerServiceUUID sets SellerServiceUUID field to given value.

### HasSellerServiceUUID

`func (o *GETConnectionByUuidResponse) HasSellerServiceUUID() bool`

HasSellerServiceUUID returns a boolean if a field has been set.

### GetSellerOrganizationName

`func (o *GETConnectionByUuidResponse) GetSellerOrganizationName() string`

GetSellerOrganizationName returns the SellerOrganizationName field if non-nil, zero value otherwise.

### GetSellerOrganizationNameOk

`func (o *GETConnectionByUuidResponse) GetSellerOrganizationNameOk() (*string, bool)`

GetSellerOrganizationNameOk returns a tuple with the SellerOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerOrganizationName

`func (o *GETConnectionByUuidResponse) SetSellerOrganizationName(v string)`

SetSellerOrganizationName sets SellerOrganizationName field to given value.

### HasSellerOrganizationName

`func (o *GETConnectionByUuidResponse) HasSellerOrganizationName() bool`

HasSellerOrganizationName returns a boolean if a field has been set.

### GetNotifications

`func (o *GETConnectionByUuidResponse) GetNotifications() []string`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *GETConnectionByUuidResponse) GetNotificationsOk() (*[]string, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *GETConnectionByUuidResponse) SetNotifications(v []string)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *GETConnectionByUuidResponse) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *GETConnectionByUuidResponse) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *GETConnectionByUuidResponse) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *GETConnectionByUuidResponse) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *GETConnectionByUuidResponse) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetNamedTag

`func (o *GETConnectionByUuidResponse) GetNamedTag() string`

GetNamedTag returns the NamedTag field if non-nil, zero value otherwise.

### GetNamedTagOk

`func (o *GETConnectionByUuidResponse) GetNamedTagOk() (*string, bool)`

GetNamedTagOk returns a tuple with the NamedTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedTag

`func (o *GETConnectionByUuidResponse) SetNamedTag(v string)`

SetNamedTag sets NamedTag field to given value.

### HasNamedTag

`func (o *GETConnectionByUuidResponse) HasNamedTag() bool`

HasNamedTag returns a boolean if a field has been set.

### GetCreatedDate

`func (o *GETConnectionByUuidResponse) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *GETConnectionByUuidResponse) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *GETConnectionByUuidResponse) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *GETConnectionByUuidResponse) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GETConnectionByUuidResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GETConnectionByUuidResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GETConnectionByUuidResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GETConnectionByUuidResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByFullName

`func (o *GETConnectionByUuidResponse) GetCreatedByFullName() string`

GetCreatedByFullName returns the CreatedByFullName field if non-nil, zero value otherwise.

### GetCreatedByFullNameOk

`func (o *GETConnectionByUuidResponse) GetCreatedByFullNameOk() (*string, bool)`

GetCreatedByFullNameOk returns a tuple with the CreatedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByFullName

`func (o *GETConnectionByUuidResponse) SetCreatedByFullName(v string)`

SetCreatedByFullName sets CreatedByFullName field to given value.

### HasCreatedByFullName

`func (o *GETConnectionByUuidResponse) HasCreatedByFullName() bool`

HasCreatedByFullName returns a boolean if a field has been set.

### GetCreatedByEmail

`func (o *GETConnectionByUuidResponse) GetCreatedByEmail() string`

GetCreatedByEmail returns the CreatedByEmail field if non-nil, zero value otherwise.

### GetCreatedByEmailOk

`func (o *GETConnectionByUuidResponse) GetCreatedByEmailOk() (*string, bool)`

GetCreatedByEmailOk returns a tuple with the CreatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByEmail

`func (o *GETConnectionByUuidResponse) SetCreatedByEmail(v string)`

SetCreatedByEmail sets CreatedByEmail field to given value.

### HasCreatedByEmail

`func (o *GETConnectionByUuidResponse) HasCreatedByEmail() bool`

HasCreatedByEmail returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *GETConnectionByUuidResponse) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *GETConnectionByUuidResponse) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *GETConnectionByUuidResponse) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *GETConnectionByUuidResponse) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetLastUpdatedDate

`func (o *GETConnectionByUuidResponse) GetLastUpdatedDate() string`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *GETConnectionByUuidResponse) GetLastUpdatedDateOk() (*string, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *GETConnectionByUuidResponse) SetLastUpdatedDate(v string)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *GETConnectionByUuidResponse) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.

### GetLastUpdatedByFullName

`func (o *GETConnectionByUuidResponse) GetLastUpdatedByFullName() string`

GetLastUpdatedByFullName returns the LastUpdatedByFullName field if non-nil, zero value otherwise.

### GetLastUpdatedByFullNameOk

`func (o *GETConnectionByUuidResponse) GetLastUpdatedByFullNameOk() (*string, bool)`

GetLastUpdatedByFullNameOk returns a tuple with the LastUpdatedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedByFullName

`func (o *GETConnectionByUuidResponse) SetLastUpdatedByFullName(v string)`

SetLastUpdatedByFullName sets LastUpdatedByFullName field to given value.

### HasLastUpdatedByFullName

`func (o *GETConnectionByUuidResponse) HasLastUpdatedByFullName() bool`

HasLastUpdatedByFullName returns a boolean if a field has been set.

### GetLastUpdatedByEmail

`func (o *GETConnectionByUuidResponse) GetLastUpdatedByEmail() string`

GetLastUpdatedByEmail returns the LastUpdatedByEmail field if non-nil, zero value otherwise.

### GetLastUpdatedByEmailOk

`func (o *GETConnectionByUuidResponse) GetLastUpdatedByEmailOk() (*string, bool)`

GetLastUpdatedByEmailOk returns a tuple with the LastUpdatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedByEmail

`func (o *GETConnectionByUuidResponse) SetLastUpdatedByEmail(v string)`

SetLastUpdatedByEmail sets LastUpdatedByEmail field to given value.

### HasLastUpdatedByEmail

`func (o *GETConnectionByUuidResponse) HasLastUpdatedByEmail() bool`

HasLastUpdatedByEmail returns a boolean if a field has been set.

### GetZSidePortName

`func (o *GETConnectionByUuidResponse) GetZSidePortName() string`

GetZSidePortName returns the ZSidePortName field if non-nil, zero value otherwise.

### GetZSidePortNameOk

`func (o *GETConnectionByUuidResponse) GetZSidePortNameOk() (*string, bool)`

GetZSidePortNameOk returns a tuple with the ZSidePortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSidePortName

`func (o *GETConnectionByUuidResponse) SetZSidePortName(v string)`

SetZSidePortName sets ZSidePortName field to given value.

### HasZSidePortName

`func (o *GETConnectionByUuidResponse) HasZSidePortName() bool`

HasZSidePortName returns a boolean if a field has been set.

### GetZSidePortUUID

`func (o *GETConnectionByUuidResponse) GetZSidePortUUID() string`

GetZSidePortUUID returns the ZSidePortUUID field if non-nil, zero value otherwise.

### GetZSidePortUUIDOk

`func (o *GETConnectionByUuidResponse) GetZSidePortUUIDOk() (*string, bool)`

GetZSidePortUUIDOk returns a tuple with the ZSidePortUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSidePortUUID

`func (o *GETConnectionByUuidResponse) SetZSidePortUUID(v string)`

SetZSidePortUUID sets ZSidePortUUID field to given value.

### HasZSidePortUUID

`func (o *GETConnectionByUuidResponse) HasZSidePortUUID() bool`

HasZSidePortUUID returns a boolean if a field has been set.

### GetZSideVlanCTag

`func (o *GETConnectionByUuidResponse) GetZSideVlanCTag() int32`

GetZSideVlanCTag returns the ZSideVlanCTag field if non-nil, zero value otherwise.

### GetZSideVlanCTagOk

`func (o *GETConnectionByUuidResponse) GetZSideVlanCTagOk() (*int32, bool)`

GetZSideVlanCTagOk returns a tuple with the ZSideVlanCTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSideVlanCTag

`func (o *GETConnectionByUuidResponse) SetZSideVlanCTag(v int32)`

SetZSideVlanCTag sets ZSideVlanCTag field to given value.

### HasZSideVlanCTag

`func (o *GETConnectionByUuidResponse) HasZSideVlanCTag() bool`

HasZSideVlanCTag returns a boolean if a field has been set.

### GetZSideVlanSTag

`func (o *GETConnectionByUuidResponse) GetZSideVlanSTag() int32`

GetZSideVlanSTag returns the ZSideVlanSTag field if non-nil, zero value otherwise.

### GetZSideVlanSTagOk

`func (o *GETConnectionByUuidResponse) GetZSideVlanSTagOk() (*int32, bool)`

GetZSideVlanSTagOk returns a tuple with the ZSideVlanSTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSideVlanSTag

`func (o *GETConnectionByUuidResponse) SetZSideVlanSTag(v int32)`

SetZSideVlanSTag sets ZSideVlanSTag field to given value.

### HasZSideVlanSTag

`func (o *GETConnectionByUuidResponse) HasZSideVlanSTag() bool`

HasZSideVlanSTag returns a boolean if a field has been set.

### GetRemote

`func (o *GETConnectionByUuidResponse) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *GETConnectionByUuidResponse) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *GETConnectionByUuidResponse) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *GETConnectionByUuidResponse) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetPrivate

`func (o *GETConnectionByUuidResponse) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *GETConnectionByUuidResponse) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *GETConnectionByUuidResponse) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *GETConnectionByUuidResponse) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetSelf

`func (o *GETConnectionByUuidResponse) GetSelf() bool`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *GETConnectionByUuidResponse) GetSelfOk() (*bool, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *GETConnectionByUuidResponse) SetSelf(v bool)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *GETConnectionByUuidResponse) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetRedundantUuid

`func (o *GETConnectionByUuidResponse) GetRedundantUuid() string`

GetRedundantUuid returns the RedundantUuid field if non-nil, zero value otherwise.

### GetRedundantUuidOk

`func (o *GETConnectionByUuidResponse) GetRedundantUuidOk() (*string, bool)`

GetRedundantUuidOk returns a tuple with the RedundantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantUuid

`func (o *GETConnectionByUuidResponse) SetRedundantUuid(v string)`

SetRedundantUuid sets RedundantUuid field to given value.

### HasRedundantUuid

`func (o *GETConnectionByUuidResponse) HasRedundantUuid() bool`

HasRedundantUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


