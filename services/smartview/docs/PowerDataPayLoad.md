# PowerDataPayLoad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNo** | Pointer to **string** | customer account number | [optional] 
**Amps** | Pointer to **float32** | instantaneous current amp reading on circuits | [optional] 
**CabinetRating** | Pointer to **float32** | maximum kVA draw allowed for the cabinet|cabinet description when the levelType is cabinet. Null otherwise. | [optional] 
**ComparisonData** | Pointer to [**ComparisonData**](ComparisonData.md) |  | [optional] 
**ContractualKva** | Pointer to **float32** | The maximum power draw contractually allowable in a  private cage.   | [optional] 
**CustomerName** | Pointer to **string** |  | [optional] 
**Ibx** | Pointer to **string** | ibx code | [optional] 
**IsAlarm** | Pointer to **string** | returns boolean based on breakertip alarm | [optional] 
**Kva** | Pointer to **float32** | power consumption in kva | [optional] 
**Kw** | Pointer to **string** | measure of real power expressed in kilowatt applicable for ibxs that have capability of energy meter reading|value will be \&quot;NA\&quot; for AMER and APAC regions   | [optional] 
**LastUpdatedTime** | Pointer to **string** | date-time when the latest value was updated (epoc - milliseconds).  | [optional] 
**LevelType** | Pointer to [**PowerDataPayLoadLevelType**](PowerDataPayLoadLevelType.md) |  | [optional] 
**LevelValue** | Pointer to **string** | power hierarchy node levelValue linked to the power data | [optional] 
**PeakKvaLastSevenDays** | Pointer to **float32** |  | [optional] 
**PeakKvaLastSevenDaysContractualKva** | Pointer to **float32** |  | [optional] 
**PeakKvaLastSevenDaysPercentage** | Pointer to **float32** |  | [optional] 
**PeakKvaLastSevenDaysTime** | Pointer to **int32** |  | [optional] 
**PercentageKva** | Pointer to **float32** | calculated field kva / contractualKva | [optional] 
**PowerFactor** | Pointer to **string** | The ratio between real power and apparent power in a circuit.(kW/kVA)|value will be \&quot;NA\&quot; for AMER and APAC regions  | [optional] 
**PrimaryKva** | Pointer to **float32** | the sum of instantaneous power draw reading on all the primary  circuits within the levelType.  | [optional] 
**ReadingTime** | Pointer to **string** | date-time when the latest value was read in (epoc - milliseconds).  | [optional] 
**RedundantKva** | Pointer to **float32** | the sum of instantaneous power draw reading on all the redundant  circuits within the levelType.  | [optional] 
**SoldAmps** | Pointer to **int32** | circuit description when the levelType is circuit. Null otherwise. | [optional] 
**SoldKva** | Pointer to **float32** | maximum amp draw allowable on a circuit | [optional] 

## Methods

### NewPowerDataPayLoad

`func NewPowerDataPayLoad() *PowerDataPayLoad`

NewPowerDataPayLoad instantiates a new PowerDataPayLoad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerDataPayLoadWithDefaults

`func NewPowerDataPayLoadWithDefaults() *PowerDataPayLoad`

NewPowerDataPayLoadWithDefaults instantiates a new PowerDataPayLoad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNo

`func (o *PowerDataPayLoad) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *PowerDataPayLoad) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *PowerDataPayLoad) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.

### HasAccountNo

`func (o *PowerDataPayLoad) HasAccountNo() bool`

HasAccountNo returns a boolean if a field has been set.

### GetAmps

`func (o *PowerDataPayLoad) GetAmps() float32`

GetAmps returns the Amps field if non-nil, zero value otherwise.

### GetAmpsOk

`func (o *PowerDataPayLoad) GetAmpsOk() (*float32, bool)`

GetAmpsOk returns a tuple with the Amps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmps

`func (o *PowerDataPayLoad) SetAmps(v float32)`

SetAmps sets Amps field to given value.

### HasAmps

`func (o *PowerDataPayLoad) HasAmps() bool`

HasAmps returns a boolean if a field has been set.

### GetCabinetRating

`func (o *PowerDataPayLoad) GetCabinetRating() float32`

GetCabinetRating returns the CabinetRating field if non-nil, zero value otherwise.

### GetCabinetRatingOk

`func (o *PowerDataPayLoad) GetCabinetRatingOk() (*float32, bool)`

GetCabinetRatingOk returns a tuple with the CabinetRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinetRating

`func (o *PowerDataPayLoad) SetCabinetRating(v float32)`

SetCabinetRating sets CabinetRating field to given value.

### HasCabinetRating

`func (o *PowerDataPayLoad) HasCabinetRating() bool`

HasCabinetRating returns a boolean if a field has been set.

### GetComparisonData

`func (o *PowerDataPayLoad) GetComparisonData() ComparisonData`

GetComparisonData returns the ComparisonData field if non-nil, zero value otherwise.

### GetComparisonDataOk

`func (o *PowerDataPayLoad) GetComparisonDataOk() (*ComparisonData, bool)`

GetComparisonDataOk returns a tuple with the ComparisonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonData

`func (o *PowerDataPayLoad) SetComparisonData(v ComparisonData)`

SetComparisonData sets ComparisonData field to given value.

### HasComparisonData

`func (o *PowerDataPayLoad) HasComparisonData() bool`

HasComparisonData returns a boolean if a field has been set.

### GetContractualKva

`func (o *PowerDataPayLoad) GetContractualKva() float32`

GetContractualKva returns the ContractualKva field if non-nil, zero value otherwise.

### GetContractualKvaOk

`func (o *PowerDataPayLoad) GetContractualKvaOk() (*float32, bool)`

GetContractualKvaOk returns a tuple with the ContractualKva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractualKva

`func (o *PowerDataPayLoad) SetContractualKva(v float32)`

SetContractualKva sets ContractualKva field to given value.

### HasContractualKva

`func (o *PowerDataPayLoad) HasContractualKva() bool`

HasContractualKva returns a boolean if a field has been set.

### GetCustomerName

`func (o *PowerDataPayLoad) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *PowerDataPayLoad) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *PowerDataPayLoad) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *PowerDataPayLoad) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetIbx

`func (o *PowerDataPayLoad) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *PowerDataPayLoad) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *PowerDataPayLoad) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *PowerDataPayLoad) HasIbx() bool`

HasIbx returns a boolean if a field has been set.

### GetIsAlarm

`func (o *PowerDataPayLoad) GetIsAlarm() string`

GetIsAlarm returns the IsAlarm field if non-nil, zero value otherwise.

### GetIsAlarmOk

`func (o *PowerDataPayLoad) GetIsAlarmOk() (*string, bool)`

GetIsAlarmOk returns a tuple with the IsAlarm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAlarm

`func (o *PowerDataPayLoad) SetIsAlarm(v string)`

SetIsAlarm sets IsAlarm field to given value.

### HasIsAlarm

`func (o *PowerDataPayLoad) HasIsAlarm() bool`

HasIsAlarm returns a boolean if a field has been set.

### GetKva

`func (o *PowerDataPayLoad) GetKva() float32`

GetKva returns the Kva field if non-nil, zero value otherwise.

### GetKvaOk

`func (o *PowerDataPayLoad) GetKvaOk() (*float32, bool)`

GetKvaOk returns a tuple with the Kva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKva

`func (o *PowerDataPayLoad) SetKva(v float32)`

SetKva sets Kva field to given value.

### HasKva

`func (o *PowerDataPayLoad) HasKva() bool`

HasKva returns a boolean if a field has been set.

### GetKw

`func (o *PowerDataPayLoad) GetKw() string`

GetKw returns the Kw field if non-nil, zero value otherwise.

### GetKwOk

`func (o *PowerDataPayLoad) GetKwOk() (*string, bool)`

GetKwOk returns a tuple with the Kw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKw

`func (o *PowerDataPayLoad) SetKw(v string)`

SetKw sets Kw field to given value.

### HasKw

`func (o *PowerDataPayLoad) HasKw() bool`

HasKw returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *PowerDataPayLoad) GetLastUpdatedTime() string`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *PowerDataPayLoad) GetLastUpdatedTimeOk() (*string, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *PowerDataPayLoad) SetLastUpdatedTime(v string)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *PowerDataPayLoad) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLevelType

`func (o *PowerDataPayLoad) GetLevelType() PowerDataPayLoadLevelType`

GetLevelType returns the LevelType field if non-nil, zero value otherwise.

### GetLevelTypeOk

`func (o *PowerDataPayLoad) GetLevelTypeOk() (*PowerDataPayLoadLevelType, bool)`

GetLevelTypeOk returns a tuple with the LevelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelType

`func (o *PowerDataPayLoad) SetLevelType(v PowerDataPayLoadLevelType)`

SetLevelType sets LevelType field to given value.

### HasLevelType

`func (o *PowerDataPayLoad) HasLevelType() bool`

HasLevelType returns a boolean if a field has been set.

### GetLevelValue

`func (o *PowerDataPayLoad) GetLevelValue() string`

GetLevelValue returns the LevelValue field if non-nil, zero value otherwise.

### GetLevelValueOk

`func (o *PowerDataPayLoad) GetLevelValueOk() (*string, bool)`

GetLevelValueOk returns a tuple with the LevelValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelValue

`func (o *PowerDataPayLoad) SetLevelValue(v string)`

SetLevelValue sets LevelValue field to given value.

### HasLevelValue

`func (o *PowerDataPayLoad) HasLevelValue() bool`

HasLevelValue returns a boolean if a field has been set.

### GetPeakKvaLastSevenDays

`func (o *PowerDataPayLoad) GetPeakKvaLastSevenDays() float32`

GetPeakKvaLastSevenDays returns the PeakKvaLastSevenDays field if non-nil, zero value otherwise.

### GetPeakKvaLastSevenDaysOk

`func (o *PowerDataPayLoad) GetPeakKvaLastSevenDaysOk() (*float32, bool)`

GetPeakKvaLastSevenDaysOk returns a tuple with the PeakKvaLastSevenDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakKvaLastSevenDays

`func (o *PowerDataPayLoad) SetPeakKvaLastSevenDays(v float32)`

SetPeakKvaLastSevenDays sets PeakKvaLastSevenDays field to given value.

### HasPeakKvaLastSevenDays

`func (o *PowerDataPayLoad) HasPeakKvaLastSevenDays() bool`

HasPeakKvaLastSevenDays returns a boolean if a field has been set.

### GetPeakKvaLastSevenDaysContractualKva

`func (o *PowerDataPayLoad) GetPeakKvaLastSevenDaysContractualKva() float32`

GetPeakKvaLastSevenDaysContractualKva returns the PeakKvaLastSevenDaysContractualKva field if non-nil, zero value otherwise.

### GetPeakKvaLastSevenDaysContractualKvaOk

`func (o *PowerDataPayLoad) GetPeakKvaLastSevenDaysContractualKvaOk() (*float32, bool)`

GetPeakKvaLastSevenDaysContractualKvaOk returns a tuple with the PeakKvaLastSevenDaysContractualKva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakKvaLastSevenDaysContractualKva

`func (o *PowerDataPayLoad) SetPeakKvaLastSevenDaysContractualKva(v float32)`

SetPeakKvaLastSevenDaysContractualKva sets PeakKvaLastSevenDaysContractualKva field to given value.

### HasPeakKvaLastSevenDaysContractualKva

`func (o *PowerDataPayLoad) HasPeakKvaLastSevenDaysContractualKva() bool`

HasPeakKvaLastSevenDaysContractualKva returns a boolean if a field has been set.

### GetPeakKvaLastSevenDaysPercentage

`func (o *PowerDataPayLoad) GetPeakKvaLastSevenDaysPercentage() float32`

GetPeakKvaLastSevenDaysPercentage returns the PeakKvaLastSevenDaysPercentage field if non-nil, zero value otherwise.

### GetPeakKvaLastSevenDaysPercentageOk

`func (o *PowerDataPayLoad) GetPeakKvaLastSevenDaysPercentageOk() (*float32, bool)`

GetPeakKvaLastSevenDaysPercentageOk returns a tuple with the PeakKvaLastSevenDaysPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakKvaLastSevenDaysPercentage

`func (o *PowerDataPayLoad) SetPeakKvaLastSevenDaysPercentage(v float32)`

SetPeakKvaLastSevenDaysPercentage sets PeakKvaLastSevenDaysPercentage field to given value.

### HasPeakKvaLastSevenDaysPercentage

`func (o *PowerDataPayLoad) HasPeakKvaLastSevenDaysPercentage() bool`

HasPeakKvaLastSevenDaysPercentage returns a boolean if a field has been set.

### GetPeakKvaLastSevenDaysTime

`func (o *PowerDataPayLoad) GetPeakKvaLastSevenDaysTime() int32`

GetPeakKvaLastSevenDaysTime returns the PeakKvaLastSevenDaysTime field if non-nil, zero value otherwise.

### GetPeakKvaLastSevenDaysTimeOk

`func (o *PowerDataPayLoad) GetPeakKvaLastSevenDaysTimeOk() (*int32, bool)`

GetPeakKvaLastSevenDaysTimeOk returns a tuple with the PeakKvaLastSevenDaysTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakKvaLastSevenDaysTime

`func (o *PowerDataPayLoad) SetPeakKvaLastSevenDaysTime(v int32)`

SetPeakKvaLastSevenDaysTime sets PeakKvaLastSevenDaysTime field to given value.

### HasPeakKvaLastSevenDaysTime

`func (o *PowerDataPayLoad) HasPeakKvaLastSevenDaysTime() bool`

HasPeakKvaLastSevenDaysTime returns a boolean if a field has been set.

### GetPercentageKva

`func (o *PowerDataPayLoad) GetPercentageKva() float32`

GetPercentageKva returns the PercentageKva field if non-nil, zero value otherwise.

### GetPercentageKvaOk

`func (o *PowerDataPayLoad) GetPercentageKvaOk() (*float32, bool)`

GetPercentageKvaOk returns a tuple with the PercentageKva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageKva

`func (o *PowerDataPayLoad) SetPercentageKva(v float32)`

SetPercentageKva sets PercentageKva field to given value.

### HasPercentageKva

`func (o *PowerDataPayLoad) HasPercentageKva() bool`

HasPercentageKva returns a boolean if a field has been set.

### GetPowerFactor

`func (o *PowerDataPayLoad) GetPowerFactor() string`

GetPowerFactor returns the PowerFactor field if non-nil, zero value otherwise.

### GetPowerFactorOk

`func (o *PowerDataPayLoad) GetPowerFactorOk() (*string, bool)`

GetPowerFactorOk returns a tuple with the PowerFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFactor

`func (o *PowerDataPayLoad) SetPowerFactor(v string)`

SetPowerFactor sets PowerFactor field to given value.

### HasPowerFactor

`func (o *PowerDataPayLoad) HasPowerFactor() bool`

HasPowerFactor returns a boolean if a field has been set.

### GetPrimaryKva

`func (o *PowerDataPayLoad) GetPrimaryKva() float32`

GetPrimaryKva returns the PrimaryKva field if non-nil, zero value otherwise.

### GetPrimaryKvaOk

`func (o *PowerDataPayLoad) GetPrimaryKvaOk() (*float32, bool)`

GetPrimaryKvaOk returns a tuple with the PrimaryKva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKva

`func (o *PowerDataPayLoad) SetPrimaryKva(v float32)`

SetPrimaryKva sets PrimaryKva field to given value.

### HasPrimaryKva

`func (o *PowerDataPayLoad) HasPrimaryKva() bool`

HasPrimaryKva returns a boolean if a field has been set.

### GetReadingTime

`func (o *PowerDataPayLoad) GetReadingTime() string`

GetReadingTime returns the ReadingTime field if non-nil, zero value otherwise.

### GetReadingTimeOk

`func (o *PowerDataPayLoad) GetReadingTimeOk() (*string, bool)`

GetReadingTimeOk returns a tuple with the ReadingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingTime

`func (o *PowerDataPayLoad) SetReadingTime(v string)`

SetReadingTime sets ReadingTime field to given value.

### HasReadingTime

`func (o *PowerDataPayLoad) HasReadingTime() bool`

HasReadingTime returns a boolean if a field has been set.

### GetRedundantKva

`func (o *PowerDataPayLoad) GetRedundantKva() float32`

GetRedundantKva returns the RedundantKva field if non-nil, zero value otherwise.

### GetRedundantKvaOk

`func (o *PowerDataPayLoad) GetRedundantKvaOk() (*float32, bool)`

GetRedundantKvaOk returns a tuple with the RedundantKva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantKva

`func (o *PowerDataPayLoad) SetRedundantKva(v float32)`

SetRedundantKva sets RedundantKva field to given value.

### HasRedundantKva

`func (o *PowerDataPayLoad) HasRedundantKva() bool`

HasRedundantKva returns a boolean if a field has been set.

### GetSoldAmps

`func (o *PowerDataPayLoad) GetSoldAmps() int32`

GetSoldAmps returns the SoldAmps field if non-nil, zero value otherwise.

### GetSoldAmpsOk

`func (o *PowerDataPayLoad) GetSoldAmpsOk() (*int32, bool)`

GetSoldAmpsOk returns a tuple with the SoldAmps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldAmps

`func (o *PowerDataPayLoad) SetSoldAmps(v int32)`

SetSoldAmps sets SoldAmps field to given value.

### HasSoldAmps

`func (o *PowerDataPayLoad) HasSoldAmps() bool`

HasSoldAmps returns a boolean if a field has been set.

### GetSoldKva

`func (o *PowerDataPayLoad) GetSoldKva() float32`

GetSoldKva returns the SoldKva field if non-nil, zero value otherwise.

### GetSoldKvaOk

`func (o *PowerDataPayLoad) GetSoldKvaOk() (*float32, bool)`

GetSoldKvaOk returns a tuple with the SoldKva field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldKva

`func (o *PowerDataPayLoad) SetSoldKva(v float32)`

SetSoldKva sets SoldKva field to given value.

### HasSoldKva

`func (o *PowerDataPayLoad) HasSoldKva() bool`

HasSoldKva returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


