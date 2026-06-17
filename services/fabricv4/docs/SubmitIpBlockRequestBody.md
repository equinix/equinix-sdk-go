# SubmitIpBlockRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**IpBlockLocation**](IpBlockLocation.md) |  | [optional] 
**Type** | [**TypeOfIpBlockProduct**](TypeOfIpBlockProduct.md) |  | 
**Project** | [**IpBlockProjectRequest**](IpBlockProjectRequest.md) |  | 
**Account** | Pointer to [**IpBlockAccount**](IpBlockAccount.md) |  | [optional] 
**Order** | Pointer to [**IpBlockOrderRequest**](IpBlockOrderRequest.md) |  | [optional] 
**Regulations** | Pointer to [**IpBlockRegulations**](IpBlockRegulations.md) |  | [optional] 
**PrefixLength** | Pointer to **int32** | IpBlockPrefix length | [optional] 
**Prefix** | Pointer to **string** | CIDR prefix | [optional] 

## Methods

### NewSubmitIpBlockRequestBody

`func NewSubmitIpBlockRequestBody(type_ TypeOfIpBlockProduct, project IpBlockProjectRequest, ) *SubmitIpBlockRequestBody`

NewSubmitIpBlockRequestBody instantiates a new SubmitIpBlockRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitIpBlockRequestBodyWithDefaults

`func NewSubmitIpBlockRequestBodyWithDefaults() *SubmitIpBlockRequestBody`

NewSubmitIpBlockRequestBodyWithDefaults instantiates a new SubmitIpBlockRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *SubmitIpBlockRequestBody) GetLocation() IpBlockLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SubmitIpBlockRequestBody) GetLocationOk() (*IpBlockLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SubmitIpBlockRequestBody) SetLocation(v IpBlockLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SubmitIpBlockRequestBody) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetType

`func (o *SubmitIpBlockRequestBody) GetType() TypeOfIpBlockProduct`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubmitIpBlockRequestBody) GetTypeOk() (*TypeOfIpBlockProduct, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubmitIpBlockRequestBody) SetType(v TypeOfIpBlockProduct)`

SetType sets Type field to given value.


### GetProject

`func (o *SubmitIpBlockRequestBody) GetProject() IpBlockProjectRequest`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SubmitIpBlockRequestBody) GetProjectOk() (*IpBlockProjectRequest, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SubmitIpBlockRequestBody) SetProject(v IpBlockProjectRequest)`

SetProject sets Project field to given value.


### GetAccount

`func (o *SubmitIpBlockRequestBody) GetAccount() IpBlockAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SubmitIpBlockRequestBody) GetAccountOk() (*IpBlockAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SubmitIpBlockRequestBody) SetAccount(v IpBlockAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SubmitIpBlockRequestBody) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOrder

`func (o *SubmitIpBlockRequestBody) GetOrder() IpBlockOrderRequest`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SubmitIpBlockRequestBody) GetOrderOk() (*IpBlockOrderRequest, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SubmitIpBlockRequestBody) SetOrder(v IpBlockOrderRequest)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *SubmitIpBlockRequestBody) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetRegulations

`func (o *SubmitIpBlockRequestBody) GetRegulations() IpBlockRegulations`

GetRegulations returns the Regulations field if non-nil, zero value otherwise.

### GetRegulationsOk

`func (o *SubmitIpBlockRequestBody) GetRegulationsOk() (*IpBlockRegulations, bool)`

GetRegulationsOk returns a tuple with the Regulations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulations

`func (o *SubmitIpBlockRequestBody) SetRegulations(v IpBlockRegulations)`

SetRegulations sets Regulations field to given value.

### HasRegulations

`func (o *SubmitIpBlockRequestBody) HasRegulations() bool`

HasRegulations returns a boolean if a field has been set.

### GetPrefixLength

`func (o *SubmitIpBlockRequestBody) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *SubmitIpBlockRequestBody) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *SubmitIpBlockRequestBody) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *SubmitIpBlockRequestBody) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.

### GetPrefix

`func (o *SubmitIpBlockRequestBody) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *SubmitIpBlockRequestBody) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *SubmitIpBlockRequestBody) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *SubmitIpBlockRequestBody) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


