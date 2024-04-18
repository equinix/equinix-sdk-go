# LinkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** | Account number. Either an account number or an accountreferenceId is required to create a link group. | [optional] 
**Throughput** | Pointer to **string** | Metro Throughput. | [optional] 
**ThroughputUnit** | Pointer to **string** | Throughput unit. | [optional] 
**MetroCode** | Pointer to **string** | Metro you want to link. | [optional] 

## Methods

### NewLinkInfo

`func NewLinkInfo() *LinkInfo`

NewLinkInfo instantiates a new LinkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkInfoWithDefaults

`func NewLinkInfoWithDefaults() *LinkInfo`

NewLinkInfoWithDefaults instantiates a new LinkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *LinkInfo) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *LinkInfo) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *LinkInfo) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *LinkInfo) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetThroughput

`func (o *LinkInfo) GetThroughput() string`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *LinkInfo) GetThroughputOk() (*string, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *LinkInfo) SetThroughput(v string)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *LinkInfo) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetThroughputUnit

`func (o *LinkInfo) GetThroughputUnit() string`

GetThroughputUnit returns the ThroughputUnit field if non-nil, zero value otherwise.

### GetThroughputUnitOk

`func (o *LinkInfo) GetThroughputUnitOk() (*string, bool)`

GetThroughputUnitOk returns a tuple with the ThroughputUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputUnit

`func (o *LinkInfo) SetThroughputUnit(v string)`

SetThroughputUnit sets ThroughputUnit field to given value.

### HasThroughputUnit

`func (o *LinkInfo) HasThroughputUnit() bool`

HasThroughputUnit returns a boolean if a field has been set.

### GetMetroCode

`func (o *LinkInfo) GetMetroCode() string`

GetMetroCode returns the MetroCode field if non-nil, zero value otherwise.

### GetMetroCodeOk

`func (o *LinkInfo) GetMetroCodeOk() (*string, bool)`

GetMetroCodeOk returns a tuple with the MetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCode

`func (o *LinkInfo) SetMetroCode(v string)`

SetMetroCode sets MetroCode field to given value.

### HasMetroCode

`func (o *LinkInfo) HasMetroCode() bool`

HasMetroCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


