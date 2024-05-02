# MetroInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metro** | **string** | Metro code or ID of where the device should be provisioned in, or it can be instructed to create the device in the best available metro with &#x60;{ \&quot;metro\&quot;: \&quot;any\&quot; }&#x60;. The special metro value of any means anywhere, any metro. When any is chosen in the request, the metro location is picked per our scheduling algorithms that favor the following factors: hardware reservation location (if requesting reserved hardware), ip reservations, spot instances, etc. The any keyword *does not* optimize for cost, this means that usage costs (instance, transfer, other features dependent on location) will vary. Please check metro value in response to see where the device was created. Either metro or facility must be provided. | 

## Methods

### NewMetroInput

`func NewMetroInput(metro string, ) *MetroInput`

NewMetroInput instantiates a new MetroInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetroInputWithDefaults

`func NewMetroInputWithDefaults() *MetroInput`

NewMetroInputWithDefaults instantiates a new MetroInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetro

`func (o *MetroInput) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *MetroInput) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *MetroInput) SetMetro(v string)`

SetMetro sets Metro field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


