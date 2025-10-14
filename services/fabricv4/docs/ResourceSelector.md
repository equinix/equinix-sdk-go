# ResourceSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Include** | Pointer to **[]string** | ### Supported metric names to use on filters with property /subject:   * &#x60;/fabric/v4/ports/&lt;uuid&gt;&#x60; - port metrics   * &#x60;/fabric/v4/connections/&lt;uuid&gt;&#x60; - connection metrics   * &#x60;/fabric/v4/metros/&lt;metroCode&gt;&#x60; - metro latency metrics  | [optional] 

## Methods

### NewResourceSelector

`func NewResourceSelector() *ResourceSelector`

NewResourceSelector instantiates a new ResourceSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSelectorWithDefaults

`func NewResourceSelectorWithDefaults() *ResourceSelector`

NewResourceSelectorWithDefaults instantiates a new ResourceSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInclude

`func (o *ResourceSelector) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *ResourceSelector) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *ResourceSelector) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *ResourceSelector) HasInclude() bool`

HasInclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


