# InternetAccessRoutingProtocolBgp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportPolicy** | [**InternetAccessExportPolicy**](InternetAccessExportPolicy.md) |  | 
**CustomerAsn** | Pointer to **int64** | Customer ASN. Valid range is 1-64495 or 131072-4199999999. Currently this option is only available for EIA over dedicated port. | [optional] 
**BgpAuthKey** | Pointer to **string** | BGP authentication key | [optional] 
**CustomerAsnRange** | [**InternetAccessCustomerAsnRange**](InternetAccessCustomerAsnRange.md) |  | 

## Methods

### NewInternetAccessRoutingProtocolBgp

`func NewInternetAccessRoutingProtocolBgp(exportPolicy InternetAccessExportPolicy, customerAsnRange InternetAccessCustomerAsnRange, ) *InternetAccessRoutingProtocolBgp`

NewInternetAccessRoutingProtocolBgp instantiates a new InternetAccessRoutingProtocolBgp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessRoutingProtocolBgpWithDefaults

`func NewInternetAccessRoutingProtocolBgpWithDefaults() *InternetAccessRoutingProtocolBgp`

NewInternetAccessRoutingProtocolBgpWithDefaults instantiates a new InternetAccessRoutingProtocolBgp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportPolicy

`func (o *InternetAccessRoutingProtocolBgp) GetExportPolicy() InternetAccessExportPolicy`

GetExportPolicy returns the ExportPolicy field if non-nil, zero value otherwise.

### GetExportPolicyOk

`func (o *InternetAccessRoutingProtocolBgp) GetExportPolicyOk() (*InternetAccessExportPolicy, bool)`

GetExportPolicyOk returns a tuple with the ExportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicy

`func (o *InternetAccessRoutingProtocolBgp) SetExportPolicy(v InternetAccessExportPolicy)`

SetExportPolicy sets ExportPolicy field to given value.


### GetCustomerAsn

`func (o *InternetAccessRoutingProtocolBgp) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *InternetAccessRoutingProtocolBgp) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *InternetAccessRoutingProtocolBgp) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.

### HasCustomerAsn

`func (o *InternetAccessRoutingProtocolBgp) HasCustomerAsn() bool`

HasCustomerAsn returns a boolean if a field has been set.

### GetBgpAuthKey

`func (o *InternetAccessRoutingProtocolBgp) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *InternetAccessRoutingProtocolBgp) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *InternetAccessRoutingProtocolBgp) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.

### HasBgpAuthKey

`func (o *InternetAccessRoutingProtocolBgp) HasBgpAuthKey() bool`

HasBgpAuthKey returns a boolean if a field has been set.

### GetCustomerAsnRange

`func (o *InternetAccessRoutingProtocolBgp) GetCustomerAsnRange() InternetAccessCustomerAsnRange`

GetCustomerAsnRange returns the CustomerAsnRange field if non-nil, zero value otherwise.

### GetCustomerAsnRangeOk

`func (o *InternetAccessRoutingProtocolBgp) GetCustomerAsnRangeOk() (*InternetAccessCustomerAsnRange, bool)`

GetCustomerAsnRangeOk returns a tuple with the CustomerAsnRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsnRange

`func (o *InternetAccessRoutingProtocolBgp) SetCustomerAsnRange(v InternetAccessCustomerAsnRange)`

SetCustomerAsnRange sets CustomerAsnRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


