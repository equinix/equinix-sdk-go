# GenericError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]string** |  | [optional] 
**TicketId** | Pointer to **string** |  | [optional] 

## Methods

### NewGenericError

`func NewGenericError() *GenericError`

NewGenericError instantiates a new GenericError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericErrorWithDefaults

`func NewGenericErrorWithDefaults() *GenericError`

NewGenericErrorWithDefaults instantiates a new GenericError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *GenericError) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GenericError) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GenericError) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GenericError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetTicketId

`func (o *GenericError) GetTicketId() string`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *GenericError) GetTicketIdOk() (*string, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *GenericError) SetTicketId(v string)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *GenericError) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


