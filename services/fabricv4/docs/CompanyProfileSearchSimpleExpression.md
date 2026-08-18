# CompanyProfileSearchSimpleExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | [**CompanyProfileSearchFieldName**](CompanyProfileSearchFieldName.md) | Searchable field name. Properties are grouped by their supported operators: String properties (support all operators):  * &#x60;/name&#x60; - Company profile name  * &#x60;/uuid&#x60; - Company profile UUID  * &#x60;/tags/name&#x60; - Tag name  * &#x60;/tags/uuid&#x60; - Tag UUID  * &#x60;/tags/displayName&#x60; - Tag display name  Discrete value properties (only support &#x60;&#x3D;&#x60;, &#x60;!&#x3D;&#x60;, &#x60;IN&#x60;, &#x60;NOT IN&#x60;):  * &#x60;/state&#x60; - Company profile state (&#x60;PENDING&#x60;, &#x60;PROVISIONED&#x60;, &#x60;DEPROVISIONED&#x60;, &#x60;REJECTED&#x60;)  * &#x60;/metros/metroCode&#x60; - Metro code (e.g. &#x60;SV&#x60;, &#x60;NY&#x60;, &#x60;DC&#x60;)  * &#x60;/change/status&#x60; - Change status (&#x60;PENDING&#x60;, &#x60;COMPLETED&#x60;, &#x60;REJECTED&#x60;) — seller and admin users only  | 
**Operator** | [**OperatorEnum**](OperatorEnum.md) | Comparison operator. &#x60;LIKE&#x60;, &#x60;NOT LIKE&#x60;, &#x60;ILIKE&#x60;, and &#x60;NOT ILIKE&#x60; require exactly one value. All other operators accept one or more values.  * &#x60;&#x3D;&#x60; - equal; equivalent to &#x60;IN&#x60; when multiple values are provided  * &#x60;!&#x3D;&#x60; - not equal; equivalent to &#x60;NOT IN&#x60; when multiple values are provided  * &#x60;LIKE&#x60; - case-sensitive partial match (single value)  * &#x60;NOT LIKE&#x60; - case-sensitive partial non-match (single value)  * &#x60;ILIKE&#x60; - case-insensitive partial match (single value)  * &#x60;NOT ILIKE&#x60; - case-insensitive partial non-match (single value)  * &#x60;IN&#x60; - matches any of the provided values  * &#x60;NOT IN&#x60; - does not match any of the provided values  | 
**Values** | **[]string** |  | 

## Methods

### NewCompanyProfileSearchSimpleExpression

`func NewCompanyProfileSearchSimpleExpression(property CompanyProfileSearchFieldName, operator OperatorEnum, values []string, ) *CompanyProfileSearchSimpleExpression`

NewCompanyProfileSearchSimpleExpression instantiates a new CompanyProfileSearchSimpleExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyProfileSearchSimpleExpressionWithDefaults

`func NewCompanyProfileSearchSimpleExpressionWithDefaults() *CompanyProfileSearchSimpleExpression`

NewCompanyProfileSearchSimpleExpressionWithDefaults instantiates a new CompanyProfileSearchSimpleExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *CompanyProfileSearchSimpleExpression) GetProperty() CompanyProfileSearchFieldName`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CompanyProfileSearchSimpleExpression) GetPropertyOk() (*CompanyProfileSearchFieldName, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CompanyProfileSearchSimpleExpression) SetProperty(v CompanyProfileSearchFieldName)`

SetProperty sets Property field to given value.


### GetOperator

`func (o *CompanyProfileSearchSimpleExpression) GetOperator() OperatorEnum`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CompanyProfileSearchSimpleExpression) GetOperatorOk() (*OperatorEnum, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CompanyProfileSearchSimpleExpression) SetOperator(v OperatorEnum)`

SetOperator sets Operator field to given value.


### GetValues

`func (o *CompanyProfileSearchSimpleExpression) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CompanyProfileSearchSimpleExpression) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CompanyProfileSearchSimpleExpression) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


