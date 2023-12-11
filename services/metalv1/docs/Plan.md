# Plan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableIn** | Pointer to [**[]PlanAvailableInInner**](PlanAvailableInInner.md) | Shows which facilities the plan is available in, and the facility-based price if it is different from the default price. | [optional] 
**AvailableInMetros** | Pointer to [**[]PlanAvailableInMetrosInner**](PlanAvailableInMetrosInner.md) | Shows which metros the plan is available in, and the metro-based price if it is different from the default price. | [optional] 
**Categories** | Pointer to **[]string** | Categories of the plan, like compute or storage. A Plan can belong to multiple categories. | [optional] 
**Class** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DeploymentTypes** | Pointer to [**[]PlanDeploymentTypesInner**](PlanDeploymentTypesInner.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Legacy** | Pointer to **bool** | Deprecated. Always return false | [optional] 
**Line** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Pricing** | Pointer to **map[string]interface{}** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Specs** | Pointer to [**PlanSpecs**](PlanSpecs.md) |  | [optional] 
**Type** | Pointer to [**PlanType**](PlanType.md) |  | [optional] 

## Methods

### NewPlan

`func NewPlan() *Plan`

NewPlan instantiates a new Plan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWithDefaults

`func NewPlanWithDefaults() *Plan`

NewPlanWithDefaults instantiates a new Plan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableIn

`func (o *Plan) GetAvailableIn() []PlanAvailableInInner`

GetAvailableIn returns the AvailableIn field if non-nil, zero value otherwise.

### GetAvailableInOk

`func (o *Plan) GetAvailableInOk() (*[]PlanAvailableInInner, bool)`

GetAvailableInOk returns a tuple with the AvailableIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableIn

`func (o *Plan) SetAvailableIn(v []PlanAvailableInInner)`

SetAvailableIn sets AvailableIn field to given value.

### HasAvailableIn

`func (o *Plan) HasAvailableIn() bool`

HasAvailableIn returns a boolean if a field has been set.

### GetAvailableInMetros

`func (o *Plan) GetAvailableInMetros() []PlanAvailableInMetrosInner`

GetAvailableInMetros returns the AvailableInMetros field if non-nil, zero value otherwise.

### GetAvailableInMetrosOk

`func (o *Plan) GetAvailableInMetrosOk() (*[]PlanAvailableInMetrosInner, bool)`

GetAvailableInMetrosOk returns a tuple with the AvailableInMetros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableInMetros

`func (o *Plan) SetAvailableInMetros(v []PlanAvailableInMetrosInner)`

SetAvailableInMetros sets AvailableInMetros field to given value.

### HasAvailableInMetros

`func (o *Plan) HasAvailableInMetros() bool`

HasAvailableInMetros returns a boolean if a field has been set.

### GetCategories

`func (o *Plan) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Plan) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Plan) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Plan) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetClass

`func (o *Plan) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Plan) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *Plan) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *Plan) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetDescription

`func (o *Plan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Plan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Plan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Plan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeploymentTypes

`func (o *Plan) GetDeploymentTypes() []PlanDeploymentTypesInner`

GetDeploymentTypes returns the DeploymentTypes field if non-nil, zero value otherwise.

### GetDeploymentTypesOk

`func (o *Plan) GetDeploymentTypesOk() (*[]PlanDeploymentTypesInner, bool)`

GetDeploymentTypesOk returns a tuple with the DeploymentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTypes

`func (o *Plan) SetDeploymentTypes(v []PlanDeploymentTypesInner)`

SetDeploymentTypes sets DeploymentTypes field to given value.

### HasDeploymentTypes

`func (o *Plan) HasDeploymentTypes() bool`

HasDeploymentTypes returns a boolean if a field has been set.

### GetId

`func (o *Plan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Plan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLegacy

`func (o *Plan) GetLegacy() bool`

GetLegacy returns the Legacy field if non-nil, zero value otherwise.

### GetLegacyOk

`func (o *Plan) GetLegacyOk() (*bool, bool)`

GetLegacyOk returns a tuple with the Legacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacy

`func (o *Plan) SetLegacy(v bool)`

SetLegacy sets Legacy field to given value.

### HasLegacy

`func (o *Plan) HasLegacy() bool`

HasLegacy returns a boolean if a field has been set.

### GetLine

`func (o *Plan) GetLine() string`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *Plan) GetLineOk() (*string, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *Plan) SetLine(v string)`

SetLine sets Line field to given value.

### HasLine

`func (o *Plan) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetName

`func (o *Plan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Plan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Plan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Plan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPricing

`func (o *Plan) GetPricing() map[string]interface{}`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *Plan) GetPricingOk() (*map[string]interface{}, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *Plan) SetPricing(v map[string]interface{})`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *Plan) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetSlug

`func (o *Plan) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Plan) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Plan) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Plan) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSpecs

`func (o *Plan) GetSpecs() PlanSpecs`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *Plan) GetSpecsOk() (*PlanSpecs, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *Plan) SetSpecs(v PlanSpecs)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *Plan) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetType

`func (o *Plan) GetType() PlanType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Plan) GetTypeOk() (*PlanType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Plan) SetType(v PlanType)`

SetType sets Type field to given value.

### HasType

`func (o *Plan) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


