# RouteFilterRulesChangeOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of  route filter rule | 
**Description** | Pointer to **string** | cust provided description | [optional] 
**Type** | [**RouteFilterRulesChangeOperationType**](RouteFilterRulesChangeOperationType.md) |  | 
**Prefix** | **string** | given prefix (does not change) | 
**Action** | [**RouteFilterRulesChangeOperationAction**](RouteFilterRulesChangeOperationAction.md) |  | 

## Methods

### NewRouteFilterRulesChangeOperation

`func NewRouteFilterRulesChangeOperation(name string, type_ RouteFilterRulesChangeOperationType, prefix string, action RouteFilterRulesChangeOperationAction, ) *RouteFilterRulesChangeOperation`

NewRouteFilterRulesChangeOperation instantiates a new RouteFilterRulesChangeOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFilterRulesChangeOperationWithDefaults

`func NewRouteFilterRulesChangeOperationWithDefaults() *RouteFilterRulesChangeOperation`

NewRouteFilterRulesChangeOperationWithDefaults instantiates a new RouteFilterRulesChangeOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RouteFilterRulesChangeOperation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteFilterRulesChangeOperation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteFilterRulesChangeOperation) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RouteFilterRulesChangeOperation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RouteFilterRulesChangeOperation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RouteFilterRulesChangeOperation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RouteFilterRulesChangeOperation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *RouteFilterRulesChangeOperation) GetType() RouteFilterRulesChangeOperationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteFilterRulesChangeOperation) GetTypeOk() (*RouteFilterRulesChangeOperationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteFilterRulesChangeOperation) SetType(v RouteFilterRulesChangeOperationType)`

SetType sets Type field to given value.


### GetPrefix

`func (o *RouteFilterRulesChangeOperation) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RouteFilterRulesChangeOperation) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RouteFilterRulesChangeOperation) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetAction

`func (o *RouteFilterRulesChangeOperation) GetAction() RouteFilterRulesChangeOperationAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RouteFilterRulesChangeOperation) GetActionOk() (*RouteFilterRulesChangeOperationAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RouteFilterRulesChangeOperation) SetAction(v RouteFilterRulesChangeOperationAction)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


