# TimeServicePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**PrecisionTimeServiceRequestType**](PrecisionTimeServiceRequestType.md) |  | [optional] 
**Package** | Pointer to [**PrecisionTimePackageRequest**](PrecisionTimePackageRequest.md) |  | [optional] 
**Connection** | Pointer to [**TimeServicePriceConnection**](TimeServicePriceConnection.md) |  | [optional] 

## Methods

### NewTimeServicePrice

`func NewTimeServicePrice() *TimeServicePrice`

NewTimeServicePrice instantiates a new TimeServicePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeServicePriceWithDefaults

`func NewTimeServicePriceWithDefaults() *TimeServicePrice`

NewTimeServicePriceWithDefaults instantiates a new TimeServicePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TimeServicePrice) GetType() PrecisionTimeServiceRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeServicePrice) GetTypeOk() (*PrecisionTimeServiceRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeServicePrice) SetType(v PrecisionTimeServiceRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *TimeServicePrice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPackage

`func (o *TimeServicePrice) GetPackage() PrecisionTimePackageRequest`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *TimeServicePrice) GetPackageOk() (*PrecisionTimePackageRequest, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *TimeServicePrice) SetPackage(v PrecisionTimePackageRequest)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *TimeServicePrice) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetConnection

`func (o *TimeServicePrice) GetConnection() TimeServicePriceConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *TimeServicePrice) GetConnectionOk() (*TimeServicePriceConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *TimeServicePrice) SetConnection(v TimeServicePriceConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *TimeServicePrice) HasConnection() bool`

HasConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


