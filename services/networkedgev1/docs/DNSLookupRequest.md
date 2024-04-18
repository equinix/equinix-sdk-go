# DNSLookupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdns** | Pointer to **[]string** | Domain names | [optional] 
**Uuid** | Pointer to **string** | The unique Id of a virtual device | [optional] 
**MetroCode** | Pointer to **string** | Metro code | [optional] 

## Methods

### NewDNSLookupRequest

`func NewDNSLookupRequest() *DNSLookupRequest`

NewDNSLookupRequest instantiates a new DNSLookupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSLookupRequestWithDefaults

`func NewDNSLookupRequestWithDefaults() *DNSLookupRequest`

NewDNSLookupRequestWithDefaults instantiates a new DNSLookupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdns

`func (o *DNSLookupRequest) GetFqdns() []string`

GetFqdns returns the Fqdns field if non-nil, zero value otherwise.

### GetFqdnsOk

`func (o *DNSLookupRequest) GetFqdnsOk() (*[]string, bool)`

GetFqdnsOk returns a tuple with the Fqdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdns

`func (o *DNSLookupRequest) SetFqdns(v []string)`

SetFqdns sets Fqdns field to given value.

### HasFqdns

`func (o *DNSLookupRequest) HasFqdns() bool`

HasFqdns returns a boolean if a field has been set.

### GetUuid

`func (o *DNSLookupRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DNSLookupRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DNSLookupRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DNSLookupRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetMetroCode

`func (o *DNSLookupRequest) GetMetroCode() string`

GetMetroCode returns the MetroCode field if non-nil, zero value otherwise.

### GetMetroCodeOk

`func (o *DNSLookupRequest) GetMetroCodeOk() (*string, bool)`

GetMetroCodeOk returns a tuple with the MetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCode

`func (o *DNSLookupRequest) SetMetroCode(v string)`

SetMetroCode sets MetroCode field to given value.

### HasMetroCode

`func (o *DNSLookupRequest) HasMetroCode() bool`

HasMetroCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


