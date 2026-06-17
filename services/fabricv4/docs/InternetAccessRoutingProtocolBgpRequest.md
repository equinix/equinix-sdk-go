# InternetAccessRoutingProtocolBgpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | [**[]InternetAccessConnectionBgpRequest**](InternetAccessConnectionBgpRequest.md) |  | 
**ExportPolicy** | [**InternetAccessExportPolicy**](InternetAccessExportPolicy.md) |  | 
**CustomerAsn** | Pointer to **int64** | Customer ASN. Valid range is 1-64495 or 65536-4199999999. | [optional] 
**BgpAuthKey** | Pointer to **string** | BGP authentication key | [optional] 
**CustomerAsnRange** | Pointer to [**InternetAccessCustomerAsnRange**](InternetAccessCustomerAsnRange.md) |  | [optional] 

## Methods

### NewInternetAccessRoutingProtocolBgpRequest

`func NewInternetAccessRoutingProtocolBgpRequest(connections []InternetAccessConnectionBgpRequest, exportPolicy InternetAccessExportPolicy, ) *InternetAccessRoutingProtocolBgpRequest`

NewInternetAccessRoutingProtocolBgpRequest instantiates a new InternetAccessRoutingProtocolBgpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessRoutingProtocolBgpRequestWithDefaults

`func NewInternetAccessRoutingProtocolBgpRequestWithDefaults() *InternetAccessRoutingProtocolBgpRequest`

NewInternetAccessRoutingProtocolBgpRequestWithDefaults instantiates a new InternetAccessRoutingProtocolBgpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *InternetAccessRoutingProtocolBgpRequest) GetConnections() []InternetAccessConnectionBgpRequest`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *InternetAccessRoutingProtocolBgpRequest) GetConnectionsOk() (*[]InternetAccessConnectionBgpRequest, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *InternetAccessRoutingProtocolBgpRequest) SetConnections(v []InternetAccessConnectionBgpRequest)`

SetConnections sets Connections field to given value.


### GetExportPolicy

`func (o *InternetAccessRoutingProtocolBgpRequest) GetExportPolicy() InternetAccessExportPolicy`

GetExportPolicy returns the ExportPolicy field if non-nil, zero value otherwise.

### GetExportPolicyOk

`func (o *InternetAccessRoutingProtocolBgpRequest) GetExportPolicyOk() (*InternetAccessExportPolicy, bool)`

GetExportPolicyOk returns a tuple with the ExportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicy

`func (o *InternetAccessRoutingProtocolBgpRequest) SetExportPolicy(v InternetAccessExportPolicy)`

SetExportPolicy sets ExportPolicy field to given value.


### GetCustomerAsn

`func (o *InternetAccessRoutingProtocolBgpRequest) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *InternetAccessRoutingProtocolBgpRequest) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *InternetAccessRoutingProtocolBgpRequest) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.

### HasCustomerAsn

`func (o *InternetAccessRoutingProtocolBgpRequest) HasCustomerAsn() bool`

HasCustomerAsn returns a boolean if a field has been set.

### GetBgpAuthKey

`func (o *InternetAccessRoutingProtocolBgpRequest) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *InternetAccessRoutingProtocolBgpRequest) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *InternetAccessRoutingProtocolBgpRequest) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.

### HasBgpAuthKey

`func (o *InternetAccessRoutingProtocolBgpRequest) HasBgpAuthKey() bool`

HasBgpAuthKey returns a boolean if a field has been set.

### GetCustomerAsnRange

`func (o *InternetAccessRoutingProtocolBgpRequest) GetCustomerAsnRange() InternetAccessCustomerAsnRange`

GetCustomerAsnRange returns the CustomerAsnRange field if non-nil, zero value otherwise.

### GetCustomerAsnRangeOk

`func (o *InternetAccessRoutingProtocolBgpRequest) GetCustomerAsnRangeOk() (*InternetAccessCustomerAsnRange, bool)`

GetCustomerAsnRangeOk returns a tuple with the CustomerAsnRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsnRange

`func (o *InternetAccessRoutingProtocolBgpRequest) SetCustomerAsnRange(v InternetAccessCustomerAsnRange)`

SetCustomerAsnRange sets CustomerAsnRange field to given value.

### HasCustomerAsnRange

`func (o *InternetAccessRoutingProtocolBgpRequest) HasCustomerAsnRange() bool`

HasCustomerAsnRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


