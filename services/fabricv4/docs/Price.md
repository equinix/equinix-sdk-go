# Price

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | An absolute URL that returns specified pricing data | [optional] 
**Type** | Pointer to [**ProductType**](ProductType.md) |  | [optional] 
**Code** | Pointer to **string** | Equinix-assigned product code | [optional] 
**Name** | Pointer to **string** | Full product name | [optional] 
**Description** | Pointer to **string** | Product description | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**Charges** | Pointer to [**[]PriceCharge**](PriceCharge.md) |  | [optional] 
**Currency** | Pointer to **string** | Product offering price currency | [optional] 
**TermLength** | Pointer to [**PriceTermLength**](PriceTermLength.md) |  | [optional] 
**Catgory** | Pointer to [**PriceCategory**](PriceCategory.md) |  | [optional] 
**Connection** | Pointer to [**VirtualConnectionPrice**](VirtualConnectionPrice.md) |  | [optional] 
**IpBlock** | Pointer to [**IpBlockPrice**](IpBlockPrice.md) |  | [optional] 
**Router** | Pointer to [**FabricCloudRouterPrice**](FabricCloudRouterPrice.md) |  | [optional] 
**Port** | Pointer to [**VirtualPortPrice**](VirtualPortPrice.md) |  | [optional] 

## Methods

### NewPrice

`func NewPrice() *Price`

NewPrice instantiates a new Price object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithDefaults

`func NewPriceWithDefaults() *Price`

NewPriceWithDefaults instantiates a new Price object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Price) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Price) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Price) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Price) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *Price) GetType() ProductType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Price) GetTypeOk() (*ProductType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Price) SetType(v ProductType)`

SetType sets Type field to given value.

### HasType

`func (o *Price) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *Price) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Price) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Price) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Price) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *Price) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Price) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Price) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Price) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Price) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Price) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Price) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Price) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccount

`func (o *Price) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Price) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Price) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Price) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCharges

`func (o *Price) GetCharges() []PriceCharge`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *Price) GetChargesOk() (*[]PriceCharge, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *Price) SetCharges(v []PriceCharge)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *Price) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetCurrency

`func (o *Price) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Price) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Price) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Price) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTermLength

`func (o *Price) GetTermLength() PriceTermLength`

GetTermLength returns the TermLength field if non-nil, zero value otherwise.

### GetTermLengthOk

`func (o *Price) GetTermLengthOk() (*PriceTermLength, bool)`

GetTermLengthOk returns a tuple with the TermLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermLength

`func (o *Price) SetTermLength(v PriceTermLength)`

SetTermLength sets TermLength field to given value.

### HasTermLength

`func (o *Price) HasTermLength() bool`

HasTermLength returns a boolean if a field has been set.

### GetCatgory

`func (o *Price) GetCatgory() PriceCategory`

GetCatgory returns the Catgory field if non-nil, zero value otherwise.

### GetCatgoryOk

`func (o *Price) GetCatgoryOk() (*PriceCategory, bool)`

GetCatgoryOk returns a tuple with the Catgory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatgory

`func (o *Price) SetCatgory(v PriceCategory)`

SetCatgory sets Catgory field to given value.

### HasCatgory

`func (o *Price) HasCatgory() bool`

HasCatgory returns a boolean if a field has been set.

### GetConnection

`func (o *Price) GetConnection() VirtualConnectionPrice`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *Price) GetConnectionOk() (*VirtualConnectionPrice, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *Price) SetConnection(v VirtualConnectionPrice)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *Price) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetIpBlock

`func (o *Price) GetIpBlock() IpBlockPrice`

GetIpBlock returns the IpBlock field if non-nil, zero value otherwise.

### GetIpBlockOk

`func (o *Price) GetIpBlockOk() (*IpBlockPrice, bool)`

GetIpBlockOk returns a tuple with the IpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlock

`func (o *Price) SetIpBlock(v IpBlockPrice)`

SetIpBlock sets IpBlock field to given value.

### HasIpBlock

`func (o *Price) HasIpBlock() bool`

HasIpBlock returns a boolean if a field has been set.

### GetRouter

`func (o *Price) GetRouter() FabricCloudRouterPrice`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *Price) GetRouterOk() (*FabricCloudRouterPrice, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *Price) SetRouter(v FabricCloudRouterPrice)`

SetRouter sets Router field to given value.

### HasRouter

`func (o *Price) HasRouter() bool`

HasRouter returns a boolean if a field has been set.

### GetPort

`func (o *Price) GetPort() VirtualPortPrice`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Price) GetPortOk() (*VirtualPortPrice, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Price) SetPort(v VirtualPortPrice)`

SetPort sets Port field to given value.

### HasPort

`func (o *Price) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


