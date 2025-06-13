# CloudRouterCommandTracerouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Output** | Pointer to **string** |  | [optional] 
**OutputStructuredTraceroute** | Pointer to [**OutputStructuredTraceroute**](OutputStructuredTraceroute.md) |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) |  | [optional] 

## Methods

### NewCloudRouterCommandTracerouteResponse

`func NewCloudRouterCommandTracerouteResponse() *CloudRouterCommandTracerouteResponse`

NewCloudRouterCommandTracerouteResponse instantiates a new CloudRouterCommandTracerouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterCommandTracerouteResponseWithDefaults

`func NewCloudRouterCommandTracerouteResponseWithDefaults() *CloudRouterCommandTracerouteResponse`

NewCloudRouterCommandTracerouteResponseWithDefaults instantiates a new CloudRouterCommandTracerouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutput

`func (o *CloudRouterCommandTracerouteResponse) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *CloudRouterCommandTracerouteResponse) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *CloudRouterCommandTracerouteResponse) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *CloudRouterCommandTracerouteResponse) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetOutputStructuredTraceroute

`func (o *CloudRouterCommandTracerouteResponse) GetOutputStructuredTraceroute() OutputStructuredTraceroute`

GetOutputStructuredTraceroute returns the OutputStructuredTraceroute field if non-nil, zero value otherwise.

### GetOutputStructuredTracerouteOk

`func (o *CloudRouterCommandTracerouteResponse) GetOutputStructuredTracerouteOk() (*OutputStructuredTraceroute, bool)`

GetOutputStructuredTracerouteOk returns a tuple with the OutputStructuredTraceroute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputStructuredTraceroute

`func (o *CloudRouterCommandTracerouteResponse) SetOutputStructuredTraceroute(v OutputStructuredTraceroute)`

SetOutputStructuredTraceroute sets OutputStructuredTraceroute field to given value.

### HasOutputStructuredTraceroute

`func (o *CloudRouterCommandTracerouteResponse) HasOutputStructuredTraceroute() bool`

HasOutputStructuredTraceroute returns a boolean if a field has been set.

### GetErrors

`func (o *CloudRouterCommandTracerouteResponse) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CloudRouterCommandTracerouteResponse) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CloudRouterCommandTracerouteResponse) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CloudRouterCommandTracerouteResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


