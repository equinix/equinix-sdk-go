# CloudRouterCommandResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Output** | Pointer to **string** |  | [optional] 
**OutputStructuredPing** | Pointer to [**OutputStructuredPing**](OutputStructuredPing.md) |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) |  | [optional] 
**OutputStructuredTraceroute** | Pointer to [**OutputStructuredTraceroute**](OutputStructuredTraceroute.md) |  | [optional] 

## Methods

### NewCloudRouterCommandResponse

`func NewCloudRouterCommandResponse() *CloudRouterCommandResponse`

NewCloudRouterCommandResponse instantiates a new CloudRouterCommandResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterCommandResponseWithDefaults

`func NewCloudRouterCommandResponseWithDefaults() *CloudRouterCommandResponse`

NewCloudRouterCommandResponseWithDefaults instantiates a new CloudRouterCommandResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutput

`func (o *CloudRouterCommandResponse) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *CloudRouterCommandResponse) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *CloudRouterCommandResponse) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *CloudRouterCommandResponse) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetOutputStructuredPing

`func (o *CloudRouterCommandResponse) GetOutputStructuredPing() OutputStructuredPing`

GetOutputStructuredPing returns the OutputStructuredPing field if non-nil, zero value otherwise.

### GetOutputStructuredPingOk

`func (o *CloudRouterCommandResponse) GetOutputStructuredPingOk() (*OutputStructuredPing, bool)`

GetOutputStructuredPingOk returns a tuple with the OutputStructuredPing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputStructuredPing

`func (o *CloudRouterCommandResponse) SetOutputStructuredPing(v OutputStructuredPing)`

SetOutputStructuredPing sets OutputStructuredPing field to given value.

### HasOutputStructuredPing

`func (o *CloudRouterCommandResponse) HasOutputStructuredPing() bool`

HasOutputStructuredPing returns a boolean if a field has been set.

### GetErrors

`func (o *CloudRouterCommandResponse) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CloudRouterCommandResponse) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CloudRouterCommandResponse) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CloudRouterCommandResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetOutputStructuredTraceroute

`func (o *CloudRouterCommandResponse) GetOutputStructuredTraceroute() OutputStructuredTraceroute`

GetOutputStructuredTraceroute returns the OutputStructuredTraceroute field if non-nil, zero value otherwise.

### GetOutputStructuredTracerouteOk

`func (o *CloudRouterCommandResponse) GetOutputStructuredTracerouteOk() (*OutputStructuredTraceroute, bool)`

GetOutputStructuredTracerouteOk returns a tuple with the OutputStructuredTraceroute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputStructuredTraceroute

`func (o *CloudRouterCommandResponse) SetOutputStructuredTraceroute(v OutputStructuredTraceroute)`

SetOutputStructuredTraceroute sets OutputStructuredTraceroute field to given value.

### HasOutputStructuredTraceroute

`func (o *CloudRouterCommandResponse) HasOutputStructuredTraceroute() bool`

HasOutputStructuredTraceroute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


