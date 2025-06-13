# CloudRouterCommandPingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Output** | Pointer to **string** |  | [optional] 
**OutputStructuredPing** | Pointer to [**OutputStructuredPing**](OutputStructuredPing.md) |  | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) |  | [optional] 

## Methods

### NewCloudRouterCommandPingResponse

`func NewCloudRouterCommandPingResponse() *CloudRouterCommandPingResponse`

NewCloudRouterCommandPingResponse instantiates a new CloudRouterCommandPingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterCommandPingResponseWithDefaults

`func NewCloudRouterCommandPingResponseWithDefaults() *CloudRouterCommandPingResponse`

NewCloudRouterCommandPingResponseWithDefaults instantiates a new CloudRouterCommandPingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutput

`func (o *CloudRouterCommandPingResponse) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *CloudRouterCommandPingResponse) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *CloudRouterCommandPingResponse) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *CloudRouterCommandPingResponse) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetOutputStructuredPing

`func (o *CloudRouterCommandPingResponse) GetOutputStructuredPing() OutputStructuredPing`

GetOutputStructuredPing returns the OutputStructuredPing field if non-nil, zero value otherwise.

### GetOutputStructuredPingOk

`func (o *CloudRouterCommandPingResponse) GetOutputStructuredPingOk() (*OutputStructuredPing, bool)`

GetOutputStructuredPingOk returns a tuple with the OutputStructuredPing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputStructuredPing

`func (o *CloudRouterCommandPingResponse) SetOutputStructuredPing(v OutputStructuredPing)`

SetOutputStructuredPing sets OutputStructuredPing field to given value.

### HasOutputStructuredPing

`func (o *CloudRouterCommandPingResponse) HasOutputStructuredPing() bool`

HasOutputStructuredPing returns a boolean if a field has been set.

### GetErrors

`func (o *CloudRouterCommandPingResponse) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CloudRouterCommandPingResponse) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CloudRouterCommandPingResponse) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CloudRouterCommandPingResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


