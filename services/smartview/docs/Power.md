# Power

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** | account number | [optional] 
**ApparentPower** | Pointer to [**PowerApparentPowerValueWithUnit**](PowerApparentPowerValueWithUnit.md) |  | [optional] 
**Asset** | [**PowerAssetDetails**](PowerAssetDetails.md) |  | 
**Cabinet** | Pointer to **string** | cabinet | [optional] 
**CabinetRating** | Pointer to [**PowerCabinetRatingValueWithUnit**](PowerCabinetRatingValueWithUnit.md) |  | [optional] 
**Cage** | Pointer to **string** | cage | [optional] 
**CircuitType** | Pointer to **string** | circuit type | [optional] 
**ContractualPower** | Pointer to [**PowerContractualPowerValueWithUnit**](PowerContractualPowerValueWithUnit.md) |  | [optional] 
**Current** | Pointer to [**PowerCurrentValueWithUnit**](PowerCurrentValueWithUnit.md) |  | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**Ibx** | **string** | ibx | 
**LastUpdated** | Pointer to **string** | last updated time | [optional] 
**Oid** | Pointer to **string** | oid | [optional] 
**PeakLastSevenDays** | Pointer to [**PowerPeakLastSevenDaysValueWithUnit**](PowerPeakLastSevenDaysValueWithUnit.md) |  | [optional] 
**PeakLastSevenDaysContractualPower** | Pointer to [**PowerPeakLastSevenDaysContractualPowerValueWithUnit**](PowerPeakLastSevenDaysContractualPowerValueWithUnit.md) |  | [optional] 
**PeakLastSevenDaysRatio** | Pointer to [**PowerPeakLastSevenDaysRatioValueWithUnit**](PowerPeakLastSevenDaysRatioValueWithUnit.md) |  | [optional] 
**PeakLastSevenDaysTime** | Pointer to **string** | peak last seven days time | [optional] 
**PowerConsumptionToContractual** | Pointer to [**PowerPowerConsumptionToContractualValueWithUnit**](PowerPowerConsumptionToContractualValueWithUnit.md) |  | [optional] 
**PowerFactor** | Pointer to [**PowerPowerFactorValueWithUnit**](PowerPowerFactorValueWithUnit.md) |  | [optional] 
**ReadingTime** | Pointer to **string** | message reading time | [optional] 
**RealPower** | Pointer to [**PowerRealPowerValueWithUnit**](PowerRealPowerValueWithUnit.md) |  | [optional] 
**SoldCurrent** | Pointer to [**PowerSoldCurrentValueWithUnit**](PowerSoldCurrentValueWithUnit.md) |  | [optional] 
**SoldPower** | Pointer to [**PowerSoldPowerValueWithUnit**](PowerSoldPowerValueWithUnit.md) |  | [optional] 
**StreamId** | **string** | unique message id | 

## Methods

### NewPower

`func NewPower(asset PowerAssetDetails, ibx string, streamId string, ) *Power`

NewPower instantiates a new Power object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerWithDefaults

`func NewPowerWithDefaults() *Power`

NewPowerWithDefaults instantiates a new Power object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *Power) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Power) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Power) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *Power) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetApparentPower

`func (o *Power) GetApparentPower() PowerApparentPowerValueWithUnit`

GetApparentPower returns the ApparentPower field if non-nil, zero value otherwise.

### GetApparentPowerOk

`func (o *Power) GetApparentPowerOk() (*PowerApparentPowerValueWithUnit, bool)`

GetApparentPowerOk returns a tuple with the ApparentPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApparentPower

`func (o *Power) SetApparentPower(v PowerApparentPowerValueWithUnit)`

SetApparentPower sets ApparentPower field to given value.

### HasApparentPower

`func (o *Power) HasApparentPower() bool`

HasApparentPower returns a boolean if a field has been set.

### GetAsset

`func (o *Power) GetAsset() PowerAssetDetails`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *Power) GetAssetOk() (*PowerAssetDetails, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *Power) SetAsset(v PowerAssetDetails)`

SetAsset sets Asset field to given value.


### GetCabinet

`func (o *Power) GetCabinet() string`

GetCabinet returns the Cabinet field if non-nil, zero value otherwise.

### GetCabinetOk

`func (o *Power) GetCabinetOk() (*string, bool)`

GetCabinetOk returns a tuple with the Cabinet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinet

`func (o *Power) SetCabinet(v string)`

SetCabinet sets Cabinet field to given value.

### HasCabinet

`func (o *Power) HasCabinet() bool`

HasCabinet returns a boolean if a field has been set.

### GetCabinetRating

`func (o *Power) GetCabinetRating() PowerCabinetRatingValueWithUnit`

GetCabinetRating returns the CabinetRating field if non-nil, zero value otherwise.

### GetCabinetRatingOk

`func (o *Power) GetCabinetRatingOk() (*PowerCabinetRatingValueWithUnit, bool)`

GetCabinetRatingOk returns a tuple with the CabinetRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinetRating

`func (o *Power) SetCabinetRating(v PowerCabinetRatingValueWithUnit)`

SetCabinetRating sets CabinetRating field to given value.

### HasCabinetRating

`func (o *Power) HasCabinetRating() bool`

HasCabinetRating returns a boolean if a field has been set.

### GetCage

`func (o *Power) GetCage() string`

GetCage returns the Cage field if non-nil, zero value otherwise.

### GetCageOk

`func (o *Power) GetCageOk() (*string, bool)`

GetCageOk returns a tuple with the Cage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCage

`func (o *Power) SetCage(v string)`

SetCage sets Cage field to given value.

### HasCage

`func (o *Power) HasCage() bool`

HasCage returns a boolean if a field has been set.

### GetCircuitType

`func (o *Power) GetCircuitType() string`

GetCircuitType returns the CircuitType field if non-nil, zero value otherwise.

### GetCircuitTypeOk

`func (o *Power) GetCircuitTypeOk() (*string, bool)`

GetCircuitTypeOk returns a tuple with the CircuitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitType

`func (o *Power) SetCircuitType(v string)`

SetCircuitType sets CircuitType field to given value.

### HasCircuitType

`func (o *Power) HasCircuitType() bool`

HasCircuitType returns a boolean if a field has been set.

### GetContractualPower

`func (o *Power) GetContractualPower() PowerContractualPowerValueWithUnit`

GetContractualPower returns the ContractualPower field if non-nil, zero value otherwise.

### GetContractualPowerOk

`func (o *Power) GetContractualPowerOk() (*PowerContractualPowerValueWithUnit, bool)`

GetContractualPowerOk returns a tuple with the ContractualPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractualPower

`func (o *Power) SetContractualPower(v PowerContractualPowerValueWithUnit)`

SetContractualPower sets ContractualPower field to given value.

### HasContractualPower

`func (o *Power) HasContractualPower() bool`

HasContractualPower returns a boolean if a field has been set.

### GetCurrent

`func (o *Power) GetCurrent() PowerCurrentValueWithUnit`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *Power) GetCurrentOk() (*PowerCurrentValueWithUnit, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *Power) SetCurrent(v PowerCurrentValueWithUnit)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *Power) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetDescription

`func (o *Power) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Power) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Power) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Power) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIbx

`func (o *Power) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *Power) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *Power) SetIbx(v string)`

SetIbx sets Ibx field to given value.


### GetLastUpdated

`func (o *Power) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Power) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Power) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Power) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetOid

`func (o *Power) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *Power) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *Power) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *Power) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetPeakLastSevenDays

`func (o *Power) GetPeakLastSevenDays() PowerPeakLastSevenDaysValueWithUnit`

GetPeakLastSevenDays returns the PeakLastSevenDays field if non-nil, zero value otherwise.

### GetPeakLastSevenDaysOk

`func (o *Power) GetPeakLastSevenDaysOk() (*PowerPeakLastSevenDaysValueWithUnit, bool)`

GetPeakLastSevenDaysOk returns a tuple with the PeakLastSevenDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLastSevenDays

`func (o *Power) SetPeakLastSevenDays(v PowerPeakLastSevenDaysValueWithUnit)`

SetPeakLastSevenDays sets PeakLastSevenDays field to given value.

### HasPeakLastSevenDays

`func (o *Power) HasPeakLastSevenDays() bool`

HasPeakLastSevenDays returns a boolean if a field has been set.

### GetPeakLastSevenDaysContractualPower

`func (o *Power) GetPeakLastSevenDaysContractualPower() PowerPeakLastSevenDaysContractualPowerValueWithUnit`

GetPeakLastSevenDaysContractualPower returns the PeakLastSevenDaysContractualPower field if non-nil, zero value otherwise.

### GetPeakLastSevenDaysContractualPowerOk

`func (o *Power) GetPeakLastSevenDaysContractualPowerOk() (*PowerPeakLastSevenDaysContractualPowerValueWithUnit, bool)`

GetPeakLastSevenDaysContractualPowerOk returns a tuple with the PeakLastSevenDaysContractualPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLastSevenDaysContractualPower

`func (o *Power) SetPeakLastSevenDaysContractualPower(v PowerPeakLastSevenDaysContractualPowerValueWithUnit)`

SetPeakLastSevenDaysContractualPower sets PeakLastSevenDaysContractualPower field to given value.

### HasPeakLastSevenDaysContractualPower

`func (o *Power) HasPeakLastSevenDaysContractualPower() bool`

HasPeakLastSevenDaysContractualPower returns a boolean if a field has been set.

### GetPeakLastSevenDaysRatio

`func (o *Power) GetPeakLastSevenDaysRatio() PowerPeakLastSevenDaysRatioValueWithUnit`

GetPeakLastSevenDaysRatio returns the PeakLastSevenDaysRatio field if non-nil, zero value otherwise.

### GetPeakLastSevenDaysRatioOk

`func (o *Power) GetPeakLastSevenDaysRatioOk() (*PowerPeakLastSevenDaysRatioValueWithUnit, bool)`

GetPeakLastSevenDaysRatioOk returns a tuple with the PeakLastSevenDaysRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLastSevenDaysRatio

`func (o *Power) SetPeakLastSevenDaysRatio(v PowerPeakLastSevenDaysRatioValueWithUnit)`

SetPeakLastSevenDaysRatio sets PeakLastSevenDaysRatio field to given value.

### HasPeakLastSevenDaysRatio

`func (o *Power) HasPeakLastSevenDaysRatio() bool`

HasPeakLastSevenDaysRatio returns a boolean if a field has been set.

### GetPeakLastSevenDaysTime

`func (o *Power) GetPeakLastSevenDaysTime() string`

GetPeakLastSevenDaysTime returns the PeakLastSevenDaysTime field if non-nil, zero value otherwise.

### GetPeakLastSevenDaysTimeOk

`func (o *Power) GetPeakLastSevenDaysTimeOk() (*string, bool)`

GetPeakLastSevenDaysTimeOk returns a tuple with the PeakLastSevenDaysTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakLastSevenDaysTime

`func (o *Power) SetPeakLastSevenDaysTime(v string)`

SetPeakLastSevenDaysTime sets PeakLastSevenDaysTime field to given value.

### HasPeakLastSevenDaysTime

`func (o *Power) HasPeakLastSevenDaysTime() bool`

HasPeakLastSevenDaysTime returns a boolean if a field has been set.

### GetPowerConsumptionToContractual

`func (o *Power) GetPowerConsumptionToContractual() PowerPowerConsumptionToContractualValueWithUnit`

GetPowerConsumptionToContractual returns the PowerConsumptionToContractual field if non-nil, zero value otherwise.

### GetPowerConsumptionToContractualOk

`func (o *Power) GetPowerConsumptionToContractualOk() (*PowerPowerConsumptionToContractualValueWithUnit, bool)`

GetPowerConsumptionToContractualOk returns a tuple with the PowerConsumptionToContractual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerConsumptionToContractual

`func (o *Power) SetPowerConsumptionToContractual(v PowerPowerConsumptionToContractualValueWithUnit)`

SetPowerConsumptionToContractual sets PowerConsumptionToContractual field to given value.

### HasPowerConsumptionToContractual

`func (o *Power) HasPowerConsumptionToContractual() bool`

HasPowerConsumptionToContractual returns a boolean if a field has been set.

### GetPowerFactor

`func (o *Power) GetPowerFactor() PowerPowerFactorValueWithUnit`

GetPowerFactor returns the PowerFactor field if non-nil, zero value otherwise.

### GetPowerFactorOk

`func (o *Power) GetPowerFactorOk() (*PowerPowerFactorValueWithUnit, bool)`

GetPowerFactorOk returns a tuple with the PowerFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFactor

`func (o *Power) SetPowerFactor(v PowerPowerFactorValueWithUnit)`

SetPowerFactor sets PowerFactor field to given value.

### HasPowerFactor

`func (o *Power) HasPowerFactor() bool`

HasPowerFactor returns a boolean if a field has been set.

### GetReadingTime

`func (o *Power) GetReadingTime() string`

GetReadingTime returns the ReadingTime field if non-nil, zero value otherwise.

### GetReadingTimeOk

`func (o *Power) GetReadingTimeOk() (*string, bool)`

GetReadingTimeOk returns a tuple with the ReadingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingTime

`func (o *Power) SetReadingTime(v string)`

SetReadingTime sets ReadingTime field to given value.

### HasReadingTime

`func (o *Power) HasReadingTime() bool`

HasReadingTime returns a boolean if a field has been set.

### GetRealPower

`func (o *Power) GetRealPower() PowerRealPowerValueWithUnit`

GetRealPower returns the RealPower field if non-nil, zero value otherwise.

### GetRealPowerOk

`func (o *Power) GetRealPowerOk() (*PowerRealPowerValueWithUnit, bool)`

GetRealPowerOk returns a tuple with the RealPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealPower

`func (o *Power) SetRealPower(v PowerRealPowerValueWithUnit)`

SetRealPower sets RealPower field to given value.

### HasRealPower

`func (o *Power) HasRealPower() bool`

HasRealPower returns a boolean if a field has been set.

### GetSoldCurrent

`func (o *Power) GetSoldCurrent() PowerSoldCurrentValueWithUnit`

GetSoldCurrent returns the SoldCurrent field if non-nil, zero value otherwise.

### GetSoldCurrentOk

`func (o *Power) GetSoldCurrentOk() (*PowerSoldCurrentValueWithUnit, bool)`

GetSoldCurrentOk returns a tuple with the SoldCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldCurrent

`func (o *Power) SetSoldCurrent(v PowerSoldCurrentValueWithUnit)`

SetSoldCurrent sets SoldCurrent field to given value.

### HasSoldCurrent

`func (o *Power) HasSoldCurrent() bool`

HasSoldCurrent returns a boolean if a field has been set.

### GetSoldPower

`func (o *Power) GetSoldPower() PowerSoldPowerValueWithUnit`

GetSoldPower returns the SoldPower field if non-nil, zero value otherwise.

### GetSoldPowerOk

`func (o *Power) GetSoldPowerOk() (*PowerSoldPowerValueWithUnit, bool)`

GetSoldPowerOk returns a tuple with the SoldPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldPower

`func (o *Power) SetSoldPower(v PowerSoldPowerValueWithUnit)`

SetSoldPower sets SoldPower field to given value.

### HasSoldPower

`func (o *Power) HasSoldPower() bool`

HasSoldPower returns a boolean if a field has been set.

### GetStreamId

`func (o *Power) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *Power) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *Power) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


