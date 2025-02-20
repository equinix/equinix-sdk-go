# PowerCurrentPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNo** | Pointer to **string** |  Customer Account Number | [optional] 
**Ibx** | Pointer to **string** | trending values | [optional] 
**LevelType** | Pointer to [**PowerCurrentPostRequestLevelType**](PowerCurrentPostRequestLevelType.md) |  | [optional] 

## Methods

### NewPowerCurrentPostRequest

`func NewPowerCurrentPostRequest() *PowerCurrentPostRequest`

NewPowerCurrentPostRequest instantiates a new PowerCurrentPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerCurrentPostRequestWithDefaults

`func NewPowerCurrentPostRequestWithDefaults() *PowerCurrentPostRequest`

NewPowerCurrentPostRequestWithDefaults instantiates a new PowerCurrentPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNo

`func (o *PowerCurrentPostRequest) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *PowerCurrentPostRequest) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *PowerCurrentPostRequest) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.

### HasAccountNo

`func (o *PowerCurrentPostRequest) HasAccountNo() bool`

HasAccountNo returns a boolean if a field has been set.

### GetIbx

`func (o *PowerCurrentPostRequest) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *PowerCurrentPostRequest) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *PowerCurrentPostRequest) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *PowerCurrentPostRequest) HasIbx() bool`

HasIbx returns a boolean if a field has been set.

### GetLevelType

`func (o *PowerCurrentPostRequest) GetLevelType() PowerCurrentPostRequestLevelType`

GetLevelType returns the LevelType field if non-nil, zero value otherwise.

### GetLevelTypeOk

`func (o *PowerCurrentPostRequest) GetLevelTypeOk() (*PowerCurrentPostRequestLevelType, bool)`

GetLevelTypeOk returns a tuple with the LevelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelType

`func (o *PowerCurrentPostRequest) SetLevelType(v PowerCurrentPostRequestLevelType)`

SetLevelType sets LevelType field to given value.

### HasLevelType

`func (o *PowerCurrentPostRequest) HasLevelType() bool`

HasLevelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


