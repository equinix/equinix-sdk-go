# OrderSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **int32** |  | [optional] 
**AgreementId** | Pointer to **string** |  | [optional] 
**Charges** | Pointer to [**[]DeviceElement**](DeviceElement.md) |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**EsignAgreementId** | Pointer to **string** |  | [optional] 
**IbxCountry** | Pointer to **string** |  | [optional] 
**IbxRegion** | Pointer to **string** |  | [optional] 
**InitialTerm** | Pointer to **int32** |  | [optional] 
**Metro** | Pointer to **string** |  | [optional] 
**MonthlyRecurringCharges** | Pointer to **float64** |  | [optional] 
**NonRecurringCharges** | Pointer to **float64** |  | [optional] 
**NonRenewalNotice** | Pointer to **string** |  | [optional] 
**OrderTerms** | Pointer to **string** |  | [optional] 
**PiPercentage** | Pointer to **string** |  | [optional] 
**ProductDescription** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**QuoteContentType** | Pointer to **string** |  | [optional] 
**QuoteFileName** | Pointer to **string** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**RenewalPeriod** | Pointer to **int32** |  | [optional] 
**RequestSignType** | Pointer to **string** |  | [optional] 
**SignStatus** | Pointer to **string** |  | [optional] 
**SignType** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalCharges** | Pointer to **float64** |  | [optional] 

## Methods

### NewOrderSummaryResponse

`func NewOrderSummaryResponse() *OrderSummaryResponse`

NewOrderSummaryResponse instantiates a new OrderSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSummaryResponseWithDefaults

`func NewOrderSummaryResponseWithDefaults() *OrderSummaryResponse`

NewOrderSummaryResponseWithDefaults instantiates a new OrderSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *OrderSummaryResponse) GetAccountNumber() int32`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *OrderSummaryResponse) GetAccountNumberOk() (*int32, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *OrderSummaryResponse) SetAccountNumber(v int32)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *OrderSummaryResponse) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAgreementId

`func (o *OrderSummaryResponse) GetAgreementId() string`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *OrderSummaryResponse) GetAgreementIdOk() (*string, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *OrderSummaryResponse) SetAgreementId(v string)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *OrderSummaryResponse) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### GetCharges

`func (o *OrderSummaryResponse) GetCharges() []DeviceElement`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *OrderSummaryResponse) GetChargesOk() (*[]DeviceElement, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *OrderSummaryResponse) SetCharges(v []DeviceElement)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *OrderSummaryResponse) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetCurrency

`func (o *OrderSummaryResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrderSummaryResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrderSummaryResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrderSummaryResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetErrorCode

`func (o *OrderSummaryResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OrderSummaryResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OrderSummaryResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OrderSummaryResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *OrderSummaryResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *OrderSummaryResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *OrderSummaryResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *OrderSummaryResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetEsignAgreementId

`func (o *OrderSummaryResponse) GetEsignAgreementId() string`

GetEsignAgreementId returns the EsignAgreementId field if non-nil, zero value otherwise.

### GetEsignAgreementIdOk

`func (o *OrderSummaryResponse) GetEsignAgreementIdOk() (*string, bool)`

GetEsignAgreementIdOk returns a tuple with the EsignAgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsignAgreementId

`func (o *OrderSummaryResponse) SetEsignAgreementId(v string)`

SetEsignAgreementId sets EsignAgreementId field to given value.

### HasEsignAgreementId

`func (o *OrderSummaryResponse) HasEsignAgreementId() bool`

HasEsignAgreementId returns a boolean if a field has been set.

### GetIbxCountry

`func (o *OrderSummaryResponse) GetIbxCountry() string`

GetIbxCountry returns the IbxCountry field if non-nil, zero value otherwise.

### GetIbxCountryOk

`func (o *OrderSummaryResponse) GetIbxCountryOk() (*string, bool)`

GetIbxCountryOk returns a tuple with the IbxCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbxCountry

`func (o *OrderSummaryResponse) SetIbxCountry(v string)`

SetIbxCountry sets IbxCountry field to given value.

### HasIbxCountry

`func (o *OrderSummaryResponse) HasIbxCountry() bool`

HasIbxCountry returns a boolean if a field has been set.

### GetIbxRegion

`func (o *OrderSummaryResponse) GetIbxRegion() string`

GetIbxRegion returns the IbxRegion field if non-nil, zero value otherwise.

### GetIbxRegionOk

`func (o *OrderSummaryResponse) GetIbxRegionOk() (*string, bool)`

GetIbxRegionOk returns a tuple with the IbxRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbxRegion

`func (o *OrderSummaryResponse) SetIbxRegion(v string)`

SetIbxRegion sets IbxRegion field to given value.

### HasIbxRegion

`func (o *OrderSummaryResponse) HasIbxRegion() bool`

HasIbxRegion returns a boolean if a field has been set.

### GetInitialTerm

`func (o *OrderSummaryResponse) GetInitialTerm() int32`

GetInitialTerm returns the InitialTerm field if non-nil, zero value otherwise.

### GetInitialTermOk

`func (o *OrderSummaryResponse) GetInitialTermOk() (*int32, bool)`

GetInitialTermOk returns a tuple with the InitialTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialTerm

`func (o *OrderSummaryResponse) SetInitialTerm(v int32)`

SetInitialTerm sets InitialTerm field to given value.

### HasInitialTerm

`func (o *OrderSummaryResponse) HasInitialTerm() bool`

HasInitialTerm returns a boolean if a field has been set.

### GetMetro

`func (o *OrderSummaryResponse) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *OrderSummaryResponse) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *OrderSummaryResponse) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *OrderSummaryResponse) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetMonthlyRecurringCharges

`func (o *OrderSummaryResponse) GetMonthlyRecurringCharges() float64`

GetMonthlyRecurringCharges returns the MonthlyRecurringCharges field if non-nil, zero value otherwise.

### GetMonthlyRecurringChargesOk

`func (o *OrderSummaryResponse) GetMonthlyRecurringChargesOk() (*float64, bool)`

GetMonthlyRecurringChargesOk returns a tuple with the MonthlyRecurringCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyRecurringCharges

`func (o *OrderSummaryResponse) SetMonthlyRecurringCharges(v float64)`

SetMonthlyRecurringCharges sets MonthlyRecurringCharges field to given value.

### HasMonthlyRecurringCharges

`func (o *OrderSummaryResponse) HasMonthlyRecurringCharges() bool`

HasMonthlyRecurringCharges returns a boolean if a field has been set.

### GetNonRecurringCharges

`func (o *OrderSummaryResponse) GetNonRecurringCharges() float64`

GetNonRecurringCharges returns the NonRecurringCharges field if non-nil, zero value otherwise.

### GetNonRecurringChargesOk

`func (o *OrderSummaryResponse) GetNonRecurringChargesOk() (*float64, bool)`

GetNonRecurringChargesOk returns a tuple with the NonRecurringCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonRecurringCharges

`func (o *OrderSummaryResponse) SetNonRecurringCharges(v float64)`

SetNonRecurringCharges sets NonRecurringCharges field to given value.

### HasNonRecurringCharges

`func (o *OrderSummaryResponse) HasNonRecurringCharges() bool`

HasNonRecurringCharges returns a boolean if a field has been set.

### GetNonRenewalNotice

`func (o *OrderSummaryResponse) GetNonRenewalNotice() string`

GetNonRenewalNotice returns the NonRenewalNotice field if non-nil, zero value otherwise.

### GetNonRenewalNoticeOk

`func (o *OrderSummaryResponse) GetNonRenewalNoticeOk() (*string, bool)`

GetNonRenewalNoticeOk returns a tuple with the NonRenewalNotice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonRenewalNotice

`func (o *OrderSummaryResponse) SetNonRenewalNotice(v string)`

SetNonRenewalNotice sets NonRenewalNotice field to given value.

### HasNonRenewalNotice

`func (o *OrderSummaryResponse) HasNonRenewalNotice() bool`

HasNonRenewalNotice returns a boolean if a field has been set.

### GetOrderTerms

`func (o *OrderSummaryResponse) GetOrderTerms() string`

GetOrderTerms returns the OrderTerms field if non-nil, zero value otherwise.

### GetOrderTermsOk

`func (o *OrderSummaryResponse) GetOrderTermsOk() (*string, bool)`

GetOrderTermsOk returns a tuple with the OrderTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTerms

`func (o *OrderSummaryResponse) SetOrderTerms(v string)`

SetOrderTerms sets OrderTerms field to given value.

### HasOrderTerms

`func (o *OrderSummaryResponse) HasOrderTerms() bool`

HasOrderTerms returns a boolean if a field has been set.

### GetPiPercentage

`func (o *OrderSummaryResponse) GetPiPercentage() string`

GetPiPercentage returns the PiPercentage field if non-nil, zero value otherwise.

### GetPiPercentageOk

`func (o *OrderSummaryResponse) GetPiPercentageOk() (*string, bool)`

GetPiPercentageOk returns a tuple with the PiPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiPercentage

`func (o *OrderSummaryResponse) SetPiPercentage(v string)`

SetPiPercentage sets PiPercentage field to given value.

### HasPiPercentage

`func (o *OrderSummaryResponse) HasPiPercentage() bool`

HasPiPercentage returns a boolean if a field has been set.

### GetProductDescription

`func (o *OrderSummaryResponse) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *OrderSummaryResponse) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *OrderSummaryResponse) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.

### HasProductDescription

`func (o *OrderSummaryResponse) HasProductDescription() bool`

HasProductDescription returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderSummaryResponse) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderSummaryResponse) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderSummaryResponse) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderSummaryResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuoteContentType

`func (o *OrderSummaryResponse) GetQuoteContentType() string`

GetQuoteContentType returns the QuoteContentType field if non-nil, zero value otherwise.

### GetQuoteContentTypeOk

`func (o *OrderSummaryResponse) GetQuoteContentTypeOk() (*string, bool)`

GetQuoteContentTypeOk returns a tuple with the QuoteContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteContentType

`func (o *OrderSummaryResponse) SetQuoteContentType(v string)`

SetQuoteContentType sets QuoteContentType field to given value.

### HasQuoteContentType

`func (o *OrderSummaryResponse) HasQuoteContentType() bool`

HasQuoteContentType returns a boolean if a field has been set.

### GetQuoteFileName

`func (o *OrderSummaryResponse) GetQuoteFileName() string`

GetQuoteFileName returns the QuoteFileName field if non-nil, zero value otherwise.

### GetQuoteFileNameOk

`func (o *OrderSummaryResponse) GetQuoteFileNameOk() (*string, bool)`

GetQuoteFileNameOk returns a tuple with the QuoteFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteFileName

`func (o *OrderSummaryResponse) SetQuoteFileName(v string)`

SetQuoteFileName sets QuoteFileName field to given value.

### HasQuoteFileName

`func (o *OrderSummaryResponse) HasQuoteFileName() bool`

HasQuoteFileName returns a boolean if a field has been set.

### GetReferenceId

`func (o *OrderSummaryResponse) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *OrderSummaryResponse) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *OrderSummaryResponse) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *OrderSummaryResponse) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetRenewalPeriod

`func (o *OrderSummaryResponse) GetRenewalPeriod() int32`

GetRenewalPeriod returns the RenewalPeriod field if non-nil, zero value otherwise.

### GetRenewalPeriodOk

`func (o *OrderSummaryResponse) GetRenewalPeriodOk() (*int32, bool)`

GetRenewalPeriodOk returns a tuple with the RenewalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPeriod

`func (o *OrderSummaryResponse) SetRenewalPeriod(v int32)`

SetRenewalPeriod sets RenewalPeriod field to given value.

### HasRenewalPeriod

`func (o *OrderSummaryResponse) HasRenewalPeriod() bool`

HasRenewalPeriod returns a boolean if a field has been set.

### GetRequestSignType

`func (o *OrderSummaryResponse) GetRequestSignType() string`

GetRequestSignType returns the RequestSignType field if non-nil, zero value otherwise.

### GetRequestSignTypeOk

`func (o *OrderSummaryResponse) GetRequestSignTypeOk() (*string, bool)`

GetRequestSignTypeOk returns a tuple with the RequestSignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSignType

`func (o *OrderSummaryResponse) SetRequestSignType(v string)`

SetRequestSignType sets RequestSignType field to given value.

### HasRequestSignType

`func (o *OrderSummaryResponse) HasRequestSignType() bool`

HasRequestSignType returns a boolean if a field has been set.

### GetSignStatus

`func (o *OrderSummaryResponse) GetSignStatus() string`

GetSignStatus returns the SignStatus field if non-nil, zero value otherwise.

### GetSignStatusOk

`func (o *OrderSummaryResponse) GetSignStatusOk() (*string, bool)`

GetSignStatusOk returns a tuple with the SignStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignStatus

`func (o *OrderSummaryResponse) SetSignStatus(v string)`

SetSignStatus sets SignStatus field to given value.

### HasSignStatus

`func (o *OrderSummaryResponse) HasSignStatus() bool`

HasSignStatus returns a boolean if a field has been set.

### GetSignType

`func (o *OrderSummaryResponse) GetSignType() string`

GetSignType returns the SignType field if non-nil, zero value otherwise.

### GetSignTypeOk

`func (o *OrderSummaryResponse) GetSignTypeOk() (*string, bool)`

GetSignTypeOk returns a tuple with the SignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignType

`func (o *OrderSummaryResponse) SetSignType(v string)`

SetSignType sets SignType field to given value.

### HasSignType

`func (o *OrderSummaryResponse) HasSignType() bool`

HasSignType returns a boolean if a field has been set.

### GetSpeed

`func (o *OrderSummaryResponse) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *OrderSummaryResponse) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *OrderSummaryResponse) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *OrderSummaryResponse) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *OrderSummaryResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderSummaryResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderSummaryResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderSummaryResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalCharges

`func (o *OrderSummaryResponse) GetTotalCharges() float64`

GetTotalCharges returns the TotalCharges field if non-nil, zero value otherwise.

### GetTotalChargesOk

`func (o *OrderSummaryResponse) GetTotalChargesOk() (*float64, bool)`

GetTotalChargesOk returns a tuple with the TotalCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCharges

`func (o *OrderSummaryResponse) SetTotalCharges(v float64)`

SetTotalCharges sets TotalCharges field to given value.

### HasTotalCharges

`func (o *OrderSummaryResponse) HasTotalCharges() bool`

HasTotalCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


