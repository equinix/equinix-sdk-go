# PlanSpecsDrivesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** | Values may include &#39;HDD&#39;, &#39;SSD&#39;, &#39;NVME&#39; | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** | Values may include &#39;boot&#39;, &#39;cache&#39;, &#39;storage&#39; | [optional] 

## Methods

### NewPlanSpecsDrivesInner

`func NewPlanSpecsDrivesInner() *PlanSpecsDrivesInner`

NewPlanSpecsDrivesInner instantiates a new PlanSpecsDrivesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanSpecsDrivesInnerWithDefaults

`func NewPlanSpecsDrivesInnerWithDefaults() *PlanSpecsDrivesInner`

NewPlanSpecsDrivesInnerWithDefaults instantiates a new PlanSpecsDrivesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PlanSpecsDrivesInner) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PlanSpecsDrivesInner) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PlanSpecsDrivesInner) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PlanSpecsDrivesInner) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetType

`func (o *PlanSpecsDrivesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlanSpecsDrivesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlanSpecsDrivesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlanSpecsDrivesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *PlanSpecsDrivesInner) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PlanSpecsDrivesInner) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PlanSpecsDrivesInner) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *PlanSpecsDrivesInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetCategory

`func (o *PlanSpecsDrivesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PlanSpecsDrivesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PlanSpecsDrivesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PlanSpecsDrivesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


