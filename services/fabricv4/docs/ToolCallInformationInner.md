# ToolCallInformationInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of tools called | [optional] 
**Input** | Pointer to **string** | Content of the tool request | [optional] 
**Response** | Pointer to **string** | Content of the tool response | [optional] 

## Methods

### NewToolCallInformationInner

`func NewToolCallInformationInner() *ToolCallInformationInner`

NewToolCallInformationInner instantiates a new ToolCallInformationInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolCallInformationInnerWithDefaults

`func NewToolCallInformationInnerWithDefaults() *ToolCallInformationInner`

NewToolCallInformationInnerWithDefaults instantiates a new ToolCallInformationInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ToolCallInformationInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolCallInformationInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolCallInformationInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ToolCallInformationInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInput

`func (o *ToolCallInformationInner) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ToolCallInformationInner) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ToolCallInformationInner) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *ToolCallInformationInner) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetResponse

`func (o *ToolCallInformationInner) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ToolCallInformationInner) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ToolCallInformationInner) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ToolCallInformationInner) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


