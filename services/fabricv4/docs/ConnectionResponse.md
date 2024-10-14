# ConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalInfo** | Pointer to [**[]ConnectionSideAdditionalInfo**](ConnectionSideAdditionalInfo.md) | Additional information | [optional] 
**Data** | Pointer to [**[]ValidateConnectionResponse**](ValidateConnectionResponse.md) | Connection response data | [optional] 

## Methods

### NewConnectionResponse

`func NewConnectionResponse() *ConnectionResponse`

NewConnectionResponse instantiates a new ConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionResponseWithDefaults

`func NewConnectionResponseWithDefaults() *ConnectionResponse`

NewConnectionResponseWithDefaults instantiates a new ConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalInfo

`func (o *ConnectionResponse) GetAdditionalInfo() []ConnectionSideAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ConnectionResponse) GetAdditionalInfoOk() (*[]ConnectionSideAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ConnectionResponse) SetAdditionalInfo(v []ConnectionSideAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *ConnectionResponse) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetData

`func (o *ConnectionResponse) GetData() []ValidateConnectionResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConnectionResponse) GetDataOk() (*[]ValidateConnectionResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConnectionResponse) SetData(v []ValidateConnectionResponse)`

SetData sets Data field to given value.

### HasData

`func (o *ConnectionResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


