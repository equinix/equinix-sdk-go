# GatewaySortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**GatewaySortDirection**](GatewaySortDirection.md) |  | [optional] [default to GATEWAYSORTDIRECTION_DESC]
**Property** | Pointer to **string** | Possible field names to use on sorting:   * &#x60;/name&#x60; - Gateway name   * &#x60;/uuid&#x60; - Gateway uuid   * &#x60;/state&#x60; - Gateway status   * &#x60;/changeLog/createdDateTime&#x60; - Date and time when change flow starts   * &#x60;/changeLog/updatedDateTime&#x60; - Date and time when change object is updated  | [optional] [default to "/changeLog/updatedDateTime"]

## Methods

### NewGatewaySortCriteria

`func NewGatewaySortCriteria() *GatewaySortCriteria`

NewGatewaySortCriteria instantiates a new GatewaySortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewaySortCriteriaWithDefaults

`func NewGatewaySortCriteriaWithDefaults() *GatewaySortCriteria`

NewGatewaySortCriteriaWithDefaults instantiates a new GatewaySortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *GatewaySortCriteria) GetDirection() GatewaySortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GatewaySortCriteria) GetDirectionOk() (*GatewaySortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GatewaySortCriteria) SetDirection(v GatewaySortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GatewaySortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *GatewaySortCriteria) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *GatewaySortCriteria) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *GatewaySortCriteria) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *GatewaySortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


