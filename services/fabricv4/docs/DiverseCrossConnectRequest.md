# DiverseCrossConnectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CreateCrossConnect**](CreateCrossConnect.md) | For a diverse pair, send two entries - one with redundancy.priority PRIMARY and one with SECONDARY - and omit redundancy.group, which Equinix assigns.  | 

## Methods

### NewDiverseCrossConnectRequest

`func NewDiverseCrossConnectRequest(data []CreateCrossConnect, ) *DiverseCrossConnectRequest`

NewDiverseCrossConnectRequest instantiates a new DiverseCrossConnectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiverseCrossConnectRequestWithDefaults

`func NewDiverseCrossConnectRequestWithDefaults() *DiverseCrossConnectRequest`

NewDiverseCrossConnectRequestWithDefaults instantiates a new DiverseCrossConnectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DiverseCrossConnectRequest) GetData() []CreateCrossConnect`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DiverseCrossConnectRequest) GetDataOk() (*[]CreateCrossConnect, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DiverseCrossConnectRequest) SetData(v []CreateCrossConnect)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


