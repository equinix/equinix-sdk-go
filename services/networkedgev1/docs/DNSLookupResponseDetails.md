# DNSLookupResponseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cdn** | Pointer to **bool** | True or false | [optional] 
**Ips** | Pointer to **[]string** | Domain names | [optional] 

## Methods

### NewDNSLookupResponseDetails

`func NewDNSLookupResponseDetails() *DNSLookupResponseDetails`

NewDNSLookupResponseDetails instantiates a new DNSLookupResponseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSLookupResponseDetailsWithDefaults

`func NewDNSLookupResponseDetailsWithDefaults() *DNSLookupResponseDetails`

NewDNSLookupResponseDetailsWithDefaults instantiates a new DNSLookupResponseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdn

`func (o *DNSLookupResponseDetails) GetCdn() bool`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *DNSLookupResponseDetails) GetCdnOk() (*bool, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *DNSLookupResponseDetails) SetCdn(v bool)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *DNSLookupResponseDetails) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### GetIps

`func (o *DNSLookupResponseDetails) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *DNSLookupResponseDetails) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *DNSLookupResponseDetails) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *DNSLookupResponseDetails) HasIps() bool`

HasIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


