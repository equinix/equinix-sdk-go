# FacilityInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facility** | **[]string** | The datacenter where the device should be created.  Either metro or facility must be provided.  The API will accept either a single facility &#x60;{ \&quot;facility\&quot;: \&quot;f1\&quot; }&#x60;, or it can be instructed to create the device in the best available datacenter &#x60;{ \&quot;facility\&quot;: \&quot;any\&quot; }&#x60;.  Additionally it is possible to set a prioritized location selection. For example &#x60;{ \&quot;facility\&quot;: [\&quot;f3\&quot;, \&quot;f2\&quot;, \&quot;any\&quot;] }&#x60; can be used to prioritize &#x60;f3&#x60; and then &#x60;f2&#x60; before accepting &#x60;any&#x60; facility. If none of the facilities provided have availability for the requested device the request will fail. | 

## Methods

### NewFacilityInput

`func NewFacilityInput(facility []string, ) *FacilityInput`

NewFacilityInput instantiates a new FacilityInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacilityInputWithDefaults

`func NewFacilityInputWithDefaults() *FacilityInput`

NewFacilityInputWithDefaults instantiates a new FacilityInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacility

`func (o *FacilityInput) GetFacility() []string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *FacilityInput) GetFacilityOk() (*[]string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *FacilityInput) SetFacility(v []string)`

SetFacility sets Facility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


