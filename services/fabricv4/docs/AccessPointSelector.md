# AccessPointSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AccessPointSelectorType**](AccessPointSelectorType.md) |  | [optional] 
**Port** | Pointer to [**SimplifiedMetadataEntity**](SimplifiedMetadataEntity.md) |  | [optional] 
**LinkProtocol** | Pointer to [**LinkProtocol**](LinkProtocol.md) |  | [optional] 

## Methods

### NewAccessPointSelector

`func NewAccessPointSelector() *AccessPointSelector`

NewAccessPointSelector instantiates a new AccessPointSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPointSelectorWithDefaults

`func NewAccessPointSelectorWithDefaults() *AccessPointSelector`

NewAccessPointSelectorWithDefaults instantiates a new AccessPointSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccessPointSelector) GetType() AccessPointSelectorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessPointSelector) GetTypeOk() (*AccessPointSelectorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessPointSelector) SetType(v AccessPointSelectorType)`

SetType sets Type field to given value.

### HasType

`func (o *AccessPointSelector) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPort

`func (o *AccessPointSelector) GetPort() SimplifiedMetadataEntity`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AccessPointSelector) GetPortOk() (*SimplifiedMetadataEntity, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AccessPointSelector) SetPort(v SimplifiedMetadataEntity)`

SetPort sets Port field to given value.

### HasPort

`func (o *AccessPointSelector) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetLinkProtocol

`func (o *AccessPointSelector) GetLinkProtocol() LinkProtocol`

GetLinkProtocol returns the LinkProtocol field if non-nil, zero value otherwise.

### GetLinkProtocolOk

`func (o *AccessPointSelector) GetLinkProtocolOk() (*LinkProtocol, bool)`

GetLinkProtocolOk returns a tuple with the LinkProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkProtocol

`func (o *AccessPointSelector) SetLinkProtocol(v LinkProtocol)`

SetLinkProtocol sets LinkProtocol field to given value.

### HasLinkProtocol

`func (o *AccessPointSelector) HasLinkProtocol() bool`

HasLinkProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


