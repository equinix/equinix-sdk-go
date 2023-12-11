# OrganizationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enforce2faAt** | Pointer to **time.Time** | Force to all members to have enabled the two factor authentication after that date, unless the value is null | [optional] 
**Logo** | Pointer to **string** | The logo for the organization; must be base64-encoded image data | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Twitter** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganizationInput

`func NewOrganizationInput() *OrganizationInput`

NewOrganizationInput instantiates a new OrganizationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInputWithDefaults

`func NewOrganizationInputWithDefaults() *OrganizationInput`

NewOrganizationInputWithDefaults instantiates a new OrganizationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *OrganizationInput) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrganizationInput) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrganizationInput) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *OrganizationInput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *OrganizationInput) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *OrganizationInput) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *OrganizationInput) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *OrganizationInput) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCustomdata

`func (o *OrganizationInput) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *OrganizationInput) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *OrganizationInput) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *OrganizationInput) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnforce2faAt

`func (o *OrganizationInput) GetEnforce2faAt() time.Time`

GetEnforce2faAt returns the Enforce2faAt field if non-nil, zero value otherwise.

### GetEnforce2faAtOk

`func (o *OrganizationInput) GetEnforce2faAtOk() (*time.Time, bool)`

GetEnforce2faAtOk returns a tuple with the Enforce2faAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforce2faAt

`func (o *OrganizationInput) SetEnforce2faAt(v time.Time)`

SetEnforce2faAt sets Enforce2faAt field to given value.

### HasEnforce2faAt

`func (o *OrganizationInput) HasEnforce2faAt() bool`

HasEnforce2faAt returns a boolean if a field has been set.

### GetLogo

`func (o *OrganizationInput) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *OrganizationInput) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *OrganizationInput) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *OrganizationInput) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetName

`func (o *OrganizationInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTwitter

`func (o *OrganizationInput) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *OrganizationInput) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *OrganizationInput) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *OrganizationInput) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### GetWebsite

`func (o *OrganizationInput) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *OrganizationInput) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *OrganizationInput) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *OrganizationInput) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


