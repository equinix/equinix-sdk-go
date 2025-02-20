# CurrentTagPointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNo** | Pointer to **string** | customer account number | [optional] 
**Ibx** | Pointer to **string** | ibx code | [optional] 
**TagIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCurrentTagPointRequest

`func NewCurrentTagPointRequest() *CurrentTagPointRequest`

NewCurrentTagPointRequest instantiates a new CurrentTagPointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentTagPointRequestWithDefaults

`func NewCurrentTagPointRequestWithDefaults() *CurrentTagPointRequest`

NewCurrentTagPointRequestWithDefaults instantiates a new CurrentTagPointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNo

`func (o *CurrentTagPointRequest) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *CurrentTagPointRequest) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *CurrentTagPointRequest) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.

### HasAccountNo

`func (o *CurrentTagPointRequest) HasAccountNo() bool`

HasAccountNo returns a boolean if a field has been set.

### GetIbx

`func (o *CurrentTagPointRequest) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *CurrentTagPointRequest) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *CurrentTagPointRequest) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *CurrentTagPointRequest) HasIbx() bool`

HasIbx returns a boolean if a field has been set.

### GetTagIds

`func (o *CurrentTagPointRequest) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *CurrentTagPointRequest) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *CurrentTagPointRequest) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *CurrentTagPointRequest) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


