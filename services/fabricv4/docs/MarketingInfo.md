# MarketingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logo** | Pointer to **string** | Logo file name | [optional] 
**Promotion** | Pointer to **bool** | Profile promotion on marketplace | [optional] 
**ProcessSteps** | Pointer to [**[]ProcessStep**](ProcessStep.md) |  | [optional] 

## Methods

### NewMarketingInfo

`func NewMarketingInfo() *MarketingInfo`

NewMarketingInfo instantiates a new MarketingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingInfoWithDefaults

`func NewMarketingInfoWithDefaults() *MarketingInfo`

NewMarketingInfoWithDefaults instantiates a new MarketingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogo

`func (o *MarketingInfo) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *MarketingInfo) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *MarketingInfo) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *MarketingInfo) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetPromotion

`func (o *MarketingInfo) GetPromotion() bool`

GetPromotion returns the Promotion field if non-nil, zero value otherwise.

### GetPromotionOk

`func (o *MarketingInfo) GetPromotionOk() (*bool, bool)`

GetPromotionOk returns a tuple with the Promotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotion

`func (o *MarketingInfo) SetPromotion(v bool)`

SetPromotion sets Promotion field to given value.

### HasPromotion

`func (o *MarketingInfo) HasPromotion() bool`

HasPromotion returns a boolean if a field has been set.

### GetProcessSteps

`func (o *MarketingInfo) GetProcessSteps() []ProcessStep`

GetProcessSteps returns the ProcessSteps field if non-nil, zero value otherwise.

### GetProcessStepsOk

`func (o *MarketingInfo) GetProcessStepsOk() (*[]ProcessStep, bool)`

GetProcessStepsOk returns a tuple with the ProcessSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessSteps

`func (o *MarketingInfo) SetProcessSteps(v []ProcessStep)`

SetProcessSteps sets ProcessSteps field to given value.

### HasProcessSteps

`func (o *MarketingInfo) HasProcessSteps() bool`

HasProcessSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


