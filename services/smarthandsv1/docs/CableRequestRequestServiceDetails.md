# CableRequestRequestServiceDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaType** | Pointer to [**CableRequestRequestServiceDetailsMediaType**](CableRequestRequestServiceDetailsMediaType.md) |  | [optional] 
**ConnectorType** | Pointer to [**CableRequestRequestServiceDetailsConnectorType**](CableRequestRequestServiceDetailsConnectorType.md) |  | [optional] 
**Length** | Pointer to **string** | Length (feet/cm). | [optional] 
**Quantity** | [**CableRequestRequestServiceDetailsQuantity**](CableRequestRequestServiceDetailsQuantity.md) |  | 
**ScopeOfWork** | **string** | Enter any additional details that will help our technicians execute your request. You may also attach your scope of work as a document if you exceed the character limit in this field. | 

## Methods

### NewCableRequestRequestServiceDetails

`func NewCableRequestRequestServiceDetails(quantity CableRequestRequestServiceDetailsQuantity, scopeOfWork string, ) *CableRequestRequestServiceDetails`

NewCableRequestRequestServiceDetails instantiates a new CableRequestRequestServiceDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCableRequestRequestServiceDetailsWithDefaults

`func NewCableRequestRequestServiceDetailsWithDefaults() *CableRequestRequestServiceDetails`

NewCableRequestRequestServiceDetailsWithDefaults instantiates a new CableRequestRequestServiceDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaType

`func (o *CableRequestRequestServiceDetails) GetMediaType() CableRequestRequestServiceDetailsMediaType`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *CableRequestRequestServiceDetails) GetMediaTypeOk() (*CableRequestRequestServiceDetailsMediaType, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *CableRequestRequestServiceDetails) SetMediaType(v CableRequestRequestServiceDetailsMediaType)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *CableRequestRequestServiceDetails) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetConnectorType

`func (o *CableRequestRequestServiceDetails) GetConnectorType() CableRequestRequestServiceDetailsConnectorType`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *CableRequestRequestServiceDetails) GetConnectorTypeOk() (*CableRequestRequestServiceDetailsConnectorType, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *CableRequestRequestServiceDetails) SetConnectorType(v CableRequestRequestServiceDetailsConnectorType)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *CableRequestRequestServiceDetails) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.

### GetLength

`func (o *CableRequestRequestServiceDetails) GetLength() string`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *CableRequestRequestServiceDetails) GetLengthOk() (*string, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *CableRequestRequestServiceDetails) SetLength(v string)`

SetLength sets Length field to given value.

### HasLength

`func (o *CableRequestRequestServiceDetails) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetQuantity

`func (o *CableRequestRequestServiceDetails) GetQuantity() CableRequestRequestServiceDetailsQuantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CableRequestRequestServiceDetails) GetQuantityOk() (*CableRequestRequestServiceDetailsQuantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CableRequestRequestServiceDetails) SetQuantity(v CableRequestRequestServiceDetailsQuantity)`

SetQuantity sets Quantity field to given value.


### GetScopeOfWork

`func (o *CableRequestRequestServiceDetails) GetScopeOfWork() string`

GetScopeOfWork returns the ScopeOfWork field if non-nil, zero value otherwise.

### GetScopeOfWorkOk

`func (o *CableRequestRequestServiceDetails) GetScopeOfWorkOk() (*string, bool)`

GetScopeOfWorkOk returns a tuple with the ScopeOfWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOfWork

`func (o *CableRequestRequestServiceDetails) SetScopeOfWork(v string)`

SetScopeOfWork sets ScopeOfWork field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


