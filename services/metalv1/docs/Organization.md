# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreditAmount** | Pointer to **float32** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enforce2faAt** | Pointer to **time.Time** | Force to all members to have enabled the two factor authentication after that date, unless the value is null | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Members** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Memberships** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Projects** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Terms** | Pointer to **int32** |  | [optional] 
**Twitter** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Organization) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Organization) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Organization) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Organization) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *Organization) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *Organization) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *Organization) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *Organization) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Organization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Organization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Organization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Organization) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreditAmount

`func (o *Organization) GetCreditAmount() float32`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *Organization) GetCreditAmountOk() (*float32, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *Organization) SetCreditAmount(v float32)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *Organization) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.

### GetCustomdata

`func (o *Organization) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *Organization) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *Organization) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *Organization) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDescription

`func (o *Organization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Organization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Organization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Organization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnforce2faAt

`func (o *Organization) GetEnforce2faAt() time.Time`

GetEnforce2faAt returns the Enforce2faAt field if non-nil, zero value otherwise.

### GetEnforce2faAtOk

`func (o *Organization) GetEnforce2faAtOk() (*time.Time, bool)`

GetEnforce2faAtOk returns a tuple with the Enforce2faAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforce2faAt

`func (o *Organization) SetEnforce2faAt(v time.Time)`

SetEnforce2faAt sets Enforce2faAt field to given value.

### HasEnforce2faAt

`func (o *Organization) HasEnforce2faAt() bool`

HasEnforce2faAt returns a boolean if a field has been set.

### GetHref

`func (o *Organization) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Organization) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Organization) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Organization) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Organization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogo

`func (o *Organization) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Organization) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Organization) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Organization) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetMembers

`func (o *Organization) GetMembers() []Href`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Organization) GetMembersOk() (*[]Href, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Organization) SetMembers(v []Href)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Organization) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMemberships

`func (o *Organization) GetMemberships() []Href`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *Organization) GetMembershipsOk() (*[]Href, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *Organization) SetMemberships(v []Href)`

SetMemberships sets Memberships field to given value.

### HasMemberships

`func (o *Organization) HasMemberships() bool`

HasMemberships returns a boolean if a field has been set.

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjects

`func (o *Organization) GetProjects() []Href`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *Organization) GetProjectsOk() (*[]Href, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *Organization) SetProjects(v []Href)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *Organization) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetTerms

`func (o *Organization) GetTerms() int32`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *Organization) GetTermsOk() (*int32, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *Organization) SetTerms(v int32)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *Organization) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetTwitter

`func (o *Organization) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *Organization) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *Organization) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *Organization) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Organization) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Organization) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Organization) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Organization) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWebsite

`func (o *Organization) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Organization) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Organization) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *Organization) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


