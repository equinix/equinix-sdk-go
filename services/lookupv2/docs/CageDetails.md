# CageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Cage Unique space id | [optional] 
**Type** | Pointer to [**CageDetailsType**](CageDetailsType.md) |  | [optional] [default to CAGEDETAILSTYPE_PRIVATE]
**AccountNumbers** | Pointer to **string** | Customer account number associated with cage | [optional] 
**CabinetId** | Pointer to **string** | Cabinet Unique space id | [optional] 
**CabinetType** | Pointer to [**CageDetailsCabinetType**](CageDetailsCabinetType.md) |  | [optional] 

## Methods

### NewCageDetails

`func NewCageDetails() *CageDetails`

NewCageDetails instantiates a new CageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCageDetailsWithDefaults

`func NewCageDetailsWithDefaults() *CageDetails`

NewCageDetailsWithDefaults instantiates a new CageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CageDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CageDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CageDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CageDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CageDetails) GetType() CageDetailsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CageDetails) GetTypeOk() (*CageDetailsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CageDetails) SetType(v CageDetailsType)`

SetType sets Type field to given value.

### HasType

`func (o *CageDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccountNumbers

`func (o *CageDetails) GetAccountNumbers() string`

GetAccountNumbers returns the AccountNumbers field if non-nil, zero value otherwise.

### GetAccountNumbersOk

`func (o *CageDetails) GetAccountNumbersOk() (*string, bool)`

GetAccountNumbersOk returns a tuple with the AccountNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumbers

`func (o *CageDetails) SetAccountNumbers(v string)`

SetAccountNumbers sets AccountNumbers field to given value.

### HasAccountNumbers

`func (o *CageDetails) HasAccountNumbers() bool`

HasAccountNumbers returns a boolean if a field has been set.

### GetCabinetId

`func (o *CageDetails) GetCabinetId() string`

GetCabinetId returns the CabinetId field if non-nil, zero value otherwise.

### GetCabinetIdOk

`func (o *CageDetails) GetCabinetIdOk() (*string, bool)`

GetCabinetIdOk returns a tuple with the CabinetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinetId

`func (o *CageDetails) SetCabinetId(v string)`

SetCabinetId sets CabinetId field to given value.

### HasCabinetId

`func (o *CageDetails) HasCabinetId() bool`

HasCabinetId returns a boolean if a field has been set.

### GetCabinetType

`func (o *CageDetails) GetCabinetType() CageDetailsCabinetType`

GetCabinetType returns the CabinetType field if non-nil, zero value otherwise.

### GetCabinetTypeOk

`func (o *CageDetails) GetCabinetTypeOk() (*CageDetailsCabinetType, bool)`

GetCabinetTypeOk returns a tuple with the CabinetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinetType

`func (o *CageDetails) SetCabinetType(v CageDetailsCabinetType)`

SetCabinetType sets CabinetType field to given value.

### HasCabinetType

`func (o *CageDetails) HasCabinetType() bool`

HasCabinetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


