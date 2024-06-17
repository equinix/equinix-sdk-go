# LineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to [**PlanIdName**](PlanIdName.md) |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**UnitPrice** | Pointer to **float32** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**ItemType** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**ProjectIdName**](ProjectIdName.md) |  | [optional] 
**Adjustments** | Pointer to [**[]LineItemAdjustment**](LineItemAdjustment.md) | Adjustments for the line item | [optional] 

## Methods

### NewLineItem

`func NewLineItem() *LineItem`

NewLineItem instantiates a new LineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemWithDefaults

`func NewLineItemWithDefaults() *LineItem`

NewLineItemWithDefaults instantiates a new LineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *LineItem) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LineItem) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LineItem) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *LineItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *LineItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LineItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LineItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *LineItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *LineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetails

`func (o *LineItem) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *LineItem) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *LineItem) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *LineItem) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetPlan

`func (o *LineItem) GetPlan() PlanIdName`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *LineItem) GetPlanOk() (*PlanIdName, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *LineItem) SetPlan(v PlanIdName)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *LineItem) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetUnit

`func (o *LineItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *LineItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *LineItem) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *LineItem) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitPrice

`func (o *LineItem) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *LineItem) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *LineItem) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *LineItem) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetHostname

`func (o *LineItem) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LineItem) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LineItem) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *LineItem) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetLocation

`func (o *LineItem) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LineItem) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LineItem) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *LineItem) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetItemType

`func (o *LineItem) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *LineItem) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *LineItem) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *LineItem) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetServiceId

`func (o *LineItem) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *LineItem) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *LineItem) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *LineItem) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetStartDate

`func (o *LineItem) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LineItem) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LineItem) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *LineItem) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *LineItem) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *LineItem) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *LineItem) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *LineItem) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetProjectId

`func (o *LineItem) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *LineItem) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *LineItem) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *LineItem) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetPlanId

`func (o *LineItem) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *LineItem) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *LineItem) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *LineItem) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProject

`func (o *LineItem) GetProject() ProjectIdName`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *LineItem) GetProjectOk() (*ProjectIdName, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *LineItem) SetProject(v ProjectIdName)`

SetProject sets Project field to given value.

### HasProject

`func (o *LineItem) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAdjustments

`func (o *LineItem) GetAdjustments() []LineItemAdjustment`

GetAdjustments returns the Adjustments field if non-nil, zero value otherwise.

### GetAdjustmentsOk

`func (o *LineItem) GetAdjustmentsOk() (*[]LineItemAdjustment, bool)`

GetAdjustmentsOk returns a tuple with the Adjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustments

`func (o *LineItem) SetAdjustments(v []LineItemAdjustment)`

SetAdjustments sets Adjustments field to given value.

### HasAdjustments

`func (o *LineItem) HasAdjustments() bool`

HasAdjustments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


