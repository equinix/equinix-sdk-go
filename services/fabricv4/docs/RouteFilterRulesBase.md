# RouteFilterRulesBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Customer-provided Route Filter Rule description | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**PrefixMatch** | Pointer to **string** |  | [optional] [default to "orlonger"]

## Methods

### NewRouteFilterRulesBase

`func NewRouteFilterRulesBase() *RouteFilterRulesBase`

NewRouteFilterRulesBase instantiates a new RouteFilterRulesBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFilterRulesBaseWithDefaults

`func NewRouteFilterRulesBaseWithDefaults() *RouteFilterRulesBase`

NewRouteFilterRulesBaseWithDefaults instantiates a new RouteFilterRulesBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RouteFilterRulesBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteFilterRulesBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteFilterRulesBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RouteFilterRulesBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RouteFilterRulesBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RouteFilterRulesBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RouteFilterRulesBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RouteFilterRulesBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrefix

`func (o *RouteFilterRulesBase) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RouteFilterRulesBase) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RouteFilterRulesBase) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *RouteFilterRulesBase) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPrefixMatch

`func (o *RouteFilterRulesBase) GetPrefixMatch() string`

GetPrefixMatch returns the PrefixMatch field if non-nil, zero value otherwise.

### GetPrefixMatchOk

`func (o *RouteFilterRulesBase) GetPrefixMatchOk() (*string, bool)`

GetPrefixMatchOk returns a tuple with the PrefixMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixMatch

`func (o *RouteFilterRulesBase) SetPrefixMatch(v string)`

SetPrefixMatch sets PrefixMatch field to given value.

### HasPrefixMatch

`func (o *RouteFilterRulesBase) HasPrefixMatch() bool`

HasPrefixMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


